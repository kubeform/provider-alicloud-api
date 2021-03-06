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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/simple/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ApplicationServerCustomImageLister helps list ApplicationServerCustomImages.
// All objects returned here must be treated as read-only.
type ApplicationServerCustomImageLister interface {
	// List lists all ApplicationServerCustomImages in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ApplicationServerCustomImage, err error)
	// ApplicationServerCustomImages returns an object that can list and get ApplicationServerCustomImages.
	ApplicationServerCustomImages(namespace string) ApplicationServerCustomImageNamespaceLister
	ApplicationServerCustomImageListerExpansion
}

// applicationServerCustomImageLister implements the ApplicationServerCustomImageLister interface.
type applicationServerCustomImageLister struct {
	indexer cache.Indexer
}

// NewApplicationServerCustomImageLister returns a new ApplicationServerCustomImageLister.
func NewApplicationServerCustomImageLister(indexer cache.Indexer) ApplicationServerCustomImageLister {
	return &applicationServerCustomImageLister{indexer: indexer}
}

// List lists all ApplicationServerCustomImages in the indexer.
func (s *applicationServerCustomImageLister) List(selector labels.Selector) (ret []*v1alpha1.ApplicationServerCustomImage, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApplicationServerCustomImage))
	})
	return ret, err
}

// ApplicationServerCustomImages returns an object that can list and get ApplicationServerCustomImages.
func (s *applicationServerCustomImageLister) ApplicationServerCustomImages(namespace string) ApplicationServerCustomImageNamespaceLister {
	return applicationServerCustomImageNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApplicationServerCustomImageNamespaceLister helps list and get ApplicationServerCustomImages.
// All objects returned here must be treated as read-only.
type ApplicationServerCustomImageNamespaceLister interface {
	// List lists all ApplicationServerCustomImages in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ApplicationServerCustomImage, err error)
	// Get retrieves the ApplicationServerCustomImage from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ApplicationServerCustomImage, error)
	ApplicationServerCustomImageNamespaceListerExpansion
}

// applicationServerCustomImageNamespaceLister implements the ApplicationServerCustomImageNamespaceLister
// interface.
type applicationServerCustomImageNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApplicationServerCustomImages in the indexer for a given namespace.
func (s applicationServerCustomImageNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ApplicationServerCustomImage, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApplicationServerCustomImage))
	})
	return ret, err
}

// Get retrieves the ApplicationServerCustomImage from the indexer for a given namespace and name.
func (s applicationServerCustomImageNamespaceLister) Get(name string) (*v1alpha1.ApplicationServerCustomImage, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("applicationservercustomimage"), name)
	}
	return obj.(*v1alpha1.ApplicationServerCustomImage), nil
}
