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

package aws

import (
	"fmt"

	"go.bytebuilders.dev/resource-model/apis/cloud"
	"go.bytebuilders.dev/resource-model/apis/cloud/v1alpha1"
	"go.bytebuilders.dev/resource-model/pkg/util"

	"github.com/aws/aws-sdk-go/service/ec2"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ParseInstance(in *Ec2Instance) (*v1alpha1.MachineType, error) {
	out := &v1alpha1.MachineType{
		ObjectMeta: metav1.ObjectMeta{
			Name: util.Sanitize(cloud.AWS + "-" + in.InstanceType),
			Labels: map[string]string{
				cloud.KeyCloudProvider: cloud.AWS,
			},
		},
		Spec: v1alpha1.MachineTypeSpec{
			SKU:         in.InstanceType,
			Description: in.InstanceType,
			Category:    in.Family,
			CPU:         util.QuantityP(resource.MustParse(in.VCPU.String())),
			RAM:         util.QuantityP(resource.MustParse(in.Memory.String() + "Gi")),
		},
	}
	if in.Storage != nil {
		out.Spec.Disk = util.QuantityP(resource.MustParse(fmt.Sprintf("%dG", in.Storage.Size)))
	}
	return out, nil
}

func ParseRegion(in *ec2.Region) *v1alpha1.Region {
	return &v1alpha1.Region{
		Region: *in.RegionName,
	}
}
