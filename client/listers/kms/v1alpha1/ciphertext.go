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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/kms/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CiphertextLister helps list Ciphertexts.
// All objects returned here must be treated as read-only.
type CiphertextLister interface {
	// List lists all Ciphertexts in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Ciphertext, err error)
	// Ciphertexts returns an object that can list and get Ciphertexts.
	Ciphertexts(namespace string) CiphertextNamespaceLister
	CiphertextListerExpansion
}

// ciphertextLister implements the CiphertextLister interface.
type ciphertextLister struct {
	indexer cache.Indexer
}

// NewCiphertextLister returns a new CiphertextLister.
func NewCiphertextLister(indexer cache.Indexer) CiphertextLister {
	return &ciphertextLister{indexer: indexer}
}

// List lists all Ciphertexts in the indexer.
func (s *ciphertextLister) List(selector labels.Selector) (ret []*v1alpha1.Ciphertext, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Ciphertext))
	})
	return ret, err
}

// Ciphertexts returns an object that can list and get Ciphertexts.
func (s *ciphertextLister) Ciphertexts(namespace string) CiphertextNamespaceLister {
	return ciphertextNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CiphertextNamespaceLister helps list and get Ciphertexts.
// All objects returned here must be treated as read-only.
type CiphertextNamespaceLister interface {
	// List lists all Ciphertexts in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Ciphertext, err error)
	// Get retrieves the Ciphertext from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Ciphertext, error)
	CiphertextNamespaceListerExpansion
}

// ciphertextNamespaceLister implements the CiphertextNamespaceLister
// interface.
type ciphertextNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Ciphertexts in the indexer for a given namespace.
func (s ciphertextNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Ciphertext, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Ciphertext))
	})
	return ret, err
}

// Get retrieves the Ciphertext from the indexer for a given namespace and name.
func (s ciphertextNamespaceLister) Get(name string) (*v1alpha1.Ciphertext, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("ciphertext"), name)
	}
	return obj.(*v1alpha1.Ciphertext), nil
}