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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/ess/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeScalingConfigurations implements ScalingConfigurationInterface
type FakeScalingConfigurations struct {
	Fake *FakeEssV1alpha1
	ns   string
}

var scalingconfigurationsResource = schema.GroupVersionResource{Group: "ess.alicloud.kubeform.com", Version: "v1alpha1", Resource: "scalingconfigurations"}

var scalingconfigurationsKind = schema.GroupVersionKind{Group: "ess.alicloud.kubeform.com", Version: "v1alpha1", Kind: "ScalingConfiguration"}

// Get takes name of the scalingConfiguration, and returns the corresponding scalingConfiguration object, and an error if there is any.
func (c *FakeScalingConfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ScalingConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(scalingconfigurationsResource, c.ns, name), &v1alpha1.ScalingConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScalingConfiguration), err
}

// List takes label and field selectors, and returns the list of ScalingConfigurations that match those selectors.
func (c *FakeScalingConfigurations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ScalingConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(scalingconfigurationsResource, scalingconfigurationsKind, c.ns, opts), &v1alpha1.ScalingConfigurationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ScalingConfigurationList{ListMeta: obj.(*v1alpha1.ScalingConfigurationList).ListMeta}
	for _, item := range obj.(*v1alpha1.ScalingConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested scalingConfigurations.
func (c *FakeScalingConfigurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(scalingconfigurationsResource, c.ns, opts))

}

// Create takes the representation of a scalingConfiguration and creates it.  Returns the server's representation of the scalingConfiguration, and an error, if there is any.
func (c *FakeScalingConfigurations) Create(ctx context.Context, scalingConfiguration *v1alpha1.ScalingConfiguration, opts v1.CreateOptions) (result *v1alpha1.ScalingConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(scalingconfigurationsResource, c.ns, scalingConfiguration), &v1alpha1.ScalingConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScalingConfiguration), err
}

// Update takes the representation of a scalingConfiguration and updates it. Returns the server's representation of the scalingConfiguration, and an error, if there is any.
func (c *FakeScalingConfigurations) Update(ctx context.Context, scalingConfiguration *v1alpha1.ScalingConfiguration, opts v1.UpdateOptions) (result *v1alpha1.ScalingConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(scalingconfigurationsResource, c.ns, scalingConfiguration), &v1alpha1.ScalingConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScalingConfiguration), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeScalingConfigurations) UpdateStatus(ctx context.Context, scalingConfiguration *v1alpha1.ScalingConfiguration, opts v1.UpdateOptions) (*v1alpha1.ScalingConfiguration, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(scalingconfigurationsResource, "status", c.ns, scalingConfiguration), &v1alpha1.ScalingConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScalingConfiguration), err
}

// Delete takes name of the scalingConfiguration and deletes it. Returns an error if one occurs.
func (c *FakeScalingConfigurations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(scalingconfigurationsResource, c.ns, name), &v1alpha1.ScalingConfiguration{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeScalingConfigurations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(scalingconfigurationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ScalingConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched scalingConfiguration.
func (c *FakeScalingConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ScalingConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(scalingconfigurationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ScalingConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScalingConfiguration), err
}
