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

// AliasLister helps list Aliases.
// All objects returned here must be treated as read-only.
type AliasLister interface {
	// List lists all Aliases in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Alias, err error)
	// Aliases returns an object that can list and get Aliases.
	Aliases(namespace string) AliasNamespaceLister
	AliasListerExpansion
}

// aliasLister implements the AliasLister interface.
type aliasLister struct {
	indexer cache.Indexer
}

// NewAliasLister returns a new AliasLister.
func NewAliasLister(indexer cache.Indexer) AliasLister {
	return &aliasLister{indexer: indexer}
}

// List lists all Aliases in the indexer.
func (s *aliasLister) List(selector labels.Selector) (ret []*v1alpha1.Alias, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Alias))
	})
	return ret, err
}

// Aliases returns an object that can list and get Aliases.
func (s *aliasLister) Aliases(namespace string) AliasNamespaceLister {
	return aliasNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AliasNamespaceLister helps list and get Aliases.
// All objects returned here must be treated as read-only.
type AliasNamespaceLister interface {
	// List lists all Aliases in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Alias, err error)
	// Get retrieves the Alias from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Alias, error)
	AliasNamespaceListerExpansion
}

// aliasNamespaceLister implements the AliasNamespaceLister
// interface.
type aliasNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Aliases in the indexer for a given namespace.
func (s aliasNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Alias, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Alias))
	})
	return ret, err
}

// Get retrieves the Alias from the indexer for a given namespace and name.
func (s aliasNamespaceLister) Get(name string) (*v1alpha1.Alias, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("alias"), name)
	}
	return obj.(*v1alpha1.Alias), nil
}
