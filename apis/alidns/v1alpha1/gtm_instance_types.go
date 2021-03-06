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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type GtmInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GtmInstanceSpec   `json:"spec,omitempty"`
	Status            GtmInstanceStatus `json:"status,omitempty"`
}

type GtmInstanceSpecAlertConfig struct {
	// +optional
	DingtalkNotice *bool `json:"dingtalkNotice,omitempty" tf:"dingtalk_notice"`
	// +optional
	EmailNotice *bool `json:"emailNotice,omitempty" tf:"email_notice"`
	// +optional
	NoticeType *string `json:"noticeType,omitempty" tf:"notice_type"`
	// +optional
	SmsNotice *bool `json:"smsNotice,omitempty" tf:"sms_notice"`
}

type GtmInstanceSpec struct {
	State *GtmInstanceSpecResource `json:"state,omitempty" tf:"-"`

	Resource GtmInstanceSpecResource `json:"resource" tf:"resource"`

	UpdatePolicy base.UpdatePolicy `json:"updatePolicy,omitempty" tf:"-"`

	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendRef *core.LocalObjectReference `json:"backendRef,omitempty" tf:"-"`
}

type GtmInstanceSpecResource struct {
	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AlertConfig []GtmInstanceSpecAlertConfig `json:"alertConfig,omitempty" tf:"alert_config"`
	// +optional
	AlertGroup []string `json:"alertGroup,omitempty" tf:"alert_group"`
	// +optional
	CnameType *string `json:"cnameType,omitempty" tf:"cname_type"`
	// +optional
	ForceUpdate          *bool   `json:"forceUpdate,omitempty" tf:"force_update"`
	HealthCheckTaskCount *int64  `json:"healthCheckTaskCount" tf:"health_check_task_count"`
	InstanceName         *string `json:"instanceName" tf:"instance_name"`
	// +optional
	Lang           *string `json:"lang,omitempty" tf:"lang"`
	PackageEdition *string `json:"packageEdition" tf:"package_edition"`
	PaymentType    *string `json:"paymentType" tf:"payment_type"`
	Period         *int64  `json:"period" tf:"period"`
	// +optional
	PublicCnameMode *string `json:"publicCnameMode,omitempty" tf:"public_cname_mode"`
	// +optional
	PublicRr *string `json:"publicRr,omitempty" tf:"public_rr"`
	// +optional
	PublicUserDomainName *string `json:"publicUserDomainName,omitempty" tf:"public_user_domain_name"`
	// +optional
	PublicZoneName *string `json:"publicZoneName,omitempty" tf:"public_zone_name"`
	// +optional
	RenewPeriod *int64 `json:"renewPeriod,omitempty" tf:"renew_period"`
	// +optional
	RenewalStatus *string `json:"renewalStatus,omitempty" tf:"renewal_status"`
	// +optional
	ResourceGroupID      *string `json:"resourceGroupID,omitempty" tf:"resource_group_id"`
	SmsNotificationCount *int64  `json:"smsNotificationCount" tf:"sms_notification_count"`
	// +optional
	StrategyMode *string `json:"strategyMode,omitempty" tf:"strategy_mode"`
	// +optional
	Ttl *int64 `json:"ttl,omitempty" tf:"ttl"`
}

type GtmInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase status.Status `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GtmInstanceList is a list of GtmInstances
type GtmInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GtmInstance CRD objects
	Items []GtmInstance `json:"items,omitempty"`
}
