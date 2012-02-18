package stringz_test

import (
  "github.com/orfjackal/gospec/src/gospec"
  "testing"
)


func TestAllSpecs(t *testing.T) {
  r := gospec.NewRunner()
  r.AddSpec(ZBoxSpec)
  r.AddSpec(ZBoxReverseSpec)
  gospec.MainGoTest(r, t)
}

