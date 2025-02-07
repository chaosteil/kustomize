// Copyright 2019 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package main_test

import (
	"testing"

	"sigs.k8s.io/kustomize/v3/api/kusttest"
)

func TestBashedConfigMapPlugin(t *testing.T) {
	tc := kusttest_test.NewPluginTestEnv(t).Set()
	defer tc.Reset()

	tc.BuildExecPlugin(
		"someteam.example.com", "v1", "BashedConfigMap")

	th := kusttest_test.NewKustTestHarnessAllowPlugins(t, "/app")

	m := th.LoadAndRunGenerator(`
apiVersion: someteam.example.com/v1
kind: BashedConfigMap
metadata:
  name: whatever
argsOneLiner: alice myMomsMaidenName
`)
	th.AssertActualEqualsExpected(m, `
apiVersion: v1
data:
  password: myMomsMaidenName
  username: alice
kind: ConfigMap
metadata:
  name: example-configmap-test
`)
	if m.Resources()[0].NeedHashSuffix() != true {
		t.Errorf("expected resource to need hashing")
	}
}
