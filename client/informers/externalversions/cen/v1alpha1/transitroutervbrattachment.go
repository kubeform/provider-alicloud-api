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

	cenv1alpha1 "kubeform.dev/provider-alicloud-api/apis/cen/v1alpha1"
	versioned "kubeform.dev/provider-alicloud-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-alicloud-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-alicloud-api/client/listers/cen/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TransitRouterVbrAttachmentInformer provides access to a shared informer and lister for
// TransitRouterVbrAttachments.
type TransitRouterVbrAttachmentInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.TransitRouterVbrAttachmentLister
}

type transitRouterVbrAttachmentInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTransitRouterVbrAttachmentInformer constructs a new informer for TransitRouterVbrAttachment type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTransitRouterVbrAttachmentInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTransitRouterVbrAttachmentInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTransitRouterVbrAttachmentInformer constructs a new informer for TransitRouterVbrAttachment type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTransitRouterVbrAttachmentInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CenV1alpha1().TransitRouterVbrAttachments(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CenV1alpha1().TransitRouterVbrAttachments(namespace).Watch(context.TODO(), options)
			},
		},
		&cenv1alpha1.TransitRouterVbrAttachment{},
		resyncPeriod,
		indexers,
	)
}

func (f *transitRouterVbrAttachmentInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTransitRouterVbrAttachmentInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *transitRouterVbrAttachmentInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cenv1alpha1.TransitRouterVbrAttachment{}, f.defaultInformer)
}

func (f *transitRouterVbrAttachmentInformer) Lister() v1alpha1.TransitRouterVbrAttachmentLister {
	return v1alpha1.NewTransitRouterVbrAttachmentLister(f.Informer().GetIndexer())
}
