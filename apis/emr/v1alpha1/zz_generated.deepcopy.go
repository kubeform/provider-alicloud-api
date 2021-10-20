//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	apiv1alpha1 "kubeform.dev/apimachinery/api/v1alpha1"

	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Cluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterList) DeepCopyInto(out *ClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Cluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterList.
func (in *ClusterList) DeepCopy() *ClusterList {
	if in == nil {
		return nil
	}
	out := new(ClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ClusterSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecBootstrapAction) DeepCopyInto(out *ClusterSpecBootstrapAction) {
	*out = *in
	if in.Arg != nil {
		in, out := &in.Arg, &out.Arg
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecBootstrapAction.
func (in *ClusterSpecBootstrapAction) DeepCopy() *ClusterSpecBootstrapAction {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecBootstrapAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecHostGroup) DeepCopyInto(out *ClusterSpecHostGroup) {
	*out = *in
	if in.AutoRenew != nil {
		in, out := &in.AutoRenew, &out.AutoRenew
		*out = new(bool)
		**out = **in
	}
	if in.ChargeType != nil {
		in, out := &in.ChargeType, &out.ChargeType
		*out = new(string)
		**out = **in
	}
	if in.DiskCapacity != nil {
		in, out := &in.DiskCapacity, &out.DiskCapacity
		*out = new(string)
		**out = **in
	}
	if in.DiskCount != nil {
		in, out := &in.DiskCount, &out.DiskCount
		*out = new(string)
		**out = **in
	}
	if in.DiskType != nil {
		in, out := &in.DiskType, &out.DiskType
		*out = new(string)
		**out = **in
	}
	if in.GpuDriver != nil {
		in, out := &in.GpuDriver, &out.GpuDriver
		*out = new(string)
		**out = **in
	}
	if in.HostGroupName != nil {
		in, out := &in.HostGroupName, &out.HostGroupName
		*out = new(string)
		**out = **in
	}
	if in.HostGroupType != nil {
		in, out := &in.HostGroupType, &out.HostGroupType
		*out = new(string)
		**out = **in
	}
	if in.InstanceList != nil {
		in, out := &in.InstanceList, &out.InstanceList
		*out = new(string)
		**out = **in
	}
	if in.InstanceType != nil {
		in, out := &in.InstanceType, &out.InstanceType
		*out = new(string)
		**out = **in
	}
	if in.NodeCount != nil {
		in, out := &in.NodeCount, &out.NodeCount
		*out = new(string)
		**out = **in
	}
	if in.Period != nil {
		in, out := &in.Period, &out.Period
		*out = new(int64)
		**out = **in
	}
	if in.SysDiskCapacity != nil {
		in, out := &in.SysDiskCapacity, &out.SysDiskCapacity
		*out = new(string)
		**out = **in
	}
	if in.SysDiskType != nil {
		in, out := &in.SysDiskType, &out.SysDiskType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecHostGroup.
func (in *ClusterSpecHostGroup) DeepCopy() *ClusterSpecHostGroup {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecHostGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpecResource) DeepCopyInto(out *ClusterSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.BootstrapAction != nil {
		in, out := &in.BootstrapAction, &out.BootstrapAction
		*out = make([]ClusterSpecBootstrapAction, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ChargeType != nil {
		in, out := &in.ChargeType, &out.ChargeType
		*out = new(string)
		**out = **in
	}
	if in.ClusterType != nil {
		in, out := &in.ClusterType, &out.ClusterType
		*out = new(string)
		**out = **in
	}
	if in.DepositType != nil {
		in, out := &in.DepositType, &out.DepositType
		*out = new(string)
		**out = **in
	}
	if in.EasEnable != nil {
		in, out := &in.EasEnable, &out.EasEnable
		*out = new(bool)
		**out = **in
	}
	if in.EmrVer != nil {
		in, out := &in.EmrVer, &out.EmrVer
		*out = new(string)
		**out = **in
	}
	if in.HighAvailabilityEnable != nil {
		in, out := &in.HighAvailabilityEnable, &out.HighAvailabilityEnable
		*out = new(bool)
		**out = **in
	}
	if in.HostGroup != nil {
		in, out := &in.HostGroup, &out.HostGroup
		*out = make([]ClusterSpecHostGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IsOpenPublicIP != nil {
		in, out := &in.IsOpenPublicIP, &out.IsOpenPublicIP
		*out = new(bool)
		**out = **in
	}
	if in.KeyPairName != nil {
		in, out := &in.KeyPairName, &out.KeyPairName
		*out = new(string)
		**out = **in
	}
	if in.MasterPwd != nil {
		in, out := &in.MasterPwd, &out.MasterPwd
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.OptionSoftwareList != nil {
		in, out := &in.OptionSoftwareList, &out.OptionSoftwareList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Period != nil {
		in, out := &in.Period, &out.Period
		*out = new(int64)
		**out = **in
	}
	if in.RelatedClusterID != nil {
		in, out := &in.RelatedClusterID, &out.RelatedClusterID
		*out = new(string)
		**out = **in
	}
	if in.SecurityGroupID != nil {
		in, out := &in.SecurityGroupID, &out.SecurityGroupID
		*out = new(string)
		**out = **in
	}
	if in.SshEnable != nil {
		in, out := &in.SshEnable, &out.SshEnable
		*out = new(bool)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.UseLocalMetadb != nil {
		in, out := &in.UseLocalMetadb, &out.UseLocalMetadb
		*out = new(bool)
		**out = **in
	}
	if in.UserDefinedEmrEcsRole != nil {
		in, out := &in.UserDefinedEmrEcsRole, &out.UserDefinedEmrEcsRole
		*out = new(string)
		**out = **in
	}
	if in.VswitchID != nil {
		in, out := &in.VswitchID, &out.VswitchID
		*out = new(string)
		**out = **in
	}
	if in.ZoneID != nil {
		in, out := &in.ZoneID, &out.ZoneID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpecResource.
func (in *ClusterSpecResource) DeepCopy() *ClusterSpecResource {
	if in == nil {
		return nil
	}
	out := new(ClusterSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStatus) DeepCopyInto(out *ClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStatus.
func (in *ClusterStatus) DeepCopy() *ClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterStatus)
	in.DeepCopyInto(out)
	return out
}
