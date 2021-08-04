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
	// Applications returns a ApplicationInformer.
	Applications() ApplicationInformer
	// ApplicationDeployments returns a ApplicationDeploymentInformer.
	ApplicationDeployments() ApplicationDeploymentInformer
	// ApplicationScales returns a ApplicationScaleInformer.
	ApplicationScales() ApplicationScaleInformer
	// Clusters returns a ClusterInformer.
	Clusters() ClusterInformer
	// DeployGroups returns a DeployGroupInformer.
	DeployGroups() DeployGroupInformer
	// InstanceClusterAttachments returns a InstanceClusterAttachmentInformer.
	InstanceClusterAttachments() InstanceClusterAttachmentInformer
	// K8sApplications returns a K8sApplicationInformer.
	K8sApplications() K8sApplicationInformer
	// K8sClusters returns a K8sClusterInformer.
	K8sClusters() K8sClusterInformer
	// SlbAttachments returns a SlbAttachmentInformer.
	SlbAttachments() SlbAttachmentInformer
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

// Applications returns a ApplicationInformer.
func (v *version) Applications() ApplicationInformer {
	return &applicationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ApplicationDeployments returns a ApplicationDeploymentInformer.
func (v *version) ApplicationDeployments() ApplicationDeploymentInformer {
	return &applicationDeploymentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ApplicationScales returns a ApplicationScaleInformer.
func (v *version) ApplicationScales() ApplicationScaleInformer {
	return &applicationScaleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Clusters returns a ClusterInformer.
func (v *version) Clusters() ClusterInformer {
	return &clusterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DeployGroups returns a DeployGroupInformer.
func (v *version) DeployGroups() DeployGroupInformer {
	return &deployGroupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// InstanceClusterAttachments returns a InstanceClusterAttachmentInformer.
func (v *version) InstanceClusterAttachments() InstanceClusterAttachmentInformer {
	return &instanceClusterAttachmentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// K8sApplications returns a K8sApplicationInformer.
func (v *version) K8sApplications() K8sApplicationInformer {
	return &k8sApplicationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// K8sClusters returns a K8sClusterInformer.
func (v *version) K8sClusters() K8sClusterInformer {
	return &k8sClusterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SlbAttachments returns a SlbAttachmentInformer.
func (v *version) SlbAttachments() SlbAttachmentInformer {
	return &slbAttachmentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}