// Breadth First Search - shortest path. From root of graph to every depth level orderly

package main

import(
	"fmt"
)

// init test graph implementation
var graph = map[int][]int{
	0: {1, 2},
	1: {0, 2, 3},
	2: {0, 1, 4},
	3: {1, 4},
	4: {2, 3, 5},
	5: {4},
}

func bfs(graph map[int][]int, a int, b int) []int {

	// child: parent of nodes (source node have value -1)
	parents := map[int]int{a: -1}
	// show the current nodes that are checked (queue)
	tocheck := []int{}
	tocheck = append(tocheck, a)

	for len(tocheck) > 0{
		// moving first node from common to visited
		node := tocheck[0]
		tocheck = tocheck[1:]
		// add current node as parent to children
		for _, child := range graph[node] {
			if _, ok := parents[child]; !ok{
				parents[child] = node
				tocheck = append(tocheck, child)
				if node == b {
					break
				}
			}
		}
	}

	// go from target to source by their parents
	path := []int{b}
	// while target becomes source (-1)
	for{
		if parents[b] != -1{
			path = append([]int{parents[b]}, path...)
			b = parents[b]
		}else{
			break
		}
	}
	return path
}
	

func main(){
	path := bfs(graph, 0, 5)
	fmt.Println(path)
}
