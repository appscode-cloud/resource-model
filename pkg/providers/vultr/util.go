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

package vultr

import (
	"strconv"

	"go.bytebuilders.dev/resource-model/apis/cloud"
	"go.bytebuilders.dev/resource-model/apis/cloud/v1alpha1"
	"go.bytebuilders.dev/resource-model/pkg/util"

	vultr "github.com/JamesClonk/vultr/lib"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ParseRegion(in *vultr.Region) *v1alpha1.Region {
	return &v1alpha1.Region{
		Location: in.Name,
		Region:   strconv.Itoa(in.ID),
		Zones: []string{
			strconv.Itoa(in.ID),
		},
	}
}

func ParseInstance(in *PlanExtended) (*v1alpha1.MachineType, error) {
	out := &v1alpha1.MachineType{
		ObjectMeta: metav1.ObjectMeta{
			Name: util.Sanitize(cloud.Vultr + "-" + strconv.Itoa(in.ID)),
			Labels: map[string]string{
				cloud.KeyCloudProvider: cloud.Vultr,
			},
		},
		Spec: v1alpha1.MachineTypeSpec{
			SKU:         strconv.Itoa(in.ID),
			Description: in.Name,
			CPU:         resource.NewQuantity(int64(in.VCpus), resource.DecimalExponent),
			Category:    in.Category,
		},
	}
	if in.Deprecated {
		out.Spec.Deprecated = in.Deprecated
	}

	disk, err := resource.ParseQuantity(in.Disk + "G")
	if err != nil {
		return nil, errors.Errorf("Parse Instance failed.reason: %v.", err)
	}
	out.Spec.Disk = &disk

	ram, err := resource.ParseQuantity(in.RAM + "M")
	if err != nil {
		return nil, errors.Errorf("Parse Instance failed.reason: %v.", err)
	}
	out.Spec.RAM = &ram

	out.Spec.Zones = []string{}
	for _, r := range in.Regions {
		region := strconv.Itoa(r)
		out.Spec.Zones = append(out.Spec.Zones, region)
	}
	return out, nil
}
