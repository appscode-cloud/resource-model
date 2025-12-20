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

package util

import (
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"go.bytebuilders.dev/resource-model/apis/cloud"
	"go.bytebuilders.dev/resource-model/apis/cloud/v1alpha1"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	"sigs.k8s.io/yaml"
)

func QuantityP(q resource.Quantity) *resource.Quantity {
	return &q
}

func Sanitize(s string) string {
	return strings.ReplaceAll(strings.ToLower(strings.TrimSpace(s)), "_", "-")
}

func ReadFile(name string) ([]byte, error) {
	dataBytes, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return dataBytes, nil
}

func WriteFile(filename string, bytes []byte) error {
	err := os.WriteFile(filename, bytes, 0o666)
	if err != nil {
		return errors.Errorf("failed to write `%s`. Reason: %v", filename, err)
	}
	return nil
}

func MBToGB(in int64) (float64, error) {
	gb, err := strconv.ParseFloat(strconv.FormatFloat(float64(in)/1024, 'f', 2, 64), 64)
	return gb, err
}

// getting provider data from cloud.yaml file
// data contained in [path to pharmer]/data/files/[provider]/cloud.yaml
func GetDataFormFile(provider string) (*v1alpha1.CloudProvider, error) {
	data := v1alpha1.CloudProvider{}
	dir := filepath.Join(cloud.DataDir, provider, "cloud.yaml")
	dataBytes, err := ReadFile(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return &v1alpha1.CloudProvider{}, nil
		}
		return nil, err
	}
	err = yaml.UnmarshalStrict(dataBytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func SortCloudProvider(data *v1alpha1.CloudProvider) *v1alpha1.CloudProvider {
	sort.Slice(data.Spec.Regions, func(i, j int) bool {
		return data.Spec.Regions[i].Region < data.Spec.Regions[j].Region
	})
	for index := range data.Spec.Regions {
		sort.Slice(data.Spec.Regions[index].Zones, func(i, j int) bool {
			return data.Spec.Regions[index].Zones[i] < data.Spec.Regions[index].Zones[j]
		})
	}
	sort.Slice(data.Spec.MachineTypes, func(i, j int) bool {
		return data.Spec.MachineTypes[i].Spec.SKU < data.Spec.MachineTypes[j].Spec.SKU
	})
	for index := range data.Spec.MachineTypes {
		sort.Slice(data.Spec.MachineTypes[index].Spec.Zones, func(i, j int) bool {
			return data.Spec.MachineTypes[index].Spec.Zones[i] < data.Spec.MachineTypes[index].Spec.Zones[j]
		})
	}
	return data
}
