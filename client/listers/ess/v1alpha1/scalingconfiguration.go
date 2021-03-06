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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/ess/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ScalingConfigurationLister helps list ScalingConfigurations.
// All objects returned here must be treated as read-only.
type ScalingConfigurationLister interface {
	// List lists all ScalingConfigurations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ScalingConfiguration, err error)
	// ScalingConfigurations returns an object that can list and get ScalingConfigurations.
	ScalingConfigurations(namespace string) ScalingConfigurationNamespaceLister
	ScalingConfigurationListerExpansion
}

// scalingConfigurationLister implements the ScalingConfigurationLister interface.
type scalingConfigurationLister struct {
	indexer cache.Indexer
}

// NewScalingConfigurationLister returns a new ScalingConfigurationLister.
func NewScalingConfigurationLister(indexer cache.Indexer) ScalingConfigurationLister {
	return &scalingConfigurationLister{indexer: indexer}
}

// List lists all ScalingConfigurations in the indexer.
func (s *scalingConfigurationLister) List(selector labels.Selector) (ret []*v1alpha1.ScalingConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ScalingConfiguration))
	})
	return ret, err
}

// ScalingConfigurations returns an object that can list and get ScalingConfigurations.
func (s *scalingConfigurationLister) ScalingConfigurations(namespace string) ScalingConfigurationNamespaceLister {
	return scalingConfigurationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ScalingConfigurationNamespaceLister helps list and get ScalingConfigurations.
// All objects returned here must be treated as read-only.
type ScalingConfigurationNamespaceLister interface {
	// List lists all ScalingConfigurations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ScalingConfiguration, err error)
	// Get retrieves the ScalingConfiguration from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ScalingConfiguration, error)
	ScalingConfigurationNamespaceListerExpansion
}

// scalingConfigurationNamespaceLister implements the ScalingConfigurationNamespaceLister
// interface.
type scalingConfigurationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ScalingConfigurations in the indexer for a given namespace.
func (s scalingConfigurationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ScalingConfiguration, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ScalingConfiguration))
	})
	return ret, err
}

// Get retrieves the ScalingConfiguration from the indexer for a given namespace and name.
func (s scalingConfigurationNamespaceLister) Get(name string) (*v1alpha1.ScalingConfiguration, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("scalingconfiguration"), name)
	}
	return obj.(*v1alpha1.ScalingConfiguration), nil
}
