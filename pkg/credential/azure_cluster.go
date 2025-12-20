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

package credential

import (
	"go.bytebuilders.dev/resource-model/apis/cloud"
	"go.bytebuilders.dev/resource-model/apis/cloud/v1alpha1"

	"github.com/spf13/pflag"
)

type Azure struct {
	*v1alpha1.AzureCredential
}

func (c *Azure) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&c.TenantID, cloud.Azure+"."+AzureTenantID, c.TenantID, "provide this flag when provider is azure")
	fs.StringVar(&c.SubscriptionID, cloud.Azure+"."+AzureSubscriptionID, c.SubscriptionID, "provide this flag when provider is azure")
	fs.StringVar(&c.ClientID, cloud.Azure+"."+AzureClientID, c.ClientID, "provide this flag when provider is azure")
	fs.StringVar(&c.ClientSecret, cloud.Azure+"."+AzureClientSecret, c.ClientSecret, "provide this flag when provider is azure")
}

func (Azure) RequiredFlags() []string {
	return []string{
		cloud.Azure + "." + AzureTenantID,
		cloud.Azure + "." + AzureSubscriptionID,
		cloud.Azure + "." + AzureClientID,
		cloud.Azure + "." + AzureClientSecret,
	}
}
