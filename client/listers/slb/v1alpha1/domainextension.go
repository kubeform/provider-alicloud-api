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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/slb/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DomainExtensionLister helps list DomainExtensions.
// All objects returned here must be treated as read-only.
type DomainExtensionLister interface {
	// List lists all DomainExtensions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DomainExtension, err error)
	// DomainExtensions returns an object that can list and get DomainExtensions.
	DomainExtensions(namespace string) DomainExtensionNamespaceLister
	DomainExtensionListerExpansion
}

// domainExtensionLister implements the DomainExtensionLister interface.
type domainExtensionLister struct {
	indexer cache.Indexer
}

// NewDomainExtensionLister returns a new DomainExtensionLister.
func NewDomainExtensionLister(indexer cache.Indexer) DomainExtensionLister {
	return &domainExtensionLister{indexer: indexer}
}

// List lists all DomainExtensions in the indexer.
func (s *domainExtensionLister) List(selector labels.Selector) (ret []*v1alpha1.DomainExtension, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DomainExtension))
	})
	return ret, err
}

// DomainExtensions returns an object that can list and get DomainExtensions.
func (s *domainExtensionLister) DomainExtensions(namespace string) DomainExtensionNamespaceLister {
	return domainExtensionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DomainExtensionNamespaceLister helps list and get DomainExtensions.
// All objects returned here must be treated as read-only.
type DomainExtensionNamespaceLister interface {
	// List lists all DomainExtensions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DomainExtension, err error)
	// Get retrieves the DomainExtension from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DomainExtension, error)
	DomainExtensionNamespaceListerExpansion
}

// domainExtensionNamespaceLister implements the DomainExtensionNamespaceLister
// interface.
type domainExtensionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DomainExtensions in the indexer for a given namespace.
func (s domainExtensionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DomainExtension, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DomainExtension))
	})
	return ret, err
}

// Get retrieves the DomainExtension from the indexer for a given namespace and name.
func (s domainExtensionNamespaceLister) Get(name string) (*v1alpha1.DomainExtension, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("domainextension"), name)
	}
	return obj.(*v1alpha1.DomainExtension), nil
}
