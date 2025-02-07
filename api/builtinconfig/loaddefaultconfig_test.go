// Copyright 2019 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

package builtinconfig

import (
	"reflect"
	"testing"

	"sigs.k8s.io/kustomize/v3/api/resid"
	"sigs.k8s.io/kustomize/v3/api/types"
	"sigs.k8s.io/kustomize/v3/internal/loadertest"
)

func TestLoadDefaultConfigsFromFiles(t *testing.T) {
	ldr := loadertest.NewFakeLoader("/app")
	ldr.AddFile("/app/config.yaml", []byte(`
namePrefix:
- path: nameprefix/path
  kind: SomeKind
`))
	tcfg, err := loadDefaultConfig(ldr, []string{"/app/config.yaml"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := &TransformerConfig{
		NamePrefix: []types.FieldSpec{
			{
				Gvk:  resid.Gvk{Kind: "SomeKind"},
				Path: "nameprefix/path",
			},
		},
	}
	if !reflect.DeepEqual(tcfg, expected) {
		t.Fatalf("expected %v\n but go6t %v\n", expected, tcfg)
	}
}
