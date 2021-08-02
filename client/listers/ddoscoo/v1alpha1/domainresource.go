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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/ddoscoo/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DomainResourceLister helps list DomainResources.
// All objects returned here must be treated as read-only.
type DomainResourceLister interface {
	// List lists all DomainResources in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DomainResource, err error)
	// DomainResources returns an object that can list and get DomainResources.
	DomainResources(namespace string) DomainResourceNamespaceLister
	DomainResourceListerExpansion
}

// domainResourceLister implements the DomainResourceLister interface.
type domainResourceLister struct {
	indexer cache.Indexer
}

// NewDomainResourceLister returns a new DomainResourceLister.
func NewDomainResourceLister(indexer cache.Indexer) DomainResourceLister {
	return &domainResourceLister{indexer: indexer}
}

// List lists all DomainResources in the indexer.
func (s *domainResourceLister) List(selector labels.Selector) (ret []*v1alpha1.DomainResource, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DomainResource))
	})
	return ret, err
}

// DomainResources returns an object that can list and get DomainResources.
func (s *domainResourceLister) DomainResources(namespace string) DomainResourceNamespaceLister {
	return domainResourceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DomainResourceNamespaceLister helps list and get DomainResources.
// All objects returned here must be treated as read-only.
type DomainResourceNamespaceLister interface {
	// List lists all DomainResources in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DomainResource, err error)
	// Get retrieves the DomainResource from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DomainResource, error)
	DomainResourceNamespaceListerExpansion
}

// domainResourceNamespaceLister implements the DomainResourceNamespaceLister
// interface.
type domainResourceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DomainResources in the indexer for a given namespace.
func (s domainResourceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DomainResource, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DomainResource))
	})
	return ret, err
}

// Get retrieves the DomainResource from the indexer for a given namespace and name.
func (s domainResourceNamespaceLister) Get(name string) (*v1alpha1.DomainResource, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("domainresource"), name)
	}
	return obj.(*v1alpha1.DomainResource), nil
}
