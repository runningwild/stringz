package stringz

type acNode struct {
  // Index into the acNodeArray for a given character
  next [256]int

  // failure link index
  failure int

  // This node indicates that the following elements matched
  matches []int
}

type ahBfs struct {
  node, data, index int
}

func AhoCorasickPreprocessSet(datas [][]byte) []acNode {
  total_len := 0
  for i := range datas {
    total_len += len(datas[i])
  }
  nodes := make([]acNode, total_len + 1)[0:1]
  for i,data := range datas {
    cur := 0
    for _,b := range data {
      if nodes[cur].next[b] != 0 {
        cur = nodes[cur].next[b]
        continue
      }
      nodes[cur].next[b] = len(nodes)
      cur = len(nodes)
      nodes = append(nodes, acNode{})
    }
    nodes[cur].matches = append(nodes[cur].matches, i)
  }

  // The skeleton of the graph is done, now we do a BFS on the nodes and form
  // failure links as we go.
  var q []ahBfs
  for i := range datas {
    // TODO: Figure out if this makes sense, maybe we should fix how the BFS
    // works instead?
    if len(datas[i]) > 1 {
      bfs := ahBfs{
        node:  nodes[0].next[datas[i][0]],
        data:  i,
        index: 1,
      }
      q = append(q, bfs)
    }
  }
  for len(q) > 0 {
    bfs := q[0]
    q = q[1:]
    mod := nodes[bfs.node].failure
    edge := datas[bfs.data][bfs.index]
    for mod != 0 && nodes[mod].next[edge] == 0 {
      mod = nodes[mod].failure
    }
    source := nodes[bfs.node].next[edge]
    if nodes[source].failure == 0 {
      target := nodes[mod].next[edge]
      nodes[source].failure = target
      for _, m := range nodes[target].matches {
        nodes[source].matches = append(nodes[source].matches, m)
      }
    }
    bfs.node = nodes[bfs.node].next[edge]
    bfs.index++
    if bfs.index < len(datas[bfs.data]) {
      q = append(q, bfs)
    }
  }

  return nodes
}

func AhoCorasick(datas [][]byte, t []byte) [][]int {
  nodes := AhoCorasickPreprocessSet(datas)
  cur := 0
  matches := make([][]int, len(datas))
  for i, c := range t {
    for _, m := range nodes[cur].matches {
      matches[m] = append(matches[m], i - len(datas[m]))
    }
    for nodes[cur].next[c] == 0 {
      if nodes[cur].failure != 0 {
        cur = nodes[cur].failure
      } else {
        cur = 0
        break
      }
    }
    cur = nodes[cur].next[c]
  }
  for _, m := range nodes[cur].matches {
    matches[m] = append(matches[m], len(t) - len(datas[m]))
  }
  return matches
}
