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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/db/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ReadonlyInstanceLister helps list ReadonlyInstances.
// All objects returned here must be treated as read-only.
type ReadonlyInstanceLister interface {
	// List lists all ReadonlyInstances in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ReadonlyInstance, err error)
	// ReadonlyInstances returns an object that can list and get ReadonlyInstances.
	ReadonlyInstances(namespace string) ReadonlyInstanceNamespaceLister
	ReadonlyInstanceListerExpansion
}

// readonlyInstanceLister implements the ReadonlyInstanceLister interface.
type readonlyInstanceLister struct {
	indexer cache.Indexer
}

// NewReadonlyInstanceLister returns a new ReadonlyInstanceLister.
func NewReadonlyInstanceLister(indexer cache.Indexer) ReadonlyInstanceLister {
	return &readonlyInstanceLister{indexer: indexer}
}

// List lists all ReadonlyInstances in the indexer.
func (s *readonlyInstanceLister) List(selector labels.Selector) (ret []*v1alpha1.ReadonlyInstance, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ReadonlyInstance))
	})
	return ret, err
}

// ReadonlyInstances returns an object that can list and get ReadonlyInstances.
func (s *readonlyInstanceLister) ReadonlyInstances(namespace string) ReadonlyInstanceNamespaceLister {
	return readonlyInstanceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ReadonlyInstanceNamespaceLister helps list and get ReadonlyInstances.
// All objects returned here must be treated as read-only.
type ReadonlyInstanceNamespaceLister interface {
	// List lists all ReadonlyInstances in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ReadonlyInstance, err error)
	// Get retrieves the ReadonlyInstance from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ReadonlyInstance, error)
	ReadonlyInstanceNamespaceListerExpansion
}

// readonlyInstanceNamespaceLister implements the ReadonlyInstanceNamespaceLister
// interface.
type readonlyInstanceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ReadonlyInstances in the indexer for a given namespace.
func (s readonlyInstanceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ReadonlyInstance, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ReadonlyInstance))
	})
	return ret, err
}

// Get retrieves the ReadonlyInstance from the indexer for a given namespace and name.
func (s readonlyInstanceNamespaceLister) Get(name string) (*v1alpha1.ReadonlyInstance, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("readonlyinstance"), name)
	}
	return obj.(*v1alpha1.ReadonlyInstance), nil
}
