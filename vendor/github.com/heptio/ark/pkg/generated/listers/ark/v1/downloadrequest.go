/*
Copyright 2017 the Heptio Ark contributors.

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

// This file was automatically generated by lister-gen

package v1

import (
	v1 "github.com/heptio/ark/pkg/apis/ark/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DownloadRequestLister helps list DownloadRequests.
type DownloadRequestLister interface {
	// List lists all DownloadRequests in the indexer.
	List(selector labels.Selector) (ret []*v1.DownloadRequest, err error)
	// DownloadRequests returns an object that can list and get DownloadRequests.
	DownloadRequests(namespace string) DownloadRequestNamespaceLister
	DownloadRequestListerExpansion
}

// downloadRequestLister implements the DownloadRequestLister interface.
type downloadRequestLister struct {
	indexer cache.Indexer
}

// NewDownloadRequestLister returns a new DownloadRequestLister.
func NewDownloadRequestLister(indexer cache.Indexer) DownloadRequestLister {
	return &downloadRequestLister{indexer: indexer}
}

// List lists all DownloadRequests in the indexer.
func (s *downloadRequestLister) List(selector labels.Selector) (ret []*v1.DownloadRequest, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.DownloadRequest))
	})
	return ret, err
}

// DownloadRequests returns an object that can list and get DownloadRequests.
func (s *downloadRequestLister) DownloadRequests(namespace string) DownloadRequestNamespaceLister {
	return downloadRequestNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DownloadRequestNamespaceLister helps list and get DownloadRequests.
type DownloadRequestNamespaceLister interface {
	// List lists all DownloadRequests in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.DownloadRequest, err error)
	// Get retrieves the DownloadRequest from the indexer for a given namespace and name.
	Get(name string) (*v1.DownloadRequest, error)
	DownloadRequestNamespaceListerExpansion
}

// downloadRequestNamespaceLister implements the DownloadRequestNamespaceLister
// interface.
type downloadRequestNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DownloadRequests in the indexer for a given namespace.
func (s downloadRequestNamespaceLister) List(selector labels.Selector) (ret []*v1.DownloadRequest, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.DownloadRequest))
	})
	return ret, err
}

// Get retrieves the DownloadRequest from the indexer for a given namespace and name.
func (s downloadRequestNamespaceLister) Get(name string) (*v1.DownloadRequest, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("downloadrequest"), name)
	}
	return obj.(*v1.DownloadRequest), nil
}
