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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/cloud/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SsoUserAttachmentLister helps list SsoUserAttachments.
// All objects returned here must be treated as read-only.
type SsoUserAttachmentLister interface {
	// List lists all SsoUserAttachments in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SsoUserAttachment, err error)
	// SsoUserAttachments returns an object that can list and get SsoUserAttachments.
	SsoUserAttachments(namespace string) SsoUserAttachmentNamespaceLister
	SsoUserAttachmentListerExpansion
}

// ssoUserAttachmentLister implements the SsoUserAttachmentLister interface.
type ssoUserAttachmentLister struct {
	indexer cache.Indexer
}

// NewSsoUserAttachmentLister returns a new SsoUserAttachmentLister.
func NewSsoUserAttachmentLister(indexer cache.Indexer) SsoUserAttachmentLister {
	return &ssoUserAttachmentLister{indexer: indexer}
}

// List lists all SsoUserAttachments in the indexer.
func (s *ssoUserAttachmentLister) List(selector labels.Selector) (ret []*v1alpha1.SsoUserAttachment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SsoUserAttachment))
	})
	return ret, err
}

// SsoUserAttachments returns an object that can list and get SsoUserAttachments.
func (s *ssoUserAttachmentLister) SsoUserAttachments(namespace string) SsoUserAttachmentNamespaceLister {
	return ssoUserAttachmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SsoUserAttachmentNamespaceLister helps list and get SsoUserAttachments.
// All objects returned here must be treated as read-only.
type SsoUserAttachmentNamespaceLister interface {
	// List lists all SsoUserAttachments in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SsoUserAttachment, err error)
	// Get retrieves the SsoUserAttachment from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.SsoUserAttachment, error)
	SsoUserAttachmentNamespaceListerExpansion
}

// ssoUserAttachmentNamespaceLister implements the SsoUserAttachmentNamespaceLister
// interface.
type ssoUserAttachmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SsoUserAttachments in the indexer for a given namespace.
func (s ssoUserAttachmentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SsoUserAttachment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SsoUserAttachment))
	})
	return ret, err
}

// Get retrieves the SsoUserAttachment from the indexer for a given namespace and name.
func (s ssoUserAttachmentNamespaceLister) Get(name string) (*v1alpha1.SsoUserAttachment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("ssouserattachment"), name)
	}
	return obj.(*v1alpha1.SsoUserAttachment), nil
}
