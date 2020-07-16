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
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MachineTypeSpec defines the desired state of MachineType
type MachineTypeSpec struct {
	SKU         string             `json:"sku" protobuf:"bytes,1,opt,name=sku"`
	Description string             `json:"description,omitempty" protobuf:"bytes,2,opt,name=description"`
	Category    string             `json:"category,omitempty" protobuf:"bytes,3,opt,name=category"`
	CPU         *resource.Quantity `json:"cpu" protobuf:"bytes,4,opt,name=cpu"`
	RAM         *resource.Quantity `json:"ram" protobuf:"bytes,5,opt,name=ram"`
	Disk        *resource.Quantity `json:"disk,omitempty" protobuf:"bytes,6,opt,name=disk"`
	Regions     []string           `json:"regions,omitempty" protobuf:"bytes,7,rep,name=regions"`
	Zones       []string           `json:"zones,omitempty" protobuf:"bytes,8,rep,name=zones"`
	Deprecated  bool               `json:"deprecated,omitempty" protobuf:"varint,9,opt,name=deprecated"`
}

// MachineType is the Schema for the machinetypes API

// +genclient
// +genclient:nonNamespaced
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=machinetypes,singular=machinetype,scope=Cluster,categories={cloud,appscode}
// +kubebuilder:printcolumn:name="SKU",type="string",JSONPath=".spec.sku"
// +kubebuilder:printcolumn:name="CPU",type="string",JSONPath=".spec.cpu"
// +kubebuilder:printcolumn:name="RAM",type="string",JSONPath=".spec.ram"
type MachineType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec MachineTypeSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MachineTypeList contains a list of MachineType
type MachineTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []MachineType `json:"items" protobuf:"bytes,2,rep,name=items"`
}
