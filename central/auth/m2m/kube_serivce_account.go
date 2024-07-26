package m2m

import (
	"fmt"
	jose "github.com/go-jose/go-jose/v4"
	jwt "github.com/go-jose/go-jose/v4/jwt"
	"github.com/pkg/errors"
	"os"
)

func GetKubeServiceAccountIssuer() (string, error) {
	token, err := readServiceAccountToken()
	if err != nil {
		return "", errors.Wrap(err, "Failed to read kube service account token")
	}

	algos := []jose.SignatureAlgorithm{
		jose.EdDSA,
		jose.HS256,
		jose.HS384,
		jose.HS512,
		jose.RS256,
		jose.RS384,
		jose.RS512,
		jose.ES256,
		jose.ES384,
		jose.ES512,
		jose.PS256,
		jose.PS384,
		jose.PS512,
	}

	parsedJwt, err := jwt.ParseSigned(token, algos)
	if err != nil {
		return "", errors.Wrap(err, "Failed to parse service account JWT")
	}

	claims := jwt.Claims{}
	if err = parsedJwt.UnsafeClaimsWithoutVerification(&claims); err != nil {
		return "", errors.Wrap(err, "Failed to parse service account JWT claims")
	}

	return claims.Issuer, nil
}

func readServiceAccountToken() (string, error) {
	token, err := os.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/token")
	if err != nil {
		return "", fmt.Errorf("error reading service account token file: %v", err)
	}

	return string(token), nil
}
