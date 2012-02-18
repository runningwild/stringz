package stringz_test

import (
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

// Returns a string of length n consisting of random characters
func makeTestString4(n int) string {
  rand.Seed(1234)
  b := make([]byte, n)
  for i := range b {
    b[i] = byte(rand.Intn(256))
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
  p := makeTestString4(100000)
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    stringz.PrecalcZboxes(p)
  }
}

func BenchmarkZBox4_1M(b *testing.B) {
  b.StopTimer()
  p := makeTestString4(1000000)
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    stringz.PrecalcZboxes(p)
  }
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