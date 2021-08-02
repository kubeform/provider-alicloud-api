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

// StackGroupLister helps list StackGroups.
// All objects returned here must be treated as read-only.
type StackGroupLister interface {
	// List lists all StackGroups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.StackGroup, err error)
	// StackGroups returns an object that can list and get StackGroups.
	StackGroups(namespace string) StackGroupNamespaceLister
	StackGroupListerExpansion
}

// stackGroupLister implements the StackGroupLister interface.
type stackGroupLister struct {
	indexer cache.Indexer
}

// NewStackGroupLister returns a new StackGroupLister.
func NewStackGroupLister(indexer cache.Indexer) StackGroupLister {
	return &stackGroupLister{indexer: indexer}
}

// List lists all StackGroups in the indexer.
func (s *stackGroupLister) List(selector labels.Selector) (ret []*v1alpha1.StackGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StackGroup))
	})
	return ret, err
}

// StackGroups returns an object that can list and get StackGroups.
func (s *stackGroupLister) StackGroups(namespace string) StackGroupNamespaceLister {
	return stackGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StackGroupNamespaceLister helps list and get StackGroups.
// All objects returned here must be treated as read-only.
type StackGroupNamespaceLister interface {
	// List lists all StackGroups in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.StackGroup, err error)
	// Get retrieves the StackGroup from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.StackGroup, error)
	StackGroupNamespaceListerExpansion
}

// stackGroupNamespaceLister implements the StackGroupNamespaceLister
// interface.
type stackGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all StackGroups in the indexer for a given namespace.
func (s stackGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.StackGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StackGroup))
	})
	return ret, err
}

// Get retrieves the StackGroup from the indexer for a given namespace and name.
func (s stackGroupNamespaceLister) Get(name string) (*v1alpha1.StackGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("stackgroup"), name)
	}
	return obj.(*v1alpha1.StackGroup), nil
}
