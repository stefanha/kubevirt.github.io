/*
Copyright 2020 The KubeVirt Authors.

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

	v1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	scheme "kubevirt.io/client-go/generated/prometheus-operator/clientset/versioned/scheme"
)

// PodMonitorsGetter has a method to return a PodMonitorInterface.
// A group's client should implement this interface.
type PodMonitorsGetter interface {
	PodMonitors(namespace string) PodMonitorInterface
}

// PodMonitorInterface has methods to work with PodMonitor resources.
type PodMonitorInterface interface {
	Create(*v1.PodMonitor) (*v1.PodMonitor, error)
	Update(*v1.PodMonitor) (*v1.PodMonitor, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.PodMonitor, error)
	List(opts metav1.ListOptions) (*v1.PodMonitorList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.PodMonitor, err error)
	PodMonitorExpansion
}

// podMonitors implements PodMonitorInterface
type podMonitors struct {
	client rest.Interface
	ns     string
}

// newPodMonitors returns a PodMonitors
func newPodMonitors(c *MonitoringV1Client, namespace string) *podMonitors {
	return &podMonitors{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the podMonitor, and returns the corresponding podMonitor object, and an error if there is any.
func (c *podMonitors) Get(name string, options metav1.GetOptions) (result *v1.PodMonitor, err error) {
	result = &v1.PodMonitor{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("podmonitors").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PodMonitors that match those selectors.
func (c *podMonitors) List(opts metav1.ListOptions) (result *v1.PodMonitorList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.PodMonitorList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("podmonitors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested podMonitors.
func (c *podMonitors) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("podmonitors").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a podMonitor and creates it.  Returns the server's representation of the podMonitor, and an error, if there is any.
func (c *podMonitors) Create(podMonitor *v1.PodMonitor) (result *v1.PodMonitor, err error) {
	result = &v1.PodMonitor{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("podmonitors").
		Body(podMonitor).
		Do().
		Into(result)
	return
}

// Update takes the representation of a podMonitor and updates it. Returns the server's representation of the podMonitor, and an error, if there is any.
func (c *podMonitors) Update(podMonitor *v1.PodMonitor) (result *v1.PodMonitor, err error) {
	result = &v1.PodMonitor{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("podmonitors").
		Name(podMonitor.Name).
		Body(podMonitor).
		Do().
		Into(result)
	return
}

// Delete takes name of the podMonitor and deletes it. Returns an error if one occurs.
func (c *podMonitors) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("podmonitors").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *podMonitors) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("podmonitors").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched podMonitor.
func (c *podMonitors) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.PodMonitor, err error) {
	result = &v1.PodMonitor{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("podmonitors").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
