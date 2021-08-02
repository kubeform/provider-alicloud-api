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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/edas/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApplicationDeployments implements ApplicationDeploymentInterface
type FakeApplicationDeployments struct {
	Fake *FakeEdasV1alpha1
	ns   string
}

var applicationdeploymentsResource = schema.GroupVersionResource{Group: "edas.alicloud.kubeform.com", Version: "v1alpha1", Resource: "applicationdeployments"}

var applicationdeploymentsKind = schema.GroupVersionKind{Group: "edas.alicloud.kubeform.com", Version: "v1alpha1", Kind: "ApplicationDeployment"}

// Get takes name of the applicationDeployment, and returns the corresponding applicationDeployment object, and an error if there is any.
func (c *FakeApplicationDeployments) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ApplicationDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(applicationdeploymentsResource, c.ns, name), &v1alpha1.ApplicationDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationDeployment), err
}

// List takes label and field selectors, and returns the list of ApplicationDeployments that match those selectors.
func (c *FakeApplicationDeployments) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ApplicationDeploymentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(applicationdeploymentsResource, applicationdeploymentsKind, c.ns, opts), &v1alpha1.ApplicationDeploymentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ApplicationDeploymentList{ListMeta: obj.(*v1alpha1.ApplicationDeploymentList).ListMeta}
	for _, item := range obj.(*v1alpha1.ApplicationDeploymentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested applicationDeployments.
func (c *FakeApplicationDeployments) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(applicationdeploymentsResource, c.ns, opts))

}

// Create takes the representation of a applicationDeployment and creates it.  Returns the server's representation of the applicationDeployment, and an error, if there is any.
func (c *FakeApplicationDeployments) Create(ctx context.Context, applicationDeployment *v1alpha1.ApplicationDeployment, opts v1.CreateOptions) (result *v1alpha1.ApplicationDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(applicationdeploymentsResource, c.ns, applicationDeployment), &v1alpha1.ApplicationDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationDeployment), err
}

// Update takes the representation of a applicationDeployment and updates it. Returns the server's representation of the applicationDeployment, and an error, if there is any.
func (c *FakeApplicationDeployments) Update(ctx context.Context, applicationDeployment *v1alpha1.ApplicationDeployment, opts v1.UpdateOptions) (result *v1alpha1.ApplicationDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(applicationdeploymentsResource, c.ns, applicationDeployment), &v1alpha1.ApplicationDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationDeployment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApplicationDeployments) UpdateStatus(ctx context.Context, applicationDeployment *v1alpha1.ApplicationDeployment, opts v1.UpdateOptions) (*v1alpha1.ApplicationDeployment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(applicationdeploymentsResource, "status", c.ns, applicationDeployment), &v1alpha1.ApplicationDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationDeployment), err
}

// Delete takes name of the applicationDeployment and deletes it. Returns an error if one occurs.
func (c *FakeApplicationDeployments) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(applicationdeploymentsResource, c.ns, name), &v1alpha1.ApplicationDeployment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApplicationDeployments) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(applicationdeploymentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ApplicationDeploymentList{})
	return err
}

// Patch applies the patch and returns the patched applicationDeployment.
func (c *FakeApplicationDeployments) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ApplicationDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(applicationdeploymentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ApplicationDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ApplicationDeployment), err
}
