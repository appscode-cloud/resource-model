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
	ResourceKindTeam = "Team"
	ResourceTeam     = "team"
	ResourceTeams    = "teams"
)

type AccessMode string

const (
	AccessModeNone  AccessMode = "none"
	AccessModeRead  AccessMode = "read"
	AccessModeWrite AccessMode = "write"
	AccessModeAdmin AccessMode = "admin"
	AccessModeOwner AccessMode = "owner"
)

// +kubebuilder:validation:Enum=Dev;Front-End;Back-End
type TeamTag string

const (
	TeamTagDev      TeamTag = "Dev"
	TeamTagFrontEnd TeamTag = "Front-End"
	TeamTagBackEnd  TeamTag = "Back-End"
)

const (
	OwnerTeamName = "Owners"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=teams,singular=team,scope=Cluster,categories={kubernetes,resource-model,appscode}
// +kubebuilder:subresource:status
type Team struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TeamSpec   `json:"spec,omitempty"`
	Status            TeamStatus `json:"status,omitempty"`
}

type TeamSpec struct {
	Name      string    `json:"name"`
	LowerName string    `json:"lowerName"`
	UID       string    `json:"uid"`
	Tags      []TeamTag `json:"tags"`
	OrgID     int64     `json:"ownerID"`
	//+optional
	OrgName     string     `json:"orgName,omitempty"`
	Description string     `json:"description,omitempty"`
	Authorize   AccessMode `json:"authorize"`
	NumMembers  int64      `json:"numMembers"`
	//+optional
	Members []TeamUser `json:"members,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type TeamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Team `json:"items,omitempty"`
}

type TeamStatus struct {
	// ObservedGeneration is the most recent generation observed for this resource. It corresponds to the
	// resource's generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
}
