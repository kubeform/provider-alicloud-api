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
	// ManagerAccounts returns a ManagerAccountInformer.
	ManagerAccounts() ManagerAccountInformer
	// ManagerControlPolicies returns a ManagerControlPolicyInformer.
	ManagerControlPolicies() ManagerControlPolicyInformer
	// ManagerControlPolicyAttachments returns a ManagerControlPolicyAttachmentInformer.
	ManagerControlPolicyAttachments() ManagerControlPolicyAttachmentInformer
	// ManagerFolders returns a ManagerFolderInformer.
	ManagerFolders() ManagerFolderInformer
	// ManagerHandshakes returns a ManagerHandshakeInformer.
	ManagerHandshakes() ManagerHandshakeInformer
	// ManagerPolicies returns a ManagerPolicyInformer.
	ManagerPolicies() ManagerPolicyInformer
	// ManagerPolicyAttachments returns a ManagerPolicyAttachmentInformer.
	ManagerPolicyAttachments() ManagerPolicyAttachmentInformer
	// ManagerPolicyVersions returns a ManagerPolicyVersionInformer.
	ManagerPolicyVersions() ManagerPolicyVersionInformer
	// ManagerResourceDirectories returns a ManagerResourceDirectoryInformer.
	ManagerResourceDirectories() ManagerResourceDirectoryInformer
	// ManagerResourceGroups returns a ManagerResourceGroupInformer.
	ManagerResourceGroups() ManagerResourceGroupInformer
	// ManagerResourceShares returns a ManagerResourceShareInformer.
	ManagerResourceShares() ManagerResourceShareInformer
	// ManagerRoles returns a ManagerRoleInformer.
	ManagerRoles() ManagerRoleInformer
	// ManagerSharedResources returns a ManagerSharedResourceInformer.
	ManagerSharedResources() ManagerSharedResourceInformer
	// ManagerSharedTargets returns a ManagerSharedTargetInformer.
	ManagerSharedTargets() ManagerSharedTargetInformer
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

// ManagerAccounts returns a ManagerAccountInformer.
func (v *version) ManagerAccounts() ManagerAccountInformer {
	return &managerAccountInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ManagerControlPolicies returns a ManagerControlPolicyInformer.
func (v *version) ManagerControlPolicies() ManagerControlPolicyInformer {
	return &managerControlPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ManagerControlPolicyAttachments returns a ManagerControlPolicyAttachmentInformer.
func (v *version) ManagerControlPolicyAttachments() ManagerControlPolicyAttachmentInformer {
	return &managerControlPolicyAttachmentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ManagerFolders returns a ManagerFolderInformer.
func (v *version) ManagerFolders() ManagerFolderInformer {
	return &managerFolderInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ManagerHandshakes returns a ManagerHandshakeInformer.
func (v *version) ManagerHandshakes() ManagerHandshakeInformer {
	return &managerHandshakeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ManagerPolicies returns a ManagerPolicyInformer.
func (v *version) ManagerPolicies() ManagerPolicyInformer {
	return &managerPolicyInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ManagerPolicyAttachments returns a ManagerPolicyAttachmentInformer.
func (v *version) ManagerPolicyAttachments() ManagerPolicyAttachmentInformer {
	return &managerPolicyAttachmentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ManagerPolicyVersions returns a ManagerPolicyVersionInformer.
func (v *version) ManagerPolicyVersions() ManagerPolicyVersionInformer {
	return &managerPolicyVersionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ManagerResourceDirectories returns a ManagerResourceDirectoryInformer.
func (v *version) ManagerResourceDirectories() ManagerResourceDirectoryInformer {
	return &managerResourceDirectoryInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ManagerResourceGroups returns a ManagerResourceGroupInformer.
func (v *version) ManagerResourceGroups() ManagerResourceGroupInformer {
	return &managerResourceGroupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ManagerResourceShares returns a ManagerResourceShareInformer.
func (v *version) ManagerResourceShares() ManagerResourceShareInformer {
	return &managerResourceShareInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ManagerRoles returns a ManagerRoleInformer.
func (v *version) ManagerRoles() ManagerRoleInformer {
	return &managerRoleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ManagerSharedResources returns a ManagerSharedResourceInformer.
func (v *version) ManagerSharedResources() ManagerSharedResourceInformer {
	return &managerSharedResourceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ManagerSharedTargets returns a ManagerSharedTargetInformer.
func (v *version) ManagerSharedTargets() ManagerSharedTargetInformer {
	return &managerSharedTargetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
