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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/hbr/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeReplicationVaults implements ReplicationVaultInterface
type FakeReplicationVaults struct {
	Fake *FakeHbrV1alpha1
	ns   string
}

var replicationvaultsResource = schema.GroupVersionResource{Group: "hbr.alicloud.kubeform.com", Version: "v1alpha1", Resource: "replicationvaults"}

var replicationvaultsKind = schema.GroupVersionKind{Group: "hbr.alicloud.kubeform.com", Version: "v1alpha1", Kind: "ReplicationVault"}

// Get takes name of the replicationVault, and returns the corresponding replicationVault object, and an error if there is any.
func (c *FakeReplicationVaults) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ReplicationVault, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(replicationvaultsResource, c.ns, name), &v1alpha1.ReplicationVault{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReplicationVault), err
}

// List takes label and field selectors, and returns the list of ReplicationVaults that match those selectors.
func (c *FakeReplicationVaults) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ReplicationVaultList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(replicationvaultsResource, replicationvaultsKind, c.ns, opts), &v1alpha1.ReplicationVaultList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ReplicationVaultList{ListMeta: obj.(*v1alpha1.ReplicationVaultList).ListMeta}
	for _, item := range obj.(*v1alpha1.ReplicationVaultList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested replicationVaults.
func (c *FakeReplicationVaults) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(replicationvaultsResource, c.ns, opts))

}

// Create takes the representation of a replicationVault and creates it.  Returns the server's representation of the replicationVault, and an error, if there is any.
func (c *FakeReplicationVaults) Create(ctx context.Context, replicationVault *v1alpha1.ReplicationVault, opts v1.CreateOptions) (result *v1alpha1.ReplicationVault, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(replicationvaultsResource, c.ns, replicationVault), &v1alpha1.ReplicationVault{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReplicationVault), err
}

// Update takes the representation of a replicationVault and updates it. Returns the server's representation of the replicationVault, and an error, if there is any.
func (c *FakeReplicationVaults) Update(ctx context.Context, replicationVault *v1alpha1.ReplicationVault, opts v1.UpdateOptions) (result *v1alpha1.ReplicationVault, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(replicationvaultsResource, c.ns, replicationVault), &v1alpha1.ReplicationVault{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReplicationVault), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeReplicationVaults) UpdateStatus(ctx context.Context, replicationVault *v1alpha1.ReplicationVault, opts v1.UpdateOptions) (*v1alpha1.ReplicationVault, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(replicationvaultsResource, "status", c.ns, replicationVault), &v1alpha1.ReplicationVault{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReplicationVault), err
}

// Delete takes name of the replicationVault and deletes it. Returns an error if one occurs.
func (c *FakeReplicationVaults) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(replicationvaultsResource, c.ns, name), &v1alpha1.ReplicationVault{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeReplicationVaults) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(replicationvaultsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ReplicationVaultList{})
	return err
}

// Patch applies the patch and returns the patched replicationVault.
func (c *FakeReplicationVaults) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ReplicationVault, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(replicationvaultsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ReplicationVault{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReplicationVault), err
}
