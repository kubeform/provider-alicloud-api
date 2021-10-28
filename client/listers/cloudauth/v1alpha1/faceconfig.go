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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/cloudauth/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FaceConfigLister helps list FaceConfigs.
// All objects returned here must be treated as read-only.
type FaceConfigLister interface {
	// List lists all FaceConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FaceConfig, err error)
	// FaceConfigs returns an object that can list and get FaceConfigs.
	FaceConfigs(namespace string) FaceConfigNamespaceLister
	FaceConfigListerExpansion
}

// faceConfigLister implements the FaceConfigLister interface.
type faceConfigLister struct {
	indexer cache.Indexer
}

// NewFaceConfigLister returns a new FaceConfigLister.
func NewFaceConfigLister(indexer cache.Indexer) FaceConfigLister {
	return &faceConfigLister{indexer: indexer}
}

// List lists all FaceConfigs in the indexer.
func (s *faceConfigLister) List(selector labels.Selector) (ret []*v1alpha1.FaceConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FaceConfig))
	})
	return ret, err
}

// FaceConfigs returns an object that can list and get FaceConfigs.
func (s *faceConfigLister) FaceConfigs(namespace string) FaceConfigNamespaceLister {
	return faceConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FaceConfigNamespaceLister helps list and get FaceConfigs.
// All objects returned here must be treated as read-only.
type FaceConfigNamespaceLister interface {
	// List lists all FaceConfigs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.FaceConfig, err error)
	// Get retrieves the FaceConfig from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.FaceConfig, error)
	FaceConfigNamespaceListerExpansion
}

// faceConfigNamespaceLister implements the FaceConfigNamespaceLister
// interface.
type faceConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FaceConfigs in the indexer for a given namespace.
func (s faceConfigNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FaceConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FaceConfig))
	})
	return ret, err
}

// Get retrieves the FaceConfig from the indexer for a given namespace and name.
func (s faceConfigNamespaceLister) Get(name string) (*v1alpha1.FaceConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("faceconfig"), name)
	}
	return obj.(*v1alpha1.FaceConfig), nil
}
