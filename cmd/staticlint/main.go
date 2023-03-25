// Static analyzer.
// Include static analytic packages:
// - golang.org/x/tools/go/analysis/passes
// - all SA/CA checks from staticcheck
// - asciicheck check that your code does not contain non-ASCII identifiers
// - containedctx check that detects struct contained context.Context field
// - ExitCheck search call os.Exit in main packages and report position
//
// Usage:
//  ./cmd/staticlint/exitcheck -exitcheck ./cmd/shortener/main.go
//  ./cmd/staticlint/exitcheck ./...

package main

import (
	"strings"

	"github.com/sivchari/containedctx"
	"github.com/tdakkota/asciicheck"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/assign"
	"golang.org/x/tools/go/analysis/passes/bools"
	"golang.org/x/tools/go/analysis/passes/httpresponse"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/shift"
	"honnef.co/go/tools/staticcheck"

	"github.com/casnerano/go-url-shortener/pkg/exitcheck"
)

func main() {
	analyzers := []*analysis.Analyzer{
		printf.Analyzer,
		shift.Analyzer,
		shadow.Analyzer,
		bools.Analyzer,
		assign.Analyzer,
		httpresponse.Analyzer,

		asciicheck.NewAnalyzer(),
		containedctx.Analyzer,

		exitcheck.Analyzer,
	}

	for _, v := range staticcheck.Analyzers {
		if strings.HasPrefix(v.Analyzer.Name, "SA") ||
			strings.HasPrefix(v.Analyzer.Name, "CA") {
			analyzers = append(analyzers, v.Analyzer)
		}
	}

	multichecker.Main(analyzers...)
}
