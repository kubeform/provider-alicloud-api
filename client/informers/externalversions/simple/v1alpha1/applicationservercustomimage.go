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

	simplev1alpha1 "kubeform.dev/provider-alicloud-api/apis/simple/v1alpha1"
	versioned "kubeform.dev/provider-alicloud-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-alicloud-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-alicloud-api/client/listers/simple/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ApplicationServerCustomImageInformer provides access to a shared informer and lister for
// ApplicationServerCustomImages.
type ApplicationServerCustomImageInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ApplicationServerCustomImageLister
}

type applicationServerCustomImageInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewApplicationServerCustomImageInformer constructs a new informer for ApplicationServerCustomImage type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewApplicationServerCustomImageInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredApplicationServerCustomImageInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredApplicationServerCustomImageInformer constructs a new informer for ApplicationServerCustomImage type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredApplicationServerCustomImageInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SimpleV1alpha1().ApplicationServerCustomImages(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SimpleV1alpha1().ApplicationServerCustomImages(namespace).Watch(context.TODO(), options)
			},
		},
		&simplev1alpha1.ApplicationServerCustomImage{},
		resyncPeriod,
		indexers,
	)
}

func (f *applicationServerCustomImageInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredApplicationServerCustomImageInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *applicationServerCustomImageInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&simplev1alpha1.ApplicationServerCustomImage{}, f.defaultInformer)
}

func (f *applicationServerCustomImageInformer) Lister() v1alpha1.ApplicationServerCustomImageLister {
	return v1alpha1.NewApplicationServerCustomImageLister(f.Informer().GetIndexer())
}
