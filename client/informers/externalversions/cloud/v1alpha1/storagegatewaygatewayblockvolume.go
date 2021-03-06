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

	cloudv1alpha1 "kubeform.dev/provider-alicloud-api/apis/cloud/v1alpha1"
	versioned "kubeform.dev/provider-alicloud-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-alicloud-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-alicloud-api/client/listers/cloud/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// StorageGatewayGatewayBlockVolumeInformer provides access to a shared informer and lister for
// StorageGatewayGatewayBlockVolumes.
type StorageGatewayGatewayBlockVolumeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.StorageGatewayGatewayBlockVolumeLister
}

type storageGatewayGatewayBlockVolumeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewStorageGatewayGatewayBlockVolumeInformer constructs a new informer for StorageGatewayGatewayBlockVolume type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewStorageGatewayGatewayBlockVolumeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredStorageGatewayGatewayBlockVolumeInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredStorageGatewayGatewayBlockVolumeInformer constructs a new informer for StorageGatewayGatewayBlockVolume type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredStorageGatewayGatewayBlockVolumeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CloudV1alpha1().StorageGatewayGatewayBlockVolumes(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CloudV1alpha1().StorageGatewayGatewayBlockVolumes(namespace).Watch(context.TODO(), options)
			},
		},
		&cloudv1alpha1.StorageGatewayGatewayBlockVolume{},
		resyncPeriod,
		indexers,
	)
}

func (f *storageGatewayGatewayBlockVolumeInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredStorageGatewayGatewayBlockVolumeInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *storageGatewayGatewayBlockVolumeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cloudv1alpha1.StorageGatewayGatewayBlockVolume{}, f.defaultInformer)
}

func (f *storageGatewayGatewayBlockVolumeInformer) Lister() v1alpha1.StorageGatewayGatewayBlockVolumeLister {
	return v1alpha1.NewStorageGatewayGatewayBlockVolumeLister(f.Informer().GetIndexer())
}
