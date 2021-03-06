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

// ConsumerChannelLister helps list ConsumerChannels.
// All objects returned here must be treated as read-only.
type ConsumerChannelLister interface {
	// List lists all ConsumerChannels in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ConsumerChannel, err error)
	// ConsumerChannels returns an object that can list and get ConsumerChannels.
	ConsumerChannels(namespace string) ConsumerChannelNamespaceLister
	ConsumerChannelListerExpansion
}

// consumerChannelLister implements the ConsumerChannelLister interface.
type consumerChannelLister struct {
	indexer cache.Indexer
}

// NewConsumerChannelLister returns a new ConsumerChannelLister.
func NewConsumerChannelLister(indexer cache.Indexer) ConsumerChannelLister {
	return &consumerChannelLister{indexer: indexer}
}

// List lists all ConsumerChannels in the indexer.
func (s *consumerChannelLister) List(selector labels.Selector) (ret []*v1alpha1.ConsumerChannel, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ConsumerChannel))
	})
	return ret, err
}

// ConsumerChannels returns an object that can list and get ConsumerChannels.
func (s *consumerChannelLister) ConsumerChannels(namespace string) ConsumerChannelNamespaceLister {
	return consumerChannelNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ConsumerChannelNamespaceLister helps list and get ConsumerChannels.
// All objects returned here must be treated as read-only.
type ConsumerChannelNamespaceLister interface {
	// List lists all ConsumerChannels in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ConsumerChannel, err error)
	// Get retrieves the ConsumerChannel from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ConsumerChannel, error)
	ConsumerChannelNamespaceListerExpansion
}

// consumerChannelNamespaceLister implements the ConsumerChannelNamespaceLister
// interface.
type consumerChannelNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ConsumerChannels in the indexer for a given namespace.
func (s consumerChannelNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ConsumerChannel, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ConsumerChannel))
	})
	return ret, err
}

// Get retrieves the ConsumerChannel from the indexer for a given namespace and name.
func (s consumerChannelNamespaceLister) Get(name string) (*v1alpha1.ConsumerChannel, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("consumerchannel"), name)
	}
	return obj.(*v1alpha1.ConsumerChannel), nil
}
