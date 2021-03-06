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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/cassandra/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BackupPlanLister helps list BackupPlans.
// All objects returned here must be treated as read-only.
type BackupPlanLister interface {
	// List lists all BackupPlans in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BackupPlan, err error)
	// BackupPlans returns an object that can list and get BackupPlans.
	BackupPlans(namespace string) BackupPlanNamespaceLister
	BackupPlanListerExpansion
}

// backupPlanLister implements the BackupPlanLister interface.
type backupPlanLister struct {
	indexer cache.Indexer
}

// NewBackupPlanLister returns a new BackupPlanLister.
func NewBackupPlanLister(indexer cache.Indexer) BackupPlanLister {
	return &backupPlanLister{indexer: indexer}
}

// List lists all BackupPlans in the indexer.
func (s *backupPlanLister) List(selector labels.Selector) (ret []*v1alpha1.BackupPlan, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BackupPlan))
	})
	return ret, err
}

// BackupPlans returns an object that can list and get BackupPlans.
func (s *backupPlanLister) BackupPlans(namespace string) BackupPlanNamespaceLister {
	return backupPlanNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BackupPlanNamespaceLister helps list and get BackupPlans.
// All objects returned here must be treated as read-only.
type BackupPlanNamespaceLister interface {
	// List lists all BackupPlans in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.BackupPlan, err error)
	// Get retrieves the BackupPlan from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.BackupPlan, error)
	BackupPlanNamespaceListerExpansion
}

// backupPlanNamespaceLister implements the BackupPlanNamespaceLister
// interface.
type backupPlanNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BackupPlans in the indexer for a given namespace.
func (s backupPlanNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.BackupPlan, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.BackupPlan))
	})
	return ret, err
}

// Get retrieves the BackupPlan from the indexer for a given namespace and name.
func (s backupPlanNamespaceLister) Get(name string) (*v1alpha1.BackupPlan, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("backupplan"), name)
	}
	return obj.(*v1alpha1.BackupPlan), nil
}
