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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/edas/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ApplicationScaleLister helps list ApplicationScales.
// All objects returned here must be treated as read-only.
type ApplicationScaleLister interface {
	// List lists all ApplicationScales in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ApplicationScale, err error)
	// ApplicationScales returns an object that can list and get ApplicationScales.
	ApplicationScales(namespace string) ApplicationScaleNamespaceLister
	ApplicationScaleListerExpansion
}

// applicationScaleLister implements the ApplicationScaleLister interface.
type applicationScaleLister struct {
	indexer cache.Indexer
}

// NewApplicationScaleLister returns a new ApplicationScaleLister.
func NewApplicationScaleLister(indexer cache.Indexer) ApplicationScaleLister {
	return &applicationScaleLister{indexer: indexer}
}

// List lists all ApplicationScales in the indexer.
func (s *applicationScaleLister) List(selector labels.Selector) (ret []*v1alpha1.ApplicationScale, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApplicationScale))
	})
	return ret, err
}

// ApplicationScales returns an object that can list and get ApplicationScales.
func (s *applicationScaleLister) ApplicationScales(namespace string) ApplicationScaleNamespaceLister {
	return applicationScaleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApplicationScaleNamespaceLister helps list and get ApplicationScales.
// All objects returned here must be treated as read-only.
type ApplicationScaleNamespaceLister interface {
	// List lists all ApplicationScales in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ApplicationScale, err error)
	// Get retrieves the ApplicationScale from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ApplicationScale, error)
	ApplicationScaleNamespaceListerExpansion
}

// applicationScaleNamespaceLister implements the ApplicationScaleNamespaceLister
// interface.
type applicationScaleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApplicationScales in the indexer for a given namespace.
func (s applicationScaleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ApplicationScale, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApplicationScale))
	})
	return ret, err
}

// Get retrieves the ApplicationScale from the indexer for a given namespace and name.
func (s applicationScaleNamespaceLister) Get(name string) (*v1alpha1.ApplicationScale, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("applicationscale"), name)
	}
	return obj.(*v1alpha1.ApplicationScale), nil
}
