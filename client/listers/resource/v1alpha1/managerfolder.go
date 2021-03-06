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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/resource/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ManagerFolderLister helps list ManagerFolders.
// All objects returned here must be treated as read-only.
type ManagerFolderLister interface {
	// List lists all ManagerFolders in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ManagerFolder, err error)
	// ManagerFolders returns an object that can list and get ManagerFolders.
	ManagerFolders(namespace string) ManagerFolderNamespaceLister
	ManagerFolderListerExpansion
}

// managerFolderLister implements the ManagerFolderLister interface.
type managerFolderLister struct {
	indexer cache.Indexer
}

// NewManagerFolderLister returns a new ManagerFolderLister.
func NewManagerFolderLister(indexer cache.Indexer) ManagerFolderLister {
	return &managerFolderLister{indexer: indexer}
}

// List lists all ManagerFolders in the indexer.
func (s *managerFolderLister) List(selector labels.Selector) (ret []*v1alpha1.ManagerFolder, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagerFolder))
	})
	return ret, err
}

// ManagerFolders returns an object that can list and get ManagerFolders.
func (s *managerFolderLister) ManagerFolders(namespace string) ManagerFolderNamespaceLister {
	return managerFolderNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ManagerFolderNamespaceLister helps list and get ManagerFolders.
// All objects returned here must be treated as read-only.
type ManagerFolderNamespaceLister interface {
	// List lists all ManagerFolders in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ManagerFolder, err error)
	// Get retrieves the ManagerFolder from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ManagerFolder, error)
	ManagerFolderNamespaceListerExpansion
}

// managerFolderNamespaceLister implements the ManagerFolderNamespaceLister
// interface.
type managerFolderNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ManagerFolders in the indexer for a given namespace.
func (s managerFolderNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ManagerFolder, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagerFolder))
	})
	return ret, err
}

// Get retrieves the ManagerFolder from the indexer for a given namespace and name.
func (s managerFolderNamespaceLister) Get(name string) (*v1alpha1.ManagerFolder, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("managerfolder"), name)
	}
	return obj.(*v1alpha1.ManagerFolder), nil
}
