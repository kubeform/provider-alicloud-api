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

package v1alpha1

import (
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/cdn/v1alpha1"
	"kubeform.dev/provider-alicloud-api/client/clientset/versioned/scheme"

	rest "k8s.io/client-go/rest"
)

type CdnV1alpha1Interface interface {
	RESTClient() rest.Interface
	DomainsGetter
	DomainConfigsGetter
	DomainNewsGetter
}

// CdnV1alpha1Client is used to interact with features provided by the cdn.alicloud.kubeform.com group.
type CdnV1alpha1Client struct {
	restClient rest.Interface
}

func (c *CdnV1alpha1Client) Domains(namespace string) DomainInterface {
	return newDomains(c, namespace)
}

func (c *CdnV1alpha1Client) DomainConfigs(namespace string) DomainConfigInterface {
	return newDomainConfigs(c, namespace)
}

func (c *CdnV1alpha1Client) DomainNews(namespace string) DomainNewInterface {
	return newDomainNews(c, namespace)
}

// NewForConfig creates a new CdnV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*CdnV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &CdnV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new CdnV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *CdnV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new CdnV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *CdnV1alpha1Client {
	return &CdnV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *CdnV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
