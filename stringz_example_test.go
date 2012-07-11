package stringz_test

import (
  "fmt"
  "github.com/runningwild/stringz"
)

func ExampleStringFinder() {
  p := []byte("foo")
  t1 := []byte("bar foo wing ding foo monkey")
  t2 := []byte("monkey foo ding wing foo bar")

  // If we want to find all occurences of p in t1 it is as simple as this:
  stringz.Find(p).In(t1)

  // We can also store the StringFinder for use later so we don't incur the
  // overhead of preprocessing p multiple times.
  find_foo := stringz.Find(p)
  hits1 := find_foo.In(t1)
  hits2 := find_foo.In(t2)
  fmt.Printf("%v %v\n", hits1, hits2)
  // Output:
  // [4 18] [7 21]
}

func ExampleStringSetFinder() {
  ps := [][]byte{
    []byte("foo"),
    []byte("bar"),
    []byte("wing"),
    []byte("ding"),
  }
  t1 := []byte("bar foo wing ding foo monkey")
  t2 := []byte("monkey foo ding wing foo bar")

  // If we want to find all occurences of any element in ps in t1 it is as
  // simple as this:
  stringz.FindSet(ps).In(t1)

  // We can also store the StringSetFinder for use later so we don't incur the
  // overhead of preprocessing ps multiple times.
  find := stringz.FindSet(ps)
  hits1 := find.In(t1)
  hits2 := find.In(t2)
  fmt.Printf("Hits1:\n")
  for i := 0; i < len(hits1); i++ {
    fmt.Printf("\"%s\": %v\n", ps[i], hits1[i])
  }
  fmt.Printf("\nHits2:\n")
  for i := 0; i < len(hits2); i++ {
    fmt.Printf("\"%s\": %v\n", ps[i], hits2[i])
  }
  // Output:
  // Hits1:
  // "foo": [4 18]
  // "bar": [0]
  // "wing": [8]
  // "ding": [13]

  // Hits2:
  // "foo": [7 21]
  // "bar": [25]
  // "wing": [16]
  // "ding": [11]
}
