package main

import "fmt"

type Node struct {
	edges []int

	counter int
	cost    int
	checked bool
}

/*
func main() {
	var Q int
	fmt.Scan(&Q)
	for i := 0; i < Q; i++ {
		fmt.Println(RoadsAndLibraries())
	}
}
*/

func DFS(nodes []Node, j, i, c int) []Node {
	if !nodes[i].checked {

		if !nodes[i].checked {
			nodes[j].cost++
		}

		nodes[i].checked = true
		nodes[i].counter = c

		for _, v := range nodes[i].edges {
			nodes = DFS(nodes, j, v, nodes[i].counter+1)
		}
	}
	return nodes
}

func RoadsAndLibraries() int {
	var N, M, clib, croad int
	fmt.Scanf("%d %d %d %d", &N, &M, &clib, &croad)

	nodes := make([]Node, N)
	for i := 0; i < M; i++ {
		var a, b int
		fmt.Scanf("%v %v", &a, &b)

		// fix the indexes
		a--
		b--

		nodes[a].edges = append(nodes[a].edges, b)
		nodes[b].edges = append(nodes[b].edges, a)
	}

	if clib <= croad {
		return clib * N
	}

	for i := 0; i < N; i++ {
		if !nodes[i].checked {
			nodes = DFS(nodes, i, i, 0)
		}
	}

	result := 0
	for i := 0; i < N; i++ {
		if nodes[i].cost != 0 {
			result += clib + (nodes[i].cost-1)*croad
		}
	}

	return result
}
