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

	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/cloudauth/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFaceConfigs implements FaceConfigInterface
type FakeFaceConfigs struct {
	Fake *FakeCloudauthV1alpha1
	ns   string
}

var faceconfigsResource = schema.GroupVersionResource{Group: "cloudauth.alicloud.kubeform.com", Version: "v1alpha1", Resource: "faceconfigs"}

var faceconfigsKind = schema.GroupVersionKind{Group: "cloudauth.alicloud.kubeform.com", Version: "v1alpha1", Kind: "FaceConfig"}

// Get takes name of the faceConfig, and returns the corresponding faceConfig object, and an error if there is any.
func (c *FakeFaceConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FaceConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(faceconfigsResource, c.ns, name), &v1alpha1.FaceConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FaceConfig), err
}

// List takes label and field selectors, and returns the list of FaceConfigs that match those selectors.
func (c *FakeFaceConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FaceConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(faceconfigsResource, faceconfigsKind, c.ns, opts), &v1alpha1.FaceConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FaceConfigList{ListMeta: obj.(*v1alpha1.FaceConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.FaceConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested faceConfigs.
func (c *FakeFaceConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(faceconfigsResource, c.ns, opts))

}

// Create takes the representation of a faceConfig and creates it.  Returns the server's representation of the faceConfig, and an error, if there is any.
func (c *FakeFaceConfigs) Create(ctx context.Context, faceConfig *v1alpha1.FaceConfig, opts v1.CreateOptions) (result *v1alpha1.FaceConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(faceconfigsResource, c.ns, faceConfig), &v1alpha1.FaceConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FaceConfig), err
}

// Update takes the representation of a faceConfig and updates it. Returns the server's representation of the faceConfig, and an error, if there is any.
func (c *FakeFaceConfigs) Update(ctx context.Context, faceConfig *v1alpha1.FaceConfig, opts v1.UpdateOptions) (result *v1alpha1.FaceConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(faceconfigsResource, c.ns, faceConfig), &v1alpha1.FaceConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FaceConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFaceConfigs) UpdateStatus(ctx context.Context, faceConfig *v1alpha1.FaceConfig, opts v1.UpdateOptions) (*v1alpha1.FaceConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(faceconfigsResource, "status", c.ns, faceConfig), &v1alpha1.FaceConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FaceConfig), err
}

// Delete takes name of the faceConfig and deletes it. Returns an error if one occurs.
func (c *FakeFaceConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(faceconfigsResource, c.ns, name), &v1alpha1.FaceConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFaceConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(faceconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.FaceConfigList{})
	return err
}

// Patch applies the patch and returns the patched faceConfig.
func (c *FakeFaceConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FaceConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(faceconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.FaceConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FaceConfig), err
}
