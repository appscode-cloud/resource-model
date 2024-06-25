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

package digitalocean

import (
	"fmt"
	"strings"

	"go.bytebuilders.dev/resource-model/apis/cloud"
	"go.bytebuilders.dev/resource-model/apis/cloud/v1alpha1"
	"go.bytebuilders.dev/resource-model/pkg/util"

	"github.com/digitalocean/godo"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ParseRegion(region *godo.Region) *v1alpha1.Region {
	return &v1alpha1.Region{
		Region: region.Slug,
		Zones: []string{
			region.Slug,
		},
		Location: region.Name,
	}
}

func ParseMachineType(sz *godo.Size) (*v1alpha1.MachineType, error) {
	return &v1alpha1.MachineType{
		ObjectMeta: metav1.ObjectMeta{
			Name: util.Sanitize(cloud.DigitalOcean + "-" + sz.Slug),
			Labels: map[string]string{
				cloud.KeyCloudProvider: cloud.DigitalOcean,
			},
		},
		Spec: v1alpha1.MachineTypeSpec{
			SKU:         sz.Slug,
			Description: sz.Slug,
			CPU:         resource.NewQuantity(int64(sz.Vcpus), resource.DecimalExponent),
			RAM:         util.QuantityP(resource.MustParse(fmt.Sprintf("%dM", sz.Memory))),
			Disk:        resource.NewScaledQuantity(int64(sz.Disk), 9),
			Category:    ParseCategoryFromSlug(sz.Slug),
			Zones:       sz.Regions,
			Deprecated:  !sz.Available,
		},
	}, nil
}

func ParseCategoryFromSlug(slug string) string {
	if strings.HasPrefix(slug, "m-") {
		return "High Memory"
	} else if strings.HasPrefix(slug, "c-") {
		return "High Cpu"
	} else {
		return "Standard Droplets"
	}
}
