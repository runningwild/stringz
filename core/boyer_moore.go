package core

import (
  "io"
)

func PrecalcZboxes(p []byte) []int {
  if len(p) == 0 {
    return nil
  }
  if len(p) == 1 {
    return []int{len(p)}
  }
  pos := 1
  for pos < len(p) && p[pos] == p[pos-1] {
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
      for pos < len(p) && p[pos] == p[pos-i] {
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
        cmp := pos - i
        for pos < len(p) && p[pos] == p[cmp] {
          pos++
          cmp++
        }
        left = i
        right = pos - 1
        zs[i] = right - left + 1
      }
    }
  }
  return zs
}

// Returns l such that l[i] is the length of the longest suffix of p[1:]
// that is a prefix of p, 0 if such a suffix does not exist.
func LongestSuffixAsPrefix(p []byte) []int {
  if len(p) == 0 {
    return nil
  }
  if len(p) == 1 {
    return []int{len(p)}
  }
  pos := 1
  for pos < len(p) && p[pos] == p[pos-1] {
    pos++
  }

  zs := make([]int, len(p))
  zs[0] = len(p)
  zs[1] = pos - 1
  ps := make([]int, len(p))
  ps[0] = zs[0]
  if pos == len(p) {
    ps[1] = zs[1]
  }

  left := 1
  right := pos - 1
  for i := 2; i < len(zs); i++ {
    if right < i {
      // We just left a zbox - so we need to see how much of a prefix we have
      // from this position
      pos := i
      for pos < len(p) && p[pos] == p[pos-i] {
        pos++
      }
      left = i
      right = pos - 1
      zs[i] = pos - i
      if pos == len(p) {
        ps[i] = zs[i]
      }
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
        cmp := pos - i
        for pos < len(p) && p[pos] == p[cmp] {
          pos++
          cmp++
        }
        left = i
        right = pos - 1

        // On the very last index we might set pos = right + 1 == len(p) + 1,
        // so we need to compare with >= or we could miss that.
        zs[i] = right - left + 1
        if pos >= len(p) {
          ps[i] = zs[i]
        }
      }
    }
  }
  for i := len(ps) - 2; i >= 0; i-- {
    if ps[i] < ps[i+1] {
      ps[i] = ps[i+1]
    }
  }
  return ps
}

func PrecalcZboxesReversed(p []byte) []int {
  n := len(p)
  if n == 0 {
    return nil
  }
  if n == 1 {
    return []int{n}
  }
  pos := n - 2
  for pos >= 0 && p[pos] == p[pos+1] {
    pos--
  }

  zs := make([]int, n)
  zs[n-1] = n
  zs[n-2] = n - pos - 2

  left := pos + 1
  right := n - 2
  for i := n - 3; i >= 0; i-- {
    if left > i {
      // We just left a zbox - so we need to see how much of a suffix we have
      // from this position
      pos := i
      for pos >= 0 && p[pos] == p[n-(i-pos)-1] {
        pos--
      }
      zs[i] = i - pos
      left = pos + 1
      right = i
    } else {
      j := right - i
      rem := i - left + 1
      zj := zs[n-j-1]
      if zj < rem {
        // The old z-value shows us that we have a prefix here that is less
        // than the length remaining in out current z-box, so we use that
        // z-value and we're done.
        zs[i] = zj
      } else {
        // We are at a prefix now that goes outside of the current z-box, so
        // we need to find how far that is, but we don't need to start
        // comparing until the end of this prefix.
        pos := left - 1
        cmp := n - (i - pos) - 1
        for pos >= 0 && p[pos] == p[cmp] {
          pos--
          cmp--
        }
        left = pos + 1
        right = i
        zs[i] = right - left + 1
      }
    }
  }
  return zs
}

// Requires knowledge of the alphabet size so that we avoid doing hashtable
// lookups.
func boyerMooreExtendedBadCharacterRule(p []byte) [256][]int {
  var m [256][]int
  for i := len(p) - 1; i >= 0; i-- {
    r := p[i]
    m[r] = append(m[r], i)
  }
  return m
}

// L here is not exactly as specified in Gusfield, since the value used when
// determining shifts is always n - L'(i) we just do that math now rather than
// later.
func BoyerMooreStrongGoodSuffixRule(p []byte) (L, l []int) {
  Z := PrecalcZboxesReversed(p)
  L = make([]int, len(p))
  for i := 0; i < len(Z)-1; i++ {
    if Z[i] == 0 {
      continue
    }
    L[len(p)-Z[i]-1] = len(p) - i - 1
  }

  l = LongestSuffixAsPrefix(p)
  for i := 0; i < len(l)-1; i++ {
    l[i] = len(p) - l[i+1]
  }
  l[len(l)-1] = 0
  return
}

type BmData struct {
  // Copy of the pattern
  p []byte

  // Output from BoyreMooreStrongGoodSuffixRule
  L, l []int

  // Output from boyerMooreExtendedBadCharacterRule
  R [256][]int
}

func BoyerMoorePreprocess(p []byte) BmData {
  var bmd BmData
  bmd.p = make([]byte, len(p))
  copy(bmd.p, p)
  bmd.L, bmd.l = BoyerMooreStrongGoodSuffixRule(p)
  bmd.R = boyerMooreExtendedBadCharacterRule(p)
  return bmd
}

func BoyerMoore(bmd BmData, t []byte) []int {
  var matches []int
  k := len(bmd.L) - 1

  // In some cases we don't need to go all the way to the left-most character
  // since we might know that a certain prefix of the current alignment
  // matches based on a previous test.
  min := 0

  for k < len(t) {
    i := len(bmd.L) - 1
    h := k
    for i >= min && bmd.p[i] == t[h] {
      i--
      h--
    }

    if i < min {
      // found a match
      matches = append(matches, k-len(bmd.L)+1)
      if bmd.l[0] > 0 {
        k += bmd.l[0]

        // Since we matched we will know some prefix of the next alignment.
        min = len(bmd.L) - bmd.l[0]
      } else {
        k++
        min = 0
      }
    } else {
      shift := 0

      // Strong good suffix rule
      if bmd.L[i] == 0 {
        shift = bmd.l[i]

        // This shift can place part of what already matched as a prefix.
        min = len(bmd.L) - bmd.l[i]
      } else {
        shift = bmd.L[i]
        min = 0
      }

      // Extended bad character rule
      bc := bmd.R[t[h]]
      if len(bc) == 0 {
        shift = i + 1
        min = 0
      } else {
        for j := range bc {
          if bc[j] < i {
            if i-bc[j] > shift {
              shift = i - bc[j]
              min = 0
            }
            break
          }
        }
      }

      // Must always shift by at least one
      if shift == 0 {
        shift = 1
        min = 0
      }

      k += shift
    }
  }
  return matches
}

// Implementation of the Boyer-Moore string search, as detailed in Gusfield.
// A detail was left out of Gusfield - in certain shifts we might know that a
// prefix of the current alignment matches, we need to keep track of that to
// avoid quadratic runtime.
func BoyerMooreFromReader(bmd BmData, in io.Reader, buf_size int) []int {
  var matches []int
  k := len(bmd.p) - 1

  if buf_size < 2 * len(bmd.p) {
    buf_size = 2 * len(bmd.p)
  }
  buf := make([]byte, buf_size)

  // In some cases we don't need to go all the way to the left-most character
  // since we might know that a certain prefix of the current alignment
  // matches based on a previous test.
  min := 0

  read := 0
  mark := 0
  for n, err := in.Read(buf[mark:]); err == nil; n, err = in.Read(buf[mark:]) {
    t := buf[:mark+n]
    for k < len(t) {
      i := len(bmd.L) - 1
      h := k
      for i >= min && bmd.p[i] == t[h] {
        i--
        h--
      }

      if i < min {
        // found a match
        matches = append(matches, k-len(bmd.L)+1+read)
        l0 := bmd.l[0]
        if l0 > 0 {
          k += l0

          // Since we matched we will know some prefix of the next alignment.
          min = len(bmd.L) - l0
        } else {
          k++
          min = 0
        }
      } else {
        shift := 0

        // Strong good suffix rule
        Li := bmd.L[i]
        if Li == 0 {
          shift = bmd.l[i]

          // This shift can place part of what already matched as a prefix.
          min = len(bmd.L) - shift
        } else {
          shift = Li
          min = 0
        }

        // Extended bad character rule
        bc := bmd.R[t[h]]
        if len(bc) == 0 {
          shift = i + 1
          min = 0
        } else {
          for _, bcj := range bc {
            if bcj < i {
              if i - shift > bcj {
                shift = i - bcj
                min = 0
              }
              break
            }
          }
        }

        // Must always shift by at least one
        if shift == 0 {
          shift = 1
          min = 0
        }

        k += shift
      }
    }
    horizon := k - len(bmd.p) + min
    if horizon < 0 {
      horizon = 0
    }
    read += horizon
    k -= horizon
    copy(buf, t[horizon:])
    mark = len(t[horizon:])
  }
  return matches
}
