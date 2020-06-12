/*
Copyright The Kubepack Authors.

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
	ResourceKindClusterCredential = "ClusterCredential"
	ResourceClusterCredential     = "clustercredential"
	ResourceClusterCredentials    = "clustercredentials"
)

type ProviderName string

const (
	ProviderDigitalOcean ProviderName = "digitalocean"
	ProviderAzure        ProviderName = "azure"
	ProviderAWS          ProviderName = "aws"
	ProviderGCE          ProviderName = "gce"
	ProviderPacket       ProviderName = "packet"
	ProviderVultr        ProviderName = "vultr"
	ProviderScaleway     ProviderName = "scaleway"
	ProviderLinode       ProviderName = "linode"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=clustercredentials,singular=clustercredential,scope=Cluster,categories={kubernetes,resource-model,appscode}
// +kubebuilder:subresource:status
type ClusterCredential struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              ClusterCredentialSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            ClusterCredentialStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type ClusterCredentialSpec struct {
	Name     string       `json:"name" protobuf:"bytes,1,opt,name=name"`
	Provider ProviderName `json:"provider" protobuf:"bytes,2,opt,name=provider"`

	//+optional
	GCECredential *GCECredential `json:"gceCredential,omitempty" protobuf:"bytes,3,opt,name=gceCredential"`
	//+optional
	DigitalOceanCredential *DigitalOceanCredential `json:"digitalOceanCredential,omitempty" protobuf:"bytes,4,opt,name=digitalOceanCredential"`
	//+optional
	AzureCredential *AzureCredential `json:"azureCredential,omitempty" protobuf:"bytes,5,opt,name=azureCredential"`
	//+optional
	AWSCredential *AWSCredential `json:"awsCredential,omitempty" protobuf:"bytes,6,opt,name=awsCredential"`
	//+optional
	PacketCredential *PacketCredential `json:"packetCredential,omitempty" protobuf:"bytes,7,opt,name=packetCredential"`
	//+optional
	ScalewayCredential *ScalewayCredential `json:"scalewayCredential,omitempty" protobuf:"bytes,8,opt,name=scalewayCredential"`
	//+optional
	LinodeCredential *LinodeCredential `json:"linodeCredential,omitempty" protobuf:"bytes,9,opt,name=linodeCredential"`
	//+optional
	VultrCredential *VultrCredential `json:"vultrCredential,omitempty" protobuf:"bytes,10,opt,name=vultrCredential"`
}

type GCECredential struct {
	ProjectID      string `json:"projectId" protobuf:"bytes,1,opt,name=projectId"`
	ServiceAccount string `json:"serviceAccount" protobuf:"bytes,2,opt,name=serviceAccount"`
}

type DigitalOceanCredential struct {
	PersonalAccessToken string `json:"personalAccessToken" protobuf:"bytes,1,opt,name="`
}

type AzureCredential struct {
	TenantID       string `json:"tenantId" protobuf:"bytes,1,opt,name=tenantId"`
	SubscriptionID string `json:"subscriptionId" protobuf:"bytes,2,opt,name=subscriptionId"`
	ClientID       string `json:"clientId" protobuf:"bytes,3,opt,name=clientId"`
	ClientSecret   string `json:"clientSecret" protobuf:"bytes,4,opt,name=clientSecret"`
}

type AWSCredential struct {
	AccessKeyID     string `json:"accessKeyId" protobuf:"bytes,1,opt,name=accessKeyId"`
	SecretAccessKey string `json:"secretAccessKey" protobuf:"bytes,2,opt,name=secretAccessKey"`
}

type PacketCredential struct {
	ProjectID string `json:"projectId" protobuf:"bytes,1,opt,name=projectId"`
	APIKey    string `json:"apiKey" protobuf:"bytes,2,opt,name=apiKey"`
}

type ScalewayCredential struct {
	Organization string `json:"organization" protobuf:"bytes,1,opt,name=organization"`
	Token        string `json:"token" protobuf:"bytes,2,opt,name=token"`
}

type LinodeCredential struct {
	Token string `json:"token" protobuf:"bytes,1,opt,name=token"`
}

type VultrCredential struct {
	PersonalAccessToken string `json:"personalAccessToken" protobuf:"bytes,1,opt,name=personalAccessToken"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
type ClusterCredentialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []ClusterCredential `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}

type ClusterCredentialStatus struct {
	// ObservedGeneration is the most recent generation observed for this resource. It corresponds to the
	// resource's generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,1,opt,name=observedGeneration"`
}
