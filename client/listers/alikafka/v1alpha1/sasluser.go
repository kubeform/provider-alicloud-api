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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/alikafka/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SaslUserLister helps list SaslUsers.
// All objects returned here must be treated as read-only.
type SaslUserLister interface {
	// List lists all SaslUsers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SaslUser, err error)
	// SaslUsers returns an object that can list and get SaslUsers.
	SaslUsers(namespace string) SaslUserNamespaceLister
	SaslUserListerExpansion
}

// saslUserLister implements the SaslUserLister interface.
type saslUserLister struct {
	indexer cache.Indexer
}

// NewSaslUserLister returns a new SaslUserLister.
func NewSaslUserLister(indexer cache.Indexer) SaslUserLister {
	return &saslUserLister{indexer: indexer}
}

// List lists all SaslUsers in the indexer.
func (s *saslUserLister) List(selector labels.Selector) (ret []*v1alpha1.SaslUser, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SaslUser))
	})
	return ret, err
}

// SaslUsers returns an object that can list and get SaslUsers.
func (s *saslUserLister) SaslUsers(namespace string) SaslUserNamespaceLister {
	return saslUserNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SaslUserNamespaceLister helps list and get SaslUsers.
// All objects returned here must be treated as read-only.
type SaslUserNamespaceLister interface {
	// List lists all SaslUsers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SaslUser, err error)
	// Get retrieves the SaslUser from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.SaslUser, error)
	SaslUserNamespaceListerExpansion
}

// saslUserNamespaceLister implements the SaslUserNamespaceLister
// interface.
type saslUserNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SaslUsers in the indexer for a given namespace.
func (s saslUserNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SaslUser, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SaslUser))
	})
	return ret, err
}

// Get retrieves the SaslUser from the indexer for a given namespace and name.
func (s saslUserNamespaceLister) Get(name string) (*v1alpha1.SaslUser, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("sasluser"), name)
	}
	return obj.(*v1alpha1.SaslUser), nil
}
