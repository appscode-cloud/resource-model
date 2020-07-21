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
	ResourceKindClusterInfo = "ClusterInfo"
	ResourceClusterInfo     = "clusterinfo"
	ResourceClusterInfos    = "clusterinfos"
)

// +kubebuilder:validation:Enum=AKS;DigitalOcean;EKS;GKE;Linode;Packet;Scaleway;Vultr
type ProviderName string

const (
	ProviderAKS          ProviderName = "AKS"
	ProviderDigitalOcean ProviderName = "DigitalOcean"
	ProviderEKS          ProviderName = "EKS"
	ProviderGKE          ProviderName = "GKE"
	ProviderLinode       ProviderName = "Linode"
	ProviderPacket       ProviderName = "Packet"
	ProviderScaleway     ProviderName = "Scaleway"
	ProviderVultr        ProviderName = "Vultr"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=clusterinfos,singular=clusterinfo,scope=Cluster,categories={kubernetes,resource-model,appscode}
// +kubebuilder:subresource:status
type ClusterInfo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              ClusterInfoSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            ClusterInfoStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type ClusterInfoSpec struct {
	Name    string `json:"name" protobuf:"bytes,1,opt,name=name"`
	UID     string `json:"uid" protobuf:"bytes,2,opt,name=uid"`
	OwnerID int64  `json:"ownerID" protobuf:"bytes,3,opt,name=ownerID"`
	//+optional
	ExternalID string `json:"externalID,omitempty" protobuf:"bytes,4,opt,name=externalID"`
	//+optional
	OwnerName string `json:"ownerName,omitempty" protobuf:"bytes,5,opt,name=ownerName"`

	//+optional
	Provider ProviderName `json:"provider,omitempty" protobuf:"bytes,6,opt,name=provider"`
	//+optional
	Endpoint string `json:"endpoint,omitempty" protobuf:"bytes,7,opt,name=endpoint"`
	//+optional
	Location string `json:"location,omitempty" protobuf:"bytes,8,opt,name=location"`
	//+optional
	Project string `json:"project,omitempty" protobuf:"bytes,9,opt,name=project"`
	//+optional
	InitialKubernetesVersion string `json:"initialKubernetesVersion" protobuf:"bytes,10,opt,name=initialKubernetesVersion"`
	//+optional
	InitialNodeCount int32 `json:"initialNodeCount" protobuf:"bytes,11,opt,name=initialNodeCount"`

	//+optional
	CreatedAt int64 `json:"createdAt,omitempty" protobuf:"bytes,12,opt,name=createdAt"`
}

type ClusterPhase string

const (
	ClusterPhaseConnected    ClusterPhase = "Connected"
	ClusterPhaseDisconnected ClusterPhase = "Disconnected"
	ClusterPhaseNotImported  ClusterPhase = "NotImported"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ClusterInfoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []ClusterInfo `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}

type ClusterInfoStatus struct {
	// ObservedGeneration is the most recent generation observed for this resource. It corresponds to the
	// resource's generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,1,opt,name=observedGeneration"`
	// Phase represents current status of the cluster
	// +optional
	Phase ClusterPhase `json:"phase,omitempty" protobuf:"varint,2,opt,name=phase"`
}
