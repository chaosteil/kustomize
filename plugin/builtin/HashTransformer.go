// Code generated by pluginator on HashTransformer; DO NOT EDIT.
package builtin

import (
	"fmt"

	"sigs.k8s.io/kustomize/v3/api/ifc"
	"sigs.k8s.io/kustomize/v3/api/resmap"
)

type HashTransformerPlugin struct {
	hasher ifc.KunstructuredHasher
}

func (p *HashTransformerPlugin) Config(
	h *resmap.PluginHelpers, config []byte) (err error) {
	p.hasher = h.ResmapFactory().RF().Hasher()
	return nil
}

// Transform appends hash to generated resources.
func (p *HashTransformerPlugin) Transform(m resmap.ResMap) error {
	for _, res := range m.Resources() {
		if res.NeedHashSuffix() {
			h, err := p.hasher.Hash(res)
			if err != nil {
				return err
			}
			res.SetName(fmt.Sprintf("%s-%s", res.GetName(), h))
		}
	}
	return nil
}

func NewHashTransformerPlugin() resmap.TransformerPlugin {
	return &HashTransformerPlugin{}
}
