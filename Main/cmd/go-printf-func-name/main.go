package main

import (
	// TODO: explore using please to solve my import issues
    "gitlab.com/boldly-go/linter-tutorial/analyzer"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}