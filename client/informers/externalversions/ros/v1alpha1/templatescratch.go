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

	rosv1alpha1 "kubeform.dev/provider-alicloud-api/apis/ros/v1alpha1"
	versioned "kubeform.dev/provider-alicloud-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-alicloud-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-alicloud-api/client/listers/ros/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TemplateScratchInformer provides access to a shared informer and lister for
// TemplateScratches.
type TemplateScratchInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.TemplateScratchLister
}

type templateScratchInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTemplateScratchInformer constructs a new informer for TemplateScratch type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTemplateScratchInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTemplateScratchInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTemplateScratchInformer constructs a new informer for TemplateScratch type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTemplateScratchInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RosV1alpha1().TemplateScratches(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RosV1alpha1().TemplateScratches(namespace).Watch(context.TODO(), options)
			},
		},
		&rosv1alpha1.TemplateScratch{},
		resyncPeriod,
		indexers,
	)
}

func (f *templateScratchInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTemplateScratchInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *templateScratchInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&rosv1alpha1.TemplateScratch{}, f.defaultInformer)
}

func (f *templateScratchInformer) Lister() v1alpha1.TemplateScratchLister {
	return v1alpha1.NewTemplateScratchLister(f.Informer().GetIndexer())
}
