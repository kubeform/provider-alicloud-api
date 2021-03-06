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
	// Accounts returns a AccountInformer.
	Accounts() AccountInformer
	// AccountPrivileges returns a AccountPrivilegeInformer.
	AccountPrivileges() AccountPrivilegeInformer
	// BackupPolicies returns a BackupPolicyInformer.
	BackupPolicies() BackupPolicyInformer
	// Clusters returns a ClusterInformer.
	Clusters() ClusterInformer
	// Databases returns a DatabaseInformer.
	Databases() DatabaseInformer
	// Endpoints returns a EndpointInformer.
	Endpoints() EndpointInformer
	// EndpointAddresses returns a EndpointAddressInformer.
	EndpointAddresses() EndpointAddressInformer
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

// Accounts returns a AccountInformer.
func (v *version) Accounts() AccountInformer {
	return &accountInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// AccountPrivileges returns a AccountPrivilegeInformer.
func (v *version) AccountPrivileges() AccountPrivilegeInformer {
	return &accountPrivilegeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BackupPolicies returns a BackupPolicyInformer.
func (v *version) BackupPolicies() BackupPolicyInformer {
	return &backupPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Clusters returns a ClusterInformer.
func (v *version) Clusters() ClusterInformer {
	return &clusterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Databases returns a DatabaseInformer.
func (v *version) Databases() DatabaseInformer {
	return &databaseInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Endpoints returns a EndpointInformer.
func (v *version) Endpoints() EndpointInformer {
	return &endpointInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// EndpointAddresses returns a EndpointAddressInformer.
func (v *version) EndpointAddresses() EndpointAddressInformer {
	return &endpointAddressInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
