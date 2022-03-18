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

// MigrationJobLister helps list MigrationJobs.
// All objects returned here must be treated as read-only.
type MigrationJobLister interface {
	// List lists all MigrationJobs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MigrationJob, err error)
	// MigrationJobs returns an object that can list and get MigrationJobs.
	MigrationJobs(namespace string) MigrationJobNamespaceLister
	MigrationJobListerExpansion
}

// migrationJobLister implements the MigrationJobLister interface.
type migrationJobLister struct {
	indexer cache.Indexer
}

// NewMigrationJobLister returns a new MigrationJobLister.
func NewMigrationJobLister(indexer cache.Indexer) MigrationJobLister {
	return &migrationJobLister{indexer: indexer}
}

// List lists all MigrationJobs in the indexer.
func (s *migrationJobLister) List(selector labels.Selector) (ret []*v1alpha1.MigrationJob, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MigrationJob))
	})
	return ret, err
}

// MigrationJobs returns an object that can list and get MigrationJobs.
func (s *migrationJobLister) MigrationJobs(namespace string) MigrationJobNamespaceLister {
	return migrationJobNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MigrationJobNamespaceLister helps list and get MigrationJobs.
// All objects returned here must be treated as read-only.
type MigrationJobNamespaceLister interface {
	// List lists all MigrationJobs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MigrationJob, err error)
	// Get retrieves the MigrationJob from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.MigrationJob, error)
	MigrationJobNamespaceListerExpansion
}

// migrationJobNamespaceLister implements the MigrationJobNamespaceLister
// interface.
type migrationJobNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MigrationJobs in the indexer for a given namespace.
func (s migrationJobNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MigrationJob, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MigrationJob))
	})
	return ret, err
}

// Get retrieves the MigrationJob from the indexer for a given namespace and name.
func (s migrationJobNamespaceLister) Get(name string) (*v1alpha1.MigrationJob, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("migrationjob"), name)
	}
	return obj.(*v1alpha1.MigrationJob), nil
}