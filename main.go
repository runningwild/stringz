package main

import (
  "fmt"
)

func precalcZboxes(p string) []int {
  if len(p) == 0 { return nil }
  if len(p) == 1 { return []int{ len(p) } }
  pos := 1
  for pos < len(p) && p[pos] == p[pos - 1] {
    pos++
  }

  zs := make([]int, len(p))
  zs[0] = len(p)
  zs[1] = pos - 1

  left := 1
  right := pos - 1
  for i := 2; i < len(zs); i++ {
    if right < i {
      // We just left a zbox - so we need to see how much of a prefix we have
      // from this position
      pos := i
      for pos < len(p) && p[pos] == p[pos - i] {
        pos++
      }
      zs[i] = pos - i
      left = i
      right = pos - i
    } else {
      j := i - left
      rem := right - i
      zj := zs[j]
      if zj < rem {
        // The old z-value shows us that we have a prefix here that is less
        // than the length remaining in out current z-box, so we use that
        // z-value and we're done.
        zs[i] = zj
      } else {
        // We are at a prefix now that goes outside of the current z-box, so
        // we need to find how far that is, but we don't need to start
        // comparing until the end of this prefix.
        pos := i + zj
        for pos < len(p) && p[pos] == p[pos - (i + zj)] {
          pos++
        }
        zs[i] = pos - i
        left = i
        right = pos - i
      }
    }
  }
  return zs
}

func main() {
  fmt.Printf("yo\n")
  pb := make([]byte, 10000000)
  for i := range pb {
    pb[i] = 'a'
  }
  precalcZboxes(string(pb))
  // for i := range zs {
  //   fmt.Printf("%c:\t%d\n", p[i], zs[i])
  // }
}