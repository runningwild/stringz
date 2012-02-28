package stringz_test

import (
  "fmt"
  . "github.com/orfjackal/gospec/src/gospec"
  "github.com/orfjackal/gospec/src/gospec"
  "github.com/runningwild/stringz"
  "testing"
)

func idiotZboxer(p string) []int {
  zs := make([]int, len(p))
  for i := range zs {
    for zs[i]+i < len(p) && p[zs[i]+i] == p[zs[i]] {
      zs[i]++
    }
  }
  return zs
}

func idiotZboxerReversed(p string) []int {
  zs := make([]int, len(p))
  for i := len(zs) - 1; i >= 0; i-- {
    for i-zs[i] >= 0 && p[i-zs[i]] == p[len(p)-zs[i]-1] {
      zs[i]++
    }
  }
  return zs
}

func idiotStringSearch(p, t string) []int {
  var matches []int
  for i := 0; i < len(t)-len(p)+1; i++ {
    good := true
    for j := 0; j < len(p) && j+i < len(t); j++ {
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

// BenchmarkBoyerMoore1* tests on a pattern and a text consisting of only one
// character.  This should be the worst case for a correct Boyer-Moore because
// it requires space to be allocated to return all of the matches.
func BenchmarkBoyerMoore1_10_100000(b *testing.B) {
  b.StopTimer()
  p := makeTestString1(10, 0)
  t := makeTestString1(100000, 0)
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    stringz.BoyerMoore(p, t)
  }
}

func BenchmarkBoyerMoore1_100_100000(b *testing.B) {
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

// BenchmarkBoyerMoore3* tests for a pattern consisting only of characters
// that are not found in the text.  These tests expect sublinear time that
// should decrease with the size of the pattern.
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

// BenchmarkBoyerMoore4* tests on random string with an alphabet size of 20,
// like protein
func BenchmarkBoyerMoore4_10_100000(b *testing.B) {
  b.StopTimer()
  p := makeTestString4(10, 20, 0)
  t := makeTestString4(100000, 20, 1)
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    stringz.BoyerMoore(p, t)
  }
}

func BenchmarkBoyerMoore4_100_100000(b *testing.B) {
  b.StopTimer()
  p := makeTestString4(100, 20, 0)
  t := makeTestString4(100000, 20, 1)
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    stringz.BoyerMoore(p, t)
  }
}

// BenchmarkBoyerMoore5* tests on random string with an alphabet size of 4,
// like DNA
func BenchmarkBoyerMoore5_10_100000(b *testing.B) {
  b.StopTimer()
  p := makeTestString4(10, 4, 0)
  t := makeTestString4(100000, 4, 1)
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    stringz.BoyerMoore(p, t)
  }
}

func BenchmarkBoyerMoore5_100_100000(b *testing.B) {
  b.StopTimer()
  p := makeTestString4(100, 4, 0)
  t := makeTestString4(100000, 4, 1)
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    stringz.BoyerMoore(p, t)
  }
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
    b := make([]byte, 17)
    for augment(b, 2) {
      p := string(b[0:5])
      t := string(b[5:])
      c.Expect(stringz.BoyerMoore(p, t), ContainsExactly, idiotStringSearch(p, t))
    }
  })

  c.Specify("Comprehensive test 3^11", func() {
    b := make([]byte, 11)
    for augment(b, 3) {
      p := string(b[0:4])
      t := string(b[4:])
      c.Expect(stringz.BoyerMoore(p, t), ContainsExactly, idiotStringSearch(p, t))
    }
  })

  c.Specify("Random test", func() {
    for i := 0; i < 10000; i += 2 {
      p := makeTestString4(15, 7, i)
      t := makeTestString4(1000, 7, i+1)
      c.Expect(stringz.BoyerMoore(p, t), ContainsExactly, idiotStringSearch(p, t))
    }
  })
}
