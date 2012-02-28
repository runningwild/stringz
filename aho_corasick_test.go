package stringz_test

import (
  . "github.com/orfjackal/gospec/src/gospec"
  "github.com/orfjackal/gospec/src/gospec"
  "github.com/runningwild/stringz"
  "testing"
)

func idiotAhoCorasick(ps [][]byte, t []byte) [][]int {
  var res [][]int
  for i := range ps {
    res = append(res, idiotStringSearch(string(ps[i]), string(t)))
  }
  return res
}

func AhoCorasickSpec(c gospec.Context) {
  c.Specify("Basic test", func() {
    strs := [][]byte{
      []byte("baa"),
      []byte("anba"),
      []byte("banana"),
    }
    res := stringz.AhoCorasick(strs, []byte("baababanananba"))
    c.Expect(res[0], ContainsExactly, []int{0})
    c.Expect(res[1], ContainsExactly, []int{10})
    c.Expect(res[2], ContainsExactly, []int{5})
  })

  c.Specify("Substrings test", func() {
    strs := [][]byte{
      []byte("baa"),
      []byte("aab"),
      []byte("aaa"),
      []byte("aa"),
    }
    res := stringz.AhoCorasick(strs, []byte("abbaababbbbaaaaaabbbbaabaabaabbb"))
    c.Expect(res[0], ContainsExactly, []int{2, 10, 20, 23, 26})
    c.Expect(res[1], ContainsExactly, []int{3, 15, 21, 24, 27})
    c.Expect(res[2], ContainsExactly, []int{11, 12, 13, 14})
    c.Expect(res[3], ContainsExactly, []int{3, 11, 12, 13, 14, 15, 21, 24, 27})
  })

  c.Specify("Comprehensive 2^4 x 2^12 test", func() {
    var ps [][]byte
    b := make([]byte, 4)
    for augment(b, 2) {
      p := make([]byte, 4)
      copy(p, b)
      ps = append(ps, p)
    }
    t := make([]byte, 12)
    for augment(t, 2) {
      acres := stringz.AhoCorasick(ps, t)
      ires := idiotAhoCorasick(ps, t)
      for i := range ps {
        c.Expect(acres[i], ContainsExactly, ires[i])
      }
    }
  })

  c.Specify("Comprehensive 3^3 x 3^8 test", func() {
    var ps [][]byte
    b := make([]byte, 3)
    for augment(b, 3) {
      p := make([]byte, 3)
      copy(p, b)
      ps = append(ps, p)
    }
    t := make([]byte, 8)
    for augment(t, 3) {
      acres := stringz.AhoCorasick(ps, t)
      ires := idiotAhoCorasick(ps, t)
      for i := range ps {
        c.Expect(acres[i], ContainsExactly, ires[i])
      }
    }
  })

  c.Specify("Larger alphabet test", func() {
    ps := [][]byte{
      []byte(makeTestString4(4, 10, 0)),
      []byte(makeTestString4(5, 10, 1)),
      []byte(makeTestString4(6, 10, 2)),
      []byte(makeTestString4(7, 10, 3)),
      []byte(makeTestString4(6, 10, 4)),
      []byte(makeTestString4(5, 10, 5)),
      []byte(makeTestString4(4, 10, 6)),
      []byte(makeTestString4(3, 10, 7)),
      []byte(makeTestString4(4, 10, 8)),
      []byte(makeTestString4(5, 10, 9)),
    }
    for seed := 10; seed < 30; seed++ {
      t := []byte(makeTestString4(10000, 11, seed))
      acres := stringz.AhoCorasick(ps, t)
      ires := idiotAhoCorasick(ps, t)
      for i := range ps {
        c.Expect(acres[i], ContainsExactly, ires[i])
      }
    }
  })
}

func BenchmarkAhoCorasick4_10x10_100000(b *testing.B) {
  b.StopTimer()
  ps := [][]byte{
    []byte(makeTestString4(5, 10, 0)),
    []byte(makeTestString4(5, 10, 1)),
    []byte(makeTestString4(5, 10, 2)),
    []byte(makeTestString4(5, 10, 3)),
    []byte(makeTestString4(5, 10, 4)),
    []byte(makeTestString4(5, 10, 5)),
    []byte(makeTestString4(5, 10, 6)),
    []byte(makeTestString4(5, 10, 7)),
    []byte(makeTestString4(5, 10, 8)),
    []byte(makeTestString4(5, 10, 9)),
  }
  t := []byte(makeTestString4(100000, 10, 10))
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    stringz.AhoCorasick(ps, t)
  }
}

func BenchmarkAhoCorasick4_100x10_100000(b *testing.B) {
  b.StopTimer()
  ps := make([][]byte, 100)
  for i := range ps {
    ps[i] = []byte(makeTestString4(5, 10, i))
  }
  t := []byte(makeTestString4(100000, 10, len(ps)))
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    stringz.AhoCorasick(ps, t)
  }
}

func BenchmarkAhoCorasick4_10x10_1000000(b *testing.B) {
  b.StopTimer()
  ps := [][]byte{
    []byte(makeTestString4(5, 10, 0)),
    []byte(makeTestString4(5, 10, 1)),
    []byte(makeTestString4(5, 10, 2)),
    []byte(makeTestString4(5, 10, 3)),
    []byte(makeTestString4(5, 10, 4)),
    []byte(makeTestString4(5, 10, 5)),
    []byte(makeTestString4(5, 10, 6)),
    []byte(makeTestString4(5, 10, 7)),
    []byte(makeTestString4(5, 10, 8)),
    []byte(makeTestString4(5, 10, 9)),
  }
  t := []byte(makeTestString4(1000000, 10, 10))
  b.StartTimer()
  for i := 0; i < b.N; i++ {
    stringz.AhoCorasick(ps, t)
  }
}

