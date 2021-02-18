/*
Copyright The Kubernetes Authors.

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

package v1alpha

import (
	"context"
	"time"

	v1alpha "github.com/EdgeNet-project/edgenet/pkg/apis/apps/v1alpha"
	scheme "github.com/EdgeNet-project/edgenet/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SlicesGetter has a method to return a SliceInterface.
// A group's client should implement this interface.
type SlicesGetter interface {
	Slices(namespace string) SliceInterface
}

// SliceInterface has methods to work with Slice resources.
type SliceInterface interface {
	Create(ctx context.Context, slice *v1alpha.Slice, opts v1.CreateOptions) (*v1alpha.Slice, error)
	Update(ctx context.Context, slice *v1alpha.Slice, opts v1.UpdateOptions) (*v1alpha.Slice, error)
	UpdateStatus(ctx context.Context, slice *v1alpha.Slice, opts v1.UpdateOptions) (*v1alpha.Slice, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha.Slice, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha.SliceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha.Slice, err error)
	SliceExpansion
}

// slices implements SliceInterface
type slices struct {
	client rest.Interface
	ns     string
}

// newSlices returns a Slices
func newSlices(c *AppsV1alphaClient, namespace string) *slices {
	return &slices{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the slice, and returns the corresponding slice object, and an error if there is any.
func (c *slices) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha.Slice, err error) {
	result = &v1alpha.Slice{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("slices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Slices that match those selectors.
func (c *slices) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha.SliceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha.SliceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("slices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested slices.
func (c *slices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("slices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a slice and creates it.  Returns the server's representation of the slice, and an error, if there is any.
func (c *slices) Create(ctx context.Context, slice *v1alpha.Slice, opts v1.CreateOptions) (result *v1alpha.Slice, err error) {
	result = &v1alpha.Slice{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("slices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(slice).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a slice and updates it. Returns the server's representation of the slice, and an error, if there is any.
func (c *slices) Update(ctx context.Context, slice *v1alpha.Slice, opts v1.UpdateOptions) (result *v1alpha.Slice, err error) {
	result = &v1alpha.Slice{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("slices").
		Name(slice.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(slice).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *slices) UpdateStatus(ctx context.Context, slice *v1alpha.Slice, opts v1.UpdateOptions) (result *v1alpha.Slice, err error) {
	result = &v1alpha.Slice{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("slices").
		Name(slice.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(slice).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the slice and deletes it. Returns an error if one occurs.
func (c *slices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("slices").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *slices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("slices").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched slice.
func (c *slices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha.Slice, err error) {
	result = &v1alpha.Slice{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("slices").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
