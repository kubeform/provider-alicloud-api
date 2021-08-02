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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/cen/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TransitRouterVpcAttachmentLister helps list TransitRouterVpcAttachments.
// All objects returned here must be treated as read-only.
type TransitRouterVpcAttachmentLister interface {
	// List lists all TransitRouterVpcAttachments in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TransitRouterVpcAttachment, err error)
	// TransitRouterVpcAttachments returns an object that can list and get TransitRouterVpcAttachments.
	TransitRouterVpcAttachments(namespace string) TransitRouterVpcAttachmentNamespaceLister
	TransitRouterVpcAttachmentListerExpansion
}

// transitRouterVpcAttachmentLister implements the TransitRouterVpcAttachmentLister interface.
type transitRouterVpcAttachmentLister struct {
	indexer cache.Indexer
}

// NewTransitRouterVpcAttachmentLister returns a new TransitRouterVpcAttachmentLister.
func NewTransitRouterVpcAttachmentLister(indexer cache.Indexer) TransitRouterVpcAttachmentLister {
	return &transitRouterVpcAttachmentLister{indexer: indexer}
}

// List lists all TransitRouterVpcAttachments in the indexer.
func (s *transitRouterVpcAttachmentLister) List(selector labels.Selector) (ret []*v1alpha1.TransitRouterVpcAttachment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TransitRouterVpcAttachment))
	})
	return ret, err
}

// TransitRouterVpcAttachments returns an object that can list and get TransitRouterVpcAttachments.
func (s *transitRouterVpcAttachmentLister) TransitRouterVpcAttachments(namespace string) TransitRouterVpcAttachmentNamespaceLister {
	return transitRouterVpcAttachmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TransitRouterVpcAttachmentNamespaceLister helps list and get TransitRouterVpcAttachments.
// All objects returned here must be treated as read-only.
type TransitRouterVpcAttachmentNamespaceLister interface {
	// List lists all TransitRouterVpcAttachments in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.TransitRouterVpcAttachment, err error)
	// Get retrieves the TransitRouterVpcAttachment from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.TransitRouterVpcAttachment, error)
	TransitRouterVpcAttachmentNamespaceListerExpansion
}

// transitRouterVpcAttachmentNamespaceLister implements the TransitRouterVpcAttachmentNamespaceLister
// interface.
type transitRouterVpcAttachmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TransitRouterVpcAttachments in the indexer for a given namespace.
func (s transitRouterVpcAttachmentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TransitRouterVpcAttachment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TransitRouterVpcAttachment))
	})
	return ret, err
}

// Get retrieves the TransitRouterVpcAttachment from the indexer for a given namespace and name.
func (s transitRouterVpcAttachmentNamespaceLister) Get(name string) (*v1alpha1.TransitRouterVpcAttachment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("transitroutervpcattachment"), name)
	}
	return obj.(*v1alpha1.TransitRouterVpcAttachment), nil
}
