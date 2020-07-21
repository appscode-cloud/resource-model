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

// CloudProviderSpec defines the desired state of CloudProvider
type CloudProviderSpec struct {
	Regions      []Region      `json:"regions,omitempty" protobuf:"bytes,1,rep,name=regions"`
	MachineTypes []MachineType `json:"machineTypes,omitempty" protobuf:"bytes,2,rep,name=machineTypes"`
}

// CloudProvider is the Schema for the cloudproviders API

// +genclient
// +genclient:nonNamespaced
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=cloudproviders,singular=cloudprovider,scope=Cluster,categories={cloud,appscode}
type CloudProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec CloudProviderSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CloudProviderList contains a list of CloudProvider
type CloudProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []CloudProvider `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// Region defines the desired state of Region
type Region struct {
	Region   string   `json:"region" protobuf:"bytes,1,opt,name=region"`
	Zones    []string `json:"zones,omitempty" protobuf:"bytes,2,rep,name=zones"`
	Location string   `json:"location,omitempty" protobuf:"bytes,3,opt,name=location"`
}
