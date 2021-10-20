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
func (in *AnycastEipAddress) DeepCopyInto(out *AnycastEipAddress) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnycastEipAddress.
func (in *AnycastEipAddress) DeepCopy() *AnycastEipAddress {
	if in == nil {
		return nil
	}
	out := new(AnycastEipAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnycastEipAddress) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnycastEipAddressAttachment) DeepCopyInto(out *AnycastEipAddressAttachment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnycastEipAddressAttachment.
func (in *AnycastEipAddressAttachment) DeepCopy() *AnycastEipAddressAttachment {
	if in == nil {
		return nil
	}
	out := new(AnycastEipAddressAttachment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnycastEipAddressAttachment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnycastEipAddressAttachmentList) DeepCopyInto(out *AnycastEipAddressAttachmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AnycastEipAddressAttachment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnycastEipAddressAttachmentList.
func (in *AnycastEipAddressAttachmentList) DeepCopy() *AnycastEipAddressAttachmentList {
	if in == nil {
		return nil
	}
	out := new(AnycastEipAddressAttachmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnycastEipAddressAttachmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnycastEipAddressAttachmentSpec) DeepCopyInto(out *AnycastEipAddressAttachmentSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(AnycastEipAddressAttachmentSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnycastEipAddressAttachmentSpec.
func (in *AnycastEipAddressAttachmentSpec) DeepCopy() *AnycastEipAddressAttachmentSpec {
	if in == nil {
		return nil
	}
	out := new(AnycastEipAddressAttachmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnycastEipAddressAttachmentSpecResource) DeepCopyInto(out *AnycastEipAddressAttachmentSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.AnycastID != nil {
		in, out := &in.AnycastID, &out.AnycastID
		*out = new(string)
		**out = **in
	}
	if in.BindInstanceID != nil {
		in, out := &in.BindInstanceID, &out.BindInstanceID
		*out = new(string)
		**out = **in
	}
	if in.BindInstanceRegionID != nil {
		in, out := &in.BindInstanceRegionID, &out.BindInstanceRegionID
		*out = new(string)
		**out = **in
	}
	if in.BindInstanceType != nil {
		in, out := &in.BindInstanceType, &out.BindInstanceType
		*out = new(string)
		**out = **in
	}
	if in.BindTime != nil {
		in, out := &in.BindTime, &out.BindTime
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnycastEipAddressAttachmentSpecResource.
func (in *AnycastEipAddressAttachmentSpecResource) DeepCopy() *AnycastEipAddressAttachmentSpecResource {
	if in == nil {
		return nil
	}
	out := new(AnycastEipAddressAttachmentSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnycastEipAddressAttachmentStatus) DeepCopyInto(out *AnycastEipAddressAttachmentStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnycastEipAddressAttachmentStatus.
func (in *AnycastEipAddressAttachmentStatus) DeepCopy() *AnycastEipAddressAttachmentStatus {
	if in == nil {
		return nil
	}
	out := new(AnycastEipAddressAttachmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnycastEipAddressList) DeepCopyInto(out *AnycastEipAddressList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AnycastEipAddress, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnycastEipAddressList.
func (in *AnycastEipAddressList) DeepCopy() *AnycastEipAddressList {
	if in == nil {
		return nil
	}
	out := new(AnycastEipAddressList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AnycastEipAddressList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnycastEipAddressSpec) DeepCopyInto(out *AnycastEipAddressSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(AnycastEipAddressSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnycastEipAddressSpec.
func (in *AnycastEipAddressSpec) DeepCopy() *AnycastEipAddressSpec {
	if in == nil {
		return nil
	}
	out := new(AnycastEipAddressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnycastEipAddressSpecResource) DeepCopyInto(out *AnycastEipAddressSpecResource) {
	*out = *in
	if in.Timeouts != nil {
		in, out := &in.Timeouts, &out.Timeouts
		*out = new(apiv1alpha1.ResourceTimeout)
		(*in).DeepCopyInto(*out)
	}
	if in.AnycastEipAddressName != nil {
		in, out := &in.AnycastEipAddressName, &out.AnycastEipAddressName
		*out = new(string)
		**out = **in
	}
	if in.Bandwidth != nil {
		in, out := &in.Bandwidth, &out.Bandwidth
		*out = new(int64)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.InternetChargeType != nil {
		in, out := &in.InternetChargeType, &out.InternetChargeType
		*out = new(string)
		**out = **in
	}
	if in.PaymentType != nil {
		in, out := &in.PaymentType, &out.PaymentType
		*out = new(string)
		**out = **in
	}
	if in.ServiceLocation != nil {
		in, out := &in.ServiceLocation, &out.ServiceLocation
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnycastEipAddressSpecResource.
func (in *AnycastEipAddressSpecResource) DeepCopy() *AnycastEipAddressSpecResource {
	if in == nil {
		return nil
	}
	out := new(AnycastEipAddressSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AnycastEipAddressStatus) DeepCopyInto(out *AnycastEipAddressStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AnycastEipAddressStatus.
func (in *AnycastEipAddressStatus) DeepCopy() *AnycastEipAddressStatus {
	if in == nil {
		return nil
	}
	out := new(AnycastEipAddressStatus)
	in.DeepCopyInto(out)
	return out
}
