package authproviders

import (
	"context"

	"github.com/stackrox/rox/generated/storage"
	"github.com/stackrox/rox/pkg/auth/tokens"
	"github.com/stackrox/rox/pkg/dberrors"
	"github.com/stackrox/rox/pkg/protocompat"
)

// Commands that providers can execute.
// So that we can keep Provider opaque, and to decouple store operations from the registry,
// These commands are temporarily applied as options.
/////////////////////////////////////////////////////

// DefaultAddToStore adds the providers stored data to the input store.
func DefaultAddToStore(ctx context.Context, store Store) ProviderOption {
	return func(pr *providerImpl) error {
		log.Info("DefaultAddToStore provider with name ", pr.Name(), " and ID ", pr.ID())
		if pr.doNotStore {
			return nil
		}
		if pr.storedInfo.LastUpdated == nil {
			pr.storedInfo.LastUpdated = protocompat.TimestampNow()
		}
		log.Info("storing provider with name ", pr.Name(), " and ID ", pr.ID())
		err := store.AddAuthProvider(ctx, pr.storedInfo)
		if err == nil {
			log.Info("status storing provider with name ", pr.Name(), " and ID ", pr.ID(), " no error")
		} else {
			log.Info("status storing provider with name ", pr.Name(), " and ID ", pr.ID(), " error ", err.Error())
		}
		return err
	}
}

// UpdateStore updates the stored value for the provider in the input store.
func UpdateStore(ctx context.Context, store Store) ProviderOption {
	return func(pr *providerImpl) error {
		log.Info("providerCommand UpdateStore")
		if pr.doNotStore {
			return nil
		}
		pr.storedInfo.LastUpdated = protocompat.TimestampNow()
		return store.UpdateAuthProvider(ctx, pr.storedInfo)
	}
}

// DeleteFromStore removes the providers stored data from the input store.
func DeleteFromStore(ctx context.Context, store Store, providerID string, force bool) ProviderOption {
	return func(pr *providerImpl) error {
		err := store.RemoveAuthProvider(ctx, providerID, force)
		if err != nil {
			// If it's a type we don't want to store, then we're okay with it not existing.
			// We do this in case it was stored in the DB in a previous version.
			if pr.doNotStore && dberrors.IsNotFound(err) {
				return nil
			}
			return err
		}
		// a deleted provider should no longer be accessible, but it's still cached as a token source so mark it as
		// no longer valid
		pr.storedInfo = &storage.AuthProvider{
			Id:      pr.storedInfo.GetId(),
			Enabled: false,
		}
		return nil
	}
}

// UnregisterSource unregisters the token source from the source factory
func UnregisterSource(factory tokens.IssuerFactory) ProviderOption {
	return func(pr *providerImpl) error {
		err := factory.UnregisterSource(pr)
		// both DeleteFromStore and UnregisterSource mutate external stores, so regardless of order the second one
		// can't return err and fail the change.
		if err != nil {
			log.Warnf("Unable to unregister token source: %v", err)
		}
		return nil
	}
}
