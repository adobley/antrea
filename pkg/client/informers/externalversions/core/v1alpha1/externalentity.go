// Copyright 2020 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	corev1alpha1 "github.com/vmware-tanzu/antrea/pkg/apis/core/v1alpha1"
	versioned "github.com/vmware-tanzu/antrea/pkg/client/clientset/versioned"
	internalinterfaces "github.com/vmware-tanzu/antrea/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/vmware-tanzu/antrea/pkg/client/listers/core/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ExternalEntityInformer provides access to a shared informer and lister for
// ExternalEntities.
type ExternalEntityInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ExternalEntityLister
}

type externalEntityInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewExternalEntityInformer constructs a new informer for ExternalEntity type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewExternalEntityInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredExternalEntityInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredExternalEntityInformer constructs a new informer for ExternalEntity type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredExternalEntityInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1alpha1().ExternalEntities(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1alpha1().ExternalEntities(namespace).Watch(options)
			},
		},
		&corev1alpha1.ExternalEntity{},
		resyncPeriod,
		indexers,
	)
}

func (f *externalEntityInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredExternalEntityInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *externalEntityInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&corev1alpha1.ExternalEntity{}, f.defaultInformer)
}

func (f *externalEntityInformer) Lister() v1alpha1.ExternalEntityLister {
	return v1alpha1.NewExternalEntityLister(f.Informer().GetIndexer())
}