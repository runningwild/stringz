// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example demonstrates decoding a JPEG image and examining its pixels.
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
	fmt.Printf("%v\n%v\n", hits1, hits2)
	// Output:
	// [[4 18] [0] [8] [13]]
	// [[7 21] [25] [16] [11]]
}
