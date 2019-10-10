package main

import (
	"github.com/reillywatson/unsafestrconv"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(unsafestrconv.Analyzer) }
