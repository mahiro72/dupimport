package dupimport_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/mahiro72/dupimport"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, dupimport.Analyzer, "dupimport")
}
