package stringz

import (
  "fmt"
)

func PrecalcZboxes(p string) []int {
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
      right = pos - 1
    } else {
      j := i - left
      rem := right - i + 1
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
        pos := right + 1
        for pos < len(p) && p[pos] == p[j + (pos - i) - 1] {
          pos++
        }
        left = i
        right = pos - 1
        zs[i] = right - left + 1
      }
    }
  }
  return zs
}

func precalcZboxesReversed(p string) []int {
  n := len(p)
  if n == 0 { return nil }
  if n == 1 { return []int{ n } }
  pos := n - 2
  for pos >= 0 && p[pos] == p[pos + 1] {
    fmt.Printf("cmp: %c %c\n", p[pos], p[pos+1])
    pos--
  }

  zs := make([]int, n)
  zs[n - 1] = n
  zs[n - 2] = n - pos - 2

  left := pos + 1
  right := n - 2
  fmt.Printf("l/r: %d/%d\n", left, right)
  for i := n - 3; i >= 0; i-- {
    fmt.Printf("i: %d\n", i)
    fmt.Printf("left/right: %d/%d\n", left, right)
    if left > i {
      fmt.Printf("left\n")
      // We just left a zbox - so we need to see how much of a suffix we have
      // from this position
      pos := i
      for pos >= 0 && p[pos] == p[n - (i - pos) - 1] {
        pos--
      }
      zs[i] = i - pos
      left = pos
      right = i
    } else {
      fmt.Printf("else\n")
      j := n - (right - i) - 1
      rem := i - left
      zj := zs[j]
      if zj < rem {
      fmt.Printf("else2\n")
        // The old z-value shows us that we have a prefix here that is less
        // than the length remaining in out current z-box, so we use that
        // z-value and we're done.
        zs[i] = zj
      } else {
      fmt.Printf("else3\n")
        // We are at a prefix now that goes outside of the current z-box, so
        // we need to find how far that is, but we don't need to start
        // comparing until the end of this prefix.
        pos := left - 1
        fmt.Printf("%d %d %d\n", left, right, pos)
        for pos >= 0 && p[pos] == p[n - right + pos] {
          pos--
        }
        zs[i] = i - pos
        left = pos
        right = i
      }
    }
  }
  return zs
}

// Might want to have a specialized boyer-moore for small alphabets like
// ascii, dna, protein.
func boyerMooreExtendedBadCharacterRule(p string) map[byte][]int {
  m := make(map[byte][]int)
  for i := len(p) - 1; i >= 0; i-- {
    r := p[i]
    m[r] = append(m[r], i)
  }
  return m
}

func boyerMooreStrongGoodSuffixRule(p string) []int {
  L := make([]int, len(p))
  for i := range L {
    j := len(L) - L[i] - 1
    L[j] = i
  }
  return L
}

func revString(s string) string {
  b := []byte(s)
  for i := 0; i < len(b) / 2; i++ {
    b[i], b[len(b) - i - 1] = b[len(b) - i -1], b[i]
  }
  return string(b)
}
