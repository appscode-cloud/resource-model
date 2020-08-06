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

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=teams,singular=team,scope=Cluster,categories={kubernetes,resource-model,appscode}
// +kubebuilder:subresource:status
type Team struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              TeamSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            TeamStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type TeamSpec struct {
	Name      string `json:"name" protobuf:"bytes,1,opt,name=name"`
	LowerName string `json:"lowerName" protobuf:"bytes,2,opt,name=lowerName"`
	UID       string `json:"uid" protobuf:"bytes,3,opt,name=uid"`
	OrgID     int64  `json:"ownerID" protobuf:"varint,4,opt,name=ownerID"`
	//+optional
	OrgName     string     `json:"orgName,omitempty" protobuf:"bytes,5,opt,name=orgName"`
	Description string     `json:"description" protobuf:"bytes,6,opt,name=description"`
	Authorize   AccessMode `json:"authorize" protobuf:"bytes,7,opt,name=authorize,casttype=AccessMode"`
	Members     struct{}   `json:"-"`
	NumMembers  int64      `json:"numMembers" protobuf:"varint,8,opt,name=numMembers"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type TeamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []Team `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}

type TeamStatus struct {
	// ObservedGeneration is the most recent generation observed for this resource. It corresponds to the
	// resource's generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,1,opt,name=observedGeneration"`
}
