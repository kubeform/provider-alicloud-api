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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/pvtz/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RuleAttachmentLister helps list RuleAttachments.
// All objects returned here must be treated as read-only.
type RuleAttachmentLister interface {
	// List lists all RuleAttachments in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RuleAttachment, err error)
	// RuleAttachments returns an object that can list and get RuleAttachments.
	RuleAttachments(namespace string) RuleAttachmentNamespaceLister
	RuleAttachmentListerExpansion
}

// ruleAttachmentLister implements the RuleAttachmentLister interface.
type ruleAttachmentLister struct {
	indexer cache.Indexer
}

// NewRuleAttachmentLister returns a new RuleAttachmentLister.
func NewRuleAttachmentLister(indexer cache.Indexer) RuleAttachmentLister {
	return &ruleAttachmentLister{indexer: indexer}
}

// List lists all RuleAttachments in the indexer.
func (s *ruleAttachmentLister) List(selector labels.Selector) (ret []*v1alpha1.RuleAttachment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RuleAttachment))
	})
	return ret, err
}

// RuleAttachments returns an object that can list and get RuleAttachments.
func (s *ruleAttachmentLister) RuleAttachments(namespace string) RuleAttachmentNamespaceLister {
	return ruleAttachmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RuleAttachmentNamespaceLister helps list and get RuleAttachments.
// All objects returned here must be treated as read-only.
type RuleAttachmentNamespaceLister interface {
	// List lists all RuleAttachments in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.RuleAttachment, err error)
	// Get retrieves the RuleAttachment from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.RuleAttachment, error)
	RuleAttachmentNamespaceListerExpansion
}

// ruleAttachmentNamespaceLister implements the RuleAttachmentNamespaceLister
// interface.
type ruleAttachmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RuleAttachments in the indexer for a given namespace.
func (s ruleAttachmentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RuleAttachment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RuleAttachment))
	})
	return ret, err
}

// Get retrieves the RuleAttachment from the indexer for a given namespace and name.
func (s ruleAttachmentNamespaceLister) Get(name string) (*v1alpha1.RuleAttachment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("ruleattachment"), name)
	}
	return obj.(*v1alpha1.RuleAttachment), nil
}
