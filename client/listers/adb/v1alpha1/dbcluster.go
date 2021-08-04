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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/adb/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DbClusterLister helps list DbClusters.
// All objects returned here must be treated as read-only.
type DbClusterLister interface {
	// List lists all DbClusters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DbCluster, err error)
	// DbClusters returns an object that can list and get DbClusters.
	DbClusters(namespace string) DbClusterNamespaceLister
	DbClusterListerExpansion
}

// dbClusterLister implements the DbClusterLister interface.
type dbClusterLister struct {
	indexer cache.Indexer
}

// NewDbClusterLister returns a new DbClusterLister.
func NewDbClusterLister(indexer cache.Indexer) DbClusterLister {
	return &dbClusterLister{indexer: indexer}
}

// List lists all DbClusters in the indexer.
func (s *dbClusterLister) List(selector labels.Selector) (ret []*v1alpha1.DbCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DbCluster))
	})
	return ret, err
}

// DbClusters returns an object that can list and get DbClusters.
func (s *dbClusterLister) DbClusters(namespace string) DbClusterNamespaceLister {
	return dbClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DbClusterNamespaceLister helps list and get DbClusters.
// All objects returned here must be treated as read-only.
type DbClusterNamespaceLister interface {
	// List lists all DbClusters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DbCluster, err error)
	// Get retrieves the DbCluster from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DbCluster, error)
	DbClusterNamespaceListerExpansion
}

// dbClusterNamespaceLister implements the DbClusterNamespaceLister
// interface.
type dbClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DbClusters in the indexer for a given namespace.
func (s dbClusterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DbCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DbCluster))
	})
	return ret, err
}

// Get retrieves the DbCluster from the indexer for a given namespace and name.
func (s dbClusterNamespaceLister) Get(name string) (*v1alpha1.DbCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dbcluster"), name)
	}
	return obj.(*v1alpha1.DbCluster), nil
}