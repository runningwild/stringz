package stringz_test

import (
  "github.com/orfjackal/gospec/src/gospec"
  "math/rand"
  "testing"
)

// List of all specs here
func TestAllSpecs(t *testing.T) {
  r := gospec.NewRunner()
  // r.AddSpec(ZBoxSpec)
  // r.AddSpec(ZBoxReverseSpec)
  // r.AddSpec(LongestSuffixAsPrefixSpec)
  // r.AddSpec(BoyerMooreSpec)
  r.AddSpec(AhoCorasickSpec)
  gospec.MainGoTest(r, t)
}

// The rest of this file is utility functions for testing

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
  for i := n / 2; i < n; i++ {
    b[i] = 1
  }
  return string(b)
}

// Returns a string of length n, cycling through the number 0-(r-1)
func makeTestString3(n, r int) string {
  b := make([]byte, n)
  for i := range b {
    b[i] = byte(i % r)
  }
  return string(b)
}

// Returns a string of length n consisting of random characters less than r,
// and using seed s
func makeTestString4(n, r, s int) string {
  rand.Seed(int64(s))
  b := make([]byte, n)
  for i := range b {
    b[i] = byte(rand.Intn(256) % r)
  }
  return string(b)
}

func augment(b []byte, radix int) bool {
  for i := range b {
    if int(b[i]) < radix-1 {
      b[i]++
      return true
    } else {
      b[i] = 0
    }
  }
  return false
}
