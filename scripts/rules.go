// Package gorules defines custom lint rules for ruleguard.
//
// golangci-lint runs these rules via go-critic, which includes support
// for ruleguard. All Go files in this directory define lint rules
// in the Ruleguard DSL; see:
//
// - https://go-ruleguard.github.io/by-example/
// - https://pkg.go.dev/github.com/quasilyte/go-ruleguard/dsl
//
// You run one of the following commands to execute your go rules only:
//
//	golangci-lint run
//	golangci-lint run --disable-all --enable=gocritic
//
// Note: don't forget to run `golangci-lint cache clean`!
package gorules

import "github.com/quasilyte/go-ruleguard/dsl"

// asStringsIsDangerous checks for the use of AsString() on cty.Value.
// This function can panic if not used correctly, so the cty.Type must be known
// before calling. Ignore this lint if you are confident in your usage.
func asStringsIsDangerous(m dsl.Matcher) {
	m.Import("github.com/zclconf/go-cty/cty")

	m.Match(
		`$v.AsString()`,
	).
		Where(
			m["v"].Type.Is("cty.Value") &&
				// Ignore unit tests
				!m.File().Name.Matches(`_test\.go$`),
		).
		Report("'AsStrings()' can result in a panic if the type is not known. Ignore this linter with caution")
}
