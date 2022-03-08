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
	// ConnectNetworks returns a ConnectNetworkInformer.
	ConnectNetworks() ConnectNetworkInformer
	// ConnectNetworkAttachments returns a ConnectNetworkAttachmentInformer.
	ConnectNetworkAttachments() ConnectNetworkAttachmentInformer
	// ConnectNetworkGrants returns a ConnectNetworkGrantInformer.
	ConnectNetworkGrants() ConnectNetworkGrantInformer
	// FirewallControlPolicies returns a FirewallControlPolicyInformer.
	FirewallControlPolicies() FirewallControlPolicyInformer
	// FirewallControlPolicyOrders returns a FirewallControlPolicyOrderInformer.
	FirewallControlPolicyOrders() FirewallControlPolicyOrderInformer
	// FirewallInstances returns a FirewallInstanceInformer.
	FirewallInstances() FirewallInstanceInformer
	// SsoAccessAssignments returns a SsoAccessAssignmentInformer.
	SsoAccessAssignments() SsoAccessAssignmentInformer
	// SsoAccessConfigurations returns a SsoAccessConfigurationInformer.
	SsoAccessConfigurations() SsoAccessConfigurationInformer
	// SsoAccessConfigurationProvisionings returns a SsoAccessConfigurationProvisioningInformer.
	SsoAccessConfigurationProvisionings() SsoAccessConfigurationProvisioningInformer
	// SsoDirectories returns a SsoDirectoryInformer.
	SsoDirectories() SsoDirectoryInformer
	// SsoGroups returns a SsoGroupInformer.
	SsoGroups() SsoGroupInformer
	// SsoScimServerCredentials returns a SsoScimServerCredentialInformer.
	SsoScimServerCredentials() SsoScimServerCredentialInformer
	// SsoUsers returns a SsoUserInformer.
	SsoUsers() SsoUserInformer
	// SsoUserAttachments returns a SsoUserAttachmentInformer.
	SsoUserAttachments() SsoUserAttachmentInformer
	// StorageGatewayExpressSyncs returns a StorageGatewayExpressSyncInformer.
	StorageGatewayExpressSyncs() StorageGatewayExpressSyncInformer
	// StorageGatewayExpressSyncShareAttachments returns a StorageGatewayExpressSyncShareAttachmentInformer.
	StorageGatewayExpressSyncShareAttachments() StorageGatewayExpressSyncShareAttachmentInformer
	// StorageGatewayGateways returns a StorageGatewayGatewayInformer.
	StorageGatewayGateways() StorageGatewayGatewayInformer
	// StorageGatewayGatewayBlockVolumes returns a StorageGatewayGatewayBlockVolumeInformer.
	StorageGatewayGatewayBlockVolumes() StorageGatewayGatewayBlockVolumeInformer
	// StorageGatewayGatewayCacheDisks returns a StorageGatewayGatewayCacheDiskInformer.
	StorageGatewayGatewayCacheDisks() StorageGatewayGatewayCacheDiskInformer
	// StorageGatewayGatewayFileShares returns a StorageGatewayGatewayFileShareInformer.
	StorageGatewayGatewayFileShares() StorageGatewayGatewayFileShareInformer
	// StorageGatewayGatewayLoggings returns a StorageGatewayGatewayLoggingInformer.
	StorageGatewayGatewayLoggings() StorageGatewayGatewayLoggingInformer
	// StorageGatewayGatewaySmbUsers returns a StorageGatewayGatewaySmbUserInformer.
	StorageGatewayGatewaySmbUsers() StorageGatewayGatewaySmbUserInformer
	// StorageGatewayStorageBundles returns a StorageGatewayStorageBundleInformer.
	StorageGatewayStorageBundles() StorageGatewayStorageBundleInformer
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

// ConnectNetworks returns a ConnectNetworkInformer.
func (v *version) ConnectNetworks() ConnectNetworkInformer {
	return &connectNetworkInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ConnectNetworkAttachments returns a ConnectNetworkAttachmentInformer.
func (v *version) ConnectNetworkAttachments() ConnectNetworkAttachmentInformer {
	return &connectNetworkAttachmentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ConnectNetworkGrants returns a ConnectNetworkGrantInformer.
func (v *version) ConnectNetworkGrants() ConnectNetworkGrantInformer {
	return &connectNetworkGrantInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FirewallControlPolicies returns a FirewallControlPolicyInformer.
func (v *version) FirewallControlPolicies() FirewallControlPolicyInformer {
	return &firewallControlPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FirewallControlPolicyOrders returns a FirewallControlPolicyOrderInformer.
func (v *version) FirewallControlPolicyOrders() FirewallControlPolicyOrderInformer {
	return &firewallControlPolicyOrderInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FirewallInstances returns a FirewallInstanceInformer.
func (v *version) FirewallInstances() FirewallInstanceInformer {
	return &firewallInstanceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SsoAccessAssignments returns a SsoAccessAssignmentInformer.
func (v *version) SsoAccessAssignments() SsoAccessAssignmentInformer {
	return &ssoAccessAssignmentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SsoAccessConfigurations returns a SsoAccessConfigurationInformer.
func (v *version) SsoAccessConfigurations() SsoAccessConfigurationInformer {
	return &ssoAccessConfigurationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SsoAccessConfigurationProvisionings returns a SsoAccessConfigurationProvisioningInformer.
func (v *version) SsoAccessConfigurationProvisionings() SsoAccessConfigurationProvisioningInformer {
	return &ssoAccessConfigurationProvisioningInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SsoDirectories returns a SsoDirectoryInformer.
func (v *version) SsoDirectories() SsoDirectoryInformer {
	return &ssoDirectoryInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SsoGroups returns a SsoGroupInformer.
func (v *version) SsoGroups() SsoGroupInformer {
	return &ssoGroupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SsoScimServerCredentials returns a SsoScimServerCredentialInformer.
func (v *version) SsoScimServerCredentials() SsoScimServerCredentialInformer {
	return &ssoScimServerCredentialInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SsoUsers returns a SsoUserInformer.
func (v *version) SsoUsers() SsoUserInformer {
	return &ssoUserInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SsoUserAttachments returns a SsoUserAttachmentInformer.
func (v *version) SsoUserAttachments() SsoUserAttachmentInformer {
	return &ssoUserAttachmentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// StorageGatewayExpressSyncs returns a StorageGatewayExpressSyncInformer.
func (v *version) StorageGatewayExpressSyncs() StorageGatewayExpressSyncInformer {
	return &storageGatewayExpressSyncInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// StorageGatewayExpressSyncShareAttachments returns a StorageGatewayExpressSyncShareAttachmentInformer.
func (v *version) StorageGatewayExpressSyncShareAttachments() StorageGatewayExpressSyncShareAttachmentInformer {
	return &storageGatewayExpressSyncShareAttachmentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// StorageGatewayGateways returns a StorageGatewayGatewayInformer.
func (v *version) StorageGatewayGateways() StorageGatewayGatewayInformer {
	return &storageGatewayGatewayInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// StorageGatewayGatewayBlockVolumes returns a StorageGatewayGatewayBlockVolumeInformer.
func (v *version) StorageGatewayGatewayBlockVolumes() StorageGatewayGatewayBlockVolumeInformer {
	return &storageGatewayGatewayBlockVolumeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// StorageGatewayGatewayCacheDisks returns a StorageGatewayGatewayCacheDiskInformer.
func (v *version) StorageGatewayGatewayCacheDisks() StorageGatewayGatewayCacheDiskInformer {
	return &storageGatewayGatewayCacheDiskInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// StorageGatewayGatewayFileShares returns a StorageGatewayGatewayFileShareInformer.
func (v *version) StorageGatewayGatewayFileShares() StorageGatewayGatewayFileShareInformer {
	return &storageGatewayGatewayFileShareInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// StorageGatewayGatewayLoggings returns a StorageGatewayGatewayLoggingInformer.
func (v *version) StorageGatewayGatewayLoggings() StorageGatewayGatewayLoggingInformer {
	return &storageGatewayGatewayLoggingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// StorageGatewayGatewaySmbUsers returns a StorageGatewayGatewaySmbUserInformer.
func (v *version) StorageGatewayGatewaySmbUsers() StorageGatewayGatewaySmbUserInformer {
	return &storageGatewayGatewaySmbUserInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// StorageGatewayStorageBundles returns a StorageGatewayStorageBundleInformer.
func (v *version) StorageGatewayStorageBundles() StorageGatewayStorageBundleInformer {
	return &storageGatewayStorageBundleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
