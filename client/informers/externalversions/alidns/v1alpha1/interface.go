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
	// Domains returns a DomainInformer.
	Domains() DomainInformer
	// DomainAttachments returns a DomainAttachmentInformer.
	DomainAttachments() DomainAttachmentInformer
	// DomainGroups returns a DomainGroupInformer.
	DomainGroups() DomainGroupInformer
	// Instances returns a InstanceInformer.
	Instances() InstanceInformer
	// Records returns a RecordInformer.
	Records() RecordInformer
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

// Domains returns a DomainInformer.
func (v *version) Domains() DomainInformer {
	return &domainInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DomainAttachments returns a DomainAttachmentInformer.
func (v *version) DomainAttachments() DomainAttachmentInformer {
	return &domainAttachmentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DomainGroups returns a DomainGroupInformer.
func (v *version) DomainGroups() DomainGroupInformer {
	return &domainGroupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Instances returns a InstanceInformer.
func (v *version) Instances() InstanceInformer {
	return &instanceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Records returns a RecordInformer.
func (v *version) Records() RecordInformer {
	return &recordInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
