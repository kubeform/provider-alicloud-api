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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/cdn/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DomainNewLister helps list DomainNews.
// All objects returned here must be treated as read-only.
type DomainNewLister interface {
	// List lists all DomainNews in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DomainNew, err error)
	// DomainNews returns an object that can list and get DomainNews.
	DomainNews(namespace string) DomainNewNamespaceLister
	DomainNewListerExpansion
}

// domainNewLister implements the DomainNewLister interface.
type domainNewLister struct {
	indexer cache.Indexer
}

// NewDomainNewLister returns a new DomainNewLister.
func NewDomainNewLister(indexer cache.Indexer) DomainNewLister {
	return &domainNewLister{indexer: indexer}
}

// List lists all DomainNews in the indexer.
func (s *domainNewLister) List(selector labels.Selector) (ret []*v1alpha1.DomainNew, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DomainNew))
	})
	return ret, err
}

// DomainNews returns an object that can list and get DomainNews.
func (s *domainNewLister) DomainNews(namespace string) DomainNewNamespaceLister {
	return domainNewNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DomainNewNamespaceLister helps list and get DomainNews.
// All objects returned here must be treated as read-only.
type DomainNewNamespaceLister interface {
	// List lists all DomainNews in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DomainNew, err error)
	// Get retrieves the DomainNew from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DomainNew, error)
	DomainNewNamespaceListerExpansion
}

// domainNewNamespaceLister implements the DomainNewNamespaceLister
// interface.
type domainNewNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DomainNews in the indexer for a given namespace.
func (s domainNewNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DomainNew, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DomainNew))
	})
	return ret, err
}

// Get retrieves the DomainNew from the indexer for a given namespace and name.
func (s domainNewNamespaceLister) Get(name string) (*v1alpha1.DomainNew, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("domainnew"), name)
	}
	return obj.(*v1alpha1.DomainNew), nil
}