/*
Copyright 2020 AppsCode Inc.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	uiv1alpha1 "go.bytebuilders.dev/resource-model/apis/ui/v1alpha1"
	versioned "go.bytebuilders.dev/resource-model/client/clientset/versioned"
	internalinterfaces "go.bytebuilders.dev/resource-model/client/informers/externalversions/internalinterfaces"
	v1alpha1 "go.bytebuilders.dev/resource-model/client/listers/ui/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// EditorOptionInformer provides access to a shared informer and lister for
// EditorOptions.
type EditorOptionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.EditorOptionLister
}

type editorOptionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewEditorOptionInformer constructs a new informer for EditorOption type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewEditorOptionInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredEditorOptionInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredEditorOptionInformer constructs a new informer for EditorOption type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredEditorOptionInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.UiV1alpha1().EditorOptions().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.UiV1alpha1().EditorOptions().Watch(context.TODO(), options)
			},
		},
		&uiv1alpha1.EditorOption{},
		resyncPeriod,
		indexers,
	)
}

func (f *editorOptionInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredEditorOptionInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *editorOptionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&uiv1alpha1.EditorOption{}, f.defaultInformer)
}

func (f *editorOptionInformer) Lister() v1alpha1.EditorOptionLister {
	return v1alpha1.NewEditorOptionLister(f.Informer().GetIndexer())
}