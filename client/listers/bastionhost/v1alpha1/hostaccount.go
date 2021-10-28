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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/bastionhost/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// HostAccountLister helps list HostAccounts.
// All objects returned here must be treated as read-only.
type HostAccountLister interface {
	// List lists all HostAccounts in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.HostAccount, err error)
	// HostAccounts returns an object that can list and get HostAccounts.
	HostAccounts(namespace string) HostAccountNamespaceLister
	HostAccountListerExpansion
}

// hostAccountLister implements the HostAccountLister interface.
type hostAccountLister struct {
	indexer cache.Indexer
}

// NewHostAccountLister returns a new HostAccountLister.
func NewHostAccountLister(indexer cache.Indexer) HostAccountLister {
	return &hostAccountLister{indexer: indexer}
}

// List lists all HostAccounts in the indexer.
func (s *hostAccountLister) List(selector labels.Selector) (ret []*v1alpha1.HostAccount, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HostAccount))
	})
	return ret, err
}

// HostAccounts returns an object that can list and get HostAccounts.
func (s *hostAccountLister) HostAccounts(namespace string) HostAccountNamespaceLister {
	return hostAccountNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HostAccountNamespaceLister helps list and get HostAccounts.
// All objects returned here must be treated as read-only.
type HostAccountNamespaceLister interface {
	// List lists all HostAccounts in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.HostAccount, err error)
	// Get retrieves the HostAccount from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.HostAccount, error)
	HostAccountNamespaceListerExpansion
}

// hostAccountNamespaceLister implements the HostAccountNamespaceLister
// interface.
type hostAccountNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all HostAccounts in the indexer for a given namespace.
func (s hostAccountNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.HostAccount, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HostAccount))
	})
	return ret, err
}

// Get retrieves the HostAccount from the indexer for a given namespace and name.
func (s hostAccountNamespaceLister) Get(name string) (*v1alpha1.HostAccount, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("hostaccount"), name)
	}
	return obj.(*v1alpha1.HostAccount), nil
}