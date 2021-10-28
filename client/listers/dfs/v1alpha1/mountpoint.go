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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/dfs/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MountPointLister helps list MountPoints.
// All objects returned here must be treated as read-only.
type MountPointLister interface {
	// List lists all MountPoints in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MountPoint, err error)
	// MountPoints returns an object that can list and get MountPoints.
	MountPoints(namespace string) MountPointNamespaceLister
	MountPointListerExpansion
}

// mountPointLister implements the MountPointLister interface.
type mountPointLister struct {
	indexer cache.Indexer
}

// NewMountPointLister returns a new MountPointLister.
func NewMountPointLister(indexer cache.Indexer) MountPointLister {
	return &mountPointLister{indexer: indexer}
}

// List lists all MountPoints in the indexer.
func (s *mountPointLister) List(selector labels.Selector) (ret []*v1alpha1.MountPoint, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MountPoint))
	})
	return ret, err
}

// MountPoints returns an object that can list and get MountPoints.
func (s *mountPointLister) MountPoints(namespace string) MountPointNamespaceLister {
	return mountPointNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MountPointNamespaceLister helps list and get MountPoints.
// All objects returned here must be treated as read-only.
type MountPointNamespaceLister interface {
	// List lists all MountPoints in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MountPoint, err error)
	// Get retrieves the MountPoint from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.MountPoint, error)
	MountPointNamespaceListerExpansion
}

// mountPointNamespaceLister implements the MountPointNamespaceLister
// interface.
type mountPointNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MountPoints in the indexer for a given namespace.
func (s mountPointNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MountPoint, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MountPoint))
	})
	return ret, err
}

// Get retrieves the MountPoint from the indexer for a given namespace and name.
func (s mountPointNamespaceLister) Get(name string) (*v1alpha1.MountPoint, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("mountpoint"), name)
	}
	return obj.(*v1alpha1.MountPoint), nil
}