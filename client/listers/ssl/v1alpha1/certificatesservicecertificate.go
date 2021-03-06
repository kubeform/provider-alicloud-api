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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "kubeform.dev/provider-alicloud-api/apis/ssl/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CertificatesServiceCertificateLister helps list CertificatesServiceCertificates.
// All objects returned here must be treated as read-only.
type CertificatesServiceCertificateLister interface {
	// List lists all CertificatesServiceCertificates in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CertificatesServiceCertificate, err error)
	// CertificatesServiceCertificates returns an object that can list and get CertificatesServiceCertificates.
	CertificatesServiceCertificates(namespace string) CertificatesServiceCertificateNamespaceLister
	CertificatesServiceCertificateListerExpansion
}

// certificatesServiceCertificateLister implements the CertificatesServiceCertificateLister interface.
type certificatesServiceCertificateLister struct {
	indexer cache.Indexer
}

// NewCertificatesServiceCertificateLister returns a new CertificatesServiceCertificateLister.
func NewCertificatesServiceCertificateLister(indexer cache.Indexer) CertificatesServiceCertificateLister {
	return &certificatesServiceCertificateLister{indexer: indexer}
}

// List lists all CertificatesServiceCertificates in the indexer.
func (s *certificatesServiceCertificateLister) List(selector labels.Selector) (ret []*v1alpha1.CertificatesServiceCertificate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CertificatesServiceCertificate))
	})
	return ret, err
}

// CertificatesServiceCertificates returns an object that can list and get CertificatesServiceCertificates.
func (s *certificatesServiceCertificateLister) CertificatesServiceCertificates(namespace string) CertificatesServiceCertificateNamespaceLister {
	return certificatesServiceCertificateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CertificatesServiceCertificateNamespaceLister helps list and get CertificatesServiceCertificates.
// All objects returned here must be treated as read-only.
type CertificatesServiceCertificateNamespaceLister interface {
	// List lists all CertificatesServiceCertificates in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.CertificatesServiceCertificate, err error)
	// Get retrieves the CertificatesServiceCertificate from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.CertificatesServiceCertificate, error)
	CertificatesServiceCertificateNamespaceListerExpansion
}

// certificatesServiceCertificateNamespaceLister implements the CertificatesServiceCertificateNamespaceLister
// interface.
type certificatesServiceCertificateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CertificatesServiceCertificates in the indexer for a given namespace.
func (s certificatesServiceCertificateNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.CertificatesServiceCertificate, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.CertificatesServiceCertificate))
	})
	return ret, err
}

// Get retrieves the CertificatesServiceCertificate from the indexer for a given namespace and name.
func (s certificatesServiceCertificateNamespaceLister) Get(name string) (*v1alpha1.CertificatesServiceCertificate, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("certificatesservicecertificate"), name)
	}
	return obj.(*v1alpha1.CertificatesServiceCertificate), nil
}
