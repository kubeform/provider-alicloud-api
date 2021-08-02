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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/brain/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// IndustrialPidProjectLister helps list IndustrialPidProjects.
// All objects returned here must be treated as read-only.
type IndustrialPidProjectLister interface {
	// List lists all IndustrialPidProjects in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IndustrialPidProject, err error)
	// IndustrialPidProjects returns an object that can list and get IndustrialPidProjects.
	IndustrialPidProjects(namespace string) IndustrialPidProjectNamespaceLister
	IndustrialPidProjectListerExpansion
}

// industrialPidProjectLister implements the IndustrialPidProjectLister interface.
type industrialPidProjectLister struct {
	indexer cache.Indexer
}

// NewIndustrialPidProjectLister returns a new IndustrialPidProjectLister.
func NewIndustrialPidProjectLister(indexer cache.Indexer) IndustrialPidProjectLister {
	return &industrialPidProjectLister{indexer: indexer}
}

// List lists all IndustrialPidProjects in the indexer.
func (s *industrialPidProjectLister) List(selector labels.Selector) (ret []*v1alpha1.IndustrialPidProject, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IndustrialPidProject))
	})
	return ret, err
}

// IndustrialPidProjects returns an object that can list and get IndustrialPidProjects.
func (s *industrialPidProjectLister) IndustrialPidProjects(namespace string) IndustrialPidProjectNamespaceLister {
	return industrialPidProjectNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IndustrialPidProjectNamespaceLister helps list and get IndustrialPidProjects.
// All objects returned here must be treated as read-only.
type IndustrialPidProjectNamespaceLister interface {
	// List lists all IndustrialPidProjects in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.IndustrialPidProject, err error)
	// Get retrieves the IndustrialPidProject from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.IndustrialPidProject, error)
	IndustrialPidProjectNamespaceListerExpansion
}

// industrialPidProjectNamespaceLister implements the IndustrialPidProjectNamespaceLister
// interface.
type industrialPidProjectNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IndustrialPidProjects in the indexer for a given namespace.
func (s industrialPidProjectNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.IndustrialPidProject, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.IndustrialPidProject))
	})
	return ret, err
}

// Get retrieves the IndustrialPidProject from the indexer for a given namespace and name.
func (s industrialPidProjectNamespaceLister) Get(name string) (*v1alpha1.IndustrialPidProject, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("industrialpidproject"), name)
	}
	return obj.(*v1alpha1.IndustrialPidProject), nil
}
