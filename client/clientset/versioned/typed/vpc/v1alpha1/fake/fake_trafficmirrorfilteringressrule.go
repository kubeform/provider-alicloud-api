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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/vpc/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTrafficMirrorFilterIngressRules implements TrafficMirrorFilterIngressRuleInterface
type FakeTrafficMirrorFilterIngressRules struct {
	Fake *FakeVpcV1alpha1
	ns   string
}

var trafficmirrorfilteringressrulesResource = schema.GroupVersionResource{Group: "vpc.alicloud.kubeform.com", Version: "v1alpha1", Resource: "trafficmirrorfilteringressrules"}

var trafficmirrorfilteringressrulesKind = schema.GroupVersionKind{Group: "vpc.alicloud.kubeform.com", Version: "v1alpha1", Kind: "TrafficMirrorFilterIngressRule"}

// Get takes name of the trafficMirrorFilterIngressRule, and returns the corresponding trafficMirrorFilterIngressRule object, and an error if there is any.
func (c *FakeTrafficMirrorFilterIngressRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.TrafficMirrorFilterIngressRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(trafficmirrorfilteringressrulesResource, c.ns, name), &v1alpha1.TrafficMirrorFilterIngressRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TrafficMirrorFilterIngressRule), err
}

// List takes label and field selectors, and returns the list of TrafficMirrorFilterIngressRules that match those selectors.
func (c *FakeTrafficMirrorFilterIngressRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.TrafficMirrorFilterIngressRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(trafficmirrorfilteringressrulesResource, trafficmirrorfilteringressrulesKind, c.ns, opts), &v1alpha1.TrafficMirrorFilterIngressRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TrafficMirrorFilterIngressRuleList{ListMeta: obj.(*v1alpha1.TrafficMirrorFilterIngressRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.TrafficMirrorFilterIngressRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested trafficMirrorFilterIngressRules.
func (c *FakeTrafficMirrorFilterIngressRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(trafficmirrorfilteringressrulesResource, c.ns, opts))

}

// Create takes the representation of a trafficMirrorFilterIngressRule and creates it.  Returns the server's representation of the trafficMirrorFilterIngressRule, and an error, if there is any.
func (c *FakeTrafficMirrorFilterIngressRules) Create(ctx context.Context, trafficMirrorFilterIngressRule *v1alpha1.TrafficMirrorFilterIngressRule, opts v1.CreateOptions) (result *v1alpha1.TrafficMirrorFilterIngressRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(trafficmirrorfilteringressrulesResource, c.ns, trafficMirrorFilterIngressRule), &v1alpha1.TrafficMirrorFilterIngressRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TrafficMirrorFilterIngressRule), err
}

// Update takes the representation of a trafficMirrorFilterIngressRule and updates it. Returns the server's representation of the trafficMirrorFilterIngressRule, and an error, if there is any.
func (c *FakeTrafficMirrorFilterIngressRules) Update(ctx context.Context, trafficMirrorFilterIngressRule *v1alpha1.TrafficMirrorFilterIngressRule, opts v1.UpdateOptions) (result *v1alpha1.TrafficMirrorFilterIngressRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(trafficmirrorfilteringressrulesResource, c.ns, trafficMirrorFilterIngressRule), &v1alpha1.TrafficMirrorFilterIngressRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TrafficMirrorFilterIngressRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTrafficMirrorFilterIngressRules) UpdateStatus(ctx context.Context, trafficMirrorFilterIngressRule *v1alpha1.TrafficMirrorFilterIngressRule, opts v1.UpdateOptions) (*v1alpha1.TrafficMirrorFilterIngressRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(trafficmirrorfilteringressrulesResource, "status", c.ns, trafficMirrorFilterIngressRule), &v1alpha1.TrafficMirrorFilterIngressRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TrafficMirrorFilterIngressRule), err
}

// Delete takes name of the trafficMirrorFilterIngressRule and deletes it. Returns an error if one occurs.
func (c *FakeTrafficMirrorFilterIngressRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(trafficmirrorfilteringressrulesResource, c.ns, name), &v1alpha1.TrafficMirrorFilterIngressRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTrafficMirrorFilterIngressRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(trafficmirrorfilteringressrulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.TrafficMirrorFilterIngressRuleList{})
	return err
}

// Patch applies the patch and returns the patched trafficMirrorFilterIngressRule.
func (c *FakeTrafficMirrorFilterIngressRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.TrafficMirrorFilterIngressRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(trafficmirrorfilteringressrulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.TrafficMirrorFilterIngressRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TrafficMirrorFilterIngressRule), err
}