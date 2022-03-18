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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/vpc/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VbrHaLister helps list VbrHas.
// All objects returned here must be treated as read-only.
type VbrHaLister interface {
	// List lists all VbrHas in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VbrHa, err error)
	// VbrHas returns an object that can list and get VbrHas.
	VbrHas(namespace string) VbrHaNamespaceLister
	VbrHaListerExpansion
}

// vbrHaLister implements the VbrHaLister interface.
type vbrHaLister struct {
	indexer cache.Indexer
}

// NewVbrHaLister returns a new VbrHaLister.
func NewVbrHaLister(indexer cache.Indexer) VbrHaLister {
	return &vbrHaLister{indexer: indexer}
}

// List lists all VbrHas in the indexer.
func (s *vbrHaLister) List(selector labels.Selector) (ret []*v1alpha1.VbrHa, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VbrHa))
	})
	return ret, err
}

// VbrHas returns an object that can list and get VbrHas.
func (s *vbrHaLister) VbrHas(namespace string) VbrHaNamespaceLister {
	return vbrHaNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VbrHaNamespaceLister helps list and get VbrHas.
// All objects returned here must be treated as read-only.
type VbrHaNamespaceLister interface {
	// List lists all VbrHas in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VbrHa, err error)
	// Get retrieves the VbrHa from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.VbrHa, error)
	VbrHaNamespaceListerExpansion
}

// vbrHaNamespaceLister implements the VbrHaNamespaceLister
// interface.
type vbrHaNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VbrHas in the indexer for a given namespace.
func (s vbrHaNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.VbrHa, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VbrHa))
	})
	return ret, err
}

// Get retrieves the VbrHa from the indexer for a given namespace and name.
func (s vbrHaNamespaceLister) Get(name string) (*v1alpha1.VbrHa, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("vbrha"), name)
	}
	return obj.(*v1alpha1.VbrHa), nil
}