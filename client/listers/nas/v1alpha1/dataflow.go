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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/nas/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DataFlowLister helps list DataFlows.
// All objects returned here must be treated as read-only.
type DataFlowLister interface {
	// List lists all DataFlows in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DataFlow, err error)
	// DataFlows returns an object that can list and get DataFlows.
	DataFlows(namespace string) DataFlowNamespaceLister
	DataFlowListerExpansion
}

// dataFlowLister implements the DataFlowLister interface.
type dataFlowLister struct {
	indexer cache.Indexer
}

// NewDataFlowLister returns a new DataFlowLister.
func NewDataFlowLister(indexer cache.Indexer) DataFlowLister {
	return &dataFlowLister{indexer: indexer}
}

// List lists all DataFlows in the indexer.
func (s *dataFlowLister) List(selector labels.Selector) (ret []*v1alpha1.DataFlow, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DataFlow))
	})
	return ret, err
}

// DataFlows returns an object that can list and get DataFlows.
func (s *dataFlowLister) DataFlows(namespace string) DataFlowNamespaceLister {
	return dataFlowNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DataFlowNamespaceLister helps list and get DataFlows.
// All objects returned here must be treated as read-only.
type DataFlowNamespaceLister interface {
	// List lists all DataFlows in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DataFlow, err error)
	// Get retrieves the DataFlow from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DataFlow, error)
	DataFlowNamespaceListerExpansion
}

// dataFlowNamespaceLister implements the DataFlowNamespaceLister
// interface.
type dataFlowNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DataFlows in the indexer for a given namespace.
func (s dataFlowNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DataFlow, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DataFlow))
	})
	return ret, err
}

// Get retrieves the DataFlow from the indexer for a given namespace and name.
func (s dataFlowNamespaceLister) Get(name string) (*v1alpha1.DataFlow, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dataflow"), name)
	}
	return obj.(*v1alpha1.DataFlow), nil
}
