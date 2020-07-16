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
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-03-01/compute"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-06-01/subscriptions"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
)

//https://docs.microsoft.com/en-us/rest/api/compute/virtualmachines/virtualmachines-list-sizes-region
//https://docs.microsoft.com/en-us/rest/api/compute/virtualmachines/virtualmachines-list-sizes-for-resizing

type credInfo struct {
	ClientId       string `json:"clientID"`
	ClientSecret   string `json:"clientSecret"`
	SubscriptionId string `json:"subscriptionID"`
	TenantId       string `json:"tenantID"`
}

func TestRegion(t *testing.T) {
	cred, err := getCredential()
	if err != nil {
		t.Error(err)
	}
	baseURI := azure.PublicCloud.ResourceManagerEndpoint
	config, err := adal.NewOAuthConfig(azure.PublicCloud.ActiveDirectoryEndpoint, cred.TenantId)
	if err != nil {
		t.Error(err)
	}

	spt, err := adal.NewServicePrincipalToken(*config, cred.ClientId, cred.ClientSecret, baseURI)
	if err != nil {
		t.Error(err)
	}
	groupsClient := subscriptions.NewClient()
	groupsClient.Authorizer = autorest.NewBearerAuthorizer(spt)
	g := Client{
		GroupsClient:   groupsClient,
		SubscriptionId: cred.SubscriptionId,
	}
	r, err := g.ListRegions()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)
}

func TestInstances(t *testing.T) {
	cred, err := getCredential()
	if err != nil {
		t.Error(err)
	}
	baseURI := azure.PublicCloud.ResourceManagerEndpoint
	config, err := adal.NewOAuthConfig(azure.PublicCloud.ActiveDirectoryEndpoint, cred.TenantId)
	if err != nil {
		t.Error(err)
	}

	spt, err := adal.NewServicePrincipalToken(*config, cred.ClientId, cred.ClientSecret, baseURI)
	if err != nil {
		t.Error(err)
	}
	vmSzClient := compute.NewVirtualMachineSizesClient(cred.SubscriptionId)
	vmSzClient.Authorizer = autorest.NewBearerAuthorizer(spt)
	groupsClient := subscriptions.NewClient()
	groupsClient.Authorizer = autorest.NewBearerAuthorizer(spt)
	g := Client{
		VmSizesClient:  vmSzClient,
		GroupsClient:   groupsClient,
		SubscriptionId: cred.SubscriptionId,
	}
	r, err := g.ListMachineTypes()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(r)
}

func getCredential() (*credInfo, error) {
	var cred credInfo
	bytes, err := ioutil.ReadFile(filepath.Join(
		os.Getenv("HOME"), ".pharmer", "store.d", os.Getenv("USER"), "credentials", "azure.json"))

	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &cred)
	if err != nil {
		return nil, err
	}
	return &cred, nil
}
