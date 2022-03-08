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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/cddc/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DedicatedHostAccountLister helps list DedicatedHostAccounts.
// All objects returned here must be treated as read-only.
type DedicatedHostAccountLister interface {
	// List lists all DedicatedHostAccounts in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DedicatedHostAccount, err error)
	// DedicatedHostAccounts returns an object that can list and get DedicatedHostAccounts.
	DedicatedHostAccounts(namespace string) DedicatedHostAccountNamespaceLister
	DedicatedHostAccountListerExpansion
}

// dedicatedHostAccountLister implements the DedicatedHostAccountLister interface.
type dedicatedHostAccountLister struct {
	indexer cache.Indexer
}

// NewDedicatedHostAccountLister returns a new DedicatedHostAccountLister.
func NewDedicatedHostAccountLister(indexer cache.Indexer) DedicatedHostAccountLister {
	return &dedicatedHostAccountLister{indexer: indexer}
}

// List lists all DedicatedHostAccounts in the indexer.
func (s *dedicatedHostAccountLister) List(selector labels.Selector) (ret []*v1alpha1.DedicatedHostAccount, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DedicatedHostAccount))
	})
	return ret, err
}

// DedicatedHostAccounts returns an object that can list and get DedicatedHostAccounts.
func (s *dedicatedHostAccountLister) DedicatedHostAccounts(namespace string) DedicatedHostAccountNamespaceLister {
	return dedicatedHostAccountNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DedicatedHostAccountNamespaceLister helps list and get DedicatedHostAccounts.
// All objects returned here must be treated as read-only.
type DedicatedHostAccountNamespaceLister interface {
	// List lists all DedicatedHostAccounts in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DedicatedHostAccount, err error)
	// Get retrieves the DedicatedHostAccount from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DedicatedHostAccount, error)
	DedicatedHostAccountNamespaceListerExpansion
}

// dedicatedHostAccountNamespaceLister implements the DedicatedHostAccountNamespaceLister
// interface.
type dedicatedHostAccountNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DedicatedHostAccounts in the indexer for a given namespace.
func (s dedicatedHostAccountNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DedicatedHostAccount, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DedicatedHostAccount))
	})
	return ret, err
}

// Get retrieves the DedicatedHostAccount from the indexer for a given namespace and name.
func (s dedicatedHostAccountNamespaceLister) Get(name string) (*v1alpha1.DedicatedHostAccount, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dedicatedhostaccount"), name)
	}
	return obj.(*v1alpha1.DedicatedHostAccount), nil
}
