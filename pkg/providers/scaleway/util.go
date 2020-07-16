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

package scaleway

import (
	"go.bytebuilders.dev/resource-model/apis/cloud"
	"go.bytebuilders.dev/resource-model/apis/cloud/v1alpha1"
	"go.bytebuilders.dev/resource-model/pkg/util"

	scaleway "github.com/scaleway/scaleway-cli/pkg/api"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ParseInstance(name string, in *scaleway.ProductServer) (*v1alpha1.MachineType, error) {
	out := &v1alpha1.MachineType{
		ObjectMeta: metav1.ObjectMeta{
			Name: util.Sanitize(cloud.Packet + "-" + name),
			Labels: map[string]string{
				cloud.KeyCloudProvider: cloud.Packet,
			},
		},
		Spec: v1alpha1.MachineTypeSpec{
			SKU:         name,
			Description: in.Arch,
			CPU:         resource.NewQuantity(int64(in.Ncpus), resource.DecimalExponent),
			RAM:         resource.NewQuantity(int64(in.Ram), resource.BinarySI),
			Disk:        resource.NewQuantity(int64(in.VolumesConstraint.MinSize), resource.DecimalSI),
		},
	}
	//if in.Baremetal {
	//	out.Category = "BareMetal"
	//} else {
	//	out.Category = "Cloud Servers"
	//}
	return out, nil
}
