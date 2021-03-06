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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/ecs/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AutoSnapshotPolicyAttachmentLister helps list AutoSnapshotPolicyAttachments.
// All objects returned here must be treated as read-only.
type AutoSnapshotPolicyAttachmentLister interface {
	// List lists all AutoSnapshotPolicyAttachments in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AutoSnapshotPolicyAttachment, err error)
	// AutoSnapshotPolicyAttachments returns an object that can list and get AutoSnapshotPolicyAttachments.
	AutoSnapshotPolicyAttachments(namespace string) AutoSnapshotPolicyAttachmentNamespaceLister
	AutoSnapshotPolicyAttachmentListerExpansion
}

// autoSnapshotPolicyAttachmentLister implements the AutoSnapshotPolicyAttachmentLister interface.
type autoSnapshotPolicyAttachmentLister struct {
	indexer cache.Indexer
}

// NewAutoSnapshotPolicyAttachmentLister returns a new AutoSnapshotPolicyAttachmentLister.
func NewAutoSnapshotPolicyAttachmentLister(indexer cache.Indexer) AutoSnapshotPolicyAttachmentLister {
	return &autoSnapshotPolicyAttachmentLister{indexer: indexer}
}

// List lists all AutoSnapshotPolicyAttachments in the indexer.
func (s *autoSnapshotPolicyAttachmentLister) List(selector labels.Selector) (ret []*v1alpha1.AutoSnapshotPolicyAttachment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AutoSnapshotPolicyAttachment))
	})
	return ret, err
}

// AutoSnapshotPolicyAttachments returns an object that can list and get AutoSnapshotPolicyAttachments.
func (s *autoSnapshotPolicyAttachmentLister) AutoSnapshotPolicyAttachments(namespace string) AutoSnapshotPolicyAttachmentNamespaceLister {
	return autoSnapshotPolicyAttachmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AutoSnapshotPolicyAttachmentNamespaceLister helps list and get AutoSnapshotPolicyAttachments.
// All objects returned here must be treated as read-only.
type AutoSnapshotPolicyAttachmentNamespaceLister interface {
	// List lists all AutoSnapshotPolicyAttachments in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AutoSnapshotPolicyAttachment, err error)
	// Get retrieves the AutoSnapshotPolicyAttachment from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AutoSnapshotPolicyAttachment, error)
	AutoSnapshotPolicyAttachmentNamespaceListerExpansion
}

// autoSnapshotPolicyAttachmentNamespaceLister implements the AutoSnapshotPolicyAttachmentNamespaceLister
// interface.
type autoSnapshotPolicyAttachmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AutoSnapshotPolicyAttachments in the indexer for a given namespace.
func (s autoSnapshotPolicyAttachmentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AutoSnapshotPolicyAttachment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AutoSnapshotPolicyAttachment))
	})
	return ret, err
}

// Get retrieves the AutoSnapshotPolicyAttachment from the indexer for a given namespace and name.
func (s autoSnapshotPolicyAttachmentNamespaceLister) Get(name string) (*v1alpha1.AutoSnapshotPolicyAttachment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("autosnapshotpolicyattachment"), name)
	}
	return obj.(*v1alpha1.AutoSnapshotPolicyAttachment), nil
}
