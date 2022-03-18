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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/vpc/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BgpGroupLister helps list BgpGroups.
// All objects returned here must be treated as read-only.
type BgpGroupLister interface {
	// List lists all BgpGroups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BgpGroup, err error)
	// BgpGroups returns an object that can list and get BgpGroups.
	BgpGroups(namespace string) BgpGroupNamespaceLister
	BgpGroupListerExpansion
}

// bgpGroupLister implements the BgpGroupLister interface.
type bgpGroupLister struct {
	indexer cache.Indexer
}

// NewBgpGroupLister returns a new BgpGroupLister.
func NewBgpGroupLister(indexer cache.Indexer) BgpGroupLister {
	return &bgpGroupLister{indexer: indexer}
}

// List lists all BgpGroups in the indexer.
func (s *bgpGroupLister) List(selector labels.Selector) (ret []*v1alpha1.BgpGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BgpGroup))
	})
	return ret, err
}

// BgpGroups returns an object that can list and get BgpGroups.
func (s *bgpGroupLister) BgpGroups(namespace string) BgpGroupNamespaceLister {
	return bgpGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BgpGroupNamespaceLister helps list and get BgpGroups.
// All objects returned here must be treated as read-only.
type BgpGroupNamespaceLister interface {
	// List lists all BgpGroups in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BgpGroup, err error)
	// Get retrieves the BgpGroup from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.BgpGroup, error)
	BgpGroupNamespaceListerExpansion
}

// bgpGroupNamespaceLister implements the BgpGroupNamespaceLister
// interface.
type bgpGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BgpGroups in the indexer for a given namespace.
func (s bgpGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BgpGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BgpGroup))
	})
	return ret, err
}

// Get retrieves the BgpGroup from the indexer for a given namespace and name.
func (s bgpGroupNamespaceLister) Get(name string) (*v1alpha1.BgpGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("bgpgroup"), name)
	}
	return obj.(*v1alpha1.BgpGroup), nil
}