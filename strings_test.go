package stringz_test

import (
  . "github.com/orfjackal/gospec/src/gospec"
  "github.com/orfjackal/gospec/src/gospec"
  "runningwild/strings"
)

func idiotZboxer(p string) []int {
  zs := make([]int, len(p))
  for i := range zs {
    for zs[i] + i < len(p) && p[zs[i] + i] == p[zs[i]] {
      zs[i]++
    }
  }
  return zs
}

func ZBoxSpec(c gospec.Context) {
  c.Specify("Basic test.", func() {
    p := ""
    c.Expect(stringz.PrecalcZboxes(p), ContainsExactly, idiotZboxer(p))
    p = "a"
    c.Expect(stringz.PrecalcZboxes(p), ContainsExactly, idiotZboxer(p))
    p = "abcabc"
    c.Expect(stringz.PrecalcZboxes(p), ContainsExactly, idiotZboxer(p))
    p = "aabcabc"
    c.Expect(stringz.PrecalcZboxes(p), ContainsExactly, idiotZboxer(p))
    p = "aaabcabc"
    c.Expect(stringz.PrecalcZboxes(p), ContainsExactly, idiotZboxer(p))
    p = "aabcaabc"
    c.Expect(stringz.PrecalcZboxes(p), ContainsExactly, idiotZboxer(p))
    p = "aabcaaabc"
    c.Expect(stringz.PrecalcZboxes(p), ContainsExactly, idiotZboxer(p))
    p = "aabcaaaabc"
    c.Expect(stringz.PrecalcZboxes(p), ContainsExactly, idiotZboxer(p))
    p = "aabcaaaabc*aabcaaaabc"
    c.Expect(stringz.PrecalcZboxes(p), ContainsExactly, idiotZboxer(p))
    p = "aabcaaaabc*aabcaaaabc*aabcaaaabc*aabcaaaabc*aabcaaaabc"
    c.Expect(stringz.PrecalcZboxes(p), ContainsExactly, idiotZboxer(p))
    p = "abcdefghijklmnopq"
    c.Expect(stringz.PrecalcZboxes(p), ContainsExactly, idiotZboxer(p))
    p = "abcdefghijklmnopqabcdefghijklmnopqabcdefghijklmnopq"
    c.Expect(stringz.PrecalcZboxes(p), ContainsExactly, idiotZboxer(p))
    p = "abcdefghijklmnopqabc efghijklmnopqab cdefghijklmnopq"
    c.Expect(stringz.PrecalcZboxes(p), ContainsExactly, idiotZboxer(p))
    p = "aaaaaaaaaaaaaaaaaa"
    c.Expect(stringz.PrecalcZboxes(p), ContainsExactly, idiotZboxer(p))
    p = "bbbbbbbbbbbaaaaaaaaaaaa"
    c.Expect(stringz.PrecalcZboxes(p), ContainsExactly, idiotZboxer(p))
    p = "bbbbbbbaaaaaaaaabbbbbbbbbbbbbbaaaabbbbbbbbbaaaaabbaaaaaaaaabbaa"
    c.Expect(stringz.PrecalcZboxes(p), ContainsExactly, idiotZboxer(p))
  })
}