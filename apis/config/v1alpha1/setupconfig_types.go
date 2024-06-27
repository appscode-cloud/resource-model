/*
Copyright AppsCode Inc. and Contributors

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
	cloudv1alpha1 "go.bytebuilders.dev/resource-model/apis/cloud/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AceSetupConfig is the Schema for the kubestashconfigs API

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type AceSetupConfig struct {
	metav1.TypeMeta `json:",inline"`

	DeploymentType         string     `json:"deploymentType,omitempty"`
	Nats                   NatsConfig `json:"nats"`
	ImporterServiceAccount string     `json:"importerServiceAccount,omitempty"`
	AceSetupInlineConfig   `json:",inline"`
}

// NatsConfig holds the NATS-related fields.
type NatsConfig struct {
	Exports              bool `json:"exports"`
	ReloadNatsAccounts   bool `json:"reloadNatsAccounts"`
	CreateNatsStream     bool `json:"createNatsStream,omitempty"`
	RefactorNatsAccounts bool `json:"refactorNatsAccounts,omitempty"`
	Migrate              bool `json:"migrate,omitempty"`
}

type AceSetupInlineConfig struct {
	// +optional
	Admin AcePlatformAdmin `json:"admin"`
	// +optional
	SelfManagement SelfManagement `json:"selfManagement"`
	// +optional
	CloudCredential cloudv1alpha1.Credential `json:"cloudCredential"`
	// +optional
	Cluster CAPIClusterConfig `json:"cluster,omitempty"`
	// +optional
	Subscription MarketplaceSubscriptionInfo `json:"subscription,omitempty"`
}

type AcePlatformAdmin struct {
	// +optional
	Username string `json:"username"`
	// +optional
	Password string `json:"password"`
	// +optional
	Email string `json:"email"`
	// +optional
	DisplayName string `json:"displayName"`
	// +optional
	Orgname string `json:"orgname"`
}

type SelfManagement struct {
	// +optional
	Import bool `json:"import"`
	// +optional
	EnableFeatures []string `json:"enableFeatures"`
	// +optional
	DisableFeatures []string `json:"disableFeatures"`
}

type CAPIClusterConfig struct {
	ClusterName       string        `json:"clusterName,omitempty"`
	Region            string        `json:"region,omitempty"`
	NetworkCIDR       string        `json:"networkCIDR,omitempty"`
	KubernetesVersion string        `json:"kubernetesVersion,omitempty"`
	GoogleProjectID   string        `json:"googleProjectID,omitempty"`
	WorkerPools       []MachinePool `json:"workerPools,omitempty"`
}

type MachinePool struct {
	MachineType  string `json:"machineType"`
	MachineCount int    `json:"machineCount"`
}

type MarketplaceSubscriptionInfo struct {
	AWS   *AWSMarSubscriptionInfo `json:"aws,omitempty"`
	Azure *AzureSubscriptionInfo  `json:"azure,omitempty"`
	GCP   *GCPSubscriptionInfo    `json:"gcp,omitempty"`
}

type AWSMarSubscriptionInfo struct {
	CustomerIdentifier string `json:"customer-identifier"`
	ProductCode        string `json:"product-code"`
	OfferIdentifier    string `json:"offer-identifier"`
}

type AzureSubscriptionInfo struct {
	ApplicationID  string              `json:"applicationId"`
	BillingDetails AzureBillingDetails `json:"billingDetails"`
	Plan           AzurePlan           `json:"plan"`
}

type AzureBillingDetails struct {
	ResourceUsageID string `json:"resourceUsageId"`
}

type AzurePlan struct {
	Publisher string `json:"publisher"`
	Product   string `json:"product"`
	Name      string `json:"name"`
	Version   string `json:"version"`
}

type GCPSubscriptionInfo struct{}
