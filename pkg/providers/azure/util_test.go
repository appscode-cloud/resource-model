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

package azure

import (
	"encoding/json"
	"reflect"
	"testing"

	v1 "go.bytebuilders.dev/resource-model/apis/cloud/v1alpha1"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-03-01/compute"
	"github.com/appscode/go/types"
	"github.com/davecgh/go-spew/spew"
)

func TestParseInstance(t *testing.T) {
	type args struct {
		in *compute.VirtualMachineSize
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "azure-a",
			args: args{
				in: &compute.VirtualMachineSize{
					Name:                 types.StringP("Standard_A0"),
					NumberOfCores:        types.Int32P(1),
					OsDiskSizeInMB:       types.Int32P(1047552),
					ResourceDiskSizeInMB: types.Int32P(20480),
					MemoryInMB:           types.Int32P(768),
					MaxDataDiskCount:     types.Int32P(1),
				},
			},
			want: `{
  "cpu": "1",
  "description": "Standard_A0",
  "disk": "1047552M",
  "ram": "768M",
  "sku": "Standard_A0"
}`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseInstance(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseInstance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			var wantMachine v1.MachineTypeSpec
			err = json.Unmarshal([]byte(tt.want), &wantMachine)
			if err != nil {
				t.Fatalf(err.Error())
			}
			spew.Dump(wantMachine)
			spew.Dump(got.Spec)
			if !reflect.DeepEqual(got.Spec, wantMachine) {
				// oneliners.PrettyJson(got.Spec, "got")
				// oneliners.PrettyJson(wantMachine, "expected")
				t.Errorf("specs didn't match")
			}
		})
	}
}
