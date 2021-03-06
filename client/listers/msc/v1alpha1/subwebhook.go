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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/msc/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SubWebhookLister helps list SubWebhooks.
// All objects returned here must be treated as read-only.
type SubWebhookLister interface {
	// List lists all SubWebhooks in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SubWebhook, err error)
	// SubWebhooks returns an object that can list and get SubWebhooks.
	SubWebhooks(namespace string) SubWebhookNamespaceLister
	SubWebhookListerExpansion
}

// subWebhookLister implements the SubWebhookLister interface.
type subWebhookLister struct {
	indexer cache.Indexer
}

// NewSubWebhookLister returns a new SubWebhookLister.
func NewSubWebhookLister(indexer cache.Indexer) SubWebhookLister {
	return &subWebhookLister{indexer: indexer}
}

// List lists all SubWebhooks in the indexer.
func (s *subWebhookLister) List(selector labels.Selector) (ret []*v1alpha1.SubWebhook, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SubWebhook))
	})
	return ret, err
}

// SubWebhooks returns an object that can list and get SubWebhooks.
func (s *subWebhookLister) SubWebhooks(namespace string) SubWebhookNamespaceLister {
	return subWebhookNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SubWebhookNamespaceLister helps list and get SubWebhooks.
// All objects returned here must be treated as read-only.
type SubWebhookNamespaceLister interface {
	// List lists all SubWebhooks in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SubWebhook, err error)
	// Get retrieves the SubWebhook from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.SubWebhook, error)
	SubWebhookNamespaceListerExpansion
}

// subWebhookNamespaceLister implements the SubWebhookNamespaceLister
// interface.
type subWebhookNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SubWebhooks in the indexer for a given namespace.
func (s subWebhookNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SubWebhook, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SubWebhook))
	})
	return ret, err
}

// Get retrieves the SubWebhook from the indexer for a given namespace and name.
func (s subWebhookNamespaceLister) Get(name string) (*v1alpha1.SubWebhook, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("subwebhook"), name)
	}
	return obj.(*v1alpha1.SubWebhook), nil
}
