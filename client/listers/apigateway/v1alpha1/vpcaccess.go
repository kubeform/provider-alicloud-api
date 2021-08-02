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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/apigateway/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VpcAccessLister helps list VpcAccesses.
// All objects returned here must be treated as read-only.
type VpcAccessLister interface {
	// List lists all VpcAccesses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VpcAccess, err error)
	// VpcAccesses returns an object that can list and get VpcAccesses.
	VpcAccesses(namespace string) VpcAccessNamespaceLister
	VpcAccessListerExpansion
}

// vpcAccessLister implements the VpcAccessLister interface.
type vpcAccessLister struct {
	indexer cache.Indexer
}

// NewVpcAccessLister returns a new VpcAccessLister.
func NewVpcAccessLister(indexer cache.Indexer) VpcAccessLister {
	return &vpcAccessLister{indexer: indexer}
}

// List lists all VpcAccesses in the indexer.
func (s *vpcAccessLister) List(selector labels.Selector) (ret []*v1alpha1.VpcAccess, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VpcAccess))
	})
	return ret, err
}

// VpcAccesses returns an object that can list and get VpcAccesses.
func (s *vpcAccessLister) VpcAccesses(namespace string) VpcAccessNamespaceLister {
	return vpcAccessNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VpcAccessNamespaceLister helps list and get VpcAccesses.
// All objects returned here must be treated as read-only.
type VpcAccessNamespaceLister interface {
	// List lists all VpcAccesses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VpcAccess, err error)
	// Get retrieves the VpcAccess from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.VpcAccess, error)
	VpcAccessNamespaceListerExpansion
}

// vpcAccessNamespaceLister implements the VpcAccessNamespaceLister
// interface.
type vpcAccessNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VpcAccesses in the indexer for a given namespace.
func (s vpcAccessNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.VpcAccess, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VpcAccess))
	})
	return ret, err
}

// Get retrieves the VpcAccess from the indexer for a given namespace and name.
func (s vpcAccessNamespaceLister) Get(name string) (*v1alpha1.VpcAccess, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("vpcaccess"), name)
	}
	return obj.(*v1alpha1.VpcAccess), nil
}
