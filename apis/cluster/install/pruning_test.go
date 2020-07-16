/*
Copyright The KubeVault Authors.

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

package install

import (
	"testing"

	"go.bytebuilders.dev/resource-model/apis/cluster/fuzzer"
	"go.bytebuilders.dev/resource-model/apis/cluster/v1alpha1"

	clientsetscheme "k8s.io/client-go/kubernetes/scheme"
	crdfuzz "kmodules.xyz/crd-schema-fuzz"
)

func TestPruneTypes(t *testing.T) {
	Install(clientsetscheme.Scheme)

	// v1alpha1
	if crd := (v1alpha1.ClusterInfo{}).CustomResourceDefinition(); crd.V1 != nil {
		crdfuzz.SchemaFuzzTestForV1CRD(t, clientsetscheme.Scheme, crd.V1, fuzzer.Funcs)
	}
	if crd := (v1alpha1.ClusterAuthInfoTemplate{}).CustomResourceDefinition(); crd.V1 != nil {
		crdfuzz.SchemaFuzzTestForV1CRD(t, clientsetscheme.Scheme, crd.V1, fuzzer.Funcs)
	}
	if crd := (v1alpha1.ClusterUserAuth{}).CustomResourceDefinition(); crd.V1 != nil {
		crdfuzz.SchemaFuzzTestForV1CRD(t, clientsetscheme.Scheme, crd.V1, fuzzer.Funcs)
	}
	if crd := (v1alpha1.CloudCredential{}).CustomResourceDefinition(); crd.V1 != nil {
		crdfuzz.SchemaFuzzTestForV1CRD(t, clientsetscheme.Scheme, crd.V1, fuzzer.Funcs)
	}

	// v1alpha1
	if crd := (v1alpha1.ClusterInfo{}).CustomResourceDefinition(); crd.V1beta1 != nil {
		crdfuzz.SchemaFuzzTestForV1beta1CRD(t, clientsetscheme.Scheme, crd.V1beta1, fuzzer.Funcs)
	}
	if crd := (v1alpha1.ClusterAuthInfoTemplate{}).CustomResourceDefinition(); crd.V1beta1 != nil {
		crdfuzz.SchemaFuzzTestForV1beta1CRD(t, clientsetscheme.Scheme, crd.V1beta1, fuzzer.Funcs)
	}
	if crd := (v1alpha1.ClusterUserAuth{}).CustomResourceDefinition(); crd.V1beta1 != nil {
		crdfuzz.SchemaFuzzTestForV1beta1CRD(t, clientsetscheme.Scheme, crd.V1beta1, fuzzer.Funcs)
	}
	if crd := (v1alpha1.CloudCredential{}).CustomResourceDefinition(); crd.V1beta1 != nil {
		crdfuzz.SchemaFuzzTestForV1beta1CRD(t, clientsetscheme.Scheme, crd.V1beta1, fuzzer.Funcs)
	}
}
