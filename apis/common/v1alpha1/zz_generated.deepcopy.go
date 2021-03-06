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
func (in *BandwidthPackage) DeepCopyInto(out *BandwidthPackage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BandwidthPackage.
func (in *BandwidthPackage) DeepCopy() *BandwidthPackage {
	if in == nil {
		return nil
	}
	out := new(BandwidthPackage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BandwidthPackage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BandwidthPackageAttachment) DeepCopyInto(out *BandwidthPackageAttachment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BandwidthPackageAttachment.
func (in *BandwidthPackageAttachment) DeepCopy() *BandwidthPackageAttachment {
	if in == nil {
		return nil
	}
	out := new(BandwidthPackageAttachment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BandwidthPackageAttachment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BandwidthPackageAttachmentList) DeepCopyInto(out *BandwidthPackageAttachmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BandwidthPackageAttachment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BandwidthPackageAttachmentList.
func (in *BandwidthPackageAttachmentList) DeepCopy() *BandwidthPackageAttachmentList {
	if in == nil {
		return nil
	}
	out := new(BandwidthPackageAttachmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BandwidthPackageAttachmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BandwidthPackageAttachmentSpec) DeepCopyInto(out *BandwidthPackageAttachmentSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(BandwidthPackageAttachmentSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BandwidthPackageAttachmentSpec.
func (in *BandwidthPackageAttachmentSpec) DeepCopy() *BandwidthPackageAttachmentSpec {
	if in == nil {
		return nil
	}
	out := new(BandwidthPackageAttachmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BandwidthPackageAttachmentSpecResource) DeepCopyInto(out *BandwidthPackageAttachmentSpecResource) {
	*out = *in
	if in.BandwidthPackageID != nil {
		in, out := &in.BandwidthPackageID, &out.BandwidthPackageID
		*out = new(string)
		**out = **in
	}
	if in.InstanceID != nil {
		in, out := &in.InstanceID, &out.InstanceID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BandwidthPackageAttachmentSpecResource.
func (in *BandwidthPackageAttachmentSpecResource) DeepCopy() *BandwidthPackageAttachmentSpecResource {
	if in == nil {
		return nil
	}
	out := new(BandwidthPackageAttachmentSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BandwidthPackageAttachmentStatus) DeepCopyInto(out *BandwidthPackageAttachmentStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BandwidthPackageAttachmentStatus.
func (in *BandwidthPackageAttachmentStatus) DeepCopy() *BandwidthPackageAttachmentStatus {
	if in == nil {
		return nil
	}
	out := new(BandwidthPackageAttachmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BandwidthPackageList) DeepCopyInto(out *BandwidthPackageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BandwidthPackage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BandwidthPackageList.
func (in *BandwidthPackageList) DeepCopy() *BandwidthPackageList {
	if in == nil {
		return nil
	}
	out := new(BandwidthPackageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BandwidthPackageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BandwidthPackageSpec) DeepCopyInto(out *BandwidthPackageSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(BandwidthPackageSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BandwidthPackageSpec.
func (in *BandwidthPackageSpec) DeepCopy() *BandwidthPackageSpec {
	if in == nil {
		return nil
	}
	out := new(BandwidthPackageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BandwidthPackageSpecResource) DeepCopyInto(out *BandwidthPackageSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.Bandwidth != nil {
		in, out := &in.Bandwidth, &out.Bandwidth
		*out = new(string)
		**out = **in
	}
	if in.BandwidthPackageName != nil {
		in, out := &in.BandwidthPackageName, &out.BandwidthPackageName
		*out = new(string)
		**out = **in
	}
	if in.DeletionProtection != nil {
		in, out := &in.DeletionProtection, &out.DeletionProtection
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Force != nil {
		in, out := &in.Force, &out.Force
		*out = new(string)
		**out = **in
	}
	if in.InternetChargeType != nil {
		in, out := &in.InternetChargeType, &out.InternetChargeType
		*out = new(string)
		**out = **in
	}
	if in.Isp != nil {
		in, out := &in.Isp, &out.Isp
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Ratio != nil {
		in, out := &in.Ratio, &out.Ratio
		*out = new(int64)
		**out = **in
	}
	if in.ResourceGroupID != nil {
		in, out := &in.ResourceGroupID, &out.ResourceGroupID
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BandwidthPackageSpecResource.
func (in *BandwidthPackageSpecResource) DeepCopy() *BandwidthPackageSpecResource {
	if in == nil {
		return nil
	}
	out := new(BandwidthPackageSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BandwidthPackageStatus) DeepCopyInto(out *BandwidthPackageStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BandwidthPackageStatus.
func (in *BandwidthPackageStatus) DeepCopy() *BandwidthPackageStatus {
	if in == nil {
		return nil
	}
	out := new(BandwidthPackageStatus)
	in.DeepCopyInto(out)
	return out
}
