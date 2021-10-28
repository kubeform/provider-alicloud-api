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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/data/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// WorksFolderLister helps list WorksFolders.
// All objects returned here must be treated as read-only.
type WorksFolderLister interface {
	// List lists all WorksFolders in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WorksFolder, err error)
	// WorksFolders returns an object that can list and get WorksFolders.
	WorksFolders(namespace string) WorksFolderNamespaceLister
	WorksFolderListerExpansion
}

// worksFolderLister implements the WorksFolderLister interface.
type worksFolderLister struct {
	indexer cache.Indexer
}

// NewWorksFolderLister returns a new WorksFolderLister.
func NewWorksFolderLister(indexer cache.Indexer) WorksFolderLister {
	return &worksFolderLister{indexer: indexer}
}

// List lists all WorksFolders in the indexer.
func (s *worksFolderLister) List(selector labels.Selector) (ret []*v1alpha1.WorksFolder, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WorksFolder))
	})
	return ret, err
}

// WorksFolders returns an object that can list and get WorksFolders.
func (s *worksFolderLister) WorksFolders(namespace string) WorksFolderNamespaceLister {
	return worksFolderNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WorksFolderNamespaceLister helps list and get WorksFolders.
// All objects returned here must be treated as read-only.
type WorksFolderNamespaceLister interface {
	// List lists all WorksFolders in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.WorksFolder, err error)
	// Get retrieves the WorksFolder from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.WorksFolder, error)
	WorksFolderNamespaceListerExpansion
}

// worksFolderNamespaceLister implements the WorksFolderNamespaceLister
// interface.
type worksFolderNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all WorksFolders in the indexer for a given namespace.
func (s worksFolderNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.WorksFolder, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WorksFolder))
	})
	return ret, err
}

// Get retrieves the WorksFolder from the indexer for a given namespace and name.
func (s worksFolderNamespaceLister) Get(name string) (*v1alpha1.WorksFolder, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("worksfolder"), name)
	}
	return obj.(*v1alpha1.WorksFolder), nil
}