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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	essv1alpha1 "kubeform.dev/provider-alicloud-api/apis/ess/v1alpha1"
	versioned "kubeform.dev/provider-alicloud-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-alicloud-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-alicloud-api/client/listers/ess/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ScheduledTaskInformer provides access to a shared informer and lister for
// ScheduledTasks.
type ScheduledTaskInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ScheduledTaskLister
}

type scheduledTaskInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewScheduledTaskInformer constructs a new informer for ScheduledTask type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewScheduledTaskInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredScheduledTaskInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredScheduledTaskInformer constructs a new informer for ScheduledTask type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredScheduledTaskInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EssV1alpha1().ScheduledTasks(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EssV1alpha1().ScheduledTasks(namespace).Watch(context.TODO(), options)
			},
		},
		&essv1alpha1.ScheduledTask{},
		resyncPeriod,
		indexers,
	)
}

func (f *scheduledTaskInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredScheduledTaskInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *scheduledTaskInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&essv1alpha1.ScheduledTask{}, f.defaultInformer)
}

func (f *scheduledTaskInformer) Lister() v1alpha1.ScheduledTaskLister {
	return v1alpha1.NewScheduledTaskLister(f.Informer().GetIndexer())
}
