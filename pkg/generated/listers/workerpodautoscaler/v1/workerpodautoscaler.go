/*
Copyright 2019 Practo Authors.

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

package v1

import (
	v1 "github.com/practo/k8s-worker-pod-autoscaler/pkg/apis/workerpodautoscaler/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// WorkerPodAutoScalerLister helps list WorkerPodAutoScalers.
// All objects returned here must be treated as read-only.
type WorkerPodAutoScalerLister interface {
	// List lists all WorkerPodAutoScalers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.WorkerPodAutoScaler, err error)
	// WorkerPodAutoScalers returns an object that can list and get WorkerPodAutoScalers.
	WorkerPodAutoScalers(namespace string) WorkerPodAutoScalerNamespaceLister
	WorkerPodAutoScalerListerExpansion
}

// workerPodAutoScalerLister implements the WorkerPodAutoScalerLister interface.
type workerPodAutoScalerLister struct {
	indexer cache.Indexer
}

// NewWorkerPodAutoScalerLister returns a new WorkerPodAutoScalerLister.
func NewWorkerPodAutoScalerLister(indexer cache.Indexer) WorkerPodAutoScalerLister {
	return &workerPodAutoScalerLister{indexer: indexer}
}

// List lists all WorkerPodAutoScalers in the indexer.
func (s *workerPodAutoScalerLister) List(selector labels.Selector) (ret []*v1.WorkerPodAutoScaler, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.WorkerPodAutoScaler))
	})
	return ret, err
}

// WorkerPodAutoScalers returns an object that can list and get WorkerPodAutoScalers.
func (s *workerPodAutoScalerLister) WorkerPodAutoScalers(namespace string) WorkerPodAutoScalerNamespaceLister {
	return workerPodAutoScalerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WorkerPodAutoScalerNamespaceLister helps list and get WorkerPodAutoScalers.
// All objects returned here must be treated as read-only.
type WorkerPodAutoScalerNamespaceLister interface {
	// List lists all WorkerPodAutoScalers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.WorkerPodAutoScaler, err error)
	// Get retrieves the WorkerPodAutoScaler from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.WorkerPodAutoScaler, error)
	WorkerPodAutoScalerNamespaceListerExpansion
}

// workerPodAutoScalerNamespaceLister implements the WorkerPodAutoScalerNamespaceLister
// interface.
type workerPodAutoScalerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all WorkerPodAutoScalers in the indexer for a given namespace.
func (s workerPodAutoScalerNamespaceLister) List(selector labels.Selector) (ret []*v1.WorkerPodAutoScaler, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.WorkerPodAutoScaler))
	})
	return ret, err
}

// Get retrieves the WorkerPodAutoScaler from the indexer for a given namespace and name.
func (s workerPodAutoScalerNamespaceLister) Get(name string) (*v1.WorkerPodAutoScaler, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("workerpodautoscaler"), name)
	}
	return obj.(*v1.WorkerPodAutoScaler), nil
}
