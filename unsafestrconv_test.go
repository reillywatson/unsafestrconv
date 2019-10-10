package unsafestrconv_test

import (
	"testing"

	"github.com/reillywatson/unsafestrconv"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, unsafestrconv.Analyzer, "a")
}
