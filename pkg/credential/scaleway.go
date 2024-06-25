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

type Scaleway struct {
	*v1alpha1.ScalewayCredential
}

func (c *Scaleway) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&c.Token, cloud.Scaleway+"."+ScalewayToken, c.Token, "Scaleway token")
	fs.StringVar(&c.Organization, cloud.Scaleway+"."+ScalewayOrganization, c.Organization, "Scaleway organization")
}

func (_ Scaleway) RequiredFlags() []string {
	return []string{
		cloud.Scaleway + "." + ScalewayToken,
		cloud.Scaleway + "." + ScalewayOrganization,
	}
}
