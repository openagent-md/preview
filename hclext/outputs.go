package hclext

import (
	"github.com/aquasecurity/trivy/pkg/iac/terraform"
	"github.com/zclconf/go-cty/cty"
)

func ExportOutputs(modules terraform.Modules) cty.Value {
	data := make(map[string]cty.Value)
	for _, block := range modules.GetBlocks().OfType("output") {
		attr := block.GetAttribute("value")
		if attr.IsNil() {
			continue
		}
		data[block.Label()] = attr.Value()
	}
	return cty.ObjectVal(data)
}
