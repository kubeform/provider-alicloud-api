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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/mongodb/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ServerlessInstanceLister helps list ServerlessInstances.
// All objects returned here must be treated as read-only.
type ServerlessInstanceLister interface {
	// List lists all ServerlessInstances in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServerlessInstance, err error)
	// ServerlessInstances returns an object that can list and get ServerlessInstances.
	ServerlessInstances(namespace string) ServerlessInstanceNamespaceLister
	ServerlessInstanceListerExpansion
}

// serverlessInstanceLister implements the ServerlessInstanceLister interface.
type serverlessInstanceLister struct {
	indexer cache.Indexer
}

// NewServerlessInstanceLister returns a new ServerlessInstanceLister.
func NewServerlessInstanceLister(indexer cache.Indexer) ServerlessInstanceLister {
	return &serverlessInstanceLister{indexer: indexer}
}

// List lists all ServerlessInstances in the indexer.
func (s *serverlessInstanceLister) List(selector labels.Selector) (ret []*v1alpha1.ServerlessInstance, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServerlessInstance))
	})
	return ret, err
}

// ServerlessInstances returns an object that can list and get ServerlessInstances.
func (s *serverlessInstanceLister) ServerlessInstances(namespace string) ServerlessInstanceNamespaceLister {
	return serverlessInstanceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServerlessInstanceNamespaceLister helps list and get ServerlessInstances.
// All objects returned here must be treated as read-only.
type ServerlessInstanceNamespaceLister interface {
	// List lists all ServerlessInstances in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ServerlessInstance, err error)
	// Get retrieves the ServerlessInstance from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ServerlessInstance, error)
	ServerlessInstanceNamespaceListerExpansion
}

// serverlessInstanceNamespaceLister implements the ServerlessInstanceNamespaceLister
// interface.
type serverlessInstanceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ServerlessInstances in the indexer for a given namespace.
func (s serverlessInstanceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ServerlessInstance, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ServerlessInstance))
	})
	return ret, err
}

// Get retrieves the ServerlessInstance from the indexer for a given namespace and name.
func (s serverlessInstanceNamespaceLister) Get(name string) (*v1alpha1.ServerlessInstance, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("serverlessinstance"), name)
	}
	return obj.(*v1alpha1.ServerlessInstance), nil
}
