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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/polardb/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BackupPolicyLister helps list BackupPolicies.
// All objects returned here must be treated as read-only.
type BackupPolicyLister interface {
	// List lists all BackupPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BackupPolicy, err error)
	// BackupPolicies returns an object that can list and get BackupPolicies.
	BackupPolicies(namespace string) BackupPolicyNamespaceLister
	BackupPolicyListerExpansion
}

// backupPolicyLister implements the BackupPolicyLister interface.
type backupPolicyLister struct {
	indexer cache.Indexer
}

// NewBackupPolicyLister returns a new BackupPolicyLister.
func NewBackupPolicyLister(indexer cache.Indexer) BackupPolicyLister {
	return &backupPolicyLister{indexer: indexer}
}

// List lists all BackupPolicies in the indexer.
func (s *backupPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.BackupPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BackupPolicy))
	})
	return ret, err
}

// BackupPolicies returns an object that can list and get BackupPolicies.
func (s *backupPolicyLister) BackupPolicies(namespace string) BackupPolicyNamespaceLister {
	return backupPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BackupPolicyNamespaceLister helps list and get BackupPolicies.
// All objects returned here must be treated as read-only.
type BackupPolicyNamespaceLister interface {
	// List lists all BackupPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BackupPolicy, err error)
	// Get retrieves the BackupPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.BackupPolicy, error)
	BackupPolicyNamespaceListerExpansion
}

// backupPolicyNamespaceLister implements the BackupPolicyNamespaceLister
// interface.
type backupPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BackupPolicies in the indexer for a given namespace.
func (s backupPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BackupPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BackupPolicy))
	})
	return ret, err
}

// Get retrieves the BackupPolicy from the indexer for a given namespace and name.
func (s backupPolicyNamespaceLister) Get(name string) (*v1alpha1.BackupPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("backuppolicy"), name)
	}
	return obj.(*v1alpha1.BackupPolicy), nil
}
