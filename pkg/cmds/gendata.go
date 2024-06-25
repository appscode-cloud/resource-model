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

package cmds

import (
	"go.bytebuilders.dev/resource-model/pkg/providers"

	"github.com/spf13/cobra"
	"gomodules.xyz/x/term"
)

func NewCmdGenData() *cobra.Command {
	opts := providers.NewOptions()
	cmd := &cobra.Command{
		Use:               "gendata",
		Short:             "Load Kubernetes cluster data for a given cloud provider",
		Example:           "",
		DisableAutoGenTag: true,
		Run: func(cmd *cobra.Command, args []string) {
			if err := opts.ValidateFlags(cmd, args); err != nil {
				term.Fatalln(err)
			}
			cloudProvider, err := providers.NewCloudProvider(*opts)
			if err != nil {
				term.Fatalln(err)
			}
			err = providers.MergeAndWriteCloudProvider(cloudProvider)
			if err != nil {
				term.Fatalln(err)
			} else {
				term.Successln("Data successfully written for ", opts.Provider)
			}
		},
	}
	opts.AddFlags(cmd.Flags())
	return cmd
}
