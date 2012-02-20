package stringz_test

import (
  "github.com/orfjackal/gospec/src/gospec"
  "testing"
)

func TestAllSpecs(t *testing.T) {
  r := gospec.NewRunner()
  r.AddSpec(ZBoxSpec)
  r.AddSpec(ZBoxReverseSpec)
  r.AddSpec(LongestSuffixAsPrefixSpec)
  r.AddSpec(BoyerMooreSpec)
  gospec.MainGoTest(r, t)
}
