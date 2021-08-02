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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/event/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BridgeEventBusLister helps list BridgeEventBuses.
// All objects returned here must be treated as read-only.
type BridgeEventBusLister interface {
	// List lists all BridgeEventBuses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BridgeEventBus, err error)
	// BridgeEventBuses returns an object that can list and get BridgeEventBuses.
	BridgeEventBuses(namespace string) BridgeEventBusNamespaceLister
	BridgeEventBusListerExpansion
}

// bridgeEventBusLister implements the BridgeEventBusLister interface.
type bridgeEventBusLister struct {
	indexer cache.Indexer
}

// NewBridgeEventBusLister returns a new BridgeEventBusLister.
func NewBridgeEventBusLister(indexer cache.Indexer) BridgeEventBusLister {
	return &bridgeEventBusLister{indexer: indexer}
}

// List lists all BridgeEventBuses in the indexer.
func (s *bridgeEventBusLister) List(selector labels.Selector) (ret []*v1alpha1.BridgeEventBus, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BridgeEventBus))
	})
	return ret, err
}

// BridgeEventBuses returns an object that can list and get BridgeEventBuses.
func (s *bridgeEventBusLister) BridgeEventBuses(namespace string) BridgeEventBusNamespaceLister {
	return bridgeEventBusNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BridgeEventBusNamespaceLister helps list and get BridgeEventBuses.
// All objects returned here must be treated as read-only.
type BridgeEventBusNamespaceLister interface {
	// List lists all BridgeEventBuses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BridgeEventBus, err error)
	// Get retrieves the BridgeEventBus from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.BridgeEventBus, error)
	BridgeEventBusNamespaceListerExpansion
}

// bridgeEventBusNamespaceLister implements the BridgeEventBusNamespaceLister
// interface.
type bridgeEventBusNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BridgeEventBuses in the indexer for a given namespace.
func (s bridgeEventBusNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BridgeEventBus, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BridgeEventBus))
	})
	return ret, err
}

// Get retrieves the BridgeEventBus from the indexer for a given namespace and name.
func (s bridgeEventBusNamespaceLister) Get(name string) (*v1alpha1.BridgeEventBus, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("bridgeeventbus"), name)
	}
	return obj.(*v1alpha1.BridgeEventBus), nil
}
