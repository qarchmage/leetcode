package main

import "fmt"

func main() {
	// true false false
	fmt.Println(possibleBipartition(4, [][]int{{1, 2}, {1, 3}, {2, 4}}))
	fmt.Println(possibleBipartition(3, [][]int{{1, 2}, {1, 3}, {2, 3}}))
	fmt.Println(possibleBipartition(5, [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {1, 5}}))

}

// dfs approach
func possibleBipartition1(N int, dislikes [][]int) bool {
	graph := make([][]int, N)
	for i := range graph {
		graph[i] = make([]int, N)
	}

	for _, v := range dislikes {
		graph[v[0]-1][v[1]-1] = 1
		graph[v[1]-1][v[0]-1] = 1
	}
	group := make([]int, N)
	for i := range group {
		if group[i] == 0 && !dfs(graph, group, i, 1) {
			return false
		}
	}

	return true
}
func dfs(graph [][]int, group []int, index int, g int) bool {
	group[index] = g
	for i := range graph {
		if graph[index][i] == 1 {
			if group[i] == g {
				return false
			}
			if group[i] == 0 && !dfs(graph, group, i, -g) {
				return false
			}
		}
	}
	return true
}

// with pointers

func possibleBipartition(N int, dislikes [][]int) bool {
	g := graph{}
	g.init(N)

	op := make([]int, N+1)

	for _, d := range dislikes {
		i := d[0]
		j := d[1]
		if g.root(i) == g.root(j) {
			return false
		}

		if op[i] != 0 {
			g.union(j, op[i])
		} else {
			op[i] = g.root(j)
		}

		if op[j] != 0 {
			g.union(i, op[j])
		} else {
			op[j] = g.root(i)
		}
	}
	return true
}

type graph struct {
	n     int
	roots []int
	sz    []int
}

func (g *graph) init(n int) {
	g.n = n
	g.roots = make([]int, n+1)
	g.sz = make([]int, n+1)
	for i := 1; i <= n; i++ {
		g.roots[i] = i
		g.sz[i] = 1
	}
}

func (g *graph) root(i int) int {
	for i != g.roots[i] {
		g.roots[i] = g.roots[g.roots[i]]
		i = g.roots[i]
	}
	return i
}

func (g *graph) union(i, j int) {
	p := g.root(i)
	q := g.root(j)
	if p == q {
		return
	}

	if g.sz[p] < g.sz[q] {
		g.roots[p] = q
		g.sz[q] += g.sz[p]
	} else {
		g.roots[q] = p
		g.sz[p] += g.sz[q]
	}
}
