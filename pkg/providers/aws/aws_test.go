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

package aws

import (
	"testing"

	"go.bytebuilders.dev/resource-model/pkg/credential"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var opts = credential.AWS{
	Region: "us-east-1",
}

func TestRegion(t *testing.T) {
	g, err := NewClient(opts)
	if err != nil {
		t.Error(err)
		return
	}
	g.session, err = session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials("", "", ""),
	})
	if err != nil {
		t.Error(err)
		return
	}
	_, err = g.ListRegions()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestInstance(t *testing.T) {
	g, err := NewClient(opts)
	if err != nil {
		t.Error(err)
		return
	}
	_, err = g.ListMachineTypes()
	if err != nil {
		t.Error(err)
		return
	}
}
