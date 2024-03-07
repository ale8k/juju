// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package service

import (
	"context"
	"fmt"

	"github.com/juju/collections/transform"
	"github.com/juju/errors"

	domainstorage "github.com/juju/juju/domain/storage"
	storageerrors "github.com/juju/juju/domain/storage/errors"
	"github.com/juju/juju/internal/storage"
)

// StoragePoolState defines an interface for interacting with the underlying state.
type StoragePoolState interface {
	CreateStoragePool(ctx context.Context, pool domainstorage.StoragePoolDetails) error
	DeleteStoragePool(ctx context.Context, name string) error
	ReplaceStoragePool(ctx context.Context, pool domainstorage.StoragePoolDetails) error
	ListStoragePools(ctx context.Context, filter domainstorage.Names, providers domainstorage.Providers) ([]domainstorage.StoragePoolDetails, error)
	GetStoragePoolByName(ctx context.Context, name string) (domainstorage.StoragePoolDetails, error)
}

// StoragePoolService defines a service for interacting with the underlying state.
type StoragePoolService struct {
	st       StoragePoolState
	logger   Logger
	registry storage.ProviderRegistry
}

// PoolAttrs define the attributes of a storage pool.
type PoolAttrs map[string]any

// CreateStoragePool creates a storage pool, returning an error satisfying [errors.AlreadyExists]
// if a pool with the same name already exists.
func (s *StoragePoolService) CreateStoragePool(ctx context.Context, name string, providerType storage.ProviderType, attrs PoolAttrs) error {
	err := s.validateConfig(name, providerType, attrs)
	if err != nil {
		return errors.Trace(err)
	}

	attrsToSave := transform.Map(attrs, func(k string, v any) (string, string) { return k, fmt.Sprint(v) })
	sp := domainstorage.StoragePoolDetails{
		Name:     name,
		Provider: string(providerType),
		Attrs:    attrsToSave,
	}
	err = s.st.CreateStoragePool(ctx, sp)
	return errors.Annotatef(err, "creating storage pool %q", name)
}

func (s *StoragePoolService) validateConfig(name string, providerType storage.ProviderType, attrs map[string]interface{}) error {
	if name == "" {
		return storageerrors.MissingPoolNameError
	}
	if !storage.IsValidPoolName(name) {
		return fmt.Errorf("pool name %q not valid%w", name, errors.Hide(storageerrors.InvalidPoolNameError))
	}
	if providerType == "" {
		return storageerrors.MissingPoolTypeError
	}
	if s.registry == nil {
		return errors.Errorf("cannot validate storage provider config for %q without a registry", name)
	}

	cfg, err := storage.NewConfig(name, providerType, attrs)
	if err != nil {
		return errors.Trace(err)
	}
	p, err := s.registry.StorageProvider(providerType)
	if err != nil {
		return errors.Trace(err)
	}
	if err := p.ValidateConfig(cfg); err != nil {
		return errors.Annotate(err, "validating storage provider config")
	}
	return nil
}

// DeleteStoragePool deletes a storage pool, returning an error satisfying
// [errors.NotFound] if it doesn't exist.
func (s *StoragePoolService) DeleteStoragePool(ctx context.Context, name string) error {
	// TODO(storage) - check in use when we have storage in dqlite
	// Below is the code from state that will need to be ported.
	/*
		var inUse bool
		cfg, err := sb.config(context.Background())
		if err != nil {
			return errors.Trace(err)
		}
		operatorStorage, ok := cfg.AllAttrs()[k8sconstants.OperatorStorageKey]
		if sb.modelType == ModelTypeCAAS && ok && operatorStorage == poolName {
			apps, err := sb.allApplications()
			if err != nil {
				return errors.Trace(err)
			}
			inUse = len(apps) > 0
		} else {
			query := bson.D{{"constraints.pool", bson.D{{"$eq", poolName}}}}
			pools, err := storageCollection.Find(query).Count()
			if err != nil {
				return errors.Trace(err)
			}
			inUse = pools > 0
		}
		if inUse {
			return errors.Errorf("storage pool %q in use", poolName)
		}
	*/
	err := s.st.DeleteStoragePool(ctx, name)
	return errors.Annotatef(err, "deleting storage pool %q", name)
}

// ReplaceStoragePool replaces an existing storage pool, returning an error
// satisfying [storageerrors.PoolNotFoundError] if a pool with the name does not exist.
func (s *StoragePoolService) ReplaceStoragePool(ctx context.Context, name string, providerType storage.ProviderType, attrs PoolAttrs) error {
	// Use the existing provider type unless explicitly overwritten.
	if providerType == "" {
		existingConfig, err := s.st.GetStoragePoolByName(ctx, name)
		if err != nil {
			return errors.Trace(err)
		}
		providerType = storage.ProviderType(existingConfig.Provider)
	}

	err := s.validateConfig(name, providerType, attrs)
	if err != nil {
		return errors.Trace(err)
	}

	attrsToSave := transform.Map(attrs, func(k string, v any) (string, string) { return k, fmt.Sprint(v) })
	sp := domainstorage.StoragePoolDetails{
		Name:     name,
		Provider: string(providerType),
		Attrs:    attrsToSave,
	}
	err = s.st.ReplaceStoragePool(ctx, sp)
	return errors.Annotatef(err, "replacing storage pool %q", name)
}

// AllStoragePools returns the all storage pools.
func (s *StoragePoolService) AllStoragePools(ctx context.Context) ([]*storage.Config, error) {
	return s.ListStoragePools(ctx, domainstorage.NilNames, domainstorage.NilProviders)
}

// ListStoragePools returns the storage pools matching the specified filter.
func (s *StoragePoolService) ListStoragePools(ctx context.Context, names domainstorage.Names, providers domainstorage.Providers) ([]*storage.Config, error) {
	if err := s.validatePoolListFilterTerms(names, providers); err != nil {
		return nil, errors.Trace(err)
	}

	pools, err := domainstorage.BuiltInStoragePools()
	if err != nil {
		return nil, errors.Trace(err)
	}

	sp, err := s.st.ListStoragePools(ctx, names, providers)
	if err != nil {
		return nil, errors.Trace(err)
	}
	pools = append(pools, sp...)

	results := make([]*storage.Config, len(pools))
	for i, p := range pools {
		results[i], err = s.toStorageConfig(p)
		if err != nil {
			return nil, errors.Trace(err)
		}
	}
	return results, nil
}

func (s *StoragePoolService) validatePoolListFilterTerms(names domainstorage.Names, providers domainstorage.Providers) error {
	if err := s.validateProviderCriteria(providers); err != nil {
		return errors.Trace(err)
	}
	if err := s.validateNameCriteria(names); err != nil {
		return errors.Trace(err)
	}
	return nil
}

func (s *StoragePoolService) validateNameCriteria(names []string) error {
	for _, n := range names {
		if !storage.IsValidPoolName(n) {
			return fmt.Errorf("pool name %q not valid%w", n, errors.Hide(storageerrors.InvalidPoolNameError))
		}
	}
	return nil
}

func (s *StoragePoolService) validateProviderCriteria(providers []string) error {
	if s.registry == nil {
		return errors.New("cannot filter storage providers without a registry")
	}
	for _, p := range providers {
		_, err := s.registry.StorageProvider(storage.ProviderType(p))
		if err != nil {
			return errors.Trace(err)
		}
	}
	return nil
}

// GetStoragePoolByName returns the storage pool with the specified name, returning an error
// satisfying [storageerrors.PoolNotFoundError] if it doesn't exist.
func (s *StoragePoolService) GetStoragePoolByName(ctx context.Context, name string) (*storage.Config, error) {
	if !storage.IsValidPoolName(name) {
		return nil, fmt.Errorf("pool name %q not valid%w", name, errors.Hide(storageerrors.InvalidPoolNameError))
	}

	builtIn, err := domainstorage.BuiltInStoragePools()
	if err != nil {
		return nil, errors.Trace(err)
	}
	for _, p := range builtIn {
		if p.Name == name {
			return s.toStorageConfig(p)
		}
	}
	sp, err := s.st.GetStoragePoolByName(ctx, name)
	if err != nil {
		return nil, errors.Trace(err)
	}
	return s.toStorageConfig(sp)
}

func (s *StoragePoolService) toStorageConfig(sp domainstorage.StoragePoolDetails) (*storage.Config, error) {
	if s.registry == nil {
		return nil, errors.New("cannot load storage pools without a registry")
	}
	var attr map[string]any
	if len(sp.Attrs) > 0 {
		attr = transform.Map(sp.Attrs, func(k, v string) (string, any) { return k, v })
	}
	cfg, err := storage.NewConfig(sp.Name, storage.ProviderType(sp.Provider), attr)
	if err != nil {
		return nil, errors.Trace(err)
	}
	p, err := s.registry.StorageProvider(cfg.Provider())
	if err != nil {
		return nil, errors.Trace(err)
	}
	if err := p.ValidateConfig(cfg); err != nil {
		return nil, errors.Trace(err)
	}
	return cfg, nil
}
