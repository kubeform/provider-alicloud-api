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

	resourcev1alpha1 "kubeform.dev/provider-alicloud-api/apis/resource/v1alpha1"
	versioned "kubeform.dev/provider-alicloud-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-alicloud-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-alicloud-api/client/listers/resource/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ManagerResourceDirectoryInformer provides access to a shared informer and lister for
// ManagerResourceDirectories.
type ManagerResourceDirectoryInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ManagerResourceDirectoryLister
}

type managerResourceDirectoryInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewManagerResourceDirectoryInformer constructs a new informer for ManagerResourceDirectory type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewManagerResourceDirectoryInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredManagerResourceDirectoryInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredManagerResourceDirectoryInformer constructs a new informer for ManagerResourceDirectory type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredManagerResourceDirectoryInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResourceV1alpha1().ManagerResourceDirectories(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResourceV1alpha1().ManagerResourceDirectories(namespace).Watch(context.TODO(), options)
			},
		},
		&resourcev1alpha1.ManagerResourceDirectory{},
		resyncPeriod,
		indexers,
	)
}

func (f *managerResourceDirectoryInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredManagerResourceDirectoryInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *managerResourceDirectoryInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&resourcev1alpha1.ManagerResourceDirectory{}, f.defaultInformer)
}

func (f *managerResourceDirectoryInformer) Lister() v1alpha1.ManagerResourceDirectoryLister {
	return v1alpha1.NewManagerResourceDirectoryLister(f.Informer().GetIndexer())
}
