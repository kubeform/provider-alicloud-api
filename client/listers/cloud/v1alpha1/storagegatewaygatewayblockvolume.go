/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/cloud/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// StorageGatewayGatewayBlockVolumeLister helps list StorageGatewayGatewayBlockVolumes.
// All objects returned here must be treated as read-only.
type StorageGatewayGatewayBlockVolumeLister interface {
	// List lists all StorageGatewayGatewayBlockVolumes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.StorageGatewayGatewayBlockVolume, err error)
	// StorageGatewayGatewayBlockVolumes returns an object that can list and get StorageGatewayGatewayBlockVolumes.
	StorageGatewayGatewayBlockVolumes(namespace string) StorageGatewayGatewayBlockVolumeNamespaceLister
	StorageGatewayGatewayBlockVolumeListerExpansion
}

// storageGatewayGatewayBlockVolumeLister implements the StorageGatewayGatewayBlockVolumeLister interface.
type storageGatewayGatewayBlockVolumeLister struct {
	indexer cache.Indexer
}

// NewStorageGatewayGatewayBlockVolumeLister returns a new StorageGatewayGatewayBlockVolumeLister.
func NewStorageGatewayGatewayBlockVolumeLister(indexer cache.Indexer) StorageGatewayGatewayBlockVolumeLister {
	return &storageGatewayGatewayBlockVolumeLister{indexer: indexer}
}

// List lists all StorageGatewayGatewayBlockVolumes in the indexer.
func (s *storageGatewayGatewayBlockVolumeLister) List(selector labels.Selector) (ret []*v1alpha1.StorageGatewayGatewayBlockVolume, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StorageGatewayGatewayBlockVolume))
	})
	return ret, err
}

// StorageGatewayGatewayBlockVolumes returns an object that can list and get StorageGatewayGatewayBlockVolumes.
func (s *storageGatewayGatewayBlockVolumeLister) StorageGatewayGatewayBlockVolumes(namespace string) StorageGatewayGatewayBlockVolumeNamespaceLister {
	return storageGatewayGatewayBlockVolumeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StorageGatewayGatewayBlockVolumeNamespaceLister helps list and get StorageGatewayGatewayBlockVolumes.
// All objects returned here must be treated as read-only.
type StorageGatewayGatewayBlockVolumeNamespaceLister interface {
	// List lists all StorageGatewayGatewayBlockVolumes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.StorageGatewayGatewayBlockVolume, err error)
	// Get retrieves the StorageGatewayGatewayBlockVolume from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.StorageGatewayGatewayBlockVolume, error)
	StorageGatewayGatewayBlockVolumeNamespaceListerExpansion
}

// storageGatewayGatewayBlockVolumeNamespaceLister implements the StorageGatewayGatewayBlockVolumeNamespaceLister
// interface.
type storageGatewayGatewayBlockVolumeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StorageGatewayGatewayBlockVolumes in the indexer for a given namespace.
func (s storageGatewayGatewayBlockVolumeNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.StorageGatewayGatewayBlockVolume, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StorageGatewayGatewayBlockVolume))
	})
	return ret, err
}

// Get retrieves the StorageGatewayGatewayBlockVolume from the indexer for a given namespace and name.
func (s storageGatewayGatewayBlockVolumeNamespaceLister) Get(name string) (*v1alpha1.StorageGatewayGatewayBlockVolume, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("storagegatewaygatewayblockvolume"), name)
	}
	return obj.(*v1alpha1.StorageGatewayGatewayBlockVolume), nil
}
