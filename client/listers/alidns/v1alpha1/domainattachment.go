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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/alidns/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DomainAttachmentLister helps list DomainAttachments.
// All objects returned here must be treated as read-only.
type DomainAttachmentLister interface {
	// List lists all DomainAttachments in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DomainAttachment, err error)
	// DomainAttachments returns an object that can list and get DomainAttachments.
	DomainAttachments(namespace string) DomainAttachmentNamespaceLister
	DomainAttachmentListerExpansion
}

// domainAttachmentLister implements the DomainAttachmentLister interface.
type domainAttachmentLister struct {
	indexer cache.Indexer
}

// NewDomainAttachmentLister returns a new DomainAttachmentLister.
func NewDomainAttachmentLister(indexer cache.Indexer) DomainAttachmentLister {
	return &domainAttachmentLister{indexer: indexer}
}

// List lists all DomainAttachments in the indexer.
func (s *domainAttachmentLister) List(selector labels.Selector) (ret []*v1alpha1.DomainAttachment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DomainAttachment))
	})
	return ret, err
}

// DomainAttachments returns an object that can list and get DomainAttachments.
func (s *domainAttachmentLister) DomainAttachments(namespace string) DomainAttachmentNamespaceLister {
	return domainAttachmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DomainAttachmentNamespaceLister helps list and get DomainAttachments.
// All objects returned here must be treated as read-only.
type DomainAttachmentNamespaceLister interface {
	// List lists all DomainAttachments in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DomainAttachment, err error)
	// Get retrieves the DomainAttachment from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DomainAttachment, error)
	DomainAttachmentNamespaceListerExpansion
}

// domainAttachmentNamespaceLister implements the DomainAttachmentNamespaceLister
// interface.
type domainAttachmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DomainAttachments in the indexer for a given namespace.
func (s domainAttachmentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DomainAttachment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DomainAttachment))
	})
	return ret, err
}

// Get retrieves the DomainAttachment from the indexer for a given namespace and name.
func (s domainAttachmentNamespaceLister) Get(name string) (*v1alpha1.DomainAttachment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("domainattachment"), name)
	}
	return obj.(*v1alpha1.DomainAttachment), nil
}
