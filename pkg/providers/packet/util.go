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

package packet

import (
	"strconv"
	"strings"

	"go.bytebuilders.dev/resource-model/apis/cloud"
	"go.bytebuilders.dev/resource-model/apis/cloud/v1alpha1"
	"go.bytebuilders.dev/resource-model/pkg/util"

	"github.com/packethost/packngo"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	NumOfCore = map[string]int{
		"Intel Atom C2550 @ 2.4Ghz":                           4,
		"Intel E3-1240 v3":                                    4,
		"Intel Xeon E3-1578L v5":                              4,
		"Intel Xeon E5-2650 v4 @2.2GHz":                       24,
		"Cavium ThunderX CN8890 @2GHz":                        96,
		"Intel E5-2640 v3":                                    16,
		"Intel Xeon D-1537 @1.7GHz":                           16,
		"AMD EPYC 7401P 24-Core Processor @ 2.0GHz":           24,
		"Intel Xeon Gold 6126":                                2 * 12,
		"Intel Scalable Gold 5120 28-Core Processor @ 2.2GHz": 28,
		"32-core @ 3.0Ghz":                                    32,
	}
)

func ParseFacility(facility *packngo.Facility) *v1alpha1.Region {
	return &v1alpha1.Region{
		Region:   facility.Name,
		Location: facility.Name,
		Zones: []string{
			facility.Code,
		},
	}
}

func ParsePlan(plan *packngo.Plan) (*v1alpha1.MachineType, error) {
	ins := &v1alpha1.MachineType{
		ObjectMeta: metav1.ObjectMeta{
			Name: util.Sanitize(cloud.Packet + "-" + plan.Slug),
			Labels: map[string]string{
				cloud.KeyCloudProvider: cloud.Packet,
			},
		},
		Spec: v1alpha1.MachineTypeSpec{
			SKU:         plan.Slug,
			Description: plan.Description,
		},
	}

	mem := plan.Specs.Memory.Total
	ram, err := resource.ParseQuantity(mem[:len(mem)-1])
	if err != nil {
		return nil, err
	}
	ins.Spec.RAM = &ram

	sz := plan.Specs.Drives[0].Size
	disk, err := resource.ParseQuantity(sz[:len(sz)-1])
	if err != nil {
		return nil, err
	}
	ins.Spec.Disk = &disk

	cpu, err := GetCpuCore(plan.Specs.Cpus[0].Type)
	if err != nil {
		return nil, err
	}
	ins.Spec.CPU = resource.NewQuantity(int64(cpu), resource.DecimalExponent)
	return ins, nil
}

//formate: "/facilities/[id]"
func GetFacilityIdFromHerf(herf string) string {
	w := strings.Split(herf, "/")
	return w[len(w)-1]
}

func GetCpuCore(name string) (int, error) {
	if core, found := NumOfCore[name]; found {
		return core, nil
	} else {
		return 0, errors.Errorf("Can't find number of core for %v.", name)
	}
}

// 4GB -> 4
// 2048MB -> 2048/1024 -> 2
func RemoveUnitRetFloat64(in string) (float64, error) {
	if in[len(in)-2:] == "GB" {
		return strconv.ParseFloat(in[:len(in)-2], 64)
	} else if in[len(in)-2:] == "MB" {
		val, err := strconv.ParseInt(in[:len(in)-2], 10, 64)
		if err != nil {
			return 0, err
		}
		return util.MBToGB(val)
	} else {
		return 0, errors.Errorf("Invalid unit: %v.", in)
	}
}
func RemoveUnitRetInt(in string) (int, error) {
	val, err := RemoveUnitRetFloat64(in)
	return int(val), err
}
