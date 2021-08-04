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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/kvstore/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ConnectionLister helps list Connections.
// All objects returned here must be treated as read-only.
type ConnectionLister interface {
	// List lists all Connections in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Connection, err error)
	// Connections returns an object that can list and get Connections.
	Connections(namespace string) ConnectionNamespaceLister
	ConnectionListerExpansion
}

// connectionLister implements the ConnectionLister interface.
type connectionLister struct {
	indexer cache.Indexer
}

// NewConnectionLister returns a new ConnectionLister.
func NewConnectionLister(indexer cache.Indexer) ConnectionLister {
	return &connectionLister{indexer: indexer}
}

// List lists all Connections in the indexer.
func (s *connectionLister) List(selector labels.Selector) (ret []*v1alpha1.Connection, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Connection))
	})
	return ret, err
}

// Connections returns an object that can list and get Connections.
func (s *connectionLister) Connections(namespace string) ConnectionNamespaceLister {
	return connectionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ConnectionNamespaceLister helps list and get Connections.
// All objects returned here must be treated as read-only.
type ConnectionNamespaceLister interface {
	// List lists all Connections in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Connection, err error)
	// Get retrieves the Connection from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Connection, error)
	ConnectionNamespaceListerExpansion
}

// connectionNamespaceLister implements the ConnectionNamespaceLister
// interface.
type connectionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Connections in the indexer for a given namespace.
func (s connectionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Connection, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Connection))
	})
	return ret, err
}

// Get retrieves the Connection from the indexer for a given namespace and name.
func (s connectionNamespaceLister) Get(name string) (*v1alpha1.Connection, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("connection"), name)
	}
	return obj.(*v1alpha1.Connection), nil
}