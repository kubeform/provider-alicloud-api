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
func (in *AlertContact) DeepCopyInto(out *AlertContact) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertContact.
func (in *AlertContact) DeepCopy() *AlertContact {
	if in == nil {
		return nil
	}
	out := new(AlertContact)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlertContact) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertContactGroup) DeepCopyInto(out *AlertContactGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertContactGroup.
func (in *AlertContactGroup) DeepCopy() *AlertContactGroup {
	if in == nil {
		return nil
	}
	out := new(AlertContactGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlertContactGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertContactGroupList) DeepCopyInto(out *AlertContactGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AlertContactGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertContactGroupList.
func (in *AlertContactGroupList) DeepCopy() *AlertContactGroupList {
	if in == nil {
		return nil
	}
	out := new(AlertContactGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlertContactGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertContactGroupSpec) DeepCopyInto(out *AlertContactGroupSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(AlertContactGroupSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertContactGroupSpec.
func (in *AlertContactGroupSpec) DeepCopy() *AlertContactGroupSpec {
	if in == nil {
		return nil
	}
	out := new(AlertContactGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertContactGroupSpecResource) DeepCopyInto(out *AlertContactGroupSpecResource) {
	*out = *in
	if in.AlertContactGroupName != nil {
		in, out := &in.AlertContactGroupName, &out.AlertContactGroupName
		*out = new(string)
		**out = **in
	}
	if in.ContactIDS != nil {
		in, out := &in.ContactIDS, &out.ContactIDS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertContactGroupSpecResource.
func (in *AlertContactGroupSpecResource) DeepCopy() *AlertContactGroupSpecResource {
	if in == nil {
		return nil
	}
	out := new(AlertContactGroupSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertContactGroupStatus) DeepCopyInto(out *AlertContactGroupStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertContactGroupStatus.
func (in *AlertContactGroupStatus) DeepCopy() *AlertContactGroupStatus {
	if in == nil {
		return nil
	}
	out := new(AlertContactGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertContactList) DeepCopyInto(out *AlertContactList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AlertContact, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertContactList.
func (in *AlertContactList) DeepCopy() *AlertContactList {
	if in == nil {
		return nil
	}
	out := new(AlertContactList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlertContactList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertContactSpec) DeepCopyInto(out *AlertContactSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(AlertContactSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertContactSpec.
func (in *AlertContactSpec) DeepCopy() *AlertContactSpec {
	if in == nil {
		return nil
	}
	out := new(AlertContactSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertContactSpecResource) DeepCopyInto(out *AlertContactSpecResource) {
	*out = *in
	if in.AlertContactName != nil {
		in, out := &in.AlertContactName, &out.AlertContactName
		*out = new(string)
		**out = **in
	}
	if in.DingRobotWebhookURL != nil {
		in, out := &in.DingRobotWebhookURL, &out.DingRobotWebhookURL
		*out = new(string)
		**out = **in
	}
	if in.Email != nil {
		in, out := &in.Email, &out.Email
		*out = new(string)
		**out = **in
	}
	if in.PhoneNum != nil {
		in, out := &in.PhoneNum, &out.PhoneNum
		*out = new(string)
		**out = **in
	}
	if in.SystemNoc != nil {
		in, out := &in.SystemNoc, &out.SystemNoc
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertContactSpecResource.
func (in *AlertContactSpecResource) DeepCopy() *AlertContactSpecResource {
	if in == nil {
		return nil
	}
	out := new(AlertContactSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertContactStatus) DeepCopyInto(out *AlertContactStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertContactStatus.
func (in *AlertContactStatus) DeepCopy() *AlertContactStatus {
	if in == nil {
		return nil
	}
	out := new(AlertContactStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DispatchRule) DeepCopyInto(out *DispatchRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DispatchRule.
func (in *DispatchRule) DeepCopy() *DispatchRule {
	if in == nil {
		return nil
	}
	out := new(DispatchRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DispatchRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DispatchRuleList) DeepCopyInto(out *DispatchRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DispatchRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DispatchRuleList.
func (in *DispatchRuleList) DeepCopy() *DispatchRuleList {
	if in == nil {
		return nil
	}
	out := new(DispatchRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DispatchRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DispatchRuleSpec) DeepCopyInto(out *DispatchRuleSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(DispatchRuleSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DispatchRuleSpec.
func (in *DispatchRuleSpec) DeepCopy() *DispatchRuleSpec {
	if in == nil {
		return nil
	}
	out := new(DispatchRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DispatchRuleSpecGroupRules) DeepCopyInto(out *DispatchRuleSpecGroupRules) {
	*out = *in
	if in.GroupID != nil {
		in, out := &in.GroupID, &out.GroupID
		*out = new(int64)
		**out = **in
	}
	if in.GroupInterval != nil {
		in, out := &in.GroupInterval, &out.GroupInterval
		*out = new(int64)
		**out = **in
	}
	if in.GroupWaitTime != nil {
		in, out := &in.GroupWaitTime, &out.GroupWaitTime
		*out = new(int64)
		**out = **in
	}
	if in.GroupingFields != nil {
		in, out := &in.GroupingFields, &out.GroupingFields
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.RepeatInterval != nil {
		in, out := &in.RepeatInterval, &out.RepeatInterval
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DispatchRuleSpecGroupRules.
func (in *DispatchRuleSpecGroupRules) DeepCopy() *DispatchRuleSpecGroupRules {
	if in == nil {
		return nil
	}
	out := new(DispatchRuleSpecGroupRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DispatchRuleSpecLabelMatchExpressionGrid) DeepCopyInto(out *DispatchRuleSpecLabelMatchExpressionGrid) {
	*out = *in
	if in.LabelMatchExpressionGroups != nil {
		in, out := &in.LabelMatchExpressionGroups, &out.LabelMatchExpressionGroups
		*out = make([]DispatchRuleSpecLabelMatchExpressionGridLabelMatchExpressionGroups, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DispatchRuleSpecLabelMatchExpressionGrid.
func (in *DispatchRuleSpecLabelMatchExpressionGrid) DeepCopy() *DispatchRuleSpecLabelMatchExpressionGrid {
	if in == nil {
		return nil
	}
	out := new(DispatchRuleSpecLabelMatchExpressionGrid)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DispatchRuleSpecLabelMatchExpressionGridLabelMatchExpressionGroups) DeepCopyInto(out *DispatchRuleSpecLabelMatchExpressionGridLabelMatchExpressionGroups) {
	*out = *in
	if in.LabelMatchExpressions != nil {
		in, out := &in.LabelMatchExpressions, &out.LabelMatchExpressions
		*out = make([]DispatchRuleSpecLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DispatchRuleSpecLabelMatchExpressionGridLabelMatchExpressionGroups.
func (in *DispatchRuleSpecLabelMatchExpressionGridLabelMatchExpressionGroups) DeepCopy() *DispatchRuleSpecLabelMatchExpressionGridLabelMatchExpressionGroups {
	if in == nil {
		return nil
	}
	out := new(DispatchRuleSpecLabelMatchExpressionGridLabelMatchExpressionGroups)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DispatchRuleSpecLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) DeepCopyInto(out *DispatchRuleSpecLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = new(string)
		**out = **in
	}
	if in.Operator != nil {
		in, out := &in.Operator, &out.Operator
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DispatchRuleSpecLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions.
func (in *DispatchRuleSpecLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions) DeepCopy() *DispatchRuleSpecLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions {
	if in == nil {
		return nil
	}
	out := new(DispatchRuleSpecLabelMatchExpressionGridLabelMatchExpressionGroupsLabelMatchExpressions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DispatchRuleSpecNotifyRules) DeepCopyInto(out *DispatchRuleSpecNotifyRules) {
	*out = *in
	if in.NotifyChannels != nil {
		in, out := &in.NotifyChannels, &out.NotifyChannels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NotifyObjects != nil {
		in, out := &in.NotifyObjects, &out.NotifyObjects
		*out = make([]DispatchRuleSpecNotifyRulesNotifyObjects, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DispatchRuleSpecNotifyRules.
func (in *DispatchRuleSpecNotifyRules) DeepCopy() *DispatchRuleSpecNotifyRules {
	if in == nil {
		return nil
	}
	out := new(DispatchRuleSpecNotifyRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DispatchRuleSpecNotifyRulesNotifyObjects) DeepCopyInto(out *DispatchRuleSpecNotifyRulesNotifyObjects) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.NotifyObjectID != nil {
		in, out := &in.NotifyObjectID, &out.NotifyObjectID
		*out = new(string)
		**out = **in
	}
	if in.NotifyType != nil {
		in, out := &in.NotifyType, &out.NotifyType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DispatchRuleSpecNotifyRulesNotifyObjects.
func (in *DispatchRuleSpecNotifyRulesNotifyObjects) DeepCopy() *DispatchRuleSpecNotifyRulesNotifyObjects {
	if in == nil {
		return nil
	}
	out := new(DispatchRuleSpecNotifyRulesNotifyObjects)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DispatchRuleSpecResource) DeepCopyInto(out *DispatchRuleSpecResource) {
	*out = *in
	if in.DispatchRuleName != nil {
		in, out := &in.DispatchRuleName, &out.DispatchRuleName
		*out = new(string)
		**out = **in
	}
	if in.DispatchType != nil {
		in, out := &in.DispatchType, &out.DispatchType
		*out = new(string)
		**out = **in
	}
	if in.GroupRules != nil {
		in, out := &in.GroupRules, &out.GroupRules
		*out = make([]DispatchRuleSpecGroupRules, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IsRecover != nil {
		in, out := &in.IsRecover, &out.IsRecover
		*out = new(bool)
		**out = **in
	}
	if in.LabelMatchExpressionGrid != nil {
		in, out := &in.LabelMatchExpressionGrid, &out.LabelMatchExpressionGrid
		*out = make([]DispatchRuleSpecLabelMatchExpressionGrid, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NotifyRules != nil {
		in, out := &in.NotifyRules, &out.NotifyRules
		*out = make([]DispatchRuleSpecNotifyRules, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DispatchRuleSpecResource.
func (in *DispatchRuleSpecResource) DeepCopy() *DispatchRuleSpecResource {
	if in == nil {
		return nil
	}
	out := new(DispatchRuleSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DispatchRuleStatus) DeepCopyInto(out *DispatchRuleStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DispatchRuleStatus.
func (in *DispatchRuleStatus) DeepCopy() *DispatchRuleStatus {
	if in == nil {
		return nil
	}
	out := new(DispatchRuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusAlertRule) DeepCopyInto(out *PrometheusAlertRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusAlertRule.
func (in *PrometheusAlertRule) DeepCopy() *PrometheusAlertRule {
	if in == nil {
		return nil
	}
	out := new(PrometheusAlertRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrometheusAlertRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusAlertRuleList) DeepCopyInto(out *PrometheusAlertRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PrometheusAlertRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusAlertRuleList.
func (in *PrometheusAlertRuleList) DeepCopy() *PrometheusAlertRuleList {
	if in == nil {
		return nil
	}
	out := new(PrometheusAlertRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrometheusAlertRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusAlertRuleSpec) DeepCopyInto(out *PrometheusAlertRuleSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(PrometheusAlertRuleSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusAlertRuleSpec.
func (in *PrometheusAlertRuleSpec) DeepCopy() *PrometheusAlertRuleSpec {
	if in == nil {
		return nil
	}
	out := new(PrometheusAlertRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusAlertRuleSpecAnnotations) DeepCopyInto(out *PrometheusAlertRuleSpecAnnotations) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusAlertRuleSpecAnnotations.
func (in *PrometheusAlertRuleSpecAnnotations) DeepCopy() *PrometheusAlertRuleSpecAnnotations {
	if in == nil {
		return nil
	}
	out := new(PrometheusAlertRuleSpecAnnotations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusAlertRuleSpecLabels) DeepCopyInto(out *PrometheusAlertRuleSpecLabels) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusAlertRuleSpecLabels.
func (in *PrometheusAlertRuleSpecLabels) DeepCopy() *PrometheusAlertRuleSpecLabels {
	if in == nil {
		return nil
	}
	out := new(PrometheusAlertRuleSpecLabels)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusAlertRuleSpecResource) DeepCopyInto(out *PrometheusAlertRuleSpecResource) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make([]PrometheusAlertRuleSpecAnnotations, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClusterID != nil {
		in, out := &in.ClusterID, &out.ClusterID
		*out = new(string)
		**out = **in
	}
	if in.DispatchRuleID != nil {
		in, out := &in.DispatchRuleID, &out.DispatchRuleID
		*out = new(string)
		**out = **in
	}
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(string)
		**out = **in
	}
	if in.Expression != nil {
		in, out := &in.Expression, &out.Expression
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]PrometheusAlertRuleSpecLabels, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
	if in.NotifyType != nil {
		in, out := &in.NotifyType, &out.NotifyType
		*out = new(string)
		**out = **in
	}
	if in.PrometheusAlertRuleName != nil {
		in, out := &in.PrometheusAlertRuleName, &out.PrometheusAlertRuleName
		*out = new(string)
		**out = **in
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(int64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusAlertRuleSpecResource.
func (in *PrometheusAlertRuleSpecResource) DeepCopy() *PrometheusAlertRuleSpecResource {
	if in == nil {
		return nil
	}
	out := new(PrometheusAlertRuleSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusAlertRuleStatus) DeepCopyInto(out *PrometheusAlertRuleStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusAlertRuleStatus.
func (in *PrometheusAlertRuleStatus) DeepCopy() *PrometheusAlertRuleStatus {
	if in == nil {
		return nil
	}
	out := new(PrometheusAlertRuleStatus)
	in.DeepCopyInto(out)
	return out
}
