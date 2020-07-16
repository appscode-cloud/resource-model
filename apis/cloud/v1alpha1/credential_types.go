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
	ResourceKindCredential = "Credential"
	ResourceCredential     = "credential"
	ResourceCredentials    = "credentials"
)

// +kubebuilder:validation:Enum=Aws;Azure;AzureStorage;DigitalOcean;GoogleCloud;GoogleOAuth;Linode;Packet;Scaleway;Vultr
type CredentialType string

const (
	CredentialTypeAWS          CredentialType = "Aws"
	CredentialTypeAzure        CredentialType = "Azure"
	CredentialTypeAzureStorage CredentialType = "AzureStorage"
	CredentialTypeDigitalOcean CredentialType = "DigitalOcean"
	CredentialTypeGoogleCloud  CredentialType = "GoogleCloud"
	CredentialTypeGoogleOAuth  CredentialType = "GoogleOAuth"
	CredentialTypeLinode       CredentialType = "Linode"
	CredentialTypePacket       CredentialType = "Packet"
	CredentialTypeScaleway     CredentialType = "Scaleway"
	CredentialTypeVultr        CredentialType = "Vultr"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=credentials,singular=credential,scope=Cluster,categories={kubernetes,resource-model,appscode}
// +kubebuilder:subresource:status
type Credential struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              CredentialSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            CredentialStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type CredentialSpec struct {
	Name    string         `json:"name" protobuf:"bytes,1,opt,name=name"`
	Type    CredentialType `json:"type" protobuf:"bytes,2,opt,name=type"`
	OwnerID int64          `json:"ownerID" protobuf:"bytes,3,opt,name=ownerID"`

	//+optional
	AWS *AWSCredential `json:"aws,omitempty" protobuf:"bytes,4,opt,name=aws"`
	//+optional
	Azure *AzureCredential `json:"azure,omitempty" protobuf:"bytes,5,opt,name=azure"`
	//+optional
	AzureStorage *AzureStorageCredential `json:"azureStorage,omitempty" protobuf:"bytes,6,opt,name=azureStorage"`
	//+optional
	DigitalOcean *DigitalOceanCredential `json:"digitalocean,omitempty" protobuf:"bytes,7,opt,name=digitalocean"`
	//+optional
	GoogleCloud *GoogleCloudCredential `json:"googleCloud,omitempty" protobuf:"bytes,8,opt,name=googleCloud"`
	//+optional
	GoogleOAuth *GoogleOAuthCredential `json:"googleOAuth,omitempty" protobuf:"bytes,9,opt,name=googleOAuth"`
	//+optional
	Linode *LinodeCredential `json:"linode,omitempty" protobuf:"bytes,10,opt,name=linode"`
	//+optional
	Packet *PacketCredential `json:"packet,omitempty" protobuf:"bytes,11,opt,name=packet"`
	//+optional
	Scaleway *ScalewayCredential `json:"scaleway,omitempty" protobuf:"bytes,12,opt,name=scaleway"`
	//+optional
	Swift *SwiftCredential `json:"swift,omitempty" protobuf:"bytes,13,opt,name=swift"`
	//+optional
	Vultr *VultrCredential `json:"vultr,omitempty" protobuf:"bytes,14,opt,name=vultr"`
}

type GoogleOAuthCredential struct {
	ClientID     string `json:"clientID" protobuf:"bytes,1,opt,name=clientID"`
	ClientSecret string `json:"clientSecret" protobuf:"bytes,2,opt,name=clientSecret"`
	AccessToken  string `json:"accessToken" protobuf:"bytes,3,opt,name=accessToken"`
	// +optional
	RefreshToken string `json:"refreshToken,omitempty" protobuf:"bytes,4,opt,name=refreshToken"`
	// +optional
	Scopes []string `json:"scopes,omitempty" protobuf:"bytes,5,rep,name=scopes"`
	// +optional
	Expiry int64 `json:"expiry,omitempty" protobuf:"bytes,6,opt,name=expiry"`
}

type GoogleCloudCredential struct {
	ProjectID      string `json:"projectID" protobuf:"bytes,1,opt,name=projectID"`
	ServiceAccount string `json:"serviceAccount" protobuf:"bytes,2,opt,name=serviceAccount"`
}

type DigitalOceanCredential struct {
	Token string `json:"token" protobuf:"bytes,1,opt,name=token"`
}

type AzureCredential struct {
	TenantID       string `json:"tenantID" protobuf:"bytes,1,opt,name=tenantID"`
	SubscriptionID string `json:"subscriptionID" protobuf:"bytes,2,opt,name=subscriptionID"`
	ClientID       string `json:"clientID" protobuf:"bytes,3,opt,name=clientID"`
	ClientSecret   string `json:"clientSecret" protobuf:"bytes,4,opt,name=clientSecret"`
}

type AzureStorageCredential struct {
	Account string `json:"account" protobuf:"bytes,1,opt,name=account"`
	Key     string `json:"key" protobuf:"bytes,2,opt,name=key"`
}

type AWSCredential struct {
	AccessKeyID     string `json:"accessKeyID" protobuf:"bytes,1,opt,name=accessKeyID"`
	SecretAccessKey string `json:"secretAccessKey" protobuf:"bytes,2,opt,name=secretAccessKey"`
}

type PacketCredential struct {
	ProjectID string `json:"projectID" protobuf:"bytes,1,opt,name=projectID"`
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
	Token string `json:"token" protobuf:"bytes,1,opt,name=token"`
}

type SwiftCredential struct {
	Username      string `json:"username" protobuf:"bytes,1,opt,name=username"`
	Password      string `json:"password" protobuf:"bytes,2,opt,name=password"`
	TenantName    string `json:"tenantName,omitempty" protobuf:"bytes,3,opt,name=tenantName"`
	TenantAuthURL string `json:"tenantAuthURL,omitempty" protobuf:"bytes,4,opt,name=tenantAuthURL"`
	Domain        string `json:"domain,omitempty" protobuf:"bytes,5,opt,name=domain"`
	Region        string `json:"region,omitempty" protobuf:"bytes,6,opt,name=region"`
	TenantId      string `json:"tenantID,omitempty" protobuf:"bytes,7,opt,name=tenantID"`
	TenantDomain  string `json:"tenantDomain,omitempty" protobuf:"bytes,8,opt,name=tenantDomain"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
type CredentialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []Credential `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}

type CredentialStatus struct {
	// ObservedGeneration is the most recent generation observed for this resource. It corresponds to the
	// resource's generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,1,opt,name=observedGeneration"`
}
