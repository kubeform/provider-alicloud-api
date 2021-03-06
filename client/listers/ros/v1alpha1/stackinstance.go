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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/ros/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// StackInstanceLister helps list StackInstances.
// All objects returned here must be treated as read-only.
type StackInstanceLister interface {
	// List lists all StackInstances in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.StackInstance, err error)
	// StackInstances returns an object that can list and get StackInstances.
	StackInstances(namespace string) StackInstanceNamespaceLister
	StackInstanceListerExpansion
}

// stackInstanceLister implements the StackInstanceLister interface.
type stackInstanceLister struct {
	indexer cache.Indexer
}

// NewStackInstanceLister returns a new StackInstanceLister.
func NewStackInstanceLister(indexer cache.Indexer) StackInstanceLister {
	return &stackInstanceLister{indexer: indexer}
}

// List lists all StackInstances in the indexer.
func (s *stackInstanceLister) List(selector labels.Selector) (ret []*v1alpha1.StackInstance, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StackInstance))
	})
	return ret, err
}

// StackInstances returns an object that can list and get StackInstances.
func (s *stackInstanceLister) StackInstances(namespace string) StackInstanceNamespaceLister {
	return stackInstanceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StackInstanceNamespaceLister helps list and get StackInstances.
// All objects returned here must be treated as read-only.
type StackInstanceNamespaceLister interface {
	// List lists all StackInstances in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.StackInstance, err error)
	// Get retrieves the StackInstance from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.StackInstance, error)
	StackInstanceNamespaceListerExpansion
}

// stackInstanceNamespaceLister implements the StackInstanceNamespaceLister
// interface.
type stackInstanceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StackInstances in the indexer for a given namespace.
func (s stackInstanceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.StackInstance, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StackInstance))
	})
	return ret, err
}

// Get retrieves the StackInstance from the indexer for a given namespace and name.
func (s stackInstanceNamespaceLister) Get(name string) (*v1alpha1.StackInstance, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("stackinstance"), name)
	}
	return obj.(*v1alpha1.StackInstance), nil
}
