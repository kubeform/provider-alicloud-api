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

	dbv1alpha1 "kubeform.dev/provider-alicloud-api/apis/db/v1alpha1"
	versioned "kubeform.dev/provider-alicloud-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-alicloud-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-alicloud-api/client/listers/db/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ReadonlyInstanceInformer provides access to a shared informer and lister for
// ReadonlyInstances.
type ReadonlyInstanceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ReadonlyInstanceLister
}

type readonlyInstanceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewReadonlyInstanceInformer constructs a new informer for ReadonlyInstance type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewReadonlyInstanceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredReadonlyInstanceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredReadonlyInstanceInformer constructs a new informer for ReadonlyInstance type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredReadonlyInstanceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DbV1alpha1().ReadonlyInstances(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DbV1alpha1().ReadonlyInstances(namespace).Watch(context.TODO(), options)
			},
		},
		&dbv1alpha1.ReadonlyInstance{},
		resyncPeriod,
		indexers,
	)
}

func (f *readonlyInstanceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredReadonlyInstanceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *readonlyInstanceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&dbv1alpha1.ReadonlyInstance{}, f.defaultInformer)
}

func (f *readonlyInstanceInformer) Lister() v1alpha1.ReadonlyInstanceLister {
	return v1alpha1.NewReadonlyInstanceLister(f.Informer().GetIndexer())
}