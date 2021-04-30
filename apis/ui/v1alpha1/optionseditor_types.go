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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindOptionsEditor = "OptionsEditor"
	ResourceOptionsEditor     = "optionseditor"
	ResourceOptionsEditors    = "optionseditors"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=optionseditors,singular=optionseditor,scope=Cluster,categories={ui,appscode}
// +kubebuilder:subresource:status
type OptionsEditor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              OptionsEditorSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            OptionsEditorStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type OptionsEditorSpec struct {
	// Resource identifies the resource to which this editor applies
	Resource ResourceID `json:"resource" protobuf:"bytes,1,opt,name=resource"`
	// ChartRef references a Helm chart with the ui.json
	ChartRef ChartRepoRef `json:"chartRef" protobuf:"bytes,2,opt,name=chartRef"`
	// Category defines the category of the application that this editor is for (eg, database)
	Category string `json:"category" protobuf:"bytes,3,opt,name=category"`
	// Provider defines the name of the provider for this application (eg, kubedb.com)
	Provider string `json:"provider" protobuf:"bytes,4,opt,name=provider"`
}

// ResourceID identifies a resource
type ResourceID struct {
	Group   string `json:"group" protobuf:"bytes,1,opt,name=group"`
	Version string `json:"version" protobuf:"bytes,2,opt,name=version"`
	// Name is the plural name of the resource to serve.  It must match the name of the CustomResourceDefinition-registration
	// too: plural.group and it must be all lowercase.
	Name string `json:"name" protobuf:"bytes,3,opt,name=name"`
	// Kind is the serialized kind of the resource.  It is normally CamelCase and singular.
	Kind  string        `json:"kind" protobuf:"bytes,4,opt,name=kind"`
	Scope ResourceScope `json:"scope" protobuf:"bytes,5,opt,name=scope,casttype=ResourceScope"`
}

// ResourceScope is an enum defining the different scopes available to a custom resource
type ResourceScope string

const (
	ClusterScoped   ResourceScope = "Cluster"
	NamespaceScoped ResourceScope = "Namespaced"
)

// ChartRepoRef references to a single version of a Chart
type ChartRepoRef struct {
	Name    string `json:"name" protobuf:"bytes,1,opt,name=name"`
	URL     string `json:"url" protobuf:"bytes,2,opt,name=url"`
	Version string `json:"version" protobuf:"bytes,3,opt,name=version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type OptionsEditorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []OptionsEditor `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}

type OptionsEditorStatus struct {
	// ObservedGeneration is the most recent generation observed for this resource. It corresponds to the
	// resource's generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,1,opt,name=observedGeneration"`
}
