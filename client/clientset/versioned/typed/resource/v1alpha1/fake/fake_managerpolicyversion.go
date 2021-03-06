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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/resource/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeManagerPolicyVersions implements ManagerPolicyVersionInterface
type FakeManagerPolicyVersions struct {
	Fake *FakeResourceV1alpha1
	ns   string
}

var managerpolicyversionsResource = schema.GroupVersionResource{Group: "resource.alicloud.kubeform.com", Version: "v1alpha1", Resource: "managerpolicyversions"}

var managerpolicyversionsKind = schema.GroupVersionKind{Group: "resource.alicloud.kubeform.com", Version: "v1alpha1", Kind: "ManagerPolicyVersion"}

// Get takes name of the managerPolicyVersion, and returns the corresponding managerPolicyVersion object, and an error if there is any.
func (c *FakeManagerPolicyVersions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ManagerPolicyVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(managerpolicyversionsResource, c.ns, name), &v1alpha1.ManagerPolicyVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerPolicyVersion), err
}

// List takes label and field selectors, and returns the list of ManagerPolicyVersions that match those selectors.
func (c *FakeManagerPolicyVersions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ManagerPolicyVersionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(managerpolicyversionsResource, managerpolicyversionsKind, c.ns, opts), &v1alpha1.ManagerPolicyVersionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ManagerPolicyVersionList{ListMeta: obj.(*v1alpha1.ManagerPolicyVersionList).ListMeta}
	for _, item := range obj.(*v1alpha1.ManagerPolicyVersionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested managerPolicyVersions.
func (c *FakeManagerPolicyVersions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(managerpolicyversionsResource, c.ns, opts))

}

// Create takes the representation of a managerPolicyVersion and creates it.  Returns the server's representation of the managerPolicyVersion, and an error, if there is any.
func (c *FakeManagerPolicyVersions) Create(ctx context.Context, managerPolicyVersion *v1alpha1.ManagerPolicyVersion, opts v1.CreateOptions) (result *v1alpha1.ManagerPolicyVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(managerpolicyversionsResource, c.ns, managerPolicyVersion), &v1alpha1.ManagerPolicyVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerPolicyVersion), err
}

// Update takes the representation of a managerPolicyVersion and updates it. Returns the server's representation of the managerPolicyVersion, and an error, if there is any.
func (c *FakeManagerPolicyVersions) Update(ctx context.Context, managerPolicyVersion *v1alpha1.ManagerPolicyVersion, opts v1.UpdateOptions) (result *v1alpha1.ManagerPolicyVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(managerpolicyversionsResource, c.ns, managerPolicyVersion), &v1alpha1.ManagerPolicyVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerPolicyVersion), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeManagerPolicyVersions) UpdateStatus(ctx context.Context, managerPolicyVersion *v1alpha1.ManagerPolicyVersion, opts v1.UpdateOptions) (*v1alpha1.ManagerPolicyVersion, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(managerpolicyversionsResource, "status", c.ns, managerPolicyVersion), &v1alpha1.ManagerPolicyVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerPolicyVersion), err
}

// Delete takes name of the managerPolicyVersion and deletes it. Returns an error if one occurs.
func (c *FakeManagerPolicyVersions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(managerpolicyversionsResource, c.ns, name), &v1alpha1.ManagerPolicyVersion{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeManagerPolicyVersions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(managerpolicyversionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ManagerPolicyVersionList{})
	return err
}

// Patch applies the patch and returns the patched managerPolicyVersion.
func (c *FakeManagerPolicyVersions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ManagerPolicyVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(managerpolicyversionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ManagerPolicyVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ManagerPolicyVersion), err
}
