// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/choerodon/choerodon-agent/pkg/apis/choerodon/v1alpha1"
	scheme "github.com/choerodon/choerodon-agent/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// C7NHelmReleasesGetter has a method to return a C7NHelmReleaseInterface.
// A group's client should implement this interface.
type C7NHelmReleasesGetter interface {
	C7NHelmReleases(namespace string) C7NHelmReleaseInterface
}

// C7NHelmReleaseInterface has methods to work with C7NHelmRelease resources.
type C7NHelmReleaseInterface interface {
	Create(*v1alpha1.C7NHelmRelease) (*v1alpha1.C7NHelmRelease, error)
	Update(*v1alpha1.C7NHelmRelease) (*v1alpha1.C7NHelmRelease, error)
	UpdateStatus(*v1alpha1.C7NHelmRelease) (*v1alpha1.C7NHelmRelease, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.C7NHelmRelease, error)
	List(opts v1.ListOptions) (*v1alpha1.C7NHelmReleaseList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.C7NHelmRelease, err error)
	C7NHelmReleaseExpansion
}

// c7NHelmReleases implements C7NHelmReleaseInterface
type c7NHelmReleases struct {
	client rest.Interface
	ns     string
}

// newC7NHelmReleases returns a C7NHelmReleases
func newC7NHelmReleases(c *ChoerodonV1alpha1Client, namespace string) *c7NHelmReleases {
	return &c7NHelmReleases{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the c7NHelmRelease, and returns the corresponding c7NHelmRelease object, and an error if there is any.
func (c *c7NHelmReleases) Get(name string, options v1.GetOptions) (result *v1alpha1.C7NHelmRelease, err error) {
	result = &v1alpha1.C7NHelmRelease{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("c7nhelmreleases").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of C7NHelmReleases that match those selectors.
func (c *c7NHelmReleases) List(opts v1.ListOptions) (result *v1alpha1.C7NHelmReleaseList, err error) {
	result = &v1alpha1.C7NHelmReleaseList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("c7nhelmreleases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested c7NHelmReleases.
func (c *c7NHelmReleases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("c7nhelmreleases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a c7NHelmRelease and creates it.  Returns the server's representation of the c7NHelmRelease, and an error, if there is any.
func (c *c7NHelmReleases) Create(c7NHelmRelease *v1alpha1.C7NHelmRelease) (result *v1alpha1.C7NHelmRelease, err error) {
	result = &v1alpha1.C7NHelmRelease{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("c7nhelmreleases").
		Body(c7NHelmRelease).
		Do().
		Into(result)
	return
}

// Update takes the representation of a c7NHelmRelease and updates it. Returns the server's representation of the c7NHelmRelease, and an error, if there is any.
func (c *c7NHelmReleases) Update(c7NHelmRelease *v1alpha1.C7NHelmRelease) (result *v1alpha1.C7NHelmRelease, err error) {
	result = &v1alpha1.C7NHelmRelease{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("c7nhelmreleases").
		Name(c7NHelmRelease.Name).
		Body(c7NHelmRelease).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *c7NHelmReleases) UpdateStatus(c7NHelmRelease *v1alpha1.C7NHelmRelease) (result *v1alpha1.C7NHelmRelease, err error) {
	result = &v1alpha1.C7NHelmRelease{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("c7nhelmreleases").
		Name(c7NHelmRelease.Name).
		SubResource("status").
		Body(c7NHelmRelease).
		Do().
		Into(result)
	return
}

// Delete takes name of the c7NHelmRelease and deletes it. Returns an error if one occurs.
func (c *c7NHelmReleases) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("c7nhelmreleases").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *c7NHelmReleases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("c7nhelmreleases").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched c7NHelmRelease.
func (c *c7NHelmReleases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.C7NHelmRelease, err error) {
	result = &v1alpha1.C7NHelmRelease{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("c7nhelmreleases").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
