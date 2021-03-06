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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/polardb/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EndpointAddressLister helps list EndpointAddresses.
// All objects returned here must be treated as read-only.
type EndpointAddressLister interface {
	// List lists all EndpointAddresses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EndpointAddress, err error)
	// EndpointAddresses returns an object that can list and get EndpointAddresses.
	EndpointAddresses(namespace string) EndpointAddressNamespaceLister
	EndpointAddressListerExpansion
}

// endpointAddressLister implements the EndpointAddressLister interface.
type endpointAddressLister struct {
	indexer cache.Indexer
}

// NewEndpointAddressLister returns a new EndpointAddressLister.
func NewEndpointAddressLister(indexer cache.Indexer) EndpointAddressLister {
	return &endpointAddressLister{indexer: indexer}
}

// List lists all EndpointAddresses in the indexer.
func (s *endpointAddressLister) List(selector labels.Selector) (ret []*v1alpha1.EndpointAddress, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EndpointAddress))
	})
	return ret, err
}

// EndpointAddresses returns an object that can list and get EndpointAddresses.
func (s *endpointAddressLister) EndpointAddresses(namespace string) EndpointAddressNamespaceLister {
	return endpointAddressNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EndpointAddressNamespaceLister helps list and get EndpointAddresses.
// All objects returned here must be treated as read-only.
type EndpointAddressNamespaceLister interface {
	// List lists all EndpointAddresses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EndpointAddress, err error)
	// Get retrieves the EndpointAddress from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.EndpointAddress, error)
	EndpointAddressNamespaceListerExpansion
}

// endpointAddressNamespaceLister implements the EndpointAddressNamespaceLister
// interface.
type endpointAddressNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EndpointAddresses in the indexer for a given namespace.
func (s endpointAddressNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EndpointAddress, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EndpointAddress))
	})
	return ret, err
}

// Get retrieves the EndpointAddress from the indexer for a given namespace and name.
func (s endpointAddressNamespaceLister) Get(name string) (*v1alpha1.EndpointAddress, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("endpointaddress"), name)
	}
	return obj.(*v1alpha1.EndpointAddress), nil
}
