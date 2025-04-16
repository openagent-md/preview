package types

import (
	"fmt"

	"github.com/aquasecurity/trivy/pkg/iac/terraform"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/zclconf/go-cty/cty"
	hcty "github.com/hashicorp/go-cty/cty"
	hctyjson "github.com/hashicorp/go-cty/cty/json"
	ctyjson "github.com/zclconf/go-cty/cty/json"

	"github.com/coder/terraform-provider-coder/v2/provider"
)

func providerValidations(vals []*ParameterValidation) []provider.Validation {
	cpy := make([]provider.Validation, 0, len(vals))
	for _, val := range vals {
		cpy = append(cpy, providerValidation(val))
	}
	return cpy
}

func providerValidation(v *ParameterValidation) provider.Validation {
	return provider.Validation{
		Min:         int(orZero(v.Min)),
		MinDisabled: v.Min == nil,
		Max:         int(orZero(v.Max)),
		MaxDisabled: v.Max == nil,
		Monotonic:   orZero(v.Monotonic),
		Regex:       orZero(v.Regex),
		Error:       v.Error,
	}
}

func providerOptions(opts []*ParameterOption) []provider.Option {
	cpy := make([]provider.Option, 0, len(opts))
	for _, opt := range opts {
		cpy = append(cpy, providerOption(opt))
	}
	return cpy
}

func providerOption(opt *ParameterOption) provider.Option {
	return provider.Option{
		Name:        opt.Name,
		Description: opt.Description,
		Value:       opt.Value.AsString(),
		Icon:        opt.Icon,
	}
}

func hclDiagnostics(diagnostics diag.Diagnostics, source *terraform.Block) hcl.Diagnostics {
	cpy := make(hcl.Diagnostics, 0, len(diagnostics))
	for _, d := range diagnostics {
		cpy = append(cpy, hclDiagnostic(d, source))
	}
	return cpy
}

func hclDiagnostic(d diag.Diagnostic, source *terraform.Block) *hcl.Diagnostic {
	sev := hcl.DiagInvalid
	switch d.Severity {
	case diag.Error:
		sev = hcl.DiagError
	case diag.Warning:
		sev = hcl.DiagWarning
	}

	// This is an imperfect way to finding the source code of the error. There is 2
	// different `cty` types at place here, the hashicorp fork and the original. So a
	// more general solution is difficult. This is good enough for now to add more
	// context to an error.
	var subject *hcl.Range
	if len(d.AttributePath) == 1 && source != nil {
		if attr, ok := d.AttributePath[0].(hcty.GetAttrStep); ok {
			src := source.GetAttribute(attr.Name)
			if src != nil {
				subject = &(src.HCLAttribute().Range)
			}
		}
	}

	return &hcl.Diagnostic{
		Severity:    sev,
		Summary:     d.Summary,
		Detail:      d.Detail,
		Subject:     subject,
		Context:     nil,
		Expression:  nil,
		EvalContext: nil,
		Extra:       nil,
	}
}
