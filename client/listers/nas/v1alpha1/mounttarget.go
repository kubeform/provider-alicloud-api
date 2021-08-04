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

// MountTargetLister helps list MountTargets.
// All objects returned here must be treated as read-only.
type MountTargetLister interface {
	// List lists all MountTargets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MountTarget, err error)
	// MountTargets returns an object that can list and get MountTargets.
	MountTargets(namespace string) MountTargetNamespaceLister
	MountTargetListerExpansion
}

// mountTargetLister implements the MountTargetLister interface.
type mountTargetLister struct {
	indexer cache.Indexer
}

// NewMountTargetLister returns a new MountTargetLister.
func NewMountTargetLister(indexer cache.Indexer) MountTargetLister {
	return &mountTargetLister{indexer: indexer}
}

// List lists all MountTargets in the indexer.
func (s *mountTargetLister) List(selector labels.Selector) (ret []*v1alpha1.MountTarget, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MountTarget))
	})
	return ret, err
}

// MountTargets returns an object that can list and get MountTargets.
func (s *mountTargetLister) MountTargets(namespace string) MountTargetNamespaceLister {
	return mountTargetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MountTargetNamespaceLister helps list and get MountTargets.
// All objects returned here must be treated as read-only.
type MountTargetNamespaceLister interface {
	// List lists all MountTargets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MountTarget, err error)
	// Get retrieves the MountTarget from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.MountTarget, error)
	MountTargetNamespaceListerExpansion
}

// mountTargetNamespaceLister implements the MountTargetNamespaceLister
// interface.
type mountTargetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MountTargets in the indexer for a given namespace.
func (s mountTargetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MountTarget, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MountTarget))
	})
	return ret, err
}

// Get retrieves the MountTarget from the indexer for a given namespace and name.
func (s mountTargetNamespaceLister) Get(name string) (*v1alpha1.MountTarget, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("mounttarget"), name)
	}
	return obj.(*v1alpha1.MountTarget), nil
}