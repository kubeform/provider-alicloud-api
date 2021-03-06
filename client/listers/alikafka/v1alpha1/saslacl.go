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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/alikafka/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SaslACLLister helps list SaslACLs.
// All objects returned here must be treated as read-only.
type SaslACLLister interface {
	// List lists all SaslACLs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SaslACL, err error)
	// SaslACLs returns an object that can list and get SaslACLs.
	SaslACLs(namespace string) SaslACLNamespaceLister
	SaslACLListerExpansion
}

// saslACLLister implements the SaslACLLister interface.
type saslACLLister struct {
	indexer cache.Indexer
}

// NewSaslACLLister returns a new SaslACLLister.
func NewSaslACLLister(indexer cache.Indexer) SaslACLLister {
	return &saslACLLister{indexer: indexer}
}

// List lists all SaslACLs in the indexer.
func (s *saslACLLister) List(selector labels.Selector) (ret []*v1alpha1.SaslACL, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SaslACL))
	})
	return ret, err
}

// SaslACLs returns an object that can list and get SaslACLs.
func (s *saslACLLister) SaslACLs(namespace string) SaslACLNamespaceLister {
	return saslACLNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SaslACLNamespaceLister helps list and get SaslACLs.
// All objects returned here must be treated as read-only.
type SaslACLNamespaceLister interface {
	// List lists all SaslACLs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SaslACL, err error)
	// Get retrieves the SaslACL from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.SaslACL, error)
	SaslACLNamespaceListerExpansion
}

// saslACLNamespaceLister implements the SaslACLNamespaceLister
// interface.
type saslACLNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SaslACLs in the indexer for a given namespace.
func (s saslACLNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SaslACL, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SaslACL))
	})
	return ret, err
}

// Get retrieves the SaslACL from the indexer for a given namespace and name.
func (s saslACLNamespaceLister) Get(name string) (*v1alpha1.SaslACL, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("saslacl"), name)
	}
	return obj.(*v1alpha1.SaslACL), nil
}
