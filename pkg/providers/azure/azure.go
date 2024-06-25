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

package azure

import (
	"context"

	"go.bytebuilders.dev/resource-model/apis/cloud"
	"go.bytebuilders.dev/resource-model/apis/cloud/v1alpha1"
	"go.bytebuilders.dev/resource-model/pkg/credential"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	compute "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/compute/armcompute"
	subscriptions "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armsubscriptions"
)

type Client struct {
	SubscriptionId string
	GroupsClient   *subscriptions.Client
	VmSizesClient  *compute.VirtualMachineSizesClient
}

func NewClient(opts credential.Azure) (*Client, error) {
	g := &Client{
		SubscriptionId: opts.SubscriptionID,
	}
	var err error

	cred, err := azidentity.NewClientSecretCredential(opts.TenantID, opts.ClientID, opts.ClientSecret, nil)
	if err != nil {
		return nil, err
	}

	g.GroupsClient, err = subscriptions.NewClient(cred, nil)
	if err != nil {
		return nil, err
	}

	g.VmSizesClient, err = compute.NewVirtualMachineSizesClient(opts.SubscriptionID, cred, nil)
	if err != nil {
		return nil, err
	}

	return g, nil
}

func (g *Client) GetName() string {
	return cloud.Azure
}

func (g *Client) ListRegions() ([]v1alpha1.Region, error) {
	var regions []v1alpha1.Region
	pager := g.GroupsClient.NewListLocationsPager(g.SubscriptionId, nil)
	for pager.More() {
		regionList, err := pager.NextPage(context.TODO())
		if err != nil {
			return nil, err
		}
		for _, r := range regionList.Value {
			region := ParseRegion(r)
			regions = append(regions, *region)
		}
	}
	return regions, nil
}

func (g *Client) ListZones() ([]string, error) {
	regions, err := g.ListRegions()
	if err != nil {
		return nil, err
	}
	visZone := map[string]bool{}
	var zones []string
	for _, r := range regions {
		for _, z := range r.Zones {
			if _, found := visZone[z]; !found {
				zones = append(zones, z)
				visZone[z] = true
			}
		}
	}
	return zones, nil
}

func (g *Client) ListMachineTypes() ([]v1alpha1.MachineType, error) {
	zones, err := g.ListZones()
	if err != nil {
		return nil, err
	}
	var instances []v1alpha1.MachineType
	// to find the position in instances array
	instancePos := map[string]int{}
	for _, zone := range zones {
		pager := g.VmSizesClient.NewListPager(zone, nil)
		for pager.More() {
			instanceList, err := pager.NextPage(context.TODO())
			if err != nil {
				continue
			}
			for _, ins := range instanceList.Value {
				instance, err := ParseInstance(ins)
				if err != nil {
					return nil, err
				}
				pos, found := instancePos[instance.Spec.SKU]
				if found {
					instances[pos].Spec.Zones = append(instances[pos].Spec.Zones, zone)
				} else {
					instancePos[instance.Spec.SKU] = len(instances)
					instance.Spec.Zones = []string{zone}
					instances = append(instances, *instance)
				}
			}
		}
	}
	return instances, nil
}
