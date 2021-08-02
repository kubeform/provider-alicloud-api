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

// FakeEdgeKuberneteses implements EdgeKubernetesInterface
type FakeEdgeKuberneteses struct {
	Fake *FakeCsV1alpha1
	ns   string
}

var edgekubernetesesResource = schema.GroupVersionResource{Group: "cs.alicloud.kubeform.com", Version: "v1alpha1", Resource: "edgekuberneteses"}

var edgekubernetesesKind = schema.GroupVersionKind{Group: "cs.alicloud.kubeform.com", Version: "v1alpha1", Kind: "EdgeKubernetes"}

// Get takes name of the edgeKubernetes, and returns the corresponding edgeKubernetes object, and an error if there is any.
func (c *FakeEdgeKuberneteses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.EdgeKubernetes, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(edgekubernetesesResource, c.ns, name), &v1alpha1.EdgeKubernetes{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EdgeKubernetes), err
}

// List takes label and field selectors, and returns the list of EdgeKuberneteses that match those selectors.
func (c *FakeEdgeKuberneteses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EdgeKubernetesList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(edgekubernetesesResource, edgekubernetesesKind, c.ns, opts), &v1alpha1.EdgeKubernetesList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EdgeKubernetesList{ListMeta: obj.(*v1alpha1.EdgeKubernetesList).ListMeta}
	for _, item := range obj.(*v1alpha1.EdgeKubernetesList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested edgeKuberneteses.
func (c *FakeEdgeKuberneteses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(edgekubernetesesResource, c.ns, opts))

}

// Create takes the representation of a edgeKubernetes and creates it.  Returns the server's representation of the edgeKubernetes, and an error, if there is any.
func (c *FakeEdgeKuberneteses) Create(ctx context.Context, edgeKubernetes *v1alpha1.EdgeKubernetes, opts v1.CreateOptions) (result *v1alpha1.EdgeKubernetes, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(edgekubernetesesResource, c.ns, edgeKubernetes), &v1alpha1.EdgeKubernetes{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EdgeKubernetes), err
}

// Update takes the representation of a edgeKubernetes and updates it. Returns the server's representation of the edgeKubernetes, and an error, if there is any.
func (c *FakeEdgeKuberneteses) Update(ctx context.Context, edgeKubernetes *v1alpha1.EdgeKubernetes, opts v1.UpdateOptions) (result *v1alpha1.EdgeKubernetes, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(edgekubernetesesResource, c.ns, edgeKubernetes), &v1alpha1.EdgeKubernetes{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EdgeKubernetes), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEdgeKuberneteses) UpdateStatus(ctx context.Context, edgeKubernetes *v1alpha1.EdgeKubernetes, opts v1.UpdateOptions) (*v1alpha1.EdgeKubernetes, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(edgekubernetesesResource, "status", c.ns, edgeKubernetes), &v1alpha1.EdgeKubernetes{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EdgeKubernetes), err
}

// Delete takes name of the edgeKubernetes and deletes it. Returns an error if one occurs.
func (c *FakeEdgeKuberneteses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(edgekubernetesesResource, c.ns, name), &v1alpha1.EdgeKubernetes{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEdgeKuberneteses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(edgekubernetesesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.EdgeKubernetesList{})
	return err
}

// Patch applies the patch and returns the patched edgeKubernetes.
func (c *FakeEdgeKuberneteses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EdgeKubernetes, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(edgekubernetesesResource, c.ns, name, pt, data, subresources...), &v1alpha1.EdgeKubernetes{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EdgeKubernetes), err
}
