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
	// BridgeEventBuses returns a BridgeEventBusInformer.
	BridgeEventBuses() BridgeEventBusInformer
	// BridgeEventSources returns a BridgeEventSourceInformer.
	BridgeEventSources() BridgeEventSourceInformer
	// BridgeRules returns a BridgeRuleInformer.
	BridgeRules() BridgeRuleInformer
	// BridgeServiceLinkedRoles returns a BridgeServiceLinkedRoleInformer.
	BridgeServiceLinkedRoles() BridgeServiceLinkedRoleInformer
	// BridgeSlrs returns a BridgeSlrInformer.
	BridgeSlrs() BridgeSlrInformer
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

// BridgeEventBuses returns a BridgeEventBusInformer.
func (v *version) BridgeEventBuses() BridgeEventBusInformer {
	return &bridgeEventBusInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BridgeEventSources returns a BridgeEventSourceInformer.
func (v *version) BridgeEventSources() BridgeEventSourceInformer {
	return &bridgeEventSourceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BridgeRules returns a BridgeRuleInformer.
func (v *version) BridgeRules() BridgeRuleInformer {
	return &bridgeRuleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BridgeServiceLinkedRoles returns a BridgeServiceLinkedRoleInformer.
func (v *version) BridgeServiceLinkedRoles() BridgeServiceLinkedRoleInformer {
	return &bridgeServiceLinkedRoleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BridgeSlrs returns a BridgeSlrInformer.
func (v *version) BridgeSlrs() BridgeSlrInformer {
	return &bridgeSlrInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
