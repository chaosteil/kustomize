// Copyright 2019 The Kubernetes Authors.
// SPDX-License-Identifier: Apache-2.0

//go:generate pluginator
package main

import (
	"sigs.k8s.io/kustomize/v3/api/resmap"
	"sigs.k8s.io/kustomize/v3/api/transform"
	"sigs.k8s.io/kustomize/v3/api/types"
	"sigs.k8s.io/yaml"
)

// Add the given annotations to the given field specifications.
type plugin struct {
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	FieldSpecs  []types.FieldSpec `json:"fieldSpecs,omitempty" yaml:"fieldSpecs,omitempty"`
}

//noinspection GoUnusedGlobalVariable
var KustomizePlugin plugin

func (p *plugin) Config(
	h *resmap.PluginHelpers, c []byte) (err error) {
	p.Annotations = nil
	p.FieldSpecs = nil
	return yaml.Unmarshal(c, p)
}

func (p *plugin) Transform(m resmap.ResMap) error {
	t, err := transform.NewMapTransformer(
		p.FieldSpecs,
		p.Annotations,
	)
	if err != nil {
		return err
	}
	return t.Transform(m)
}
