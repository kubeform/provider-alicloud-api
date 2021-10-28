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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/ecd/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PolicyGroupLister helps list PolicyGroups.
// All objects returned here must be treated as read-only.
type PolicyGroupLister interface {
	// List lists all PolicyGroups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PolicyGroup, err error)
	// PolicyGroups returns an object that can list and get PolicyGroups.
	PolicyGroups(namespace string) PolicyGroupNamespaceLister
	PolicyGroupListerExpansion
}

// policyGroupLister implements the PolicyGroupLister interface.
type policyGroupLister struct {
	indexer cache.Indexer
}

// NewPolicyGroupLister returns a new PolicyGroupLister.
func NewPolicyGroupLister(indexer cache.Indexer) PolicyGroupLister {
	return &policyGroupLister{indexer: indexer}
}

// List lists all PolicyGroups in the indexer.
func (s *policyGroupLister) List(selector labels.Selector) (ret []*v1alpha1.PolicyGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PolicyGroup))
	})
	return ret, err
}

// PolicyGroups returns an object that can list and get PolicyGroups.
func (s *policyGroupLister) PolicyGroups(namespace string) PolicyGroupNamespaceLister {
	return policyGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PolicyGroupNamespaceLister helps list and get PolicyGroups.
// All objects returned here must be treated as read-only.
type PolicyGroupNamespaceLister interface {
	// List lists all PolicyGroups in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PolicyGroup, err error)
	// Get retrieves the PolicyGroup from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.PolicyGroup, error)
	PolicyGroupNamespaceListerExpansion
}

// policyGroupNamespaceLister implements the PolicyGroupNamespaceLister
// interface.
type policyGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PolicyGroups in the indexer for a given namespace.
func (s policyGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PolicyGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PolicyGroup))
	})
	return ret, err
}

// Get retrieves the PolicyGroup from the indexer for a given namespace and name.
func (s policyGroupNamespaceLister) Get(name string) (*v1alpha1.PolicyGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("policygroup"), name)
	}
	return obj.(*v1alpha1.PolicyGroup), nil
}
