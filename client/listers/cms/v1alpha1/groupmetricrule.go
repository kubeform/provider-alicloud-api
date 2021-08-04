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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/cms/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GroupMetricRuleLister helps list GroupMetricRules.
// All objects returned here must be treated as read-only.
type GroupMetricRuleLister interface {
	// List lists all GroupMetricRules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.GroupMetricRule, err error)
	// GroupMetricRules returns an object that can list and get GroupMetricRules.
	GroupMetricRules(namespace string) GroupMetricRuleNamespaceLister
	GroupMetricRuleListerExpansion
}

// groupMetricRuleLister implements the GroupMetricRuleLister interface.
type groupMetricRuleLister struct {
	indexer cache.Indexer
}

// NewGroupMetricRuleLister returns a new GroupMetricRuleLister.
func NewGroupMetricRuleLister(indexer cache.Indexer) GroupMetricRuleLister {
	return &groupMetricRuleLister{indexer: indexer}
}

// List lists all GroupMetricRules in the indexer.
func (s *groupMetricRuleLister) List(selector labels.Selector) (ret []*v1alpha1.GroupMetricRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GroupMetricRule))
	})
	return ret, err
}

// GroupMetricRules returns an object that can list and get GroupMetricRules.
func (s *groupMetricRuleLister) GroupMetricRules(namespace string) GroupMetricRuleNamespaceLister {
	return groupMetricRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GroupMetricRuleNamespaceLister helps list and get GroupMetricRules.
// All objects returned here must be treated as read-only.
type GroupMetricRuleNamespaceLister interface {
	// List lists all GroupMetricRules in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.GroupMetricRule, err error)
	// Get retrieves the GroupMetricRule from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.GroupMetricRule, error)
	GroupMetricRuleNamespaceListerExpansion
}

// groupMetricRuleNamespaceLister implements the GroupMetricRuleNamespaceLister
// interface.
type groupMetricRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GroupMetricRules in the indexer for a given namespace.
func (s groupMetricRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.GroupMetricRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GroupMetricRule))
	})
	return ret, err
}

// Get retrieves the GroupMetricRule from the indexer for a given namespace and name.
func (s groupMetricRuleNamespaceLister) Get(name string) (*v1alpha1.GroupMetricRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("groupmetricrule"), name)
	}
	return obj.(*v1alpha1.GroupMetricRule), nil
}