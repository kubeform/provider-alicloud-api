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

// StorageGatewayGatewayLister helps list StorageGatewayGateways.
// All objects returned here must be treated as read-only.
type StorageGatewayGatewayLister interface {
	// List lists all StorageGatewayGateways in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.StorageGatewayGateway, err error)
	// StorageGatewayGateways returns an object that can list and get StorageGatewayGateways.
	StorageGatewayGateways(namespace string) StorageGatewayGatewayNamespaceLister
	StorageGatewayGatewayListerExpansion
}

// storageGatewayGatewayLister implements the StorageGatewayGatewayLister interface.
type storageGatewayGatewayLister struct {
	indexer cache.Indexer
}

// NewStorageGatewayGatewayLister returns a new StorageGatewayGatewayLister.
func NewStorageGatewayGatewayLister(indexer cache.Indexer) StorageGatewayGatewayLister {
	return &storageGatewayGatewayLister{indexer: indexer}
}

// List lists all StorageGatewayGateways in the indexer.
func (s *storageGatewayGatewayLister) List(selector labels.Selector) (ret []*v1alpha1.StorageGatewayGateway, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StorageGatewayGateway))
	})
	return ret, err
}

// StorageGatewayGateways returns an object that can list and get StorageGatewayGateways.
func (s *storageGatewayGatewayLister) StorageGatewayGateways(namespace string) StorageGatewayGatewayNamespaceLister {
	return storageGatewayGatewayNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StorageGatewayGatewayNamespaceLister helps list and get StorageGatewayGateways.
// All objects returned here must be treated as read-only.
type StorageGatewayGatewayNamespaceLister interface {
	// List lists all StorageGatewayGateways in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.StorageGatewayGateway, err error)
	// Get retrieves the StorageGatewayGateway from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.StorageGatewayGateway, error)
	StorageGatewayGatewayNamespaceListerExpansion
}

// storageGatewayGatewayNamespaceLister implements the StorageGatewayGatewayNamespaceLister
// interface.
type storageGatewayGatewayNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StorageGatewayGateways in the indexer for a given namespace.
func (s storageGatewayGatewayNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.StorageGatewayGateway, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StorageGatewayGateway))
	})
	return ret, err
}

// Get retrieves the StorageGatewayGateway from the indexer for a given namespace and name.
func (s storageGatewayGatewayNamespaceLister) Get(name string) (*v1alpha1.StorageGatewayGateway, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("storagegatewaygateway"), name)
	}
	return obj.(*v1alpha1.StorageGatewayGateway), nil
}
