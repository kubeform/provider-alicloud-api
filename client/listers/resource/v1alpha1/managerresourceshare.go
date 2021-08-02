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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/resource/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ManagerResourceShareLister helps list ManagerResourceShares.
// All objects returned here must be treated as read-only.
type ManagerResourceShareLister interface {
	// List lists all ManagerResourceShares in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ManagerResourceShare, err error)
	// ManagerResourceShares returns an object that can list and get ManagerResourceShares.
	ManagerResourceShares(namespace string) ManagerResourceShareNamespaceLister
	ManagerResourceShareListerExpansion
}

// managerResourceShareLister implements the ManagerResourceShareLister interface.
type managerResourceShareLister struct {
	indexer cache.Indexer
}

// NewManagerResourceShareLister returns a new ManagerResourceShareLister.
func NewManagerResourceShareLister(indexer cache.Indexer) ManagerResourceShareLister {
	return &managerResourceShareLister{indexer: indexer}
}

// List lists all ManagerResourceShares in the indexer.
func (s *managerResourceShareLister) List(selector labels.Selector) (ret []*v1alpha1.ManagerResourceShare, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagerResourceShare))
	})
	return ret, err
}

// ManagerResourceShares returns an object that can list and get ManagerResourceShares.
func (s *managerResourceShareLister) ManagerResourceShares(namespace string) ManagerResourceShareNamespaceLister {
	return managerResourceShareNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ManagerResourceShareNamespaceLister helps list and get ManagerResourceShares.
// All objects returned here must be treated as read-only.
type ManagerResourceShareNamespaceLister interface {
	// List lists all ManagerResourceShares in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ManagerResourceShare, err error)
	// Get retrieves the ManagerResourceShare from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ManagerResourceShare, error)
	ManagerResourceShareNamespaceListerExpansion
}

// managerResourceShareNamespaceLister implements the ManagerResourceShareNamespaceLister
// interface.
type managerResourceShareNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ManagerResourceShares in the indexer for a given namespace.
func (s managerResourceShareNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ManagerResourceShare, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagerResourceShare))
	})
	return ret, err
}

// Get retrieves the ManagerResourceShare from the indexer for a given namespace and name.
func (s managerResourceShareNamespaceLister) Get(name string) (*v1alpha1.ManagerResourceShare, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("managerresourceshare"), name)
	}
	return obj.(*v1alpha1.ManagerResourceShare), nil
}
