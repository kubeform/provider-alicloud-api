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

	ramv1alpha1 "kubeform.dev/provider-alicloud-api/apis/ram/v1alpha1"
	versioned "kubeform.dev/provider-alicloud-api/client/clientset/versioned"
	internalinterfaces "kubeform.dev/provider-alicloud-api/client/informers/externalversions/internalinterfaces"
	v1alpha1 "kubeform.dev/provider-alicloud-api/client/listers/ram/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// AccountPasswordPolicyInformer provides access to a shared informer and lister for
// AccountPasswordPolicies.
type AccountPasswordPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.AccountPasswordPolicyLister
}

type accountPasswordPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAccountPasswordPolicyInformer constructs a new informer for AccountPasswordPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAccountPasswordPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAccountPasswordPolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAccountPasswordPolicyInformer constructs a new informer for AccountPasswordPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAccountPasswordPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RamV1alpha1().AccountPasswordPolicies(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RamV1alpha1().AccountPasswordPolicies(namespace).Watch(context.TODO(), options)
			},
		},
		&ramv1alpha1.AccountPasswordPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *accountPasswordPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAccountPasswordPolicyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *accountPasswordPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&ramv1alpha1.AccountPasswordPolicy{}, f.defaultInformer)
}

func (f *accountPasswordPolicyInformer) Lister() v1alpha1.AccountPasswordPolicyLister {
	return v1alpha1.NewAccountPasswordPolicyLister(f.Informer().GetIndexer())
}