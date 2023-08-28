// Breadth First Search - shortest path. From root of graph to every depth level orderly

package main

import(
	"fmt"
	"slices"
)

var graph = map[int][]int{
	0: {1, 2},
	1: {0, 2, 3}
	2: {0, 2, 4}
	3: {1, 4}
	4: {2, 3}
}

func bfs(graph map, a, b) map {

	// visited := []int{}
	parents := map[int][]int{}
	common := []int{}

	common = append(common, a)

	for len(common) > 0{
		node := common[0]
		visited = append(visited, node)
		common = common[1:]
		for val, _ := range graph[node] {
			if !slices.Contains(parents, val){
				parents[val] = node
				common = append(common, val)
				if node == b {
					break
				}
			}
		}
	}
	// insert here path restoring from parents map
}
	

func main(){
	path := bfs(graph, 1, 4)
	fmt.Println(path)
}
