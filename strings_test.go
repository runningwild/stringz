package stringz_test

import (
  // "fmt"
  . "github.com/orfjackal/gospec/src/gospec"
  "github.com/orfjackal/gospec/src/gospec"
  "runningwild/strings"
  "testing"
  "math/rand"
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

func idiotZboxerReversed(p string) []int {
  zs := make([]int, len(p))
  for i := len(zs) - 1; i >= 0; i-- {
    for i - zs[i] >= 0 && p[i - zs[i]] == p[len(p) - zs[i] - 1] {
      zs[i]++
    }
  }
  return zs
}

// Returns a string of length n of all the same character
func makeTestString1(n int) string {
  b := make([]byte, n)
  return string(b)
}

// Returns a string of length n, first half one character, second half a
// different character
func makeTestString2(n int) string {
  b := make([]byte, n)
  for i := n/2; i < n; i++ {
    b[i] = 1
  }
  return string(b)
}

// Returns a string of length n, cycling through the number 0-255
func makeTestString3(n int) string {
  b := make([]byte, n)
  for i := range b {
    b[i] = byte(i % 256)
  }
  return string(b)
}

// Returns a string of length n consisting of random characters less than r
func makeTestString4(n,r int) string {
  rand.Seed(1234)
  b := make([]byte, n)
  for i := range b {
    b[i] = byte(rand.Intn(256) % r)
  }
  return string(b)
}

func BenchmarkZBox1_100k(b *testing.B) {
  b.StopTimer()
  p := makeTestString1(100000)
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    stringz.PrecalcZboxes(p)
  }
}

func BenchmarkZBox1_1M(b *testing.B) {
  b.StopTimer()
  p := makeTestString1(1000000)
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    stringz.PrecalcZboxes(p)
  }
}

func BenchmarkZBox2_100k(b *testing.B) {
  b.StopTimer()
  p := makeTestString2(100000)
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    stringz.PrecalcZboxes(p)
  }
}

func BenchmarkZBox2_1M(b *testing.B) {
  b.StopTimer()
  p := makeTestString2(1000000)
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    stringz.PrecalcZboxes(p)
  }
}

func BenchmarkZBox3_100k(b *testing.B) {
  b.StopTimer()
  p := makeTestString3(100000)
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    stringz.PrecalcZboxes(p)
  }
}

func BenchmarkZBox3_1M(b *testing.B) {
  b.StopTimer()
  p := makeTestString3(1000000)
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    stringz.PrecalcZboxes(p)
  }
}

func BenchmarkZBox4_100k(b *testing.B) {
  b.StopTimer()
  p := makeTestString4(100000, 256)
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    stringz.PrecalcZboxes(p)
  }
}

func BenchmarkZBox4_1M(b *testing.B) {
  b.StopTimer()
  p := makeTestString4(1000000, 256)
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    stringz.PrecalcZboxes(p)
  }
}

func augment(b []byte, radix int) bool {
  for i := range b {
  if int(b[i]) < radix - 1 {
    b[i]++
      return true
    } else {
      b[i] = 0
    }
  }
  return false
}

func ZBoxSpec(c gospec.Context) {
  c.Specify("Comprehensive test 3^9", func() {
    b := make([]byte, 9)
    for augment(b, 3) {
      p := string(b)
      c.Expect(stringz.PrecalcZboxes(p), ContainsExactly, idiotZboxer(p))
    }
  })
  c.Specify("Comprehensive test 2^15", func() {
    b := make([]byte, 15)
    for augment(b, 2) {
      p := string(b)
      c.Expect(stringz.PrecalcZboxes(p), ContainsExactly, idiotZboxer(p))
    }
  })
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
    p = "aabbaaa"
    c.Expect(stringz.PrecalcZboxes(p), ContainsExactly, idiotZboxer(p))
  })
}

func ZBoxReverseSpec(c gospec.Context) {
  c.Specify("Comprehensive test 3^9", func() {
    b := make([]byte, 9)
    for augment(b, 3) {
      p := string(b)
      c.Expect(stringz.PrecalcZboxesReversed(p), ContainsExactly, idiotZboxerReversed(p))
    }
  })
  c.Specify("Comprehensive test 2^15", func() {
    b := make([]byte, 15)
    for augment(b, 2) {
      p := string(b)
      c.Expect(stringz.PrecalcZboxesReversed(p), ContainsExactly, idiotZboxerReversed(p))
    }
  })
  c.Specify("Basic test.", func() {
    p := ""
    c.Expect(stringz.PrecalcZboxesReversed(p), ContainsExactly, idiotZboxerReversed(p))
    p = "a"
    c.Expect(stringz.PrecalcZboxesReversed(p), ContainsExactly, idiotZboxerReversed(p))
    p = "abcabc"
    c.Expect(stringz.PrecalcZboxesReversed(p), ContainsExactly, idiotZboxerReversed(p))
    p = "aabcabc"
    c.Expect(stringz.PrecalcZboxesReversed(p), ContainsExactly, idiotZboxerReversed(p))
    p = "aaabcabc"
    c.Expect(stringz.PrecalcZboxesReversed(p), ContainsExactly, idiotZboxerReversed(p))
    p = "aabcaabc"
    c.Expect(stringz.PrecalcZboxesReversed(p), ContainsExactly, idiotZboxerReversed(p))
    p = "aabcaaabc"
    c.Expect(stringz.PrecalcZboxesReversed(p), ContainsExactly, idiotZboxerReversed(p))
    p = "aabcaaaabc"
    c.Expect(stringz.PrecalcZboxesReversed(p), ContainsExactly, idiotZboxerReversed(p))
    p = "aabcaaaabc*aabcaaaabc"
    c.Expect(stringz.PrecalcZboxesReversed(p), ContainsExactly, idiotZboxerReversed(p))
    p = "aabcaaaabc*aabcaaaabc*aabcaaaabc*aabcaaaabc*aabcaaaabc"
    c.Expect(stringz.PrecalcZboxesReversed(p), ContainsExactly, idiotZboxerReversed(p))
    p = "abcdefghijklmnopq"
    c.Expect(stringz.PrecalcZboxesReversed(p), ContainsExactly, idiotZboxerReversed(p))
    p = "abcdefghijklmnopqabcdefghijklmnopqabcdefghijklmnopq"
    c.Expect(stringz.PrecalcZboxesReversed(p), ContainsExactly, idiotZboxerReversed(p))
    p = "abcdefghijklmnopqabc efghijklmnopqab cdefghijklmnopq"
    c.Expect(stringz.PrecalcZboxesReversed(p), ContainsExactly, idiotZboxerReversed(p))
    p = "aaaaaaaaaaaaaaaaaa"
    c.Expect(stringz.PrecalcZboxesReversed(p), ContainsExactly, idiotZboxerReversed(p))
    p = "bbbbbbbbbbbaaaaaaaaaaaa"
    c.Expect(stringz.PrecalcZboxesReversed(p), ContainsExactly, idiotZboxerReversed(p))
    p = "bbbbbbbaaaaaaaaabbbbbbbbbbbbbbaaaabbbbbbbbbaaaaabbaaaaaaaaabbaa"
    c.Expect(stringz.PrecalcZboxesReversed(p), ContainsExactly, idiotZboxerReversed(p))
    p = "aabbaaa"
    c.Expect(stringz.PrecalcZboxesReversed(p), ContainsExactly, idiotZboxerReversed(p))
  })
}

func idiotStringSearch(p,t string) []int {
  var matches []int
  for i := 0; i < len(t) - len(p) + 1; i++ {
    good := true
    for j := 0; j < len(p) && j + i < len(t); j++ {
      if p[j] != t[i] {
        good = false
        break
      }
    }
    if good {
      matches = append(matches, i)
    }
  }
  return matches
}

func BenchmarkBoyerMoore_10_100000(b *testing.B) {
  b.StopTimer()
  p := makeTestString1(10)
  t := makeTestString1(100000)
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    stringz.BoyerMoore(p, t)
  }
}

func BenchmarkBoyerMoore_100_100000(b *testing.B) {
  b.StopTimer()
  p := makeTestString1(100)
  t := makeTestString1(100000)
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    stringz.BoyerMoore(p, t)
  }
}

func BoyerMooreSpec(c gospec.Context) {
  c.Specify("Comprehensive test 3^9", func() {
    // p := "cabdabdab"
    // L,l := stringz.BoyerMooreStrongGoodSuffixRule(p)
    // fmt.Printf("%s\n%v\n", p, L)
    // fmt.Printf("%v\n", l)
    p := makeTestString4(5, 5)
    t := makeTestString4(5, 500)
    c.Expect(stringz.BoyerMoore(p, t), ContainsExactly, idiotStringSearch(p, t))
    p = makeTestString4(4, 5)
    t = makeTestString4(5, 500)
    c.Expect(stringz.BoyerMoore(p, t), ContainsExactly, idiotStringSearch(p, t))
    p = makeTestString4(4, 15)
    t = makeTestString4(5, 500)
    c.Expect(stringz.BoyerMoore(p, t), ContainsExactly, idiotStringSearch(p, t))
  })

}
