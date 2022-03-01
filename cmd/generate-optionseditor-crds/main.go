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
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	v1alpha12 "go.bytebuilders.dev/resource-model/apis/ui/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"kmodules.xyz/resource-metadata/apis/meta/v1alpha1"
	"kmodules.xyz/resource-metadata/hub"
	"sigs.k8s.io/yaml"
)

var dir = "/home/tamal/go/src/go.bytebuilders.dev/ui-wizards/ui"

func main() {
	reg := hub.NewRegistryOfKnownResources()
	reg.Visit(func(key string, rd *v1alpha1.ResourceDescriptor) {
		if rd.Spec.UI == nil || rd.Spec.UI.Options == nil {
			return
		}

		gvr := rd.Spec.Resource.GroupVersionResource()
		filename := filepath.Join(dir, Filename(gvr))

		category := ""
		if rd.Spec.Resource.Group == "kubedb.com" {
			category = "database"
		}

		editor := v1alpha12.OptionsEditor{
			TypeMeta: v1.TypeMeta{
				APIVersion: v1alpha12.SchemeGroupVersion.String(),
				Kind:       "OptionsEditor",
			},
			ObjectMeta: v1.ObjectMeta{
				Name: Name(gvr),
			},
			Spec: v1alpha12.OptionsEditorSpec{
				Resource: v1alpha12.ResourceID{
					Group:   rd.Spec.Resource.Group,
					Version: rd.Spec.Resource.Version,
					Name:    rd.Spec.Resource.Name,
					Kind:    rd.Spec.Resource.Kind,
					Scope:   v1alpha12.ResourceScope(rd.Spec.Resource.Scope),
				},
				ChartRef: v1alpha12.ChartRepoRef{
					Name:    rd.Spec.UI.Options.Name,
					URL:     rd.Spec.UI.Options.URL,
					Version: rd.Spec.UI.Options.Version,
				},
				Category: category,
				Provider: "AppsCode",
			},
		}
		data, err := yaml.Marshal(editor)
		if err != nil {
			panic(err)
		}

		if err := os.MkdirAll(filepath.Dir(filename), 0o755); err != nil {
			panic(err)
		}
		err = ioutil.WriteFile(filename, data, 0o644)
		if err != nil {
			panic(err)
		}
	})
}

func Filename(gvr schema.GroupVersionResource) string {
	if gvr.Group == "" && gvr.Version == "v1" {
		return fmt.Sprintf("core/v1/%s.yaml", gvr.Resource)
	}
	return fmt.Sprintf("%s/%s/%s.yaml", gvr.Group, gvr.Version, gvr.Resource)
}

func Name(gvr schema.GroupVersionResource) string {
	if gvr.Group == "" && gvr.Version == "v1" {
		return fmt.Sprintf("core-v1-%s", gvr.Resource)
	}
	return fmt.Sprintf("%s-%s-%s", gvr.Group, gvr.Version, gvr.Resource)
}
