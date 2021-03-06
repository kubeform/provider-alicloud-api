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

// ManagerServiceLinkedRoleLister helps list ManagerServiceLinkedRoles.
// All objects returned here must be treated as read-only.
type ManagerServiceLinkedRoleLister interface {
	// List lists all ManagerServiceLinkedRoles in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ManagerServiceLinkedRole, err error)
	// ManagerServiceLinkedRoles returns an object that can list and get ManagerServiceLinkedRoles.
	ManagerServiceLinkedRoles(namespace string) ManagerServiceLinkedRoleNamespaceLister
	ManagerServiceLinkedRoleListerExpansion
}

// managerServiceLinkedRoleLister implements the ManagerServiceLinkedRoleLister interface.
type managerServiceLinkedRoleLister struct {
	indexer cache.Indexer
}

// NewManagerServiceLinkedRoleLister returns a new ManagerServiceLinkedRoleLister.
func NewManagerServiceLinkedRoleLister(indexer cache.Indexer) ManagerServiceLinkedRoleLister {
	return &managerServiceLinkedRoleLister{indexer: indexer}
}

// List lists all ManagerServiceLinkedRoles in the indexer.
func (s *managerServiceLinkedRoleLister) List(selector labels.Selector) (ret []*v1alpha1.ManagerServiceLinkedRole, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagerServiceLinkedRole))
	})
	return ret, err
}

// ManagerServiceLinkedRoles returns an object that can list and get ManagerServiceLinkedRoles.
func (s *managerServiceLinkedRoleLister) ManagerServiceLinkedRoles(namespace string) ManagerServiceLinkedRoleNamespaceLister {
	return managerServiceLinkedRoleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ManagerServiceLinkedRoleNamespaceLister helps list and get ManagerServiceLinkedRoles.
// All objects returned here must be treated as read-only.
type ManagerServiceLinkedRoleNamespaceLister interface {
	// List lists all ManagerServiceLinkedRoles in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ManagerServiceLinkedRole, err error)
	// Get retrieves the ManagerServiceLinkedRole from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ManagerServiceLinkedRole, error)
	ManagerServiceLinkedRoleNamespaceListerExpansion
}

// managerServiceLinkedRoleNamespaceLister implements the ManagerServiceLinkedRoleNamespaceLister
// interface.
type managerServiceLinkedRoleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ManagerServiceLinkedRoles in the indexer for a given namespace.
func (s managerServiceLinkedRoleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ManagerServiceLinkedRole, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagerServiceLinkedRole))
	})
	return ret, err
}

// Get retrieves the ManagerServiceLinkedRole from the indexer for a given namespace and name.
func (s managerServiceLinkedRoleNamespaceLister) Get(name string) (*v1alpha1.ManagerServiceLinkedRole, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("managerservicelinkedrole"), name)
	}
	return obj.(*v1alpha1.ManagerServiceLinkedRole), nil
}
