package stringz_test

import (
  "fmt"
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

// Returns a string of length n of all the same character, c
func makeTestString1(n int, c byte) string {
  b := make([]byte, n)
  for i := range b {
    b[i] = c
  }
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

// Returns a string of length n, cycling through the number 0-(r-1)
func makeTestString3(n,r int) string {
  b := make([]byte, n)
  for i := range b {
    b[i] = byte(i % r)
  }
  return string(b)
}

// Returns a string of length n consisting of random characters less than r,
// and using seed s
func makeTestString4(n,r,s int) string {
  rand.Seed(int64(s))
  b := make([]byte, n)
  for i := range b {
    b[i] = byte(rand.Intn(256) % r)
  }
  return string(b)
}

func BenchmarkZBox1_100k(b *testing.B) {
  b.StopTimer()
  p := makeTestString1(100000, 0)
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    stringz.PrecalcZboxes(p)
  }
}

func BenchmarkZBox1_1M(b *testing.B) {
  b.StopTimer()
  p := makeTestString1(1000000, 0)
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
  p := makeTestString3(100000, 255)
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    stringz.PrecalcZboxes(p)
  }
}

func BenchmarkZBox3_1M(b *testing.B) {
  b.StopTimer()
  p := makeTestString3(1000000, 255)
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    stringz.PrecalcZboxes(p)
  }
}

func BenchmarkZBox4_100k(b *testing.B) {
  b.StopTimer()
  p := makeTestString4(100000, 256, 1)
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    stringz.PrecalcZboxes(p)
  }
}

func BenchmarkZBox4_1M(b *testing.B) {
  b.StopTimer()
  p := makeTestString4(1000000, 256, 1)
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
      if p[j] != t[i+j] {
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
  p := makeTestString1(10, 0)
  t := makeTestString1(100000, 0)
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    stringz.BoyerMoore(p, t)
  }
}

func BenchmarkBoyerMoore_100_100000(b *testing.B) {
  b.StopTimer()
  p := makeTestString1(100, 0)
  t := makeTestString1(100000, 0)
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    stringz.BoyerMoore(p, t)
  }
}

// BenchmarkBoyerMoore2* tests make sure that the runtime does not go
// quadratic when search for something of the for abxb(ab)+ in (ab)+
func BenchmarkBoyerMoore2_10_100000(b *testing.B) {
  b.StopTimer()
  P := makeTestString3(10, 2)
  pb := []byte(P)
  pb[2] = 'x'
  p := string(pb)
  t := makeTestString3(100000, 2)
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    stringz.BoyerMoore(p, t)
  }
}

func BenchmarkBoyerMoore2_100_100000(b *testing.B) {
  b.StopTimer()
  P := makeTestString3(100, 2)
  pb := []byte(P)
  pb[2] = 'x'
  p := string(pb)
  t := makeTestString3(100000, 2)
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    stringz.BoyerMoore(p, t)
  }
}

func BenchmarkBoyerMoore3_10_100000(b *testing.B) {
  b.StopTimer()
  p := makeTestString1(10, 0)
  t := makeTestString1(100000, 1)
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    stringz.BoyerMoore(p, t)
  }
}

func BenchmarkBoyerMoore3_100_100000(b *testing.B) {
  b.StopTimer()
  p := makeTestString1(100, 0)
  t := makeTestString1(100000, 1)
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    stringz.BoyerMoore(p, t)
  }
}

func same(a,b []int) bool {
  if len(a) != len(b) { return false }
  for i := range a {
    if a[i] != b[i] { return false }
  }
  return true
}

func dstr(b []byte) string {
  b2 := make([]byte, len(b))
  for i := range b2 {
    b2[i] += 'a'
  }
  return string(b2)
}

func idiotLongestSuffixAsPrefix(p string) []int {
  v := make([]int, len(p))
  for i := range p {
    for j := i; j < len(p); j++ {
      s := p[j:]
      if s == p[0:len(s)] {
        v[i] = len(s)
        break
      }
    }
  }
  return v
}
func LongestSuffixAsPrefixSpec(c gospec.Context) {
  fmt.Printf("")
  c.Specify("Comprehensive test 2^15", func() {
    b := make([]byte, 15)
    for augment(b, 2) {
      p := string(b)
      c.Expect(stringz.LongestSuffixAsPrefix(p), ContainsExactly, idiotLongestSuffixAsPrefix(p))
    }
  })
  c.Specify("Comprehensive test 3^9", func() {
    b := make([]byte, 9)
    for augment(b, 3) {
      p := string(b)
      c.Expect(stringz.LongestSuffixAsPrefix(p), ContainsExactly, idiotLongestSuffixAsPrefix(p))
    }
  })
  c.Specify("Comprehensive test 4^7", func() {
    b := make([]byte, 7)
    for augment(b, 4) {
      p := string(b)
      c.Expect(stringz.LongestSuffixAsPrefix(p), ContainsExactly, idiotLongestSuffixAsPrefix(p))
    }
  })
}

func BoyerMooreSpec(c gospec.Context) {
  c.Specify("Comprehensive test 2^17", func() {
    // p := "abaa"
    // t := "aabaaaaaa"
    // zr := stringz.PrecalcZboxesReversed(p)
    // fmt.Printf("%s\nzr: %v\n", p, zr)
    // c.Expect(stringz.BoyerMoore(p, t), ContainsExactly, idiotStringSearch(p, t))
    // panic('a')
    b := make([]byte, 17)
    for augment(b, 2) {
      p := string(b[0 : 5])
      t := string(b[5 :  ])
        // fmt.Printf("p: %v\nt: %v\n", []byte(p), []byte(t))
      bm_m := stringz.BoyerMoore(p, t)
      i_m := idiotStringSearch(p, t)
      if !same(bm_m, i_m) {
        fmt.Printf("p: %v\nt: %v\n", []byte(p), []byte(t))
        fmt.Printf("b: %v\n", bm_m)
        fmt.Printf("i: %v\n", i_m)
        fmt.Printf("\n")
        panic("A")
      }
      c.Expect(stringz.BoyerMoore(p, t), ContainsExactly, idiotStringSearch(p, t))
    }
  })

  c.Specify("Comprehensive test 3^9", func() {
    b := make([]byte, 11)
    for augment(b, 3) {
      p := string(b[0 : 4])
      t := string(b[4 :  ])
      bm_m := stringz.BoyerMoore(p, t)
      i_m := idiotStringSearch(p, t)
      if !same(bm_m, i_m) {
        fmt.Printf("p: %v\n", (b[0:5]))
        fmt.Printf("t: %v\n", (b[5:]))
        fmt.Printf("b: %v\n", bm_m)
        fmt.Printf("i: %v\n", i_m)
        fmt.Printf("\n")
        panic("A")
      }
      c.Expect(stringz.BoyerMoore(p, t), ContainsExactly, idiotStringSearch(p, t))
    }
  })

  c.Specify("Random test", func() {
    for i := 0; i < 10000; i+=2 {
      p := makeTestString4(15, 7, i)
      t := makeTestString4(1000, 7, i+1)
      bm_m := stringz.BoyerMoore(p, t)
      i_m := idiotStringSearch(p, t)
      if !same(bm_m, i_m) {
        fmt.Printf("p: %s\n", p)
        fmt.Printf("t: %s\n", t)
        fmt.Printf("b: %v\n", bm_m)
        fmt.Printf("i: %v\n", i_m)
        fmt.Printf("\n")
        panic("A")
      }
      c.Expect(stringz.BoyerMoore(p, t), ContainsExactly, idiotStringSearch(p, t))
    }
  })

}
