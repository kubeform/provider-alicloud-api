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
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/oos/v1alpha1"
	"kubeform.dev/provider-alicloud-api/client/clientset/versioned/scheme"

	rest "k8s.io/client-go/rest"
)

type OosV1alpha1Interface interface {
	RESTClient() rest.Interface
	ApplicationsGetter
	ApplicationGroupsGetter
	ExecutionsGetter
	ParametersGetter
	PatchBaselinesGetter
	SecretParametersGetter
	ServiceSettingsGetter
	StateConfigurationsGetter
	TemplatesGetter
}

// OosV1alpha1Client is used to interact with features provided by the oos.alicloud.kubeform.com group.
type OosV1alpha1Client struct {
	restClient rest.Interface
}

func (c *OosV1alpha1Client) Applications(namespace string) ApplicationInterface {
	return newApplications(c, namespace)
}

func (c *OosV1alpha1Client) ApplicationGroups(namespace string) ApplicationGroupInterface {
	return newApplicationGroups(c, namespace)
}

func (c *OosV1alpha1Client) Executions(namespace string) ExecutionInterface {
	return newExecutions(c, namespace)
}

func (c *OosV1alpha1Client) Parameters(namespace string) ParameterInterface {
	return newParameters(c, namespace)
}

func (c *OosV1alpha1Client) PatchBaselines(namespace string) PatchBaselineInterface {
	return newPatchBaselines(c, namespace)
}

func (c *OosV1alpha1Client) SecretParameters(namespace string) SecretParameterInterface {
	return newSecretParameters(c, namespace)
}

func (c *OosV1alpha1Client) ServiceSettings(namespace string) ServiceSettingInterface {
	return newServiceSettings(c, namespace)
}

func (c *OosV1alpha1Client) StateConfigurations(namespace string) StateConfigurationInterface {
	return newStateConfigurations(c, namespace)
}

func (c *OosV1alpha1Client) Templates(namespace string) TemplateInterface {
	return newTemplates(c, namespace)
}

// NewForConfig creates a new OosV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*OosV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &OosV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new OosV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *OosV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new OosV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *OosV1alpha1Client {
	return &OosV1alpha1Client{c}
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
func (c *OosV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
