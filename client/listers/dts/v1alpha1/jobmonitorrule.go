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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/dts/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// JobMonitorRuleLister helps list JobMonitorRules.
// All objects returned here must be treated as read-only.
type JobMonitorRuleLister interface {
	// List lists all JobMonitorRules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.JobMonitorRule, err error)
	// JobMonitorRules returns an object that can list and get JobMonitorRules.
	JobMonitorRules(namespace string) JobMonitorRuleNamespaceLister
	JobMonitorRuleListerExpansion
}

// jobMonitorRuleLister implements the JobMonitorRuleLister interface.
type jobMonitorRuleLister struct {
	indexer cache.Indexer
}

// NewJobMonitorRuleLister returns a new JobMonitorRuleLister.
func NewJobMonitorRuleLister(indexer cache.Indexer) JobMonitorRuleLister {
	return &jobMonitorRuleLister{indexer: indexer}
}

// List lists all JobMonitorRules in the indexer.
func (s *jobMonitorRuleLister) List(selector labels.Selector) (ret []*v1alpha1.JobMonitorRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.JobMonitorRule))
	})
	return ret, err
}

// JobMonitorRules returns an object that can list and get JobMonitorRules.
func (s *jobMonitorRuleLister) JobMonitorRules(namespace string) JobMonitorRuleNamespaceLister {
	return jobMonitorRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// JobMonitorRuleNamespaceLister helps list and get JobMonitorRules.
// All objects returned here must be treated as read-only.
type JobMonitorRuleNamespaceLister interface {
	// List lists all JobMonitorRules in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.JobMonitorRule, err error)
	// Get retrieves the JobMonitorRule from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.JobMonitorRule, error)
	JobMonitorRuleNamespaceListerExpansion
}

// jobMonitorRuleNamespaceLister implements the JobMonitorRuleNamespaceLister
// interface.
type jobMonitorRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all JobMonitorRules in the indexer for a given namespace.
func (s jobMonitorRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.JobMonitorRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.JobMonitorRule))
	})
	return ret, err
}

// Get retrieves the JobMonitorRule from the indexer for a given namespace and name.
func (s jobMonitorRuleNamespaceLister) Get(name string) (*v1alpha1.JobMonitorRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("jobmonitorrule"), name)
	}
	return obj.(*v1alpha1.JobMonitorRule), nil
}
