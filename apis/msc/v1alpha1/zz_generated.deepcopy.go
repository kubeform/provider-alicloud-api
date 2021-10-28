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
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubContact) DeepCopyInto(out *SubContact) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubContact.
func (in *SubContact) DeepCopy() *SubContact {
	if in == nil {
		return nil
	}
	out := new(SubContact)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubContact) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubContactList) DeepCopyInto(out *SubContactList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SubContact, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubContactList.
func (in *SubContactList) DeepCopy() *SubContactList {
	if in == nil {
		return nil
	}
	out := new(SubContactList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubContactList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubContactSpec) DeepCopyInto(out *SubContactSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(SubContactSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubContactSpec.
func (in *SubContactSpec) DeepCopy() *SubContactSpec {
	if in == nil {
		return nil
	}
	out := new(SubContactSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubContactSpecResource) DeepCopyInto(out *SubContactSpecResource) {
	*out = *in
	if in.ContactName != nil {
		in, out := &in.ContactName, &out.ContactName
		*out = new(string)
		**out = **in
	}
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.Mobile != nil {
		in, out := &in.Mobile, &out.Mobile
		*out = new(string)
		**out = **in
	}
	if in.Position != nil {
		in, out := &in.Position, &out.Position
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubContactSpecResource.
func (in *SubContactSpecResource) DeepCopy() *SubContactSpecResource {
	if in == nil {
		return nil
	}
	out := new(SubContactSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubContactStatus) DeepCopyInto(out *SubContactStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubContactStatus.
func (in *SubContactStatus) DeepCopy() *SubContactStatus {
	if in == nil {
		return nil
	}
	out := new(SubContactStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubSubscription) DeepCopyInto(out *SubSubscription) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubSubscription.
func (in *SubSubscription) DeepCopy() *SubSubscription {
	if in == nil {
		return nil
	}
	out := new(SubSubscription)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubSubscription) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubSubscriptionList) DeepCopyInto(out *SubSubscriptionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SubSubscription, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubSubscriptionList.
func (in *SubSubscriptionList) DeepCopy() *SubSubscriptionList {
	if in == nil {
		return nil
	}
	out := new(SubSubscriptionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SubSubscriptionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubSubscriptionSpec) DeepCopyInto(out *SubSubscriptionSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(SubSubscriptionSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubSubscriptionSpec.
func (in *SubSubscriptionSpec) DeepCopy() *SubSubscriptionSpec {
	if in == nil {
		return nil
	}
	out := new(SubSubscriptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubSubscriptionSpecResource) DeepCopyInto(out *SubSubscriptionSpecResource) {
	*out = *in
	if in.Channel != nil {
		in, out := &in.Channel, &out.Channel
		*out = new(string)
		**out = **in
	}
	if in.ContactIDS != nil {
		in, out := &in.ContactIDS, &out.ContactIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.EmailStatus != nil {
		in, out := &in.EmailStatus, &out.EmailStatus
		*out = new(int64)
		**out = **in
	}
	if in.ItemName != nil {
		in, out := &in.ItemName, &out.ItemName
		*out = new(string)
		**out = **in
	}
	if in.PmsgStatus != nil {
		in, out := &in.PmsgStatus, &out.PmsgStatus
		*out = new(int64)
		**out = **in
	}
	if in.SmsStatus != nil {
		in, out := &in.SmsStatus, &out.SmsStatus
		*out = new(int64)
		**out = **in
	}
	if in.TtsStatus != nil {
		in, out := &in.TtsStatus, &out.TtsStatus
		*out = new(int64)
		**out = **in
	}
	if in.WebhookIDS != nil {
		in, out := &in.WebhookIDS, &out.WebhookIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.WebhookStatus != nil {
		in, out := &in.WebhookStatus, &out.WebhookStatus
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubSubscriptionSpecResource.
func (in *SubSubscriptionSpecResource) DeepCopy() *SubSubscriptionSpecResource {
	if in == nil {
		return nil
	}
	out := new(SubSubscriptionSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubSubscriptionStatus) DeepCopyInto(out *SubSubscriptionStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubSubscriptionStatus.
func (in *SubSubscriptionStatus) DeepCopy() *SubSubscriptionStatus {
	if in == nil {
		return nil
	}
	out := new(SubSubscriptionStatus)
	in.DeepCopyInto(out)
	return out
}