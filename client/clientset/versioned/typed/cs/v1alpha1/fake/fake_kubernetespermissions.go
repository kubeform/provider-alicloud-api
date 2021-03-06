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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/cs/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKubernetesPermissionses implements KubernetesPermissionsInterface
type FakeKubernetesPermissionses struct {
	Fake *FakeCsV1alpha1
	ns   string
}

var kubernetespermissionsesResource = schema.GroupVersionResource{Group: "cs.alicloud.kubeform.com", Version: "v1alpha1", Resource: "kubernetespermissionses"}

var kubernetespermissionsesKind = schema.GroupVersionKind{Group: "cs.alicloud.kubeform.com", Version: "v1alpha1", Kind: "KubernetesPermissions"}

// Get takes name of the kubernetesPermissions, and returns the corresponding kubernetesPermissions object, and an error if there is any.
func (c *FakeKubernetesPermissionses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.KubernetesPermissions, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kubernetespermissionsesResource, c.ns, name), &v1alpha1.KubernetesPermissions{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubernetesPermissions), err
}

// List takes label and field selectors, and returns the list of KubernetesPermissionses that match those selectors.
func (c *FakeKubernetesPermissionses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KubernetesPermissionsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kubernetespermissionsesResource, kubernetespermissionsesKind, c.ns, opts), &v1alpha1.KubernetesPermissionsList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KubernetesPermissionsList{ListMeta: obj.(*v1alpha1.KubernetesPermissionsList).ListMeta}
	for _, item := range obj.(*v1alpha1.KubernetesPermissionsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kubernetesPermissionses.
func (c *FakeKubernetesPermissionses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kubernetespermissionsesResource, c.ns, opts))

}

// Create takes the representation of a kubernetesPermissions and creates it.  Returns the server's representation of the kubernetesPermissions, and an error, if there is any.
func (c *FakeKubernetesPermissionses) Create(ctx context.Context, kubernetesPermissions *v1alpha1.KubernetesPermissions, opts v1.CreateOptions) (result *v1alpha1.KubernetesPermissions, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kubernetespermissionsesResource, c.ns, kubernetesPermissions), &v1alpha1.KubernetesPermissions{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubernetesPermissions), err
}

// Update takes the representation of a kubernetesPermissions and updates it. Returns the server's representation of the kubernetesPermissions, and an error, if there is any.
func (c *FakeKubernetesPermissionses) Update(ctx context.Context, kubernetesPermissions *v1alpha1.KubernetesPermissions, opts v1.UpdateOptions) (result *v1alpha1.KubernetesPermissions, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kubernetespermissionsesResource, c.ns, kubernetesPermissions), &v1alpha1.KubernetesPermissions{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubernetesPermissions), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKubernetesPermissionses) UpdateStatus(ctx context.Context, kubernetesPermissions *v1alpha1.KubernetesPermissions, opts v1.UpdateOptions) (*v1alpha1.KubernetesPermissions, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(kubernetespermissionsesResource, "status", c.ns, kubernetesPermissions), &v1alpha1.KubernetesPermissions{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubernetesPermissions), err
}

// Delete takes name of the kubernetesPermissions and deletes it. Returns an error if one occurs.
func (c *FakeKubernetesPermissionses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(kubernetespermissionsesResource, c.ns, name), &v1alpha1.KubernetesPermissions{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKubernetesPermissionses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kubernetespermissionsesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.KubernetesPermissionsList{})
	return err
}

// Patch applies the patch and returns the patched kubernetesPermissions.
func (c *FakeKubernetesPermissionses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KubernetesPermissions, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kubernetespermissionsesResource, c.ns, name, pt, data, subresources...), &v1alpha1.KubernetesPermissions{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubernetesPermissions), err
}
