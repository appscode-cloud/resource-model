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
	"fmt"
	"os"

	"go.bytebuilders.dev/resource-model/apis/cloud"
	"go.bytebuilders.dev/resource-model/apis/cloud/v1alpha1"
	"go.bytebuilders.dev/resource-model/pkg/credential"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gomodules.xyz/x/flags"
)

type Options struct {
	Provider string
	Do       credential.DigitalOcean
	Packet   credential.Packet
	GCE      credential.GCE
	AWS      credential.AWS
	Azure    credential.Azure
	Vultr    credential.Vultr
	Linode   credential.Linode
	Scaleway credential.Scaleway
}

func NewOptions() *Options {
	return &Options{
		Provider: "",
	}
}

func NewOptionsForCredential(c v1alpha1.Credential) Options {
	opt := Options{
		Provider: string(c.Spec.Type),
	}
	switch c.Spec.Type {
	case v1alpha1.CredentialTypeAWS:
		opt.AWS = credential.AWS{AWSCredential: c.Spec.AWS}
	case v1alpha1.CredentialTypeGoogleCloud:
		opt.GCE = credential.GCE{GoogleCloudCredential: c.Spec.GoogleCloud}
	case v1alpha1.CredentialTypeDigitalOcean:
		opt.Do = credential.DigitalOcean{DigitalOceanCredential: c.Spec.DigitalOcean}
	case v1alpha1.CredentialTypePacket:
		opt.Packet = credential.Packet{PacketCredential: c.Spec.Packet}
	case v1alpha1.CredentialTypeAzure:
		opt.Azure = credential.Azure{AzureCredential: c.Spec.Azure}
	case v1alpha1.CredentialTypeVultr:
		opt.Vultr = credential.Vultr{VultrCredential: c.Spec.Vultr}
	case v1alpha1.CredentialTypeLinode:
		opt.Linode = credential.Linode{LinodeCredential: c.Spec.Linode}
	case v1alpha1.CredentialTypeScaleway:
		opt.Scaleway = credential.Scaleway{ScalewayCredential: c.Spec.Scaleway}
	default:
		panic("unknown provider " + c.Spec.Type)
	}
	return opt
}

func (c *Options) AddFlags(fs *pflag.FlagSet) {
	fs.StringVarP(&c.Provider, "provider", "p", c.Provider, "Name of the Cloud provider")
	c.Do.AddFlags(fs)
	c.Packet.AddFlags(fs)
	c.GCE.AddFlags(fs)
	c.AWS.AddFlags(fs)
	c.Azure.AddFlags(fs)
	c.Vultr.AddFlags(fs)
	c.Linode.AddFlags(fs)
	c.Scaleway.AddFlags(fs)
}

func (c *Options) ValidateFlags(cmd *cobra.Command, args []string) error {
	var required []string

	switch c.Provider {
	case cloud.GCE:
		required = c.GCE.RequiredFlags()
	case cloud.DigitalOcean:
		required = c.Do.RequiredFlags()
	case cloud.Packet:
		required = c.Packet.RequiredFlags()
	case cloud.AWS:
		required = c.AWS.RequiredFlags()
	case cloud.Azure:
		required = c.Azure.RequiredFlags()
	case cloud.Vultr:
		required = c.Vultr.RequiredFlags()
	case cloud.Linode:
		required = c.Linode.RequiredFlags()
	case cloud.Scaleway:
		required = c.Scaleway.RequiredFlags()
	default:
		fmt.Println("missing --provider flag")
		os.Exit(1)
	}

	if len(required) > 0 {
		flags.EnsureRequiredFlags(cmd, required...)
	}
	return nil
}
