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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/mns/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TopicLister helps list Topics.
// All objects returned here must be treated as read-only.
type TopicLister interface {
	// List lists all Topics in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Topic, err error)
	// Topics returns an object that can list and get Topics.
	Topics(namespace string) TopicNamespaceLister
	TopicListerExpansion
}

// topicLister implements the TopicLister interface.
type topicLister struct {
	indexer cache.Indexer
}

// NewTopicLister returns a new TopicLister.
func NewTopicLister(indexer cache.Indexer) TopicLister {
	return &topicLister{indexer: indexer}
}

// List lists all Topics in the indexer.
func (s *topicLister) List(selector labels.Selector) (ret []*v1alpha1.Topic, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Topic))
	})
	return ret, err
}

// Topics returns an object that can list and get Topics.
func (s *topicLister) Topics(namespace string) TopicNamespaceLister {
	return topicNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TopicNamespaceLister helps list and get Topics.
// All objects returned here must be treated as read-only.
type TopicNamespaceLister interface {
	// List lists all Topics in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Topic, err error)
	// Get retrieves the Topic from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Topic, error)
	TopicNamespaceListerExpansion
}

// topicNamespaceLister implements the TopicNamespaceLister
// interface.
type topicNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Topics in the indexer for a given namespace.
func (s topicNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Topic, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Topic))
	})
	return ret, err
}

// Get retrieves the Topic from the indexer for a given namespace and name.
func (s topicNamespaceLister) Get(name string) (*v1alpha1.Topic, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("topic"), name)
	}
	return obj.(*v1alpha1.Topic), nil
}