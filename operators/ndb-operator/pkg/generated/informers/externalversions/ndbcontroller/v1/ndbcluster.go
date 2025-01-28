// Copyright (c) 2020, 2022, Oracle and/or its affiliates.
//
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl/

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	ndbcontrollerv1 "github.com/mysql/ndb-operator/pkg/apis/ndbcontroller/v1"
	versioned "github.com/mysql/ndb-operator/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/mysql/ndb-operator/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/mysql/ndb-operator/pkg/generated/listers/ndbcontroller/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// NdbClusterInformer provides access to a shared informer and lister for
// NdbClusters.
type NdbClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.NdbClusterLister
}

type ndbClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewNdbClusterInformer constructs a new informer for NdbCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNdbClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNdbClusterInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredNdbClusterInformer constructs a new informer for NdbCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNdbClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MysqlV1().NdbClusters(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MysqlV1().NdbClusters(namespace).Watch(context.TODO(), options)
			},
		},
		&ndbcontrollerv1.NdbCluster{},
		resyncPeriod,
		indexers,
	)
}

func (f *ndbClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNdbClusterInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *ndbClusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&ndbcontrollerv1.NdbCluster{}, f.defaultInformer)
}

func (f *ndbClusterInformer) Lister() v1.NdbClusterLister {
	return v1.NewNdbClusterLister(f.Informer().GetIndexer())
}
