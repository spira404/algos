package main

import(
	"fmt"
)

// using this insted of math.Inf bc it's int
var inf = 10000000

func bellmanford(graph [][]int, src int){
	
	V := len(graph)
	dist := make([]int, V-1)

	for i := range dist{
		dist[i] = inf
	}

	dist[src] = 0
	for i := 0; i < V-1; i++{
		for _, j := range graph{
			u, v, w := j[0], j[1], j[2] 
			if dist[u] != inf && dist[u] + w < dist[v]{
				dist[v] = dist[u] + w
			}
		}
	}

	for _, j := range graph{
		u, v, w := j[0], j[1], j[2] 
		if dist[u] != inf && dist[u] + w < dist[v]{
			fmt.Println("At least one negative cycle exist!")
			return
		}
	}

	for idx, weight := range dist{
		fmt.Printf("From source to vertex %d - %d \n", idx, weight)}
}


func main(){
//	            ┌───┐          ┌───┐
//	  ┌───5────►│ 1 ├────2────►│ 3 ├────2────┐
//	  │         └─┬─┘          └───┘         │
//	  │           │              ▲           ▼
//	┌─┴─┐         1              │         ┌───┐
//	│ 0 │         │             -1         │ 5 │
//	└───┘         ▼              │         └─┬─┘
//	            ┌───┐          ┌─┴─┐         │
//	            │ 2 ├────1────►│ 4 │◄───-3───┘
//	            └───┘          └───┘
	graph := [][]int{
		{0, 1, 5},
		{1, 2, 1},
		{1, 3, 2},
		{2, 4, 1},
		{3, 5, 2},
		{4, 3, -1},
		{5, 4, -3}}

	bellmanford(graph, 0)

	// change the last weight to +
	graph = [][]int{
		{0, 1, 5},
		{1, 2, 1},
		{1, 3, 2},
		{2, 4, 1},
		{3, 5, 2},
		{4, 3, -1},
		{5, 4, 3}}

	bellmanford(graph, 0)
	}


