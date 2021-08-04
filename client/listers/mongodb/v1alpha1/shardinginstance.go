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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/mongodb/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ShardingInstanceLister helps list ShardingInstances.
// All objects returned here must be treated as read-only.
type ShardingInstanceLister interface {
	// List lists all ShardingInstances in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ShardingInstance, err error)
	// ShardingInstances returns an object that can list and get ShardingInstances.
	ShardingInstances(namespace string) ShardingInstanceNamespaceLister
	ShardingInstanceListerExpansion
}

// shardingInstanceLister implements the ShardingInstanceLister interface.
type shardingInstanceLister struct {
	indexer cache.Indexer
}

// NewShardingInstanceLister returns a new ShardingInstanceLister.
func NewShardingInstanceLister(indexer cache.Indexer) ShardingInstanceLister {
	return &shardingInstanceLister{indexer: indexer}
}

// List lists all ShardingInstances in the indexer.
func (s *shardingInstanceLister) List(selector labels.Selector) (ret []*v1alpha1.ShardingInstance, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ShardingInstance))
	})
	return ret, err
}

// ShardingInstances returns an object that can list and get ShardingInstances.
func (s *shardingInstanceLister) ShardingInstances(namespace string) ShardingInstanceNamespaceLister {
	return shardingInstanceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ShardingInstanceNamespaceLister helps list and get ShardingInstances.
// All objects returned here must be treated as read-only.
type ShardingInstanceNamespaceLister interface {
	// List lists all ShardingInstances in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ShardingInstance, err error)
	// Get retrieves the ShardingInstance from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ShardingInstance, error)
	ShardingInstanceNamespaceListerExpansion
}

// shardingInstanceNamespaceLister implements the ShardingInstanceNamespaceLister
// interface.
type shardingInstanceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ShardingInstances in the indexer for a given namespace.
func (s shardingInstanceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ShardingInstance, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ShardingInstance))
	})
	return ret, err
}

// Get retrieves the ShardingInstance from the indexer for a given namespace and name.
func (s shardingInstanceNamespaceLister) Get(name string) (*v1alpha1.ShardingInstance, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("shardinginstance"), name)
	}
	return obj.(*v1alpha1.ShardingInstance), nil
}