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
	"os"

	"go.bytebuilders.dev/resource-model/apis/cloud/v1alpha1"

	"github.com/spf13/pflag"
)

type GCE struct {
	*v1alpha1.GoogleCloudCredential

	credentialFile string
}

func (c GCE) ServiceAccountJson() string {
	if c.ServiceAccount != "" {
		return c.ServiceAccount
	}

	data, err := os.ReadFile(c.credentialFile)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func (c *GCE) AddFlags(fs *pflag.FlagSet) {
	fs.StringVarP(&c.credentialFile, "gce.credential_file", "c", c.credentialFile, "Location of cloud credential file (required when --provider=gce)")
	fs.StringVar(&c.ProjectID, "gce.project_id", c.ProjectID, "provide this flag when provider is gce")
}

func (_ GCE) RequiredFlags() []string {
	return []string{
		"gce.credential_file",
		"gce.project_id",
	}
}
