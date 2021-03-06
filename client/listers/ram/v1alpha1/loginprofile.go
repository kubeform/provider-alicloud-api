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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/ram/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// LoginProfileLister helps list LoginProfiles.
// All objects returned here must be treated as read-only.
type LoginProfileLister interface {
	// List lists all LoginProfiles in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LoginProfile, err error)
	// LoginProfiles returns an object that can list and get LoginProfiles.
	LoginProfiles(namespace string) LoginProfileNamespaceLister
	LoginProfileListerExpansion
}

// loginProfileLister implements the LoginProfileLister interface.
type loginProfileLister struct {
	indexer cache.Indexer
}

// NewLoginProfileLister returns a new LoginProfileLister.
func NewLoginProfileLister(indexer cache.Indexer) LoginProfileLister {
	return &loginProfileLister{indexer: indexer}
}

// List lists all LoginProfiles in the indexer.
func (s *loginProfileLister) List(selector labels.Selector) (ret []*v1alpha1.LoginProfile, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LoginProfile))
	})
	return ret, err
}

// LoginProfiles returns an object that can list and get LoginProfiles.
func (s *loginProfileLister) LoginProfiles(namespace string) LoginProfileNamespaceLister {
	return loginProfileNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LoginProfileNamespaceLister helps list and get LoginProfiles.
// All objects returned here must be treated as read-only.
type LoginProfileNamespaceLister interface {
	// List lists all LoginProfiles in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LoginProfile, err error)
	// Get retrieves the LoginProfile from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.LoginProfile, error)
	LoginProfileNamespaceListerExpansion
}

// loginProfileNamespaceLister implements the LoginProfileNamespaceLister
// interface.
type loginProfileNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LoginProfiles in the indexer for a given namespace.
func (s loginProfileNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LoginProfile, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LoginProfile))
	})
	return ret, err
}

// Get retrieves the LoginProfile from the indexer for a given namespace and name.
func (s loginProfileNamespaceLister) Get(name string) (*v1alpha1.LoginProfile, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("loginprofile"), name)
	}
	return obj.(*v1alpha1.LoginProfile), nil
}
