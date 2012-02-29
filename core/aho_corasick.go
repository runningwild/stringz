package core

type AcNode struct {
  // Index into the AcNodeArray for a given character
  next [256]int

  // failure link index
  failure int

  // This node indicates that the following elements matched
  matches []int
}

type ahBfs struct {
  node, data, index int
}

type AcData struct {
  // Lengths of the patterns
  Lengths []int

  // The graph
  Nodes []AcNode
}

func AhoCorasickPreprocess(datas [][]byte) AcData {
  var acd AcData
  acd.Lengths = make([]int, len(datas))
  for i := range acd.Lengths {
    acd.Lengths[i] = len(datas[i])
  }

  total_len := 0
  for i := range datas {
    total_len += len(datas[i])
  }
  acd.Nodes = make([]AcNode, total_len+1)[0:1]
  for i, data := range datas {
    cur := 0
    for _, b := range data {
      if acd.Nodes[cur].next[b] != 0 {
        cur = acd.Nodes[cur].next[b]
        continue
      }
      acd.Nodes[cur].next[b] = len(acd.Nodes)
      cur = len(acd.Nodes)
      acd.Nodes = append(acd.Nodes, AcNode{})
    }
    acd.Nodes[cur].matches = append(acd.Nodes[cur].matches, i)
  }

  // The skeleton of the graph is done, now we do a BFS on the nodes and form
  // failure links as we go.
  var q []ahBfs
  for i := range datas {
    // TODO: Figure out if this makes sense, maybe we should fix how the BFS
    // works instead?
    if len(datas[i]) > 1 {
      bfs := ahBfs{
        node:  acd.Nodes[0].next[datas[i][0]],
        data:  i,
        index: 1,
      }
      q = append(q, bfs)
    }
  }
  for len(q) > 0 {
    bfs := q[0]
    q = q[1:]
    mod := acd.Nodes[bfs.node].failure
    edge := datas[bfs.data][bfs.index]
    for mod != 0 && acd.Nodes[mod].next[edge] == 0 {
      mod = acd.Nodes[mod].failure
    }
    source := acd.Nodes[bfs.node].next[edge]
    if acd.Nodes[source].failure == 0 {
      target := acd.Nodes[mod].next[edge]
      acd.Nodes[source].failure = target
      for _, m := range acd.Nodes[target].matches {
        acd.Nodes[source].matches = append(acd.Nodes[source].matches, m)
      }
    }
    bfs.node = acd.Nodes[bfs.node].next[edge]
    bfs.index++
    if bfs.index < len(datas[bfs.data]) {
      q = append(q, bfs)
    }
  }

  return acd
}

func AhoCorasick(acd AcData, t []byte) [][]int {
  cur := 0
  matches := make([][]int, len(acd.Lengths))
  for i, c := range t {
    for _, m := range acd.Nodes[cur].matches {
      matches[m] = append(matches[m], i-acd.Lengths[m])
    }
    for acd.Nodes[cur].next[c] == 0 {
      if acd.Nodes[cur].failure != 0 {
        cur = acd.Nodes[cur].failure
      } else {
        cur = 0
        break
      }
    }
    cur = acd.Nodes[cur].next[c]
  }
  for _, m := range acd.Nodes[cur].matches {
    matches[m] = append(matches[m], len(t)-acd.Lengths[m])
  }
  return matches
}
