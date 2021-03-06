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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/resource/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ManagerControlPolicyLister helps list ManagerControlPolicies.
// All objects returned here must be treated as read-only.
type ManagerControlPolicyLister interface {
	// List lists all ManagerControlPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ManagerControlPolicy, err error)
	// ManagerControlPolicies returns an object that can list and get ManagerControlPolicies.
	ManagerControlPolicies(namespace string) ManagerControlPolicyNamespaceLister
	ManagerControlPolicyListerExpansion
}

// managerControlPolicyLister implements the ManagerControlPolicyLister interface.
type managerControlPolicyLister struct {
	indexer cache.Indexer
}

// NewManagerControlPolicyLister returns a new ManagerControlPolicyLister.
func NewManagerControlPolicyLister(indexer cache.Indexer) ManagerControlPolicyLister {
	return &managerControlPolicyLister{indexer: indexer}
}

// List lists all ManagerControlPolicies in the indexer.
func (s *managerControlPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.ManagerControlPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagerControlPolicy))
	})
	return ret, err
}

// ManagerControlPolicies returns an object that can list and get ManagerControlPolicies.
func (s *managerControlPolicyLister) ManagerControlPolicies(namespace string) ManagerControlPolicyNamespaceLister {
	return managerControlPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ManagerControlPolicyNamespaceLister helps list and get ManagerControlPolicies.
// All objects returned here must be treated as read-only.
type ManagerControlPolicyNamespaceLister interface {
	// List lists all ManagerControlPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ManagerControlPolicy, err error)
	// Get retrieves the ManagerControlPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ManagerControlPolicy, error)
	ManagerControlPolicyNamespaceListerExpansion
}

// managerControlPolicyNamespaceLister implements the ManagerControlPolicyNamespaceLister
// interface.
type managerControlPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ManagerControlPolicies in the indexer for a given namespace.
func (s managerControlPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ManagerControlPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagerControlPolicy))
	})
	return ret, err
}

// Get retrieves the ManagerControlPolicy from the indexer for a given namespace and name.
func (s managerControlPolicyNamespaceLister) Get(name string) (*v1alpha1.ManagerControlPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("managercontrolpolicy"), name)
	}
	return obj.(*v1alpha1.ManagerControlPolicy), nil
}
