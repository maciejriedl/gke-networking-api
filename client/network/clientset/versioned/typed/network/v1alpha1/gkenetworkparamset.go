/*
Copyright 2024 Google LLC

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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/GoogleCloudPlatform/gke-networking-api/apis/network/v1alpha1"
	scheme "github.com/GoogleCloudPlatform/gke-networking-api/client/network/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GKENetworkParamSetsGetter has a method to return a GKENetworkParamSetInterface.
// A group's client should implement this interface.
type GKENetworkParamSetsGetter interface {
	GKENetworkParamSets() GKENetworkParamSetInterface
}

// GKENetworkParamSetInterface has methods to work with GKENetworkParamSet resources.
type GKENetworkParamSetInterface interface {
	Create(ctx context.Context, gKENetworkParamSet *v1alpha1.GKENetworkParamSet, opts v1.CreateOptions) (*v1alpha1.GKENetworkParamSet, error)
	Update(ctx context.Context, gKENetworkParamSet *v1alpha1.GKENetworkParamSet, opts v1.UpdateOptions) (*v1alpha1.GKENetworkParamSet, error)
	UpdateStatus(ctx context.Context, gKENetworkParamSet *v1alpha1.GKENetworkParamSet, opts v1.UpdateOptions) (*v1alpha1.GKENetworkParamSet, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.GKENetworkParamSet, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.GKENetworkParamSetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GKENetworkParamSet, err error)
	GKENetworkParamSetExpansion
}

// gKENetworkParamSets implements GKENetworkParamSetInterface
type gKENetworkParamSets struct {
	client rest.Interface
}

// newGKENetworkParamSets returns a GKENetworkParamSets
func newGKENetworkParamSets(c *NetworkingV1alpha1Client) *gKENetworkParamSets {
	return &gKENetworkParamSets{
		client: c.RESTClient(),
	}
}

// Get takes name of the gKENetworkParamSet, and returns the corresponding gKENetworkParamSet object, and an error if there is any.
func (c *gKENetworkParamSets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.GKENetworkParamSet, err error) {
	result = &v1alpha1.GKENetworkParamSet{}
	err = c.client.Get().
		Resource("gkenetworkparamsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GKENetworkParamSets that match those selectors.
func (c *gKENetworkParamSets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.GKENetworkParamSetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GKENetworkParamSetList{}
	err = c.client.Get().
		Resource("gkenetworkparamsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested gKENetworkParamSets.
func (c *gKENetworkParamSets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("gkenetworkparamsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a gKENetworkParamSet and creates it.  Returns the server's representation of the gKENetworkParamSet, and an error, if there is any.
func (c *gKENetworkParamSets) Create(ctx context.Context, gKENetworkParamSet *v1alpha1.GKENetworkParamSet, opts v1.CreateOptions) (result *v1alpha1.GKENetworkParamSet, err error) {
	result = &v1alpha1.GKENetworkParamSet{}
	err = c.client.Post().
		Resource("gkenetworkparamsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gKENetworkParamSet).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a gKENetworkParamSet and updates it. Returns the server's representation of the gKENetworkParamSet, and an error, if there is any.
func (c *gKENetworkParamSets) Update(ctx context.Context, gKENetworkParamSet *v1alpha1.GKENetworkParamSet, opts v1.UpdateOptions) (result *v1alpha1.GKENetworkParamSet, err error) {
	result = &v1alpha1.GKENetworkParamSet{}
	err = c.client.Put().
		Resource("gkenetworkparamsets").
		Name(gKENetworkParamSet.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gKENetworkParamSet).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *gKENetworkParamSets) UpdateStatus(ctx context.Context, gKENetworkParamSet *v1alpha1.GKENetworkParamSet, opts v1.UpdateOptions) (result *v1alpha1.GKENetworkParamSet, err error) {
	result = &v1alpha1.GKENetworkParamSet{}
	err = c.client.Put().
		Resource("gkenetworkparamsets").
		Name(gKENetworkParamSet.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gKENetworkParamSet).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the gKENetworkParamSet and deletes it. Returns an error if one occurs.
func (c *gKENetworkParamSets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("gkenetworkparamsets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *gKENetworkParamSets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("gkenetworkparamsets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched gKENetworkParamSet.
func (c *gKENetworkParamSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.GKENetworkParamSet, err error) {
	result = &v1alpha1.GKENetworkParamSet{}
	err = c.client.Patch(pt).
		Resource("gkenetworkparamsets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
