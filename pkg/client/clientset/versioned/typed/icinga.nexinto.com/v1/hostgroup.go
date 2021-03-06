/*
Copyright 2018 Nexinto

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
	v1 "github.com/Nexinto/kubernetes-icinga/pkg/apis/icinga.nexinto.com/v1"
	scheme "github.com/Nexinto/kubernetes-icinga/pkg/client/clientset/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// HostGroupsGetter has a method to return a HostGroupInterface.
// A group's client should implement this interface.
type HostGroupsGetter interface {
	HostGroups(namespace string) HostGroupInterface
}

// HostGroupInterface has methods to work with HostGroup resources.
type HostGroupInterface interface {
	Create(*v1.HostGroup) (*v1.HostGroup, error)
	Update(*v1.HostGroup) (*v1.HostGroup, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.HostGroup, error)
	List(opts meta_v1.ListOptions) (*v1.HostGroupList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.HostGroup, err error)
	HostGroupExpansion
}

// hostGroups implements HostGroupInterface
type hostGroups struct {
	client rest.Interface
	ns     string
}

// newHostGroups returns a HostGroups
func newHostGroups(c *IcingaV1Client, namespace string) *hostGroups {
	return &hostGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the hostGroup, and returns the corresponding hostGroup object, and an error if there is any.
func (c *hostGroups) Get(name string, options meta_v1.GetOptions) (result *v1.HostGroup, err error) {
	result = &v1.HostGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hostgroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of HostGroups that match those selectors.
func (c *hostGroups) List(opts meta_v1.ListOptions) (result *v1.HostGroupList, err error) {
	result = &v1.HostGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("hostgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested hostGroups.
func (c *hostGroups) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("hostgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a hostGroup and creates it.  Returns the server's representation of the hostGroup, and an error, if there is any.
func (c *hostGroups) Create(hostGroup *v1.HostGroup) (result *v1.HostGroup, err error) {
	result = &v1.HostGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("hostgroups").
		Body(hostGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a hostGroup and updates it. Returns the server's representation of the hostGroup, and an error, if there is any.
func (c *hostGroups) Update(hostGroup *v1.HostGroup) (result *v1.HostGroup, err error) {
	result = &v1.HostGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("hostgroups").
		Name(hostGroup.Name).
		Body(hostGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the hostGroup and deletes it. Returns an error if one occurs.
func (c *hostGroups) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hostgroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *hostGroups) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("hostgroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched hostGroup.
func (c *hostGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.HostGroup, err error) {
	result = &v1.HostGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("hostgroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
