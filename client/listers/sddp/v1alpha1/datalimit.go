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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/sddp/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DataLimitLister helps list DataLimits.
// All objects returned here must be treated as read-only.
type DataLimitLister interface {
	// List lists all DataLimits in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DataLimit, err error)
	// DataLimits returns an object that can list and get DataLimits.
	DataLimits(namespace string) DataLimitNamespaceLister
	DataLimitListerExpansion
}

// dataLimitLister implements the DataLimitLister interface.
type dataLimitLister struct {
	indexer cache.Indexer
}

// NewDataLimitLister returns a new DataLimitLister.
func NewDataLimitLister(indexer cache.Indexer) DataLimitLister {
	return &dataLimitLister{indexer: indexer}
}

// List lists all DataLimits in the indexer.
func (s *dataLimitLister) List(selector labels.Selector) (ret []*v1alpha1.DataLimit, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DataLimit))
	})
	return ret, err
}

// DataLimits returns an object that can list and get DataLimits.
func (s *dataLimitLister) DataLimits(namespace string) DataLimitNamespaceLister {
	return dataLimitNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DataLimitNamespaceLister helps list and get DataLimits.
// All objects returned here must be treated as read-only.
type DataLimitNamespaceLister interface {
	// List lists all DataLimits in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DataLimit, err error)
	// Get retrieves the DataLimit from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DataLimit, error)
	DataLimitNamespaceListerExpansion
}

// dataLimitNamespaceLister implements the DataLimitNamespaceLister
// interface.
type dataLimitNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DataLimits in the indexer for a given namespace.
func (s dataLimitNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DataLimit, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DataLimit))
	})
	return ret, err
}

// Get retrieves the DataLimit from the indexer for a given namespace and name.
func (s dataLimitNamespaceLister) Get(name string) (*v1alpha1.DataLimit, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("datalimit"), name)
	}
	return obj.(*v1alpha1.DataLimit), nil
}
