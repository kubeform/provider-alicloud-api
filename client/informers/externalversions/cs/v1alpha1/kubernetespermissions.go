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

	csv1alpha1 "kubeform.dev/provider-alicloud-api/apis/cs/v1alpha1"
	versioned "kubeform.dev/provider-alicloud-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-alicloud-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-alicloud-api/client/listers/cs/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// KubernetesPermissionsInformer provides access to a shared informer and lister for
// KubernetesPermissionses.
type KubernetesPermissionsInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.KubernetesPermissionsLister
}

type kubernetesPermissionsInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewKubernetesPermissionsInformer constructs a new informer for KubernetesPermissions type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKubernetesPermissionsInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKubernetesPermissionsInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredKubernetesPermissionsInformer constructs a new informer for KubernetesPermissions type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKubernetesPermissionsInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CsV1alpha1().KubernetesPermissionses(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CsV1alpha1().KubernetesPermissionses(namespace).Watch(context.TODO(), options)
			},
		},
		&csv1alpha1.KubernetesPermissions{},
		resyncPeriod,
		indexers,
	)
}

func (f *kubernetesPermissionsInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKubernetesPermissionsInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *kubernetesPermissionsInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&csv1alpha1.KubernetesPermissions{}, f.defaultInformer)
}

func (f *kubernetesPermissionsInformer) Lister() v1alpha1.KubernetesPermissionsLister {
	return v1alpha1.NewKubernetesPermissionsLister(f.Informer().GetIndexer())
}
