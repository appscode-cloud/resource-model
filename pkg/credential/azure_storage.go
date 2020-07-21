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

package credential

import (
	"go.bytebuilders.dev/resource-model/apis/cloud"
	"go.bytebuilders.dev/resource-model/apis/cloud/v1alpha1"

	"github.com/spf13/pflag"
)

type AzureStorage struct {
	*v1alpha1.AzureStorageCredential
}

func (c *AzureStorage) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&c.Account, cloud.AzureStorage+"."+AzureStorageAccount, c.Account, "Azure storage account")
	fs.StringVar(&c.Key, cloud.AzureStorage+"."+AzureStorageKey, c.Key, "Azure storage account key")
}

func (_ AzureStorage) RequiredFlags() []string {
	return []string{
		cloud.AzureStorage + "." + AzureStorageAccount,
		cloud.AzureStorage + "." + AzureStorageKey,
	}
}
