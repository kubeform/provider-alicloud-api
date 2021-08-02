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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/fc/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CustomDomainLister helps list CustomDomains.
// All objects returned here must be treated as read-only.
type CustomDomainLister interface {
	// List lists all CustomDomains in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CustomDomain, err error)
	// CustomDomains returns an object that can list and get CustomDomains.
	CustomDomains(namespace string) CustomDomainNamespaceLister
	CustomDomainListerExpansion
}

// customDomainLister implements the CustomDomainLister interface.
type customDomainLister struct {
	indexer cache.Indexer
}

// NewCustomDomainLister returns a new CustomDomainLister.
func NewCustomDomainLister(indexer cache.Indexer) CustomDomainLister {
	return &customDomainLister{indexer: indexer}
}

// List lists all CustomDomains in the indexer.
func (s *customDomainLister) List(selector labels.Selector) (ret []*v1alpha1.CustomDomain, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CustomDomain))
	})
	return ret, err
}

// CustomDomains returns an object that can list and get CustomDomains.
func (s *customDomainLister) CustomDomains(namespace string) CustomDomainNamespaceLister {
	return customDomainNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CustomDomainNamespaceLister helps list and get CustomDomains.
// All objects returned here must be treated as read-only.
type CustomDomainNamespaceLister interface {
	// List lists all CustomDomains in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CustomDomain, err error)
	// Get retrieves the CustomDomain from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.CustomDomain, error)
	CustomDomainNamespaceListerExpansion
}

// customDomainNamespaceLister implements the CustomDomainNamespaceLister
// interface.
type customDomainNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CustomDomains in the indexer for a given namespace.
func (s customDomainNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CustomDomain, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CustomDomain))
	})
	return ret, err
}

// Get retrieves the CustomDomain from the indexer for a given namespace and name.
func (s customDomainNamespaceLister) Get(name string) (*v1alpha1.CustomDomain, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("customdomain"), name)
	}
	return obj.(*v1alpha1.CustomDomain), nil
}
