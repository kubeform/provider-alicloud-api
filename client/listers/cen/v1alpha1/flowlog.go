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

// FlowlogLister helps list Flowlogs.
// All objects returned here must be treated as read-only.
type FlowlogLister interface {
	// List lists all Flowlogs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Flowlog, err error)
	// Flowlogs returns an object that can list and get Flowlogs.
	Flowlogs(namespace string) FlowlogNamespaceLister
	FlowlogListerExpansion
}

// flowlogLister implements the FlowlogLister interface.
type flowlogLister struct {
	indexer cache.Indexer
}

// NewFlowlogLister returns a new FlowlogLister.
func NewFlowlogLister(indexer cache.Indexer) FlowlogLister {
	return &flowlogLister{indexer: indexer}
}

// List lists all Flowlogs in the indexer.
func (s *flowlogLister) List(selector labels.Selector) (ret []*v1alpha1.Flowlog, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Flowlog))
	})
	return ret, err
}

// Flowlogs returns an object that can list and get Flowlogs.
func (s *flowlogLister) Flowlogs(namespace string) FlowlogNamespaceLister {
	return flowlogNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FlowlogNamespaceLister helps list and get Flowlogs.
// All objects returned here must be treated as read-only.
type FlowlogNamespaceLister interface {
	// List lists all Flowlogs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Flowlog, err error)
	// Get retrieves the Flowlog from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Flowlog, error)
	FlowlogNamespaceListerExpansion
}

// flowlogNamespaceLister implements the FlowlogNamespaceLister
// interface.
type flowlogNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Flowlogs in the indexer for a given namespace.
func (s flowlogNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Flowlog, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Flowlog))
	})
	return ret, err
}

// Get retrieves the Flowlog from the indexer for a given namespace and name.
func (s flowlogNamespaceLister) Get(name string) (*v1alpha1.Flowlog, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("flowlog"), name)
	}
	return obj.(*v1alpha1.Flowlog), nil
}
