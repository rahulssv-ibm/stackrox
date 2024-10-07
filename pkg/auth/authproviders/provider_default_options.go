package authproviders

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/stackrox/rox/pkg/auth/permissions"
	"github.com/stackrox/rox/pkg/auth/tokens"
	"github.com/stackrox/rox/pkg/uuid"
)

var (
	noStoredInfoErr = errors.New("no storage information for auth provider")
)

// Default Options which fill in values only if not already set.
////////////////////////////////////////////////////////////////

// DefaultNewID sets the id of the provider to a new value if not already set.
func DefaultNewID() ProviderOption {
	return func(pr *providerImpl) (RevertOption, error) {
		if pr.storedInfo.GetId() != "" {
			return noOpRevert, nil
		}
		if pr.storedInfo == nil {
			return noOpRevert, noStoredInfoErr
		}
		oldID := pr.storedInfo.GetId()
		revert := func(pr *providerImpl) error {
			if pr.storedInfo == nil {
				return noStoredInfoErr
			}
			pr.storedInfo.Id = oldID
			return nil
		}
		pr.storedInfo.Id = uuid.NewV4().String()
		return revert, nil
	}
}

// DefaultLoginURL fills in the login url if not set using a function that creates a url for the provider id.
func DefaultLoginURL(fn func(authProviderID string) string) ProviderOption {
	return func(pr *providerImpl) (RevertOption, error) {
		if pr.storedInfo.GetLoginUrl() != "" {
			return noOpRevert, nil
		}
		if pr.storedInfo == nil {
			return noOpRevert, noStoredInfoErr
		}
		oldLoginURL := pr.storedInfo.GetLoginUrl()
		revert := func(pr *providerImpl) error {
			if pr.storedInfo == nil {
				return noStoredInfoErr
			}
			pr.storedInfo.LoginUrl = oldLoginURL
			return nil
		}
		pr.storedInfo.LoginUrl = fn(pr.storedInfo.Id)
		return revert, nil
	}
}

const tokenTTL = 12 * time.Hour

// DefaultTokenIssuerFromFactory sets the token issuer of the provider from the factory if not already set.
func DefaultTokenIssuerFromFactory(tf tokens.IssuerFactory) ProviderOption {
	return func(pr *providerImpl) (RevertOption, error) {
		if pr.issuer != nil {
			return noOpRevert, nil
		}
		oldIssuer := pr.issuer
		revert := func(pr *providerImpl) error {
			pr.issuer = oldIssuer
			return nil
		}
		issuer, err := tf.CreateIssuer(pr, tokens.WithTTL(tokenTTL))
		if err != nil {
			return noOpRevert, errors.Wrap(err, "failed to create issuer for newly created auth provider")
		}
		pr.issuer = issuer
		return revert, nil
	}
}

// DefaultRoleMapperOption loads a role mapper from the factory if one is not set on the provider.
func DefaultRoleMapperOption(fn func(id string) permissions.RoleMapper) ProviderOption {
	return func(pr *providerImpl) (RevertOption, error) {
		if pr.roleMapper != nil {
			return noOpRevert, nil
		}
		if pr.storedInfo.GetId() == "" {
			return noOpRevert, nil
		}
		oldRoleMapper := pr.roleMapper
		revert := func(pr *providerImpl) error {
			pr.roleMapper = oldRoleMapper
			return nil
		}
		pr.roleMapper = fn(pr.storedInfo.GetId())
		return revert, nil
	}
}

// DefaultBackend sets a backend from the pool of backend factories if one is not set.
func DefaultBackend(ctx context.Context, backendFactoryPool map[string]BackendFactory) ProviderOption {
	return func(pr *providerImpl) (RevertOption, error) {
		if pr.backend != nil {
			return noOpRevert, nil
		}

		oldBackendFactory := pr.backendFactory
		oldBackend := pr.backend
		oldConfig := pr.storedInfo.GetConfig()
		revert := func(pr *providerImpl) error {
			if pr.storedInfo != nil {
				pr.storedInfo.Config = oldConfig
			}
			pr.backend = oldBackend
			backendID := pr.storedInfo.GetId()
			pr.backendFactory.CleanupBackend(backendID)
			pr.backendFactory = oldBackendFactory
			return nil
		}
		// Get the backend factory for the type of provider.
		backendFactory := backendFactoryPool[pr.storedInfo.GetType()]
		if backendFactory == nil {
			return noOpRevert, errors.Errorf("provider type %q is either unknown, no longer available, or incompatible with this installation", pr.storedInfo.GetType())
		}

		pr.backendFactory = backendFactory

		// Create the backend for the provider.
		backend, err := backendFactory.CreateBackend(ctx, pr.storedInfo.GetId(), AllUIEndpoints(pr.storedInfo), pr.storedInfo.GetConfig(), pr.storedInfo.GetClaimMappings())
		if err != nil {
			return revert, errors.Wrapf(err, "unable to create backend for provider id %s", pr.storedInfo.GetId())
		}
		pr.backend = backend
		// We can assume that pr.storedInfo is non-nil here because pr.storedInfo.GetType() referenced a valid auth provider type.
		pr.storedInfo.Config = backend.Config()
		return revert, nil
	}
}

// Pre-baked default option lists.
//////////////////////////////////

// DefaultOptionsForStoredProvider returns the default options that should be run for providers loaded from the store on initialization.
func DefaultOptionsForStoredProvider(backendFactoryPool map[string]BackendFactory, issuerFactory tokens.IssuerFactory, roleMapperFactory permissions.RoleMapperFactory, loginURLFn func(authProviderID string) string) []ProviderOption {
	return []ProviderOption{
		DefaultLoginURL(loginURLFn),
		LogOptionError(DefaultBackend(context.Background(), backendFactoryPool)), // Its ok to fail to load a backend on providers loaded from the store.
		LogOptionError(DefaultTokenIssuerFromFactory(issuerFactory)),             // Its ok to not have a token issuer.
		DefaultRoleMapperOption(roleMapperFactory.GetRoleMapper),
	}
}

// DefaultOptionsForNewProvider returns the default options that should be run for newly created providers.
func DefaultOptionsForNewProvider(ctx context.Context, store Store, backendFactoryPool map[string]BackendFactory, issuerFactory tokens.IssuerFactory, roleMapperFactory permissions.RoleMapperFactory, loginURLFn func(authProviderID string) string) []ProviderOption {
	return []ProviderOption{
		DefaultNewID(),
		DefaultLoginURL(loginURLFn),                                  // Must have id set, so do this after default id setting.
		DefaultBackend(ctx, backendFactoryPool),                      // Not ok to fail to load a backend for newly created providers.
		LogOptionError(DefaultTokenIssuerFromFactory(issuerFactory)), // Its ok to not have a token issuer.
		DefaultRoleMapperOption(roleMapperFactory.GetRoleMapper),
		DefaultAddToStore(ctx, store),
	}
}

// Helpers that modify options.
///////////////////////////////

// LogOptionError eats any error from the input option and logs it.
func LogOptionError(po ProviderOption) ProviderOption {
	return func(pr *providerImpl) (RevertOption, error) {
		revert, err := po(pr)
		if err != nil {
			log.Errorf("error adding option to provider: %s", err)
		}
		return revert, nil
	}
}
