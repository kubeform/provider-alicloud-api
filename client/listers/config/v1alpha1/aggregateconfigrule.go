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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/config/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AggregateConfigRuleLister helps list AggregateConfigRules.
// All objects returned here must be treated as read-only.
type AggregateConfigRuleLister interface {
	// List lists all AggregateConfigRules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AggregateConfigRule, err error)
	// AggregateConfigRules returns an object that can list and get AggregateConfigRules.
	AggregateConfigRules(namespace string) AggregateConfigRuleNamespaceLister
	AggregateConfigRuleListerExpansion
}

// aggregateConfigRuleLister implements the AggregateConfigRuleLister interface.
type aggregateConfigRuleLister struct {
	indexer cache.Indexer
}

// NewAggregateConfigRuleLister returns a new AggregateConfigRuleLister.
func NewAggregateConfigRuleLister(indexer cache.Indexer) AggregateConfigRuleLister {
	return &aggregateConfigRuleLister{indexer: indexer}
}

// List lists all AggregateConfigRules in the indexer.
func (s *aggregateConfigRuleLister) List(selector labels.Selector) (ret []*v1alpha1.AggregateConfigRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AggregateConfigRule))
	})
	return ret, err
}

// AggregateConfigRules returns an object that can list and get AggregateConfigRules.
func (s *aggregateConfigRuleLister) AggregateConfigRules(namespace string) AggregateConfigRuleNamespaceLister {
	return aggregateConfigRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AggregateConfigRuleNamespaceLister helps list and get AggregateConfigRules.
// All objects returned here must be treated as read-only.
type AggregateConfigRuleNamespaceLister interface {
	// List lists all AggregateConfigRules in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AggregateConfigRule, err error)
	// Get retrieves the AggregateConfigRule from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AggregateConfigRule, error)
	AggregateConfigRuleNamespaceListerExpansion
}

// aggregateConfigRuleNamespaceLister implements the AggregateConfigRuleNamespaceLister
// interface.
type aggregateConfigRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AggregateConfigRules in the indexer for a given namespace.
func (s aggregateConfigRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AggregateConfigRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AggregateConfigRule))
	})
	return ret, err
}

// Get retrieves the AggregateConfigRule from the indexer for a given namespace and name.
func (s aggregateConfigRuleNamespaceLister) Get(name string) (*v1alpha1.AggregateConfigRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("aggregateconfigrule"), name)
	}
	return obj.(*v1alpha1.AggregateConfigRule), nil
}
