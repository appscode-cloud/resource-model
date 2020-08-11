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
	ResourceKindTeamUser = "TeamUser"
	ResourceTeamUser     = "teamuser"
	ResourceTeamUsers    = "teamusers"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=teamusers,singular=teamuser,scope=Cluster,categories={kubernetes,resource-model,appscode}
// +kubebuilder:subresource:status
type TeamUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Spec              TeamUserSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status            TeamUserStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type TeamUserSpec struct {
	UserID int64 `json:"userID" protobuf:"varint,1,opt,name=userID"`
	// +optional
	UserName string `json:"userName,omitempty" protobuf:"bytes,2,opt,name=userName"`
	// +optional
	FullName string `json:"fullName,omitempty" protobuf:"bytes,3,opt,name=fullName"`
	// +optional
	Email string `json:"email,omitempty" protobuf:"bytes,4,opt,name=email"`
	// +optional
	AvatarURL string `json:"avatarURL,omitempty" protobuf:"bytes,5,opt,name=avatarURL"`
	OrgID     int64  `json:"orgID" protobuf:"varint,6,opt,name=orgID"`
	// +optional
	OrgName string `json:"orgName,omitempty" protobuf:"bytes,7,opt,name=orgName"`
	TeamID  int64  `json:"teamID" protobuf:"varint,8,opt,name=teamID"`
	// +optional
	TeamName int64 `json:"teamName" protobuf:"varint,9,opt,name=teamName"`
	IsPublic bool  `json:"isPublic" protobuf:"varint,10,opt,name=isPublic"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type TeamUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []TeamUser `json:"items,omitempty" protobuf:"bytes,2,rep,name=items"`
}

type TeamUserStatus struct {
	// ObservedGeneration is the most recent generation observed for this resource. It corresponds to the
	// resource's generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty" protobuf:"varint,1,opt,name=observedGeneration"`
}
