package core_test

import (
  "fmt"
  "bytes"
  . "github.com/orfjackal/gospec/src/gospec"
  "github.com/orfjackal/gospec/src/gospec"
  "github.com/runningwild/stringz/core"
  "testing"
)

func idiotZboxer(p []byte) []int {
  zs := make([]int, len(p))
  for i := range zs {
    for zs[i]+i < len(p) && p[zs[i]+i] == p[zs[i]] {
      zs[i]++
    }
  }
  return zs
}

func idiotZboxerReversed(p []byte) []int {
  zs := make([]int, len(p))
  for i := len(zs) - 1; i >= 0; i-- {
    for i-zs[i] >= 0 && p[i-zs[i]] == p[len(p)-zs[i]-1] {
      zs[i]++
    }
  }
  return zs
}

func idiotStringSearch(p, t []byte) []int {
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

func idiotLongestSuffixAsPrefix(p []byte) []int {
  v := make([]int, len(p))
  for i := range p {
    for j := i; j < len(p); j++ {
      s := p[j:]
      if string(s) == string(p[0:len(s)]) {
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
    core.PrecalcZboxes(p)
  }
}

func BenchmarkZBox1_1M(b *testing.B) {
  b.StopTimer()
  p := makeTestString1(1000000, 0)
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    core.PrecalcZboxes(p)
  }
}

func BenchmarkZBox2_100k(b *testing.B) {
  b.StopTimer()
  p := makeTestString2(100000)
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    core.PrecalcZboxes(p)
  }
}

func BenchmarkZBox2_1M(b *testing.B) {
  b.StopTimer()
  p := makeTestString2(1000000)
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    core.PrecalcZboxes(p)
  }
}

func BenchmarkZBox3_100k(b *testing.B) {
  b.StopTimer()
  p := makeTestString3(100000, 255)
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    core.PrecalcZboxes(p)
  }
}

func BenchmarkZBox3_1M(b *testing.B) {
  b.StopTimer()
  p := makeTestString3(1000000, 255)
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    core.PrecalcZboxes(p)
  }
}

func BenchmarkZBox4_100k(b *testing.B) {
  b.StopTimer()
  p := makeTestString4(100000, 256, 1)
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    core.PrecalcZboxes(p)
  }
}

func BenchmarkZBox4_1M(b *testing.B) {
  b.StopTimer()
  p := makeTestString4(1000000, 256, 1)
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    core.PrecalcZboxes(p)
  }
}

func ZBoxSpec(c gospec.Context) {
  c.Specify("Comprehensive test 3^9", func() {
    p := make([]byte, 9)
    for augment(p, 3) {
      c.Expect(core.PrecalcZboxes(p), ContainsExactly, idiotZboxer(p))
    }
  })
  c.Specify("Comprehensive test 2^15", func() {
    p := make([]byte, 15)
    for augment(p, 2) {
      c.Expect(core.PrecalcZboxes(p), ContainsExactly, idiotZboxer(p))
    }
  })
  c.Specify("Basic test.", func() {
    p := ""
    c.Expect(core.PrecalcZboxes([]byte(p)), ContainsExactly, idiotZboxer([]byte(p)))
    p = "a"
    c.Expect(core.PrecalcZboxes([]byte(p)), ContainsExactly, idiotZboxer([]byte(p)))
    p = "abcabc"
    c.Expect(core.PrecalcZboxes([]byte(p)), ContainsExactly, idiotZboxer([]byte(p)))
    p = "aabcabc"
    c.Expect(core.PrecalcZboxes([]byte(p)), ContainsExactly, idiotZboxer([]byte(p)))
    p = "aaabcabc"
    c.Expect(core.PrecalcZboxes([]byte(p)), ContainsExactly, idiotZboxer([]byte(p)))
    p = "aabcaabc"
    c.Expect(core.PrecalcZboxes([]byte(p)), ContainsExactly, idiotZboxer([]byte(p)))
    p = "aabcaaabc"
    c.Expect(core.PrecalcZboxes([]byte(p)), ContainsExactly, idiotZboxer([]byte(p)))
    p = "aabcaaaabc"
    c.Expect(core.PrecalcZboxes([]byte(p)), ContainsExactly, idiotZboxer([]byte(p)))
    p = "aabcaaaabc*aabcaaaabc"
    c.Expect(core.PrecalcZboxes([]byte(p)), ContainsExactly, idiotZboxer([]byte(p)))
    p = "aabcaaaabc*aabcaaaabc*aabcaaaabc*aabcaaaabc*aabcaaaabc"
    c.Expect(core.PrecalcZboxes([]byte(p)), ContainsExactly, idiotZboxer([]byte(p)))
    p = "abcdefghijklmnopq"
    c.Expect(core.PrecalcZboxes([]byte(p)), ContainsExactly, idiotZboxer([]byte(p)))
    p = "abcdefghijklmnopqabcdefghijklmnopqabcdefghijklmnopq"
    c.Expect(core.PrecalcZboxes([]byte(p)), ContainsExactly, idiotZboxer([]byte(p)))
    p = "abcdefghijklmnopqabc efghijklmnopqab cdefghijklmnopq"
    c.Expect(core.PrecalcZboxes([]byte(p)), ContainsExactly, idiotZboxer([]byte(p)))
    p = "aaaaaaaaaaaaaaaaaa"
    c.Expect(core.PrecalcZboxes([]byte(p)), ContainsExactly, idiotZboxer([]byte(p)))
    p = "bbbbbbbbbbbaaaaaaaaaaaa"
    c.Expect(core.PrecalcZboxes([]byte(p)), ContainsExactly, idiotZboxer([]byte(p)))
    p = "bbbbbbbaaaaaaaaabbbbbbbbbbbbbbaaaabbbbbbbbbaaaaabbaaaaaaaaabbaa"
    c.Expect(core.PrecalcZboxes([]byte(p)), ContainsExactly, idiotZboxer([]byte(p)))
    p = "aabbaaa"
    c.Expect(core.PrecalcZboxes([]byte(p)), ContainsExactly, idiotZboxer([]byte(p)))
  })
}

func ZBoxReverseSpec(c gospec.Context) {
  c.Specify("Comprehensive test 3^9", func() {
    p := make([]byte, 9)
    for augment(p, 3) {
      c.Expect(core.PrecalcZboxesReversed(p), ContainsExactly, idiotZboxerReversed(p))
    }
  })
  c.Specify("Comprehensive test 2^15", func() {
    p := make([]byte, 15)
    for augment(p, 2) {
      c.Expect(core.PrecalcZboxesReversed(p), ContainsExactly, idiotZboxerReversed(p))
    }
  })
  c.Specify("Basic test.", func() {
    p := ""
    c.Expect(core.PrecalcZboxesReversed([]byte(p)), ContainsExactly, idiotZboxerReversed([]byte(p)))
    p = "a"
    c.Expect(core.PrecalcZboxesReversed([]byte(p)), ContainsExactly, idiotZboxerReversed([]byte(p)))
    p = "abcabc"
    c.Expect(core.PrecalcZboxesReversed([]byte(p)), ContainsExactly, idiotZboxerReversed([]byte(p)))
    p = "aabcabc"
    c.Expect(core.PrecalcZboxesReversed([]byte(p)), ContainsExactly, idiotZboxerReversed([]byte(p)))
    p = "aaabcabc"
    c.Expect(core.PrecalcZboxesReversed([]byte(p)), ContainsExactly, idiotZboxerReversed([]byte(p)))
    p = "aabcaabc"
    c.Expect(core.PrecalcZboxesReversed([]byte(p)), ContainsExactly, idiotZboxerReversed([]byte(p)))
    p = "aabcaaabc"
    c.Expect(core.PrecalcZboxesReversed([]byte(p)), ContainsExactly, idiotZboxerReversed([]byte(p)))
    p = "aabcaaaabc"
    c.Expect(core.PrecalcZboxesReversed([]byte(p)), ContainsExactly, idiotZboxerReversed([]byte(p)))
    p = "aabcaaaabc*aabcaaaabc"
    c.Expect(core.PrecalcZboxesReversed([]byte(p)), ContainsExactly, idiotZboxerReversed([]byte(p)))
    p = "aabcaaaabc*aabcaaaabc*aabcaaaabc*aabcaaaabc*aabcaaaabc"
    c.Expect(core.PrecalcZboxesReversed([]byte(p)), ContainsExactly, idiotZboxerReversed([]byte(p)))
    p = "abcdefghijklmnopq"
    c.Expect(core.PrecalcZboxesReversed([]byte(p)), ContainsExactly, idiotZboxerReversed([]byte(p)))
    p = "abcdefghijklmnopqabcdefghijklmnopqabcdefghijklmnopq"
    c.Expect(core.PrecalcZboxesReversed([]byte(p)), ContainsExactly, idiotZboxerReversed([]byte(p)))
    p = "abcdefghijklmnopqabc efghijklmnopqab cdefghijklmnopq"
    c.Expect(core.PrecalcZboxesReversed([]byte(p)), ContainsExactly, idiotZboxerReversed([]byte(p)))
    p = "aaaaaaaaaaaaaaaaaa"
    c.Expect(core.PrecalcZboxesReversed([]byte(p)), ContainsExactly, idiotZboxerReversed([]byte(p)))
    p = "bbbbbbbbbbbaaaaaaaaaaaa"
    c.Expect(core.PrecalcZboxesReversed([]byte(p)), ContainsExactly, idiotZboxerReversed([]byte(p)))
    p = "bbbbbbbaaaaaaaaabbbbbbbbbbbbbbaaaabbbbbbbbbaaaaabbaaaaaaaaabbaa"
    c.Expect(core.PrecalcZboxesReversed([]byte(p)), ContainsExactly, idiotZboxerReversed([]byte(p)))
    p = "aabbaaa"
    c.Expect(core.PrecalcZboxesReversed([]byte(p)), ContainsExactly, idiotZboxerReversed([]byte(p)))
  })
}

// BenchmarkBoyerMoore1* tests on a pattern and a text consisting of only one
// character.  This should be the worst case for a correct Boyer-Moore because
// it requires space to be allocated to return all of the matches.
func BenchmarkBoyerMoore1_10_100000(b *testing.B) {
  b.StopTimer()
  p := makeTestString1(10, 0)
  t := makeTestString1(100000, 0)
  bmd := core.BoyerMoorePreprocess(p)
  b.StartTimer()
  var matches []int
  for i := 0; i < b.N; i++ {
    matches = matches[0:0]
    core.BoyerMoore(bmd, t, &matches)
  }
}

func BenchmarkBoyerMoore1_100_100000(b *testing.B) {
  b.StopTimer()
  p := makeTestString1(100, 0)
  t := makeTestString1(100000, 0)
  bmd := core.BoyerMoorePreprocess(p)
  b.StartTimer()
  var matches []int
  for i := 0; i < b.N; i++ {
    matches = matches[0:0]
    core.BoyerMoore(bmd, t, &matches)
  }
}

// BenchmarkBoyerMoore2* tests make sure that the runtime does not go
// quadratic when search for something of the for abxb(ab)+ in (ab)+
func BenchmarkBoyerMoore2_10_100000(b *testing.B) {
  b.StopTimer()
  p := makeTestString3(10, 2)
  p[2] = 'x'
  t := makeTestString3(100000, 2)
  bmd := core.BoyerMoorePreprocess(p)
  b.StartTimer()
  var matches []int
  for i := 0; i < b.N; i++ {
    matches = matches[0:0]
    core.BoyerMoore(bmd, t, &matches)
  }
}

func BenchmarkBoyerMoore2_100_100000(b *testing.B) {
  b.StopTimer()
  p := makeTestString3(100, 2)
  p[2] = 'x'
  t := makeTestString3(100000, 2)
  bmd := core.BoyerMoorePreprocess(p)
  b.StartTimer()
  var matches []int
  for i := 0; i < b.N; i++ {
    matches = matches[0:0]
    core.BoyerMoore(bmd, t, &matches)
  }
}

// BenchmarkBoyerMoore3* tests for a pattern consisting only of characters
// that are not found in the text.  These tests expect sublinear time that
// should decrease with the size of the pattern.
func BenchmarkBoyerMoore3_10_100000(b *testing.B) {
  b.StopTimer()
  p := makeTestString1(10, 0)
  t := makeTestString1(100000, 1)
  bmd := core.BoyerMoorePreprocess(p)
  b.StartTimer()
  var matches []int
  for i := 0; i < b.N; i++ {
    matches = matches[0:0]
    core.BoyerMoore(bmd, t, &matches)
  }
}

func BenchmarkBoyerMoore3_100_100000(b *testing.B) {
  b.StopTimer()
  p := makeTestString1(100, 0)
  t := makeTestString1(100000, 1)
  bmd := core.BoyerMoorePreprocess(p)
  b.StartTimer()
  var matches []int
  for i := 0; i < b.N; i++ {
    matches = matches[0:0]
    core.BoyerMoore(bmd, t, &matches)
  }
}

// BenchmarkBoyerMoore4* tests on random string with an alphabet size of 20,
// like protein
func BenchmarkBoyerMoore4_10_100000(b *testing.B) {
  b.StopTimer()
  p := makeTestString4(10, 20, 0)
  t := makeTestString4(100000, 20, 1)
  bmd := core.BoyerMoorePreprocess(p)
  b.StartTimer()
  var matches []int
  for i := 0; i < b.N; i++ {
    matches = matches[0:0]
    core.BoyerMoore(bmd, t, &matches)
  }
}

func BenchmarkBoyerMoore4_100_100000(b *testing.B) {
  b.StopTimer()
  p := makeTestString4(100, 20, 0)
  t := makeTestString4(100000, 20, 1)
  bmd := core.BoyerMoorePreprocess(p)
  b.StartTimer()
  var matches []int
  for i := 0; i < b.N; i++ {
    matches = matches[0:0]
    core.BoyerMoore(bmd, t, &matches)
  }
}

// BenchmarkBoyerMoore5* tests on random string with an alphabet size of 4,
// like DNA
func BenchmarkBoyerMoore5_10_100000(b *testing.B) {
  b.StopTimer()
  p := makeTestString4(10, 4, 0)
  t := makeTestString4(100000, 4, 1)
  bmd := core.BoyerMoorePreprocess(p)
  b.StartTimer()
  var matches []int
  for i := 0; i < b.N; i++ {
    matches = matches[0:0]
    core.BoyerMoore(bmd, t, &matches)
  }
}

func BenchmarkBoyerMoore5_100_100000(b *testing.B) {
  b.StopTimer()
  p := makeTestString4(100, 4, 0)
  t := makeTestString4(100000, 4, 1)
  bmd := core.BoyerMoorePreprocess(p)
  b.StartTimer()
  var matches []int
  for i := 0; i < b.N; i++ {
    matches = matches[0:0]
    core.BoyerMoore(bmd, t, &matches)
  }
}

// BenchmarkBoyerMoore5* tests on random string with an alphabet size of 4,
// like DNA, using an io.Reader instead of a slice of bytes
func BenchmarkBoyerMooreReader5_10_100000(b *testing.B) {
  b.StopTimer()
  p := makeTestString4(10, 4, 0)
  t := makeTestString4(1000, 4, 1)
  bmd := core.BoyerMoorePreprocess(p)
  b.StartTimer()
  var matches []int
  buf := make([]byte, 5000)
  for i := 0; i < b.N; i++ {
    matches = matches[0:0]
    b.StopTimer()
    in := bytes.NewBuffer(t)
    b.StartTimer()
    core.BoyerMooreFromReader(bmd, in, buf, &matches)
  }
}

func BenchmarkBoyerMooreReader5_100_100000(b *testing.B) {
  b.StopTimer()
  p := makeTestString4(10, 4, 0)
  t := makeTestString4(10000, 4, 1)
  bmd := core.BoyerMoorePreprocess(p)
  b.StartTimer()
  var matches []int
  buf := make([]byte, 5000)
  for i := 0; i < b.N; i++ {
    matches = matches[0:0]
    b.StopTimer()
    in := bytes.NewBuffer(t)
    b.StartTimer()
    core.BoyerMooreFromReader(bmd, in, buf, &matches)
  }
}

func LongestSuffixAsPrefixSpec(c gospec.Context) {
  fmt.Printf("")
  c.Specify("Comprehensive test 2^15", func() {
    p := make([]byte, 15)
    for augment(p, 2) {
      c.Expect(core.LongestSuffixAsPrefix(p), ContainsExactly, idiotLongestSuffixAsPrefix(p))
    }
  })
  c.Specify("Comprehensive test 3^9", func() {
    p := make([]byte, 9)
    for augment(p, 3) {
      c.Expect(core.LongestSuffixAsPrefix(p), ContainsExactly, idiotLongestSuffixAsPrefix(p))
    }
  })
  c.Specify("Comprehensive test 4^7", func() {
    p := make([]byte, 7)
    for augment(p, 4) {
      c.Expect(core.LongestSuffixAsPrefix(p), ContainsExactly, idiotLongestSuffixAsPrefix(p))
    }
  })
}

func BoyerMooreSpec(c gospec.Context) {
  c.Specify("Basic test", func() {
    p := []byte("a")
    t := []byte("aaaaaaaaaa")
    var matches []int
    core.BoyerMoore(core.BoyerMoorePreprocess(p), t, &matches)
    c.Expect(matches, ContainsExactly, idiotStringSearch(p, t))
    matches = matches[0:0]
    p = []byte("aa")
    t = []byte("aaaaaaaaaa")
    core.BoyerMoore(core.BoyerMoorePreprocess(p), t, &matches)
    c.Expect(matches, ContainsExactly, idiotStringSearch(p, t))
    matches = matches[0:0]
    p = []byte("aaaaaaaaaa")
    t = []byte("aaaaaaaaaa")
    core.BoyerMoore(core.BoyerMoorePreprocess(p), t, &matches)
    c.Expect(matches, ContainsExactly, idiotStringSearch(p, t))
    matches = matches[0:0]
    p = []byte("aaaaaaaaaaaaaaaaaaaaa")
    t = []byte("aaaaaaaaaa")
    core.BoyerMoore(core.BoyerMoorePreprocess(p), t, &matches)
    c.Expect(matches, ContainsExactly, idiotStringSearch(p, t))
    matches = matches[0:0]
    p = []byte("b")
    t = []byte("aaaabaaaaa")
    core.BoyerMoore(core.BoyerMoorePreprocess(p), t, &matches)
    c.Expect(matches, ContainsExactly, idiotStringSearch(p, t))
    matches = matches[0:0]
    p = []byte("ba")
    t = []byte("aaaabaaaaa")
    core.BoyerMoore(core.BoyerMoorePreprocess(p), t, &matches)
    c.Expect(matches, ContainsExactly, idiotStringSearch(p, t))
    matches = matches[0:0]
    p = []byte("aaaabaaaaa")
    t = []byte("aaaabaaaaa")
    core.BoyerMoore(core.BoyerMoorePreprocess(p), t, &matches)
    c.Expect(matches, ContainsExactly, idiotStringSearch(p, t))
    matches = matches[0:0]
    p = []byte("aaaaaaaaaaaaaaaaabaaa")
    t = []byte("aaaaaabaaa")
    core.BoyerMoore(core.BoyerMoorePreprocess(p), t, &matches)
    c.Expect(matches, ContainsExactly, idiotStringSearch(p, t))
    matches = matches[0:0]
  })

  c.Specify("Comprehensive test 2^17", func() {
    b := make([]byte, 17)
    var matches []int
    for augment(b, 2) {
      p := b[0:5]
      t := b[5:]
      bmd := core.BoyerMoorePreprocess(p)
      matches = matches[0:0]
      core.BoyerMoore(bmd, t, &matches)
      c.Expect(matches, ContainsExactly, idiotStringSearch(p, t))
    }
  })

  c.Specify("Comprehensive test 3^11", func() {
    b := make([]byte, 11)
    var matches []int
    for augment(b, 3) {
      p := b[0:4]
      t := b[4:]
      bmd := core.BoyerMoorePreprocess(p)
      matches = matches[0:0]
      core.BoyerMoore(bmd, t, &matches)
      c.Expect(matches, ContainsExactly, idiotStringSearch(p, t))
    }
  })

  c.Specify("Random test", func() {
    var matches []int
    for i := 0; i < 10000; i += 2 {
      p := makeTestString4(15, 7, i)
      t := makeTestString4(1000, 7, i+1)
      bmd := core.BoyerMoorePreprocess(p)
      matches = matches[0:0]
      core.BoyerMoore(bmd, t, &matches)
      c.Expect(matches, ContainsExactly, idiotStringSearch(p, t))
    }
  })
}

func BoyerMooreReaderSpec(c gospec.Context) {
  c.Specify("Basic test", func() {
    var matches []int
    buffer := make([]byte, 500)
    p := []byte("a")
    t := []byte("aaaaaaaaa")
    bmd := core.BoyerMoorePreprocess(p)
    matches = matches[0:0]
    core.BoyerMooreFromReader(bmd, bytes.NewBuffer(t), buffer, &matches)
    c.Expect(matches, ContainsExactly, idiotStringSearch(p, t))
    p = []byte("aa")
    t = []byte("aaaaaaaaaa")
    core.BoyerMoorePreprocess(p)
    bmd = core.BoyerMoorePreprocess(p)
    matches = matches[0:0]
    core.BoyerMooreFromReader(bmd, bytes.NewBuffer(t), buffer, &matches)
    c.Expect(matches, ContainsExactly, idiotStringSearch(p, t))
    p = []byte("aaaaaaaaaa")
    t = []byte("aaaaaaaaaa")
    core.BoyerMoorePreprocess(p)
    bmd = core.BoyerMoorePreprocess(p)
    matches = matches[0:0]
    core.BoyerMooreFromReader(bmd, bytes.NewBuffer(t), buffer, &matches)
    c.Expect(matches, ContainsExactly, idiotStringSearch(p, t))
    p = []byte("aaaaaaaaaaaaaaaaaaaaa")
    t = []byte("aaaaaaaaaa")
    core.BoyerMoorePreprocess(p)
    bmd = core.BoyerMoorePreprocess(p)
    matches = matches[0:0]
    core.BoyerMooreFromReader(bmd, bytes.NewBuffer(t), buffer, &matches)
    c.Expect(matches, ContainsExactly, idiotStringSearch(p, t))
    p = []byte("b")
    t = []byte("aaaabaaaaa")
    core.BoyerMoorePreprocess(p)
    bmd = core.BoyerMoorePreprocess(p)
    matches = matches[0:0]
    core.BoyerMooreFromReader(bmd, bytes.NewBuffer(t), buffer, &matches)
    c.Expect(matches, ContainsExactly, idiotStringSearch(p, t))
    p = []byte("ba")
    t = []byte("aaaabaaaaa")
    core.BoyerMoorePreprocess(p)
    bmd = core.BoyerMoorePreprocess(p)
    matches = matches[0:0]
    core.BoyerMooreFromReader(bmd, bytes.NewBuffer(t), buffer, &matches)
    c.Expect(matches, ContainsExactly, idiotStringSearch(p, t))
    p = []byte("aaaabaaaaa")
    t = []byte("aaaabaaaaa")
    core.BoyerMoorePreprocess(p)
    bmd = core.BoyerMoorePreprocess(p)
    matches = matches[0:0]
    core.BoyerMooreFromReader(bmd, bytes.NewBuffer(t), buffer, &matches)
    c.Expect(matches, ContainsExactly, idiotStringSearch(p, t))
    p = []byte("aaaaaaaaaaaaaaaaabaaa")
    t = []byte("aaaaaabaaa")
    core.BoyerMoorePreprocess(p)
    bmd = core.BoyerMoorePreprocess(p)
    matches = matches[0:0]
    core.BoyerMooreFromReader(bmd, bytes.NewBuffer(t), buffer, &matches)
    c.Expect(matches, ContainsExactly, idiotStringSearch(p, t))
  })

  c.Specify("Comprehensive test 2^17", func() {
    var matches []int
    buffer := make([]byte, 15)
    b := make([]byte, 17)
    for augment(b, 2) {
      p := b[0:5]
      t := b[5:]
      bmd := core.BoyerMoorePreprocess(p)
      matches = matches[0:0]
      core.BoyerMooreFromReader(bmd, bytes.NewBuffer(t), buffer, &matches)
      c.Expect(matches, ContainsExactly, idiotStringSearch(p, t))
    }
  })

  c.Specify("Comprehensive test 3^11", func() {
    var matches []int
    buffer := make([]byte, 15)
    b := make([]byte, 11)
    for augment(b, 3) {
      p := b[0:4]
      t := b[4:]
      bmd := core.BoyerMoorePreprocess(p)
      matches = matches[0:0]
      core.BoyerMooreFromReader(bmd, bytes.NewBuffer(t), buffer, &matches)
      c.Expect(matches, ContainsExactly, idiotStringSearch(p, t))
    }
  })

  c.Specify("Random test", func() {
    var matches []int
    buffer := make([]byte, 15)
    for i := 0; i < 10000; i += 2 {
      p := makeTestString4(15, 7, i)
      t := makeTestString4(1000, 7, i+1)
      bmd := core.BoyerMoorePreprocess(p)
      matches = matches[0:0]
      core.BoyerMooreFromReader(bmd, bytes.NewBuffer(t), buffer, &matches)
      c.Expect(matches, ContainsExactly, idiotStringSearch(p, t))
    }
  })
}
