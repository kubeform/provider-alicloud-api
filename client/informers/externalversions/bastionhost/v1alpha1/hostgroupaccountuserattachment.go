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

	bastionhostv1alpha1 "kubeform.dev/provider-alicloud-api/apis/bastionhost/v1alpha1"
	versioned "kubeform.dev/provider-alicloud-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-alicloud-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-alicloud-api/client/listers/bastionhost/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// HostGroupAccountUserAttachmentInformer provides access to a shared informer and lister for
// HostGroupAccountUserAttachments.
type HostGroupAccountUserAttachmentInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.HostGroupAccountUserAttachmentLister
}

type hostGroupAccountUserAttachmentInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewHostGroupAccountUserAttachmentInformer constructs a new informer for HostGroupAccountUserAttachment type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewHostGroupAccountUserAttachmentInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredHostGroupAccountUserAttachmentInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredHostGroupAccountUserAttachmentInformer constructs a new informer for HostGroupAccountUserAttachment type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredHostGroupAccountUserAttachmentInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BastionhostV1alpha1().HostGroupAccountUserAttachments(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BastionhostV1alpha1().HostGroupAccountUserAttachments(namespace).Watch(context.TODO(), options)
			},
		},
		&bastionhostv1alpha1.HostGroupAccountUserAttachment{},
		resyncPeriod,
		indexers,
	)
}

func (f *hostGroupAccountUserAttachmentInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredHostGroupAccountUserAttachmentInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *hostGroupAccountUserAttachmentInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&bastionhostv1alpha1.HostGroupAccountUserAttachment{}, f.defaultInformer)
}

func (f *hostGroupAccountUserAttachmentInformer) Lister() v1alpha1.HostGroupAccountUserAttachmentLister {
	return v1alpha1.NewHostGroupAccountUserAttachmentLister(f.Informer().GetIndexer())
}