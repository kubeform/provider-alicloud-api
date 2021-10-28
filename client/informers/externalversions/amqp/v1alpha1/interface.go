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
	internalinterfaces "kubeform.dev/provider-alicloud-api/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Bindings returns a BindingInformer.
	Bindings() BindingInformer
	// Exchanges returns a ExchangeInformer.
	Exchanges() ExchangeInformer
	// Instances returns a InstanceInformer.
	Instances() InstanceInformer
	// Queues returns a QueueInformer.
	Queues() QueueInformer
	// VirtualHosts returns a VirtualHostInformer.
	VirtualHosts() VirtualHostInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Bindings returns a BindingInformer.
func (v *version) Bindings() BindingInformer {
	return &bindingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Exchanges returns a ExchangeInformer.
func (v *version) Exchanges() ExchangeInformer {
	return &exchangeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Instances returns a InstanceInformer.
func (v *version) Instances() InstanceInformer {
	return &instanceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Queues returns a QueueInformer.
func (v *version) Queues() QueueInformer {
	return &queueInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VirtualHosts returns a VirtualHostInformer.
func (v *version) VirtualHosts() VirtualHostInformer {
	return &virtualHostInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
