/*
Copyright 2024 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

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

	v1beta1 "github.com/GoogleCloudPlatform/gke-networking-api/apis/gcpfirewall/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGCPFirewalls implements GCPFirewallInterface
type FakeGCPFirewalls struct {
	Fake *FakeNetworkingV1beta1
}

var gcpfirewallsResource = v1beta1.SchemeGroupVersion.WithResource("gcpfirewalls")

var gcpfirewallsKind = v1beta1.SchemeGroupVersion.WithKind("GCPFirewall")

// Get takes name of the gCPFirewall, and returns the corresponding gCPFirewall object, and an error if there is any.
func (c *FakeGCPFirewalls) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.GCPFirewall, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(gcpfirewallsResource, name), &v1beta1.GCPFirewall{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.GCPFirewall), err
}

// List takes label and field selectors, and returns the list of GCPFirewalls that match those selectors.
func (c *FakeGCPFirewalls) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.GCPFirewallList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(gcpfirewallsResource, gcpfirewallsKind, opts), &v1beta1.GCPFirewallList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.GCPFirewallList{ListMeta: obj.(*v1beta1.GCPFirewallList).ListMeta}
	for _, item := range obj.(*v1beta1.GCPFirewallList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested gCPFirewalls.
func (c *FakeGCPFirewalls) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(gcpfirewallsResource, opts))
}

// Create takes the representation of a gCPFirewall and creates it.  Returns the server's representation of the gCPFirewall, and an error, if there is any.
func (c *FakeGCPFirewalls) Create(ctx context.Context, gCPFirewall *v1beta1.GCPFirewall, opts v1.CreateOptions) (result *v1beta1.GCPFirewall, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(gcpfirewallsResource, gCPFirewall), &v1beta1.GCPFirewall{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.GCPFirewall), err
}

// Update takes the representation of a gCPFirewall and updates it. Returns the server's representation of the gCPFirewall, and an error, if there is any.
func (c *FakeGCPFirewalls) Update(ctx context.Context, gCPFirewall *v1beta1.GCPFirewall, opts v1.UpdateOptions) (result *v1beta1.GCPFirewall, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(gcpfirewallsResource, gCPFirewall), &v1beta1.GCPFirewall{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.GCPFirewall), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGCPFirewalls) UpdateStatus(ctx context.Context, gCPFirewall *v1beta1.GCPFirewall, opts v1.UpdateOptions) (*v1beta1.GCPFirewall, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(gcpfirewallsResource, "status", gCPFirewall), &v1beta1.GCPFirewall{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.GCPFirewall), err
}

// Delete takes name of the gCPFirewall and deletes it. Returns an error if one occurs.
func (c *FakeGCPFirewalls) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(gcpfirewallsResource, name, opts), &v1beta1.GCPFirewall{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGCPFirewalls) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(gcpfirewallsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.GCPFirewallList{})
	return err
}

// Patch applies the patch and returns the patched gCPFirewall.
func (c *FakeGCPFirewalls) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.GCPFirewall, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(gcpfirewallsResource, name, pt, data, subresources...), &v1beta1.GCPFirewall{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.GCPFirewall), err
}
