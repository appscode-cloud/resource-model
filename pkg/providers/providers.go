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

package providers

import (
	"os"
	"path/filepath"
	"strings"

	"go.bytebuilders.dev/resource-model/apis/cloud"
	"go.bytebuilders.dev/resource-model/apis/cloud/v1alpha1"
	"go.bytebuilders.dev/resource-model/pkg/providers/aws"
	"go.bytebuilders.dev/resource-model/pkg/providers/azure"
	"go.bytebuilders.dev/resource-model/pkg/providers/digitalocean"
	"go.bytebuilders.dev/resource-model/pkg/providers/gce"
	"go.bytebuilders.dev/resource-model/pkg/providers/linode"
	"go.bytebuilders.dev/resource-model/pkg/providers/packet"
	"go.bytebuilders.dev/resource-model/pkg/providers/scaleway"
	"go.bytebuilders.dev/resource-model/pkg/providers/vultr"
	"go.bytebuilders.dev/resource-model/pkg/util"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	mu "kmodules.xyz/client-go/meta"
)

var providers = []string{
	cloud.GCE,
	cloud.DigitalOcean,
	cloud.Packet,
	cloud.AWS,
	cloud.Azure,
	cloud.Vultr,
	cloud.Linode,
	cloud.Scaleway,
}

func List() []string {
	return append([]string(nil), providers...)
}

type Interface interface {
	GetName() string
	ListRegions() ([]v1alpha1.Region, error)
	ListZones() ([]string, error)
	ListMachineTypes() ([]v1alpha1.MachineType, error)
}

func NewCloudProvider(opts Options) (Interface, error) {
	switch opts.Provider {
	case cloud.GCE:
		return gce.NewClient(opts.GCE)
	case cloud.DigitalOcean:
		return digitalocean.NewClient(opts.Do)
	case cloud.Packet:
		return packet.NewClient(opts.Packet)
	case cloud.AWS:
		return aws.NewClient(opts.AWS)
	case cloud.Azure:
		return azure.NewClient(opts.Azure)
	case cloud.Vultr:
		return vultr.NewClient(opts.Vultr)
	case cloud.Linode:
		return linode.NewClient(opts.Linode)
	case cloud.Scaleway:
		return scaleway.NewClient(opts.Scaleway)
	}
	return nil, errors.Errorf("Unknown cloud provider: %s", opts.Provider)
}

// get data from api
func GetCloudProvider(i Interface) (*v1alpha1.CloudProvider, error) {
	var err error
	data := v1alpha1.CloudProvider{
		ObjectMeta: metav1.ObjectMeta{
			Name: i.GetName(),
		},
		Spec: v1alpha1.CloudProviderSpec{},
	}
	data.Spec.Regions, err = i.ListRegions()
	if err != nil {
		return nil, err
	}
	data.Spec.MachineTypes, err = i.ListMachineTypes()
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func WriteObject(obj runtime.Object) error {
	kind := mu.GetKind(obj)
	resource := strings.ToLower(kind) + "s"
	name, err := meta.NewAccessor().Name(obj)
	if err != nil {
		return err
	}

	yamlDir := filepath.Join(cloud.DataDir, "yaml", "apis", v1alpha1.SchemeGroupVersion.Group, v1alpha1.SchemeGroupVersion.Version, resource)
	err = os.MkdirAll(yamlDir, 0o755)
	if err != nil {
		return err
	}
	jsonDir := filepath.Join(cloud.DataDir, "json", "apis", v1alpha1.SchemeGroupVersion.Group, v1alpha1.SchemeGroupVersion.Version, resource)
	err = os.MkdirAll(jsonDir, 0o755)
	if err != nil {
		return err
	}

	yamlBytes, err := mu.MarshalToYAML(obj, v1alpha1.SchemeGroupVersion)
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath.Join(yamlDir, name+".yaml"), yamlBytes, 0o755)
	if err != nil {
		return err
	}

	jsonBytes, err := mu.MarshalToPrettyJson(obj, v1alpha1.SchemeGroupVersion)
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(jsonDir, name+".json"), jsonBytes, 0o755)
}

func WriteCloudProvider(data *v1alpha1.CloudProvider) error {
	data = util.SortCloudProvider(data)
	err := WriteObject(data)
	if err != nil {
		return err
	}

	for _, mt := range data.Spec.MachineTypes {
		err := WriteObject(&mt)
		if err != nil {
			return err
		}
	}
	return nil
}

// region merge rule:
//
//	if region doesn't exist in old data, but exists in cur data, then add it
//	if region exists in old data, but doesn't exists in cur data, then delete it
//	if region exist in both, then
//		if field data exists in both cur and old data , then take the cur data
//		otherwise, take data from (old or cur)whichever contains it
//
// instanceType merge rule: same as region rule, except
//
//	if instance exists in old data, but doesn't exists in cur data, then add it , set the deprecated true
//
// In MergeCloudProvider, we merge only the region and instanceType data
func MergeCloudProvider(oldData, curData *v1alpha1.CloudProvider) (*v1alpha1.CloudProvider, error) {
	// region merge
	regionIndex := map[string]int{} // keep regionName,corresponding region index in oldData.Regions[] as (key,value) pair
	for index, r := range oldData.Spec.Regions {
		regionIndex[r.Region] = index
	}
	for index := range curData.Spec.Regions {
		pos, found := regionIndex[curData.Spec.Regions[index].Region]
		if found {
			// location
			if curData.Spec.Regions[index].Location == "" && oldData.Spec.Regions[pos].Location != "" {
				curData.Spec.Regions[index].Location = oldData.Spec.Regions[pos].Location
			}
			// zones
			if len(curData.Spec.Regions[index].Zones) == 0 && len(oldData.Spec.Regions[pos].Zones) != 0 {
				curData.Spec.Regions[index].Location = oldData.Spec.Regions[pos].Location
			}
		}
	}

	// instanceType
	instanceIndex := map[string]int{} // keep SKU,corresponding instance index in oldData.MachineTypes[] as (key,value) pair
	for index, ins := range oldData.Spec.MachineTypes {
		instanceIndex[ins.Spec.SKU] = index
	}
	for index := range curData.Spec.MachineTypes {
		pos, found := instanceIndex[curData.Spec.MachineTypes[index].Spec.SKU]
		if found {
			// description
			if curData.Spec.MachineTypes[index].Spec.Description == "" && oldData.Spec.MachineTypes[pos].Spec.Description != "" {
				curData.Spec.MachineTypes[index].Spec.Description = oldData.Spec.MachineTypes[pos].Spec.Description
			}
			// zones
			if len(curData.Spec.MachineTypes[index].Spec.Zones) == 0 && len(oldData.Spec.MachineTypes[pos].Spec.Zones) == 0 {
				curData.Spec.MachineTypes[index].Spec.Zones = oldData.Spec.MachineTypes[pos].Spec.Zones
			}
			//regions
			//if len(curData.Spec.MachineTypes[index].Spec.Regions)==0 && len(oldData.Spec.MachineTypes[pos].Spec.Regions)!=0 {
			//	curData.Spec.MachineTypes[index].Spec.Regions = oldData.Spec.MachineTypes[pos].Spec.Regions
			//}
			//Disk
			if curData.Spec.MachineTypes[index].Spec.Disk == nil && oldData.Spec.MachineTypes[pos].Spec.Disk != nil {
				curData.Spec.MachineTypes[index].Spec.Disk = oldData.Spec.MachineTypes[pos].Spec.Disk
			}
			// RAM
			if curData.Spec.MachineTypes[index].Spec.RAM == nil && oldData.Spec.MachineTypes[pos].Spec.RAM != nil {
				curData.Spec.MachineTypes[index].Spec.RAM = oldData.Spec.MachineTypes[pos].Spec.RAM
			}
			// category
			if curData.Spec.MachineTypes[index].Spec.Category == "" && oldData.Spec.MachineTypes[pos].Spec.Category != "" {
				curData.Spec.MachineTypes[index].Spec.Category = oldData.Spec.MachineTypes[pos].Spec.Category
			}
			// CPU
			if curData.Spec.MachineTypes[index].Spec.CPU == nil && oldData.Spec.MachineTypes[pos].Spec.CPU != nil {
				curData.Spec.MachineTypes[index].Spec.CPU = oldData.Spec.MachineTypes[pos].Spec.CPU
			}
			// to detect it already added to curData
			instanceIndex[curData.Spec.MachineTypes[index].Spec.SKU] = -1
		}
	}
	for _, index := range instanceIndex {
		if index > -1 {
			// using regions as zones
			if len(oldData.Spec.MachineTypes[index].Spec.Regions) > 0 {
				if len(oldData.Spec.MachineTypes[index].Spec.Zones) == 0 {
					oldData.Spec.MachineTypes[index].Spec.Zones = oldData.Spec.MachineTypes[index].Spec.Regions
				}
				oldData.Spec.MachineTypes[index].Spec.Regions = nil
			}
			curData.Spec.MachineTypes = append(curData.Spec.MachineTypes, oldData.Spec.MachineTypes[index])
			curData.Spec.MachineTypes[len(curData.Spec.MachineTypes)-1].Spec.Deprecated = true
		}
	}
	return curData, nil
}

// get data from api , merge it with previous data and write the data
// previous data written in cloud_old.json
func MergeAndWriteCloudProvider(i Interface) error {
	klog.Infof("Getting cloud data for `%v` provider", i.GetName())
	curData, err := GetCloudProvider(i)
	if err != nil {
		return err
	}

	oldData, err := util.GetDataFormFile(i.GetName())
	if err != nil {
		return err
	}
	klog.Info("Merging cloud data...")
	res, err := MergeCloudProvider(oldData, curData)
	if err != nil {
		return err
	}

	//err = WriteCloudProvider(oldData,"cloud_old.json")
	//if err!=nil {
	//	return err
	//}
	klog.Info("Writing cloud data...")
	err = WriteCloudProvider(res)
	if err != nil {
		return err
	}
	return nil
}
