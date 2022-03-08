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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/hbr/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ReplicationVaultLister helps list ReplicationVaults.
// All objects returned here must be treated as read-only.
type ReplicationVaultLister interface {
	// List lists all ReplicationVaults in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ReplicationVault, err error)
	// ReplicationVaults returns an object that can list and get ReplicationVaults.
	ReplicationVaults(namespace string) ReplicationVaultNamespaceLister
	ReplicationVaultListerExpansion
}

// replicationVaultLister implements the ReplicationVaultLister interface.
type replicationVaultLister struct {
	indexer cache.Indexer
}

// NewReplicationVaultLister returns a new ReplicationVaultLister.
func NewReplicationVaultLister(indexer cache.Indexer) ReplicationVaultLister {
	return &replicationVaultLister{indexer: indexer}
}

// List lists all ReplicationVaults in the indexer.
func (s *replicationVaultLister) List(selector labels.Selector) (ret []*v1alpha1.ReplicationVault, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ReplicationVault))
	})
	return ret, err
}

// ReplicationVaults returns an object that can list and get ReplicationVaults.
func (s *replicationVaultLister) ReplicationVaults(namespace string) ReplicationVaultNamespaceLister {
	return replicationVaultNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ReplicationVaultNamespaceLister helps list and get ReplicationVaults.
// All objects returned here must be treated as read-only.
type ReplicationVaultNamespaceLister interface {
	// List lists all ReplicationVaults in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ReplicationVault, err error)
	// Get retrieves the ReplicationVault from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ReplicationVault, error)
	ReplicationVaultNamespaceListerExpansion
}

// replicationVaultNamespaceLister implements the ReplicationVaultNamespaceLister
// interface.
type replicationVaultNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ReplicationVaults in the indexer for a given namespace.
func (s replicationVaultNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ReplicationVault, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ReplicationVault))
	})
	return ret, err
}

// Get retrieves the ReplicationVault from the indexer for a given namespace and name.
func (s replicationVaultNamespaceLister) Get(name string) (*v1alpha1.ReplicationVault, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("replicationvault"), name)
	}
	return obj.(*v1alpha1.ReplicationVault), nil
}
