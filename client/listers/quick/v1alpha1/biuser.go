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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/quick/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BiUserLister helps list BiUsers.
// All objects returned here must be treated as read-only.
type BiUserLister interface {
	// List lists all BiUsers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BiUser, err error)
	// BiUsers returns an object that can list and get BiUsers.
	BiUsers(namespace string) BiUserNamespaceLister
	BiUserListerExpansion
}

// biUserLister implements the BiUserLister interface.
type biUserLister struct {
	indexer cache.Indexer
}

// NewBiUserLister returns a new BiUserLister.
func NewBiUserLister(indexer cache.Indexer) BiUserLister {
	return &biUserLister{indexer: indexer}
}

// List lists all BiUsers in the indexer.
func (s *biUserLister) List(selector labels.Selector) (ret []*v1alpha1.BiUser, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BiUser))
	})
	return ret, err
}

// BiUsers returns an object that can list and get BiUsers.
func (s *biUserLister) BiUsers(namespace string) BiUserNamespaceLister {
	return biUserNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BiUserNamespaceLister helps list and get BiUsers.
// All objects returned here must be treated as read-only.
type BiUserNamespaceLister interface {
	// List lists all BiUsers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BiUser, err error)
	// Get retrieves the BiUser from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.BiUser, error)
	BiUserNamespaceListerExpansion
}

// biUserNamespaceLister implements the BiUserNamespaceLister
// interface.
type biUserNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BiUsers in the indexer for a given namespace.
func (s biUserNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BiUser, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BiUser))
	})
	return ret, err
}

// Get retrieves the BiUser from the indexer for a given namespace and name.
func (s biUserNamespaceLister) Get(name string) (*v1alpha1.BiUser, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("biuser"), name)
	}
	return obj.(*v1alpha1.BiUser), nil
}
