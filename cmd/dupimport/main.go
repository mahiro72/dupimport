package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"

	"github.com/mahiro72/dupimport"
)

func main() {
	unitchecker.Main(dupimport.Analyzer)
}
