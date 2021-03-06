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

// ManagerRoleLister helps list ManagerRoles.
// All objects returned here must be treated as read-only.
type ManagerRoleLister interface {
	// List lists all ManagerRoles in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ManagerRole, err error)
	// ManagerRoles returns an object that can list and get ManagerRoles.
	ManagerRoles(namespace string) ManagerRoleNamespaceLister
	ManagerRoleListerExpansion
}

// managerRoleLister implements the ManagerRoleLister interface.
type managerRoleLister struct {
	indexer cache.Indexer
}

// NewManagerRoleLister returns a new ManagerRoleLister.
func NewManagerRoleLister(indexer cache.Indexer) ManagerRoleLister {
	return &managerRoleLister{indexer: indexer}
}

// List lists all ManagerRoles in the indexer.
func (s *managerRoleLister) List(selector labels.Selector) (ret []*v1alpha1.ManagerRole, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagerRole))
	})
	return ret, err
}

// ManagerRoles returns an object that can list and get ManagerRoles.
func (s *managerRoleLister) ManagerRoles(namespace string) ManagerRoleNamespaceLister {
	return managerRoleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ManagerRoleNamespaceLister helps list and get ManagerRoles.
// All objects returned here must be treated as read-only.
type ManagerRoleNamespaceLister interface {
	// List lists all ManagerRoles in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ManagerRole, err error)
	// Get retrieves the ManagerRole from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ManagerRole, error)
	ManagerRoleNamespaceListerExpansion
}

// managerRoleNamespaceLister implements the ManagerRoleNamespaceLister
// interface.
type managerRoleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ManagerRoles in the indexer for a given namespace.
func (s managerRoleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ManagerRole, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ManagerRole))
	})
	return ret, err
}

// Get retrieves the ManagerRole from the indexer for a given namespace and name.
func (s managerRoleNamespaceLister) Get(name string) (*v1alpha1.ManagerRole, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("managerrole"), name)
	}
	return obj.(*v1alpha1.ManagerRole), nil
}
