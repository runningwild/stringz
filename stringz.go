// Implementation of several algorithms for doing pattern matching with
// strings.  Preprocessing data can be maintained so that the cost of doing
// the preprocessing is not incurred multiple times for the same string or set
// of strings.
//
// This file is a high level interface to these algorithms.  This should be
// sufficient for most applications, but if necessary the underlying
// algorithms are exposed in github.com/runningwild/stringz/core.
package stringz

import (
  "github.com/runningwild/stringz/core"
  "io"
)

// This was just an experiment, probably shouldn't publish it yet
// type SerialStringFinder struct {
//   bmd core.BmData
//   buf []byte
//   res []int
// }

// func SerialFind(p []byte) *SerialStringFinder {
//   return &SerialStringFinder{
//     bmd: core.BoyerMoorePreprocess(p),
//   }
// }
// func (ssf *SerialStringFinder) In(t []byte) []int {
//   ssf.res = ssf.res[0:0]
//   core.BoyerMoore(ssf.bmd, t, &ssf.res)
//   return ssf.res
// }

type StringFinder struct {
  bmd core.BmData
}

// Preprocesses p and returns a *StringFinder that can be used to quickly
// search for occurrences of p in other strings.  Uses Boyer-Moore, which
// requires O(n) time to preprocess p, and O(n) space to store the result.
// Methods on StringFinder can be called concurrently from multiple
// go-routines.
func Find(p []byte) *StringFinder {
  return &StringFinder{bmd: core.BoyerMoorePreprocess(p)}
}

// Searches t for the pattern, p, that was used to create the StringFinder.
// Returns a list of all indices at which p occurs in t, including overlaps,
// and in ascending order.  The search takes O(m) time in the worst case, and
// O(m/n) in the best case, where m and n are the lengths of t and p,
// respectively.  The search requires O(k) space, where k is the number of
// times p occurs in t.
func (sf *StringFinder) In(t []byte) []int {
  var res []int
  core.BoyerMoore(sf.bmd, t, &res)
  return res
}

// Like In(), but searches the data from a Reader instead of a []byte.
func (sf *StringFinder) InReader(r io.Reader) []int {
  var res []int
  core.BoyerMooreFromReader(sf.bmd, r, make([]byte, 100000), &res)
  return res
}

type StringSetFinder struct {
  acd core.AcData
}

// Preprocesses ps and returns a *StringSetFinder that can be used to quickly
// search for all occurrences of all elements of ps in other strings.  Uses
// Aho-Corasick, which requires O(n) time to preprocess ps, and O(n) to store
// the result, where n is the sum of the lengths of all of the elements in ps.
// Methods on StringSetFinder can be called concurrently from multiple
// go-routines.
func FindSet(ps [][]byte) *StringSetFinder {
  return &StringSetFinder{acd: core.AhoCorasickPreprocess(ps)}
}

// Searches t for all patterns in the set of patterns, ps, that was used to
// create the StringSetFinder.  Returns a list, H, such that H[i] is a list of
// every index in t at which ps[i] occurs.  Each H[i] includes overlaps and is
// in ascending order.  The search takes O(m) time and O(k) space, where m is
// the length of t, and k is the total number of occurrences of all elements
// of ps in t.
func (ssf *StringSetFinder) In(t []byte) map[int][]int {
  return core.AhoCorasick(ssf.acd, t)
}

// Like In(), but searches the data from a Reader instead of a []byte.
func (sf *StringSetFinder) InReader(input io.Reader) map[int][]int {
  return core.AhoCorasickFromReader(sf.acd, input, 100000)
}
