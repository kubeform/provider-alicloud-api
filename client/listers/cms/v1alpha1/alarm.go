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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/cms/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AlarmLister helps list Alarms.
// All objects returned here must be treated as read-only.
type AlarmLister interface {
	// List lists all Alarms in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Alarm, err error)
	// Alarms returns an object that can list and get Alarms.
	Alarms(namespace string) AlarmNamespaceLister
	AlarmListerExpansion
}

// alarmLister implements the AlarmLister interface.
type alarmLister struct {
	indexer cache.Indexer
}

// NewAlarmLister returns a new AlarmLister.
func NewAlarmLister(indexer cache.Indexer) AlarmLister {
	return &alarmLister{indexer: indexer}
}

// List lists all Alarms in the indexer.
func (s *alarmLister) List(selector labels.Selector) (ret []*v1alpha1.Alarm, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Alarm))
	})
	return ret, err
}

// Alarms returns an object that can list and get Alarms.
func (s *alarmLister) Alarms(namespace string) AlarmNamespaceLister {
	return alarmNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AlarmNamespaceLister helps list and get Alarms.
// All objects returned here must be treated as read-only.
type AlarmNamespaceLister interface {
	// List lists all Alarms in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Alarm, err error)
	// Get retrieves the Alarm from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Alarm, error)
	AlarmNamespaceListerExpansion
}

// alarmNamespaceLister implements the AlarmNamespaceLister
// interface.
type alarmNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Alarms in the indexer for a given namespace.
func (s alarmNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Alarm, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Alarm))
	})
	return ret, err
}

// Get retrieves the Alarm from the indexer for a given namespace and name.
func (s alarmNamespaceLister) Get(name string) (*v1alpha1.Alarm, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("alarm"), name)
	}
	return obj.(*v1alpha1.Alarm), nil
}
