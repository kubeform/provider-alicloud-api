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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/arms/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AlertContactLister helps list AlertContacts.
// All objects returned here must be treated as read-only.
type AlertContactLister interface {
	// List lists all AlertContacts in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AlertContact, err error)
	// AlertContacts returns an object that can list and get AlertContacts.
	AlertContacts(namespace string) AlertContactNamespaceLister
	AlertContactListerExpansion
}

// alertContactLister implements the AlertContactLister interface.
type alertContactLister struct {
	indexer cache.Indexer
}

// NewAlertContactLister returns a new AlertContactLister.
func NewAlertContactLister(indexer cache.Indexer) AlertContactLister {
	return &alertContactLister{indexer: indexer}
}

// List lists all AlertContacts in the indexer.
func (s *alertContactLister) List(selector labels.Selector) (ret []*v1alpha1.AlertContact, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AlertContact))
	})
	return ret, err
}

// AlertContacts returns an object that can list and get AlertContacts.
func (s *alertContactLister) AlertContacts(namespace string) AlertContactNamespaceLister {
	return alertContactNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AlertContactNamespaceLister helps list and get AlertContacts.
// All objects returned here must be treated as read-only.
type AlertContactNamespaceLister interface {
	// List lists all AlertContacts in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AlertContact, err error)
	// Get retrieves the AlertContact from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AlertContact, error)
	AlertContactNamespaceListerExpansion
}

// alertContactNamespaceLister implements the AlertContactNamespaceLister
// interface.
type alertContactNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AlertContacts in the indexer for a given namespace.
func (s alertContactNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AlertContact, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AlertContact))
	})
	return ret, err
}

// Get retrieves the AlertContact from the indexer for a given namespace and name.
func (s alertContactNamespaceLister) Get(name string) (*v1alpha1.AlertContact, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("alertcontact"), name)
	}
	return obj.(*v1alpha1.AlertContact), nil
}
