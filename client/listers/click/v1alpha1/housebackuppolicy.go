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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/click/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// HouseBackupPolicyLister helps list HouseBackupPolicies.
// All objects returned here must be treated as read-only.
type HouseBackupPolicyLister interface {
	// List lists all HouseBackupPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.HouseBackupPolicy, err error)
	// HouseBackupPolicies returns an object that can list and get HouseBackupPolicies.
	HouseBackupPolicies(namespace string) HouseBackupPolicyNamespaceLister
	HouseBackupPolicyListerExpansion
}

// houseBackupPolicyLister implements the HouseBackupPolicyLister interface.
type houseBackupPolicyLister struct {
	indexer cache.Indexer
}

// NewHouseBackupPolicyLister returns a new HouseBackupPolicyLister.
func NewHouseBackupPolicyLister(indexer cache.Indexer) HouseBackupPolicyLister {
	return &houseBackupPolicyLister{indexer: indexer}
}

// List lists all HouseBackupPolicies in the indexer.
func (s *houseBackupPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.HouseBackupPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HouseBackupPolicy))
	})
	return ret, err
}

// HouseBackupPolicies returns an object that can list and get HouseBackupPolicies.
func (s *houseBackupPolicyLister) HouseBackupPolicies(namespace string) HouseBackupPolicyNamespaceLister {
	return houseBackupPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HouseBackupPolicyNamespaceLister helps list and get HouseBackupPolicies.
// All objects returned here must be treated as read-only.
type HouseBackupPolicyNamespaceLister interface {
	// List lists all HouseBackupPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.HouseBackupPolicy, err error)
	// Get retrieves the HouseBackupPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.HouseBackupPolicy, error)
	HouseBackupPolicyNamespaceListerExpansion
}

// houseBackupPolicyNamespaceLister implements the HouseBackupPolicyNamespaceLister
// interface.
type houseBackupPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all HouseBackupPolicies in the indexer for a given namespace.
func (s houseBackupPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.HouseBackupPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HouseBackupPolicy))
	})
	return ret, err
}

// Get retrieves the HouseBackupPolicy from the indexer for a given namespace and name.
func (s houseBackupPolicyNamespaceLister) Get(name string) (*v1alpha1.HouseBackupPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("housebackuppolicy"), name)
	}
	return obj.(*v1alpha1.HouseBackupPolicy), nil
}