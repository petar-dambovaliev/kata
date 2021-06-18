package kata

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/iterator"
	"gonum.org/v1/gonum/graph/path"
)

type Game struct {
	wg *wordGraph
}


// Creates a new game instance
func NewGame(words io.Reader) (*Game, error) {
	wg := newWordGraph()

	sc := bufio.NewScanner(words)
	for sc.Scan() {
		wg.include(sc.Text())
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	return &Game{wg: &wg}, nil
}

// GetShortestPath finds the shortest transformation path
// between 2 words and returns it
func (g *Game) GetShortestPath(first, last string) []string {
	pth := path.DijkstraFrom(g.wg.nodeFor(first), g.wg)
	ladder, _ := pth.To(g.wg.nodeFor(last).ID())

	res := make([]string, len(ladder))
	for i, w := range ladder {
		res[i] = fmt.Sprintf("%v", w)
	}
	return res
}

type wordGraph struct {
	words []string
	ids   map[string]int64
}

func newWordGraph() wordGraph {
	return wordGraph{ids: make(map[string]int64)}
}

func (g *wordGraph) include(word string) {
	word = strings.ToLower(word)
	if _, exists := g.ids[word]; exists {
		return
	}
	g.ids[word] = int64(len(g.words))
	g.words = append(g.words, word)
}

func (g wordGraph) nodeFor(word string) graph.Node {
	id, ok := g.ids[word]
	if !ok {
		return nil
	}
	return node{word, id}
}

func (g wordGraph) From(id int64) graph.Nodes {
	if uint64(id) >= uint64(len(g.words)) {
		return graph.Empty
	}
	return iterator.NewOrderedNodes(neighbours(g.words[id], g.ids))
}

func neighbours(word string, words map[string]int64) []graph.Node {
	var adj []graph.Node
	for j := range word {
		for d := byte('a'); d <= 'z'; d++ {
			b := make([]byte, len(word))
			for i, c := range []byte(word) {
				if i == j {
					b[i] = d
				} else {
					b[i] = c
				}
			}
			w := string(b)
			if w != word {
				if _, ok := words[w]; ok {
					adj = append(adj, node{word: w, id: words[w]})
				}
			}
		}
	}
	return adj
}

func (g wordGraph) Edge(uid, vid int64) graph.Edge {
	if uid == vid {
		return nil
	}
	if uint64(uid) >= uint64(len(g.words)) {
		return nil
	}
	if uint64(vid) >= uint64(len(g.words)) {
		return nil
	}
	u := g.words[uid]
	v := g.words[vid]
	d := distance(u, v)
	if d != 1 {
		return nil
	}
	return edge{f: node{u, uid}, t: node{v, vid}}
}

func distance(a, b string) int {
	if len(a) != len(b) {
		panic("word length mismatch")
	}
	var d int
	for i, c := range []byte(a) {
		if c != b[i] {
			d++
		}
	}
	return d
}

type node struct {
	word string
	id   int64
}

func (n node) ID() int64      { return n.id }
func (n node) String() string { return n.word }

type edge struct{ f, t node }

func (e edge) From() graph.Node         { return e.f }
func (e edge) To() graph.Node           { return e.t }
func (e edge) ReversedEdge() graph.Edge { return edge{f: e.t, t: e.f} }