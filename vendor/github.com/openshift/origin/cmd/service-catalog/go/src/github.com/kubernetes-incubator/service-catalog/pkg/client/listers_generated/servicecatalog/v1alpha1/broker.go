/*
Copyright 2017 The Kubernetes Authors.

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

package v1alpha1

import (
	v1alpha1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BrokerLister helps list Brokers.
type BrokerLister interface {
	// List lists all Brokers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Broker, err error)
	// Get retrieves the Broker from the index for a given name.
	Get(name string) (*v1alpha1.Broker, error)
	BrokerListerExpansion
}

// brokerLister implements the BrokerLister interface.
type brokerLister struct {
	indexer cache.Indexer
}

// NewBrokerLister returns a new BrokerLister.
func NewBrokerLister(indexer cache.Indexer) BrokerLister {
	return &brokerLister{indexer: indexer}
}

// List lists all Brokers in the indexer.
func (s *brokerLister) List(selector labels.Selector) (ret []*v1alpha1.Broker, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Broker))
	})
	return ret, err
}

// Get retrieves the Broker from the index for a given name.
func (s *brokerLister) Get(name string) (*v1alpha1.Broker, error) {
	key := &v1alpha1.Broker{ObjectMeta: v1.ObjectMeta{Name: name}}
	obj, exists, err := s.indexer.Get(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("broker"), name)
	}
	return obj.(*v1alpha1.Broker), nil
}