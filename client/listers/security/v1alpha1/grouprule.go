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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/security/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GroupRuleLister helps list GroupRules.
// All objects returned here must be treated as read-only.
type GroupRuleLister interface {
	// List lists all GroupRules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.GroupRule, err error)
	// GroupRules returns an object that can list and get GroupRules.
	GroupRules(namespace string) GroupRuleNamespaceLister
	GroupRuleListerExpansion
}

// groupRuleLister implements the GroupRuleLister interface.
type groupRuleLister struct {
	indexer cache.Indexer
}

// NewGroupRuleLister returns a new GroupRuleLister.
func NewGroupRuleLister(indexer cache.Indexer) GroupRuleLister {
	return &groupRuleLister{indexer: indexer}
}

// List lists all GroupRules in the indexer.
func (s *groupRuleLister) List(selector labels.Selector) (ret []*v1alpha1.GroupRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GroupRule))
	})
	return ret, err
}

// GroupRules returns an object that can list and get GroupRules.
func (s *groupRuleLister) GroupRules(namespace string) GroupRuleNamespaceLister {
	return groupRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GroupRuleNamespaceLister helps list and get GroupRules.
// All objects returned here must be treated as read-only.
type GroupRuleNamespaceLister interface {
	// List lists all GroupRules in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.GroupRule, err error)
	// Get retrieves the GroupRule from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.GroupRule, error)
	GroupRuleNamespaceListerExpansion
}

// groupRuleNamespaceLister implements the GroupRuleNamespaceLister
// interface.
type groupRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GroupRules in the indexer for a given namespace.
func (s groupRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.GroupRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GroupRule))
	})
	return ret, err
}

// Get retrieves the GroupRule from the indexer for a given namespace and name.
func (s groupRuleNamespaceLister) Get(name string) (*v1alpha1.GroupRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("grouprule"), name)
	}
	return obj.(*v1alpha1.GroupRule), nil
}