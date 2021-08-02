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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/ess/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AttachmentLister helps list Attachments.
// All objects returned here must be treated as read-only.
type AttachmentLister interface {
	// List lists all Attachments in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Attachment, err error)
	// Attachments returns an object that can list and get Attachments.
	Attachments(namespace string) AttachmentNamespaceLister
	AttachmentListerExpansion
}

// attachmentLister implements the AttachmentLister interface.
type attachmentLister struct {
	indexer cache.Indexer
}

// NewAttachmentLister returns a new AttachmentLister.
func NewAttachmentLister(indexer cache.Indexer) AttachmentLister {
	return &attachmentLister{indexer: indexer}
}

// List lists all Attachments in the indexer.
func (s *attachmentLister) List(selector labels.Selector) (ret []*v1alpha1.Attachment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Attachment))
	})
	return ret, err
}

// Attachments returns an object that can list and get Attachments.
func (s *attachmentLister) Attachments(namespace string) AttachmentNamespaceLister {
	return attachmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AttachmentNamespaceLister helps list and get Attachments.
// All objects returned here must be treated as read-only.
type AttachmentNamespaceLister interface {
	// List lists all Attachments in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Attachment, err error)
	// Get retrieves the Attachment from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Attachment, error)
	AttachmentNamespaceListerExpansion
}

// attachmentNamespaceLister implements the AttachmentNamespaceLister
// interface.
type attachmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Attachments in the indexer for a given namespace.
func (s attachmentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Attachment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Attachment))
	})
	return ret, err
}

// Get retrieves the Attachment from the indexer for a given namespace and name.
func (s attachmentNamespaceLister) Get(name string) (*v1alpha1.Attachment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("attachment"), name)
	}
	return obj.(*v1alpha1.Attachment), nil
}
