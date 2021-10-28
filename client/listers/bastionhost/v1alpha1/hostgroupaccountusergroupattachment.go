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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/bastionhost/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// HostGroupAccountUserGroupAttachmentLister helps list HostGroupAccountUserGroupAttachments.
// All objects returned here must be treated as read-only.
type HostGroupAccountUserGroupAttachmentLister interface {
	// List lists all HostGroupAccountUserGroupAttachments in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.HostGroupAccountUserGroupAttachment, err error)
	// HostGroupAccountUserGroupAttachments returns an object that can list and get HostGroupAccountUserGroupAttachments.
	HostGroupAccountUserGroupAttachments(namespace string) HostGroupAccountUserGroupAttachmentNamespaceLister
	HostGroupAccountUserGroupAttachmentListerExpansion
}

// hostGroupAccountUserGroupAttachmentLister implements the HostGroupAccountUserGroupAttachmentLister interface.
type hostGroupAccountUserGroupAttachmentLister struct {
	indexer cache.Indexer
}

// NewHostGroupAccountUserGroupAttachmentLister returns a new HostGroupAccountUserGroupAttachmentLister.
func NewHostGroupAccountUserGroupAttachmentLister(indexer cache.Indexer) HostGroupAccountUserGroupAttachmentLister {
	return &hostGroupAccountUserGroupAttachmentLister{indexer: indexer}
}

// List lists all HostGroupAccountUserGroupAttachments in the indexer.
func (s *hostGroupAccountUserGroupAttachmentLister) List(selector labels.Selector) (ret []*v1alpha1.HostGroupAccountUserGroupAttachment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HostGroupAccountUserGroupAttachment))
	})
	return ret, err
}

// HostGroupAccountUserGroupAttachments returns an object that can list and get HostGroupAccountUserGroupAttachments.
func (s *hostGroupAccountUserGroupAttachmentLister) HostGroupAccountUserGroupAttachments(namespace string) HostGroupAccountUserGroupAttachmentNamespaceLister {
	return hostGroupAccountUserGroupAttachmentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HostGroupAccountUserGroupAttachmentNamespaceLister helps list and get HostGroupAccountUserGroupAttachments.
// All objects returned here must be treated as read-only.
type HostGroupAccountUserGroupAttachmentNamespaceLister interface {
	// List lists all HostGroupAccountUserGroupAttachments in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.HostGroupAccountUserGroupAttachment, err error)
	// Get retrieves the HostGroupAccountUserGroupAttachment from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.HostGroupAccountUserGroupAttachment, error)
	HostGroupAccountUserGroupAttachmentNamespaceListerExpansion
}

// hostGroupAccountUserGroupAttachmentNamespaceLister implements the HostGroupAccountUserGroupAttachmentNamespaceLister
// interface.
type hostGroupAccountUserGroupAttachmentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all HostGroupAccountUserGroupAttachments in the indexer for a given namespace.
func (s hostGroupAccountUserGroupAttachmentNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.HostGroupAccountUserGroupAttachment, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HostGroupAccountUserGroupAttachment))
	})
	return ret, err
}

// Get retrieves the HostGroupAccountUserGroupAttachment from the indexer for a given namespace and name.
func (s hostGroupAccountUserGroupAttachmentNamespaceLister) Get(name string) (*v1alpha1.HostGroupAccountUserGroupAttachment, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("hostgroupaccountusergroupattachment"), name)
	}
	return obj.(*v1alpha1.HostGroupAccountUserGroupAttachment), nil
}
