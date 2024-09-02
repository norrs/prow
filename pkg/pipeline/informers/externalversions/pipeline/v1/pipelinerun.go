/*
Copyright The Kubernetes Authors.

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

package v1

import (
	"context"
	time "time"

	pipelinev1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	versioned "sigs.k8s.io/prow/pkg/pipeline/clientset/versioned"
	internalinterfaces "sigs.k8s.io/prow/pkg/pipeline/informers/externalversions/internalinterfaces"
	v1 "sigs.k8s.io/prow/pkg/pipeline/listers/pipeline/v1"
)

// PipelineRunInformer provides access to a shared informer and lister for
// PipelineRuns.
type PipelineRunInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.PipelineRunLister
}

type pipelineRunInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPipelineRunInformer constructs a new informer for PipelineRun type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPipelineRunInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPipelineRunInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPipelineRunInformer constructs a new informer for PipelineRun type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPipelineRunInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TektonV1().PipelineRuns(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TektonV1().PipelineRuns(namespace).Watch(context.TODO(), options)
			},
		},
		&pipelinev1.PipelineRun{},
		resyncPeriod,
		indexers,
	)
}

func (f *pipelineRunInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPipelineRunInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *pipelineRunInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&pipelinev1.PipelineRun{}, f.defaultInformer)
}

func (f *pipelineRunInformer) Lister() v1.PipelineRunLister {
	return v1.NewPipelineRunLister(f.Informer().GetIndexer())
}