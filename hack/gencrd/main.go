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

package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	clusterinstall "go.bytebuilders.dev/resource-model/apis/cluster/install"
	clusterv1alpha1 "go.bytebuilders.dev/resource-model/apis/cluster/v1alpha1"
	identityinstall "go.bytebuilders.dev/resource-model/apis/identity/install"
	identityv1alpha1 "go.bytebuilders.dev/resource-model/apis/identity/v1alpha1"

	gort "github.com/appscode/go/runtime"
	"github.com/go-openapi/spec"
	"github.com/golang/glog"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/kube-openapi/pkg/common"
	"kmodules.xyz/client-go/openapi"
)

func generateSwaggerJson() {
	var (
		Scheme = runtime.NewScheme()
		Codecs = serializer.NewCodecFactory(Scheme)
	)

	clusterinstall.Install(Scheme)
	identityinstall.Install(Scheme)

	apispec, err := openapi.RenderOpenAPISpec(openapi.Config{
		Scheme: Scheme,
		Codecs: Codecs,
		Info: spec.InfoProps{
			Title:   "ByteBuilders",
			Version: "v0.0.1",
			Contact: &spec.ContactInfo{
				Name:  "AppsCode Inc.",
				URL:   "https://appscode.com",
				Email: "hello@appscode.com",
			},
			License: &spec.License{
				Name: "Apache 2.0",
				URL:  "https://www.apache.org/licenses/LICENSE-2.0.html",
			},
		},
		OpenAPIDefinitions: []common.GetOpenAPIDefinitions{
			clusterv1alpha1.GetOpenAPIDefinitions,
			identityv1alpha1.GetOpenAPIDefinitions,
		},
		//nolint:govet
		Resources: []openapi.TypeInfo{
			// v1alpha1 resources
			{clusterv1alpha1.SchemeGroupVersion, clusterv1alpha1.ResourceClusterInfos, clusterv1alpha1.ResourceKindClusterInfo, true},
			{clusterv1alpha1.SchemeGroupVersion, clusterv1alpha1.ResourceClusterAuthInfoTemplates, clusterv1alpha1.ResourceKindClusterAuthInfoTemplate, true},
			{clusterv1alpha1.SchemeGroupVersion, clusterv1alpha1.ResourceClusterUserAuths, clusterv1alpha1.ResourceKindClusterUserAuth, true},
			{clusterv1alpha1.SchemeGroupVersion, clusterv1alpha1.ResourceCloudCredentials, clusterv1alpha1.ResourceKindCloudCredential, true},
		},
		//nolint:govet
		RDResources: []openapi.TypeInfo{
			{identityv1alpha1.SchemeGroupVersion, identityv1alpha1.ResourcePluralSnapshot, identityv1alpha1.ResourceKindSnapshot, true},
		},
	})
	if err != nil {
		glog.Fatal(err)
	}

	filename := gort.GOPath() + "/src/go.bytebuilders.dev/resource-model/openapi/swagger.json"
	err = os.MkdirAll(filepath.Dir(filename), 0755)
	if err != nil {
		glog.Fatal(err)
	}
	err = ioutil.WriteFile(filename, []byte(apispec), 0644)
	if err != nil {
		glog.Fatal(err)
	}
}

func main() {
	generateSwaggerJson()
}
