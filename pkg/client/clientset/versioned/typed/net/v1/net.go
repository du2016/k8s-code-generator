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

package v1

import (
	"time"

	v1 "github.com/du2016/code-generator/pkg/apis/net/v1"
	scheme "github.com/du2016/code-generator/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NetsGetter has a method to return a NetInterface.
// A group's client should implement this interface.
type NetsGetter interface {
	Nets(namespace string) NetInterface
}

// NetInterface has methods to work with Net resources.
type NetInterface interface {
	Create(*v1.Net) (*v1.Net, error)
	Update(*v1.Net) (*v1.Net, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Net, error)
	List(opts metav1.ListOptions) (*v1.NetList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Net, err error)
	NetExpansion
}

// nets implements NetInterface
type nets struct {
	client rest.Interface
	ns     string
}

// newNets returns a Nets
func newNets(c *NetV1Client, namespace string) *nets {
	return &nets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the net, and returns the corresponding net object, and an error if there is any.
func (c *nets) Get(name string, options metav1.GetOptions) (result *v1.Net, err error) {
	result = &v1.Net{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("nets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Nets that match those selectors.
func (c *nets) List(opts metav1.ListOptions) (result *v1.NetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.NetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("nets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested nets.
func (c *nets) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("nets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a net and creates it.  Returns the server's representation of the net, and an error, if there is any.
func (c *nets) Create(net *v1.Net) (result *v1.Net, err error) {
	result = &v1.Net{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("nets").
		Body(net).
		Do().
		Into(result)
	return
}

// Update takes the representation of a net and updates it. Returns the server's representation of the net, and an error, if there is any.
func (c *nets) Update(net *v1.Net) (result *v1.Net, err error) {
	result = &v1.Net{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("nets").
		Name(net.Name).
		Body(net).
		Do().
		Into(result)
	return
}

// Delete takes name of the net and deletes it. Returns an error if one occurs.
func (c *nets) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("nets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *nets) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("nets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched net.
func (c *nets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Net, err error) {
	result = &v1.Net{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("nets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
