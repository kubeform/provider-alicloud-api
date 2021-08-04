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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/ssl/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VpnServerLister helps list VpnServers.
// All objects returned here must be treated as read-only.
type VpnServerLister interface {
	// List lists all VpnServers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VpnServer, err error)
	// VpnServers returns an object that can list and get VpnServers.
	VpnServers(namespace string) VpnServerNamespaceLister
	VpnServerListerExpansion
}

// vpnServerLister implements the VpnServerLister interface.
type vpnServerLister struct {
	indexer cache.Indexer
}

// NewVpnServerLister returns a new VpnServerLister.
func NewVpnServerLister(indexer cache.Indexer) VpnServerLister {
	return &vpnServerLister{indexer: indexer}
}

// List lists all VpnServers in the indexer.
func (s *vpnServerLister) List(selector labels.Selector) (ret []*v1alpha1.VpnServer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VpnServer))
	})
	return ret, err
}

// VpnServers returns an object that can list and get VpnServers.
func (s *vpnServerLister) VpnServers(namespace string) VpnServerNamespaceLister {
	return vpnServerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VpnServerNamespaceLister helps list and get VpnServers.
// All objects returned here must be treated as read-only.
type VpnServerNamespaceLister interface {
	// List lists all VpnServers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VpnServer, err error)
	// Get retrieves the VpnServer from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.VpnServer, error)
	VpnServerNamespaceListerExpansion
}

// vpnServerNamespaceLister implements the VpnServerNamespaceLister
// interface.
type vpnServerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VpnServers in the indexer for a given namespace.
func (s vpnServerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.VpnServer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VpnServer))
	})
	return ret, err
}

// Get retrieves the VpnServer from the indexer for a given namespace and name.
func (s vpnServerNamespaceLister) Get(name string) (*v1alpha1.VpnServer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("vpnserver"), name)
	}
	return obj.(*v1alpha1.VpnServer), nil
}