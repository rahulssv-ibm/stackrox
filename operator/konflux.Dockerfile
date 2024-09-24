# TODO(ROX-20312): we can't pin image tag or digest because currently there's no mechanism to auto-update that.
# We're targeting a floating tag here which should be reasonably safe to do as both RHEL major 8 and Go major.minor 1.20 should provide enough stability.
FROM brew.registry.redhat.io/rh-osbs/openshift-golang-builder:rhel_8_1.22 as builder

WORKDIR /go/src/github.com/stackrox/rox/app

COPY . .

RUN .konflux/scripts/fail-build-if-git-is-dirty.sh

ARG VERSIONS_SUFFIX
ENV MAIN_TAG_SUFFIX="$VERSIONS_SUFFIX" COLLECTOR_TAG_SUFFIX="$VERSIONS_SUFFIX" SCANNER_TAG_SUFFIX="$VERSIONS_SUFFIX"

# Build the operator binary.
# TODO(ROX-24276): re-enable release builds for fast stream.
# TODO(ROX-20240): enable non-release development builds.
# GOTAGS="release"
ENV CI=1 GOFLAGS="" CGO_ENABLED=1

RUN GOOS=linux GOARCH=$(go env GOARCH) scripts/go-build-file.sh operator/cmd/main.go image/bin/operator


# TODO(ROX-20312): pin image tags when there's a process that updates them automatically.
FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

ARG MAIN_IMAGE_TAG
RUN if [[ "$MAIN_IMAGE_TAG" == "" ]]; then >&2 echo "error: required MAIN_IMAGE_TAG arg is unset"; exit 6; fi

LABEL \
    com.redhat.component="rhacs-operator-container" \
    com.redhat.license_terms="https://www.redhat.com/agreements" \
    description="Operator for Red Hat Advanced Cluster Security for Kubernetes" \
    io.k8s.description="Operator for Red Hat Advanced Cluster Security for Kubernetes" \
    io.k8s.display-name="operator" \
    io.openshift.tags="rhacs,operator,stackrox" \
    maintainer="Red Hat, Inc." \
    name="rhacs-rhel8-operator" \
    source-location="https://github.com/stackrox/stackrox" \
    summary="Operator for Red Hat Advanced Cluster Security for Kubernetes" \
    url="https://catalog.redhat.com/software/container-stacks/detail/60eefc88ee05ae7c5b8f041c" \
    # We must set version label to prevent inheriting value set in the base stage.
    version="${MAIN_IMAGE_TAG}" \
    # Release label is required by EC although has no practical semantics.
    # We also set it to not inherit one from a base stage in case it's RHEL or UBI.
    release="1"

COPY --from=builder /go/src/github.com/stackrox/rox/app/image/bin/operator /usr/local/bin/rhacs-operator

# TODO(ROX-20234): use hermetic builds when installing/updating RPMs becomes hermetic.
RUN microdnf upgrade -y --nobest && \
    microdnf clean all && \
    rpm --verbose -e --nodeps $(rpm -qa curl '*rpm*' '*dnf*' '*libsolv*' '*hawkey*' 'yum*') && \
    rm -rf /var/cache/dnf /var/cache/yum

COPY LICENSE /licenses/LICENSE

# TODO(ROX-22245): set proper image flavor for user-facing GA Fast Stream images.
ENV ROX_IMAGE_FLAVOR="development_build"

# The following are numeric uid and gid of `nobody` user in UBI.
# We can't use symbolic names because otherwise k8s will fail to start the pod with an error like this:
# Error: container has runAsNonRoot and image has non-numeric user (nobody), cannot verify user is non-root.
USER 65534:65534

ENTRYPOINT ["/usr/local/bin/rhacs-operator"]
