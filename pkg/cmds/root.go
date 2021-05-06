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

package cmds

import (
	"flag"

	"go.bytebuilders.dev/resource-model/client/clientset/versioned/scheme"

	"github.com/spf13/cobra"
	"gomodules.xyz/kglog"
	"gomodules.xyz/x/flags"
	v "gomodules.xyz/x/version"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientsetscheme "k8s.io/client-go/kubernetes/scheme"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:               "bytebuilders-crd-tools",
		Short:             `ByteBuilders CRD tools by Appscode`,
		DisableAutoGenTag: true,
		PersistentPreRun: func(c *cobra.Command, args []string) {
			flags.DumpAll(c.Flags())
			utilruntime.Must(scheme.AddToScheme(clientsetscheme.Scheme))
		},
	}
	rootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	// ref: https://github.com/kubernetes/kubernetes/issues/17162#issuecomment-225596212
	kglog.ParseFlags()

	rootCmd.AddCommand(NewCmdGenData())
	rootCmd.AddCommand(v.NewCmdVersion())

	return rootCmd
}
