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

// FakeKubernetesNodePools implements KubernetesNodePoolInterface
type FakeKubernetesNodePools struct {
	Fake *FakeCsV1alpha1
	ns   string
}

var kubernetesnodepoolsResource = schema.GroupVersionResource{Group: "cs.alicloud.kubeform.com", Version: "v1alpha1", Resource: "kubernetesnodepools"}

var kubernetesnodepoolsKind = schema.GroupVersionKind{Group: "cs.alicloud.kubeform.com", Version: "v1alpha1", Kind: "KubernetesNodePool"}

// Get takes name of the kubernetesNodePool, and returns the corresponding kubernetesNodePool object, and an error if there is any.
func (c *FakeKubernetesNodePools) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.KubernetesNodePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kubernetesnodepoolsResource, c.ns, name), &v1alpha1.KubernetesNodePool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubernetesNodePool), err
}

// List takes label and field selectors, and returns the list of KubernetesNodePools that match those selectors.
func (c *FakeKubernetesNodePools) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.KubernetesNodePoolList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kubernetesnodepoolsResource, kubernetesnodepoolsKind, c.ns, opts), &v1alpha1.KubernetesNodePoolList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.KubernetesNodePoolList{ListMeta: obj.(*v1alpha1.KubernetesNodePoolList).ListMeta}
	for _, item := range obj.(*v1alpha1.KubernetesNodePoolList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kubernetesNodePools.
func (c *FakeKubernetesNodePools) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kubernetesnodepoolsResource, c.ns, opts))

}

// Create takes the representation of a kubernetesNodePool and creates it.  Returns the server's representation of the kubernetesNodePool, and an error, if there is any.
func (c *FakeKubernetesNodePools) Create(ctx context.Context, kubernetesNodePool *v1alpha1.KubernetesNodePool, opts v1.CreateOptions) (result *v1alpha1.KubernetesNodePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kubernetesnodepoolsResource, c.ns, kubernetesNodePool), &v1alpha1.KubernetesNodePool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubernetesNodePool), err
}

// Update takes the representation of a kubernetesNodePool and updates it. Returns the server's representation of the kubernetesNodePool, and an error, if there is any.
func (c *FakeKubernetesNodePools) Update(ctx context.Context, kubernetesNodePool *v1alpha1.KubernetesNodePool, opts v1.UpdateOptions) (result *v1alpha1.KubernetesNodePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kubernetesnodepoolsResource, c.ns, kubernetesNodePool), &v1alpha1.KubernetesNodePool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubernetesNodePool), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKubernetesNodePools) UpdateStatus(ctx context.Context, kubernetesNodePool *v1alpha1.KubernetesNodePool, opts v1.UpdateOptions) (*v1alpha1.KubernetesNodePool, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(kubernetesnodepoolsResource, "status", c.ns, kubernetesNodePool), &v1alpha1.KubernetesNodePool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubernetesNodePool), err
}

// Delete takes name of the kubernetesNodePool and deletes it. Returns an error if one occurs.
func (c *FakeKubernetesNodePools) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(kubernetesnodepoolsResource, c.ns, name), &v1alpha1.KubernetesNodePool{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKubernetesNodePools) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kubernetesnodepoolsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.KubernetesNodePoolList{})
	return err
}

// Patch applies the patch and returns the patched kubernetesNodePool.
func (c *FakeKubernetesNodePools) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.KubernetesNodePool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kubernetesnodepoolsResource, c.ns, name, pt, data, subresources...), &v1alpha1.KubernetesNodePool{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.KubernetesNodePool), err
}
