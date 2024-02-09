package scan

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/daemon"
	"github.com/google/go-containerregistry/pkg/v1/tarball"
	"github.com/quay/claircore"
	"github.com/stackrox/rox/pkg/utils"
)

var _ Image = (*dockerLocalImage)(nil)

type Image interface {
	GetManifest(context.Context) (*claircore.Manifest, error)
}

type dockerLocalImage struct {
	image v1.Image
	path  string
}

func newImage(ref string) (*dockerLocalImage, error) {
	r, err := name.ParseReference(ref)
	if err != nil {
		return nil, fmt.Errorf("invalid reference %s: %w", ref, err)
	}

	f, err := os.CreateTemp("", "roxctl*.tar")
	if err != nil {
		return nil, fmt.Errorf("failed to open: %w", err)
	}
	defer utils.IgnoreError(f.Close)

	img, err := daemon.Image(r)
	if err == nil {
		tag, err := name.NewTag("roxctl-test-tag")
		utils.Must(err)
		err = tarball.MultiRefWrite(map[name.Reference]v1.Image{r: img, tag: img}, f)
		utils.Must(err) // , "could not write tarball")
		return newDockerLocalImage(f.Name())
	}
	//return newDockerLocalImage("/tmp/roxctl1634259876.tar", r)
	panic(err)

	return remoteImage(ref, err)
}

func remoteImage(ref string, err error) (*dockerLocalImage, error) {
	f, err := os.CreateTemp("", "roxctl*.tar")
	if err != nil {
		return nil, fmt.Errorf("failed to open: %w", err)
	}
	defer utils.IgnoreError(f.Close)
	desc, err := crane.Get(ref)
	var img v1.Image
	if desc.MediaType.IsSchema1() {
		img, err = desc.Schema1()
		if err != nil {
			return nil, fmt.Errorf("pulling schema 1 image %s: %w", ref, err)
		}
	} else {
		img, err = desc.Image()
		if err != nil {
			return nil, fmt.Errorf("pulling Image %s: %w", ref, err)
		}
	}

	file, err := os.Create(f.Name())
	if err != nil {
		return nil, err
	}
	r, err := name.ParseReference(ref)
	if err != nil {
		return nil, err
	}

	err = tarball.Write(r, img, file)
	if err != nil {
		return nil, err
	}
	return newDockerLocalImage(f.Name())
}

func newDockerLocalImage(imageTar string) (*dockerLocalImage, error) {
	opener := pathOpener(imageTar)
	manifest, err := tarball.LoadManifest(opener)
	utils.Must(err)
	spew.Dump(manifest)
	tag, err := name.NewTag(manifest[0].RepoTags[0])
	utils.Must(err)

	img, err := tarball.Image(opener, &tag)
	utils.Must(err) // , "can create tag")

	di := &dockerLocalImage{
		image: img,
		path:  imageTar,
	}

	return di, nil
}

func pathOpener(path string) tarball.Opener {
	return func() (io.ReadCloser, error) {
		return os.Open(path)
	}
}

func (i *dockerLocalImage) GetManifest(_ context.Context) (*claircore.Manifest, error) {
	hash, err := i.image.Digest()
	utils.Must(err)
	digest, err := claircore.ParseDigest(hash.String())
	utils.Must(err)

	l, err := i.image.Layers()
	utils.Must(err)

	layers := make([]*claircore.Layer, 0, len(l))
	for _, layer := range l {
		hash, err := layer.Digest()
		utils.Must(err)
		parseDigest, err := claircore.ParseDigest(hash.String())
		cl := claircore.Layer{
			Hash: parseDigest,
		}
		uncompressed, err := layer.Uncompressed()
		utils.Must(err)

		buff := bytes.NewBuffer([]byte{})
		_, err = io.Copy(buff, uncompressed)
		utils.Must(err)
		reader := bytes.NewReader(buff.Bytes())

		err = cl.Init(context.TODO(), &claircore.LayerDescription{
			Digest:    parseDigest.String(),
			MediaType: `application/vnd.oci.image.layer.nondistributable.v1.tar`,
		}, reader)
		utils.Must(err) // , "init layer")

		layers = append(layers, &cl)
	}

	return &claircore.Manifest{
		Hash:   digest,
		Layers: layers,
	}, nil
}
