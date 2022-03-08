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

// FakeKubernetesAddons implements KubernetesAddonInterface
type FakeKubernetesAddons struct {
	Fake *FakeCsV1alpha1
	ns   string
}

var kubernetesaddonsResource = schema.GroupVersionResource{Group: "cs.alicloud.kubeform.com", Version: "v1alpha1", Resource: "kubernetesaddons"}

var kubernetesaddonsKind = schema.GroupVersionKind{Group: "cs.alicloud.kubeform.com", Version: "v1alpha1", Kind: "KubernetesAddon"}

// Get takes name of the kubernetesAddon, and returns the corresponding kubernetesAddon object, and an error if there is any.
func (c *FakeKubernetesAddons) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.KubernetesAddon, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kubernetesaddonsResource, c.ns, name), &v1alpha1.KubernetesAddon{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubernetesAddon), err
}

// List takes label and field selectors, and returns the list of KubernetesAddons that match those selectors.
func (c *FakeKubernetesAddons) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KubernetesAddonList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kubernetesaddonsResource, kubernetesaddonsKind, c.ns, opts), &v1alpha1.KubernetesAddonList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KubernetesAddonList{ListMeta: obj.(*v1alpha1.KubernetesAddonList).ListMeta}
	for _, item := range obj.(*v1alpha1.KubernetesAddonList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kubernetesAddons.
func (c *FakeKubernetesAddons) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kubernetesaddonsResource, c.ns, opts))

}

// Create takes the representation of a kubernetesAddon and creates it.  Returns the server's representation of the kubernetesAddon, and an error, if there is any.
func (c *FakeKubernetesAddons) Create(ctx context.Context, kubernetesAddon *v1alpha1.KubernetesAddon, opts v1.CreateOptions) (result *v1alpha1.KubernetesAddon, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kubernetesaddonsResource, c.ns, kubernetesAddon), &v1alpha1.KubernetesAddon{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubernetesAddon), err
}

// Update takes the representation of a kubernetesAddon and updates it. Returns the server's representation of the kubernetesAddon, and an error, if there is any.
func (c *FakeKubernetesAddons) Update(ctx context.Context, kubernetesAddon *v1alpha1.KubernetesAddon, opts v1.UpdateOptions) (result *v1alpha1.KubernetesAddon, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kubernetesaddonsResource, c.ns, kubernetesAddon), &v1alpha1.KubernetesAddon{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubernetesAddon), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKubernetesAddons) UpdateStatus(ctx context.Context, kubernetesAddon *v1alpha1.KubernetesAddon, opts v1.UpdateOptions) (*v1alpha1.KubernetesAddon, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(kubernetesaddonsResource, "status", c.ns, kubernetesAddon), &v1alpha1.KubernetesAddon{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubernetesAddon), err
}

// Delete takes name of the kubernetesAddon and deletes it. Returns an error if one occurs.
func (c *FakeKubernetesAddons) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(kubernetesaddonsResource, c.ns, name), &v1alpha1.KubernetesAddon{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKubernetesAddons) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kubernetesaddonsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.KubernetesAddonList{})
	return err
}

// Patch applies the patch and returns the patched kubernetesAddon.
func (c *FakeKubernetesAddons) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KubernetesAddon, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kubernetesaddonsResource, c.ns, name, pt, data, subresources...), &v1alpha1.KubernetesAddon{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubernetesAddon), err
}
