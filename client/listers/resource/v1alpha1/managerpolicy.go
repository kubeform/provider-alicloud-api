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

// ManagerPolicyLister helps list ManagerPolicies.
// All objects returned here must be treated as read-only.
type ManagerPolicyLister interface {
	// List lists all ManagerPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ManagerPolicy, err error)
	// ManagerPolicies returns an object that can list and get ManagerPolicies.
	ManagerPolicies(namespace string) ManagerPolicyNamespaceLister
	ManagerPolicyListerExpansion
}

// managerPolicyLister implements the ManagerPolicyLister interface.
type managerPolicyLister struct {
	indexer cache.Indexer
}

// NewManagerPolicyLister returns a new ManagerPolicyLister.
func NewManagerPolicyLister(indexer cache.Indexer) ManagerPolicyLister {
	return &managerPolicyLister{indexer: indexer}
}

// List lists all ManagerPolicies in the indexer.
func (s *managerPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.ManagerPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagerPolicy))
	})
	return ret, err
}

// ManagerPolicies returns an object that can list and get ManagerPolicies.
func (s *managerPolicyLister) ManagerPolicies(namespace string) ManagerPolicyNamespaceLister {
	return managerPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ManagerPolicyNamespaceLister helps list and get ManagerPolicies.
// All objects returned here must be treated as read-only.
type ManagerPolicyNamespaceLister interface {
	// List lists all ManagerPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ManagerPolicy, err error)
	// Get retrieves the ManagerPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ManagerPolicy, error)
	ManagerPolicyNamespaceListerExpansion
}

// managerPolicyNamespaceLister implements the ManagerPolicyNamespaceLister
// interface.
type managerPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ManagerPolicies in the indexer for a given namespace.
func (s managerPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ManagerPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagerPolicy))
	})
	return ret, err
}

// Get retrieves the ManagerPolicy from the indexer for a given namespace and name.
func (s managerPolicyNamespaceLister) Get(name string) (*v1alpha1.ManagerPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("managerpolicy"), name)
	}
	return obj.(*v1alpha1.ManagerPolicy), nil
}