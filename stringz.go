package stringz

import (
  "github.com/runningwild/stringz/core"
)

// Sample usage:
// p := []byte("foo")
// t := []byte("bar wing foo ding foo mo cake")
// stringz.Find(p).In(t)
type StringFinder struct {
  bmd core.BmData
}

func Find(p []byte) *StringFinder {
  return &StringFinder{bmd: core.BoyerMoorePreprocess(p)}
}

func (sf *StringFinder) In(t []byte) []int {
  return core.BoyerMoore(sf.bmd, t)
}

// Sample usage:
// ps := [][]byte{ []byte("foo"), []byte("wing"), []byte("ding") }
// t := []byte("bar wing foo ding foo mo cake")
// stringz.FindSet`(p).In(t)
type StringSetFinder struct {
  acd core.AcData
}

func FindSet(ps [][]byte) *StringSetFinder {
  return &StringSetFinder{acd: core.AhoCorasickPreprocess(ps)}
}

func (ssf *StringSetFinder) In(t []byte) [][]int {
  return core.AhoCorasick(ssf.acd, t)
}
