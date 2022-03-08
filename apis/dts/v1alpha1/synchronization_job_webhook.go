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
	"fmt"
	"strings"

	base "kubeform.dev/apimachinery/api/v1alpha1"
	"kubeform.dev/apimachinery/pkg/util"

	jsoniter "github.com/json-iterator/go"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

func (r *SynchronizationJob) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:verbs=create;update;delete,path=/validate-dts-alicloud-kubeform-com-v1alpha1-synchronizationjob,mutating=false,failurePolicy=fail,groups=dts.alicloud.kubeform.com,resources=synchronizationjobs,versions=v1alpha1,name=synchronizationjob.dts.alicloud.kubeform.io,sideEffects=None,admissionReviewVersions=v1

var _ webhook.Validator = &SynchronizationJob{}

var synchronizationjobForceNewList = map[string]bool{
	"/checkpoint":                         true,
	"/data_initialization":                true,
	"/data_synchronization":               true,
	"/db_list":                            true,
	"/delay_notice":                       true,
	"/delay_phone":                        true,
	"/delay_rule_time":                    true,
	"/destination_endpoint_database_name": true,
	"/destination_endpoint_engine_name":   true,
	"/destination_endpoint_instance_id":   true,
	"/destination_endpoint_instance_type": true,
	"/destination_endpoint_ip":            true,
	"/destination_endpoint_oracle_sid":    true,
	"/destination_endpoint_port":          true,
	"/destination_endpoint_region":        true,
	"/destination_endpoint_user_name":     true,
	"/dts_instance_id":                    true,
	"/error_notice":                       true,
	"/error_phone":                        true,
	"/source_endpoint_database_name":      true,
	"/source_endpoint_engine_name":        true,
	"/source_endpoint_instance_id":        true,
	"/source_endpoint_instance_type":      true,
	"/source_endpoint_ip":                 true,
	"/source_endpoint_oracle_sid":         true,
	"/source_endpoint_owner_id":           true,
	"/source_endpoint_port":               true,
	"/source_endpoint_region":             true,
	"/source_endpoint_role":               true,
	"/source_endpoint_user_name":          true,
	"/structure_initialization":           true,
	"/synchronization_direction":          true,
}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *SynchronizationJob) ValidateCreate() error {
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *SynchronizationJob) ValidateUpdate(old runtime.Object) error {
	if r.Spec.Resource.ID == "" {
		return nil
	}
	newObj := r.Spec.Resource
	res := old.(*SynchronizationJob)
	oldObj := res.Spec.Resource

	jsnitr := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		TagKey:                 "tf",
		ValidateJsonRawMessage: true,
		TypeEncoders:           GetEncoder(),
		TypeDecoders:           GetDecoder(),
	}.Froze()

	byteNew, err := jsnitr.Marshal(newObj)
	if err != nil {
		return err
	}
	tempNew := make(map[string]interface{})
	err = jsnitr.Unmarshal(byteNew, &tempNew)
	if err != nil {
		return err
	}

	byteOld, err := jsnitr.Marshal(oldObj)
	if err != nil {
		return err
	}
	tempOld := make(map[string]interface{})
	err = jsnitr.Unmarshal(byteOld, &tempOld)
	if err != nil {
		return err
	}

	for key, _ := range synchronizationjobForceNewList {
		keySplit := strings.Split(key, "/*")
		length := len(keySplit)
		checkIfAnyDif := false
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempOld, tempOld, tempNew)
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempNew, tempOld, tempNew)

		if checkIfAnyDif && r.Spec.UpdatePolicy == base.UpdatePolicyDoNotDestroy {
			return fmt.Errorf(`synchronizationjob "%v/%v" immutable field can't be updated. To update, change spec.updatePolicy to Destroy`, r.Namespace, r.Name)
		}
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *SynchronizationJob) ValidateDelete() error {
	if r.Spec.TerminationPolicy == base.TerminationPolicyDoNotTerminate {
		return fmt.Errorf(`synchronizationjob "%v/%v" can't be terminated. To delete, change spec.terminationPolicy to Delete`, r.Namespace, r.Name)
	}
	return nil
}
