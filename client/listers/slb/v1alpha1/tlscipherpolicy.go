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

// TlsCipherPolicyLister helps list TlsCipherPolicies.
// All objects returned here must be treated as read-only.
type TlsCipherPolicyLister interface {
	// List lists all TlsCipherPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TlsCipherPolicy, err error)
	// TlsCipherPolicies returns an object that can list and get TlsCipherPolicies.
	TlsCipherPolicies(namespace string) TlsCipherPolicyNamespaceLister
	TlsCipherPolicyListerExpansion
}

// tlsCipherPolicyLister implements the TlsCipherPolicyLister interface.
type tlsCipherPolicyLister struct {
	indexer cache.Indexer
}

// NewTlsCipherPolicyLister returns a new TlsCipherPolicyLister.
func NewTlsCipherPolicyLister(indexer cache.Indexer) TlsCipherPolicyLister {
	return &tlsCipherPolicyLister{indexer: indexer}
}

// List lists all TlsCipherPolicies in the indexer.
func (s *tlsCipherPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.TlsCipherPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TlsCipherPolicy))
	})
	return ret, err
}

// TlsCipherPolicies returns an object that can list and get TlsCipherPolicies.
func (s *tlsCipherPolicyLister) TlsCipherPolicies(namespace string) TlsCipherPolicyNamespaceLister {
	return tlsCipherPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TlsCipherPolicyNamespaceLister helps list and get TlsCipherPolicies.
// All objects returned here must be treated as read-only.
type TlsCipherPolicyNamespaceLister interface {
	// List lists all TlsCipherPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TlsCipherPolicy, err error)
	// Get retrieves the TlsCipherPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.TlsCipherPolicy, error)
	TlsCipherPolicyNamespaceListerExpansion
}

// tlsCipherPolicyNamespaceLister implements the TlsCipherPolicyNamespaceLister
// interface.
type tlsCipherPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TlsCipherPolicies in the indexer for a given namespace.
func (s tlsCipherPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TlsCipherPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TlsCipherPolicy))
	})
	return ret, err
}

// Get retrieves the TlsCipherPolicy from the indexer for a given namespace and name.
func (s tlsCipherPolicyNamespaceLister) Get(name string) (*v1alpha1.TlsCipherPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("tlscipherpolicy"), name)
	}
	return obj.(*v1alpha1.TlsCipherPolicy), nil
}
