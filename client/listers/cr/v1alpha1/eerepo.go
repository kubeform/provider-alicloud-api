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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/cr/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EeRepoLister helps list EeRepos.
// All objects returned here must be treated as read-only.
type EeRepoLister interface {
	// List lists all EeRepos in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EeRepo, err error)
	// EeRepos returns an object that can list and get EeRepos.
	EeRepos(namespace string) EeRepoNamespaceLister
	EeRepoListerExpansion
}

// eeRepoLister implements the EeRepoLister interface.
type eeRepoLister struct {
	indexer cache.Indexer
}

// NewEeRepoLister returns a new EeRepoLister.
func NewEeRepoLister(indexer cache.Indexer) EeRepoLister {
	return &eeRepoLister{indexer: indexer}
}

// List lists all EeRepos in the indexer.
func (s *eeRepoLister) List(selector labels.Selector) (ret []*v1alpha1.EeRepo, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EeRepo))
	})
	return ret, err
}

// EeRepos returns an object that can list and get EeRepos.
func (s *eeRepoLister) EeRepos(namespace string) EeRepoNamespaceLister {
	return eeRepoNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EeRepoNamespaceLister helps list and get EeRepos.
// All objects returned here must be treated as read-only.
type EeRepoNamespaceLister interface {
	// List lists all EeRepos in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EeRepo, err error)
	// Get retrieves the EeRepo from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.EeRepo, error)
	EeRepoNamespaceListerExpansion
}

// eeRepoNamespaceLister implements the EeRepoNamespaceLister
// interface.
type eeRepoNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EeRepos in the indexer for a given namespace.
func (s eeRepoNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EeRepo, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EeRepo))
	})
	return ret, err
}

// Get retrieves the EeRepo from the indexer for a given namespace and name.
func (s eeRepoNamespaceLister) Get(name string) (*v1alpha1.EeRepo, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("eerepo"), name)
	}
	return obj.(*v1alpha1.EeRepo), nil
}
