package stringz

// import "fmt"

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

func PrecalcZboxesReversed(p string) []int {
  n := len(p)
  if n == 0 { return nil }
  if n == 1 { return []int{ n } }
  pos := n - 2
  for pos >= 0 && p[pos] == p[pos + 1] {
    pos--
  }

  zs := make([]int, n)
  zs[n - 1] = n
  zs[n - 2] = n - pos - 2

  left := pos + 1
  right := n - 2
  for i := n - 3; i >= 0; i-- {
    if left > i {
      // We just left a zbox - so we need to see how much of a suffix we have
      // from this position
      pos := i
      for pos >= 0 && p[pos] == p[n - (i - pos) - 1] {
        pos--
      }
      zs[i] = i - pos
      left = pos + 1
      right = i
    } else {
      j := right - i
      rem := i - left + 1
      zj := zs[n - j - 1]
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

// L here is not exactly as specified in Gusfield, since the value used when
// determining shifts is always n - L'(i) we just do that math now rather than
// later.
func BoyerMooreStrongGoodSuffixRule(p string) (L,l []int) {
  Z := PrecalcZboxesReversed(p)
  L = make([]int, len(p))
  for i := 0; i < len(Z) - 1; i++ {
    if Z[i] == 0 { continue }
    L[len(p) - Z[i] - 1] = Z[i]
  }

  l = PrecalcZboxes(p)
  for i := len(l) - 2; i >= 0; i-- {
    if l[i+1] > l[i] {
      l[i] = l[i+1]
    }
  }
  for i := 0; i < len(l) - 1; i++ {
    l[i] = len(p) - l[i+1]
  }
  l[len(l) - 1] = 0
  return
}

func BoyerMoore(p,t string) []int {
  var matches []int
  L,l := BoyerMooreStrongGoodSuffixRule(p)
  // R := boyerMooreExtendedBadCharacterRule(p)
  k := len(p) - 1
  // fmt.Printf("%v\n%v\n", L, l)

  // In some cases we don't need to go all the way to the left-most character
  // since we might know that a certain prefix of the current alignment
  // matches based on a previous test.
  min := 0
  for k < len(t) {
    i := len(p) - 1
    h := k
    // fmt.Printf("State: %d %d %d\n", i, h, k)
    for i >= min && p[i] == t[h] {
      i--
      h--
    }

    if i < min {
      // found a match
      matches = append(matches, k - len(p) + 1)
      k += l[0]
      min = len(p) - l[0]
      // i = len(p) - 1
      // h = k
    } else {
      shift := L[i]
      if shift == 0 {
        // shift = l[i]
      }
      // println("Mismatch at ", i, ",", h, ",", k, " shift L=", L[i], "/", l[i])
      // r := i
      // for _,v := range R[t[h]] {
      //   if v < i {
      //     r = i - v
      //     break
      //   }
      // }
      // if r > shift {
      //   shift = r
      // }
      if shift == 0 {
        shift = 1
        // i = len(p) - 1
        // h = k
      }
      k += shift
      // h += shift
      // fmt.Printf("Shift %d -> %d\n", i, i + shift)
      // i += shift
      min = 0
    }
  }
  return matches
}

func revString(s string) string {
  b := []byte(s)
  for i := 0; i < len(b) / 2; i++ {
    b[i], b[len(b) - i - 1] = b[len(b) - i -1], b[i]
  }
  return string(b)
}
