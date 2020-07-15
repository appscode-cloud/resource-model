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
	"strconv"

	"go.bytebuilders.dev/resource-model/apis/cluster"
	"go.bytebuilders.dev/resource-model/crds"

	"k8s.io/apimachinery/pkg/fields"
	"kmodules.xyz/client-go/apiextensions"
)

func (_ ClusterCredential) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	return crds.MustCustomResourceDefinition(SchemeGroupVersion.WithResource(ResourceClusterCredentials))
}

func (cred *ClusterCredential) SetLabels(resourceName, provider string, ownerID int64) {
	labelMap := map[string]string{
		cluster.LabelResourceName:              resourceName,
		cluster.LabelClusterCredentialOwnerID:  strconv.FormatInt(ownerID, 10),
		cluster.LabelClusterCredentialProvider: provider,
	}
	cred.ObjectMeta.SetLabels(labelMap)
}

func (_ ClusterCredential) FormatLabels(resourceName, provider string, ownerID int64) string {
	labelMap := make(map[string]string)
	if resourceName != "" {
		labelMap[cluster.LabelResourceName] = resourceName
	}
	if ownerID != 0 {
		labelMap[cluster.LabelClusterCredentialOwnerID] = strconv.FormatInt(ownerID, 10)
	}
	if provider != "" {
		labelMap[cluster.LabelClusterCredentialProvider] = provider
	}

	return fields.SelectorFromSet(labelMap).String()
}
