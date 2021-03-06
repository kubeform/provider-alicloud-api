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
	// ConsumerChannels returns a ConsumerChannelInformer.
	ConsumerChannels() ConsumerChannelInformer
	// JobMonitorRules returns a JobMonitorRuleInformer.
	JobMonitorRules() JobMonitorRuleInformer
	// MigrationInstances returns a MigrationInstanceInformer.
	MigrationInstances() MigrationInstanceInformer
	// MigrationJobs returns a MigrationJobInformer.
	MigrationJobs() MigrationJobInformer
	// SubscriptionJobs returns a SubscriptionJobInformer.
	SubscriptionJobs() SubscriptionJobInformer
	// SynchronizationInstances returns a SynchronizationInstanceInformer.
	SynchronizationInstances() SynchronizationInstanceInformer
	// SynchronizationJobs returns a SynchronizationJobInformer.
	SynchronizationJobs() SynchronizationJobInformer
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

// ConsumerChannels returns a ConsumerChannelInformer.
func (v *version) ConsumerChannels() ConsumerChannelInformer {
	return &consumerChannelInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// JobMonitorRules returns a JobMonitorRuleInformer.
func (v *version) JobMonitorRules() JobMonitorRuleInformer {
	return &jobMonitorRuleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MigrationInstances returns a MigrationInstanceInformer.
func (v *version) MigrationInstances() MigrationInstanceInformer {
	return &migrationInstanceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// MigrationJobs returns a MigrationJobInformer.
func (v *version) MigrationJobs() MigrationJobInformer {
	return &migrationJobInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SubscriptionJobs returns a SubscriptionJobInformer.
func (v *version) SubscriptionJobs() SubscriptionJobInformer {
	return &subscriptionJobInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SynchronizationInstances returns a SynchronizationInstanceInformer.
func (v *version) SynchronizationInstances() SynchronizationInstanceInformer {
	return &synchronizationInstanceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SynchronizationJobs returns a SynchronizationJobInformer.
func (v *version) SynchronizationJobs() SynchronizationJobInformer {
	return &synchronizationJobInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
