package main

import(
	"fmt"
)

// using this inf insted of math.Inf bc it's int
var inf = 10000000
var processed = map[string]int{}


func find_lowest_cost_node(costs map[string]int) string{
	lowest_cost := inf
	lowest_cost_node := ""

	for node := range costs{
		cost := costs[node]
		if cost < lowest_cost{
			if _, ok := processed[node]; !ok{
				lowest_cost = cost
				lowest_cost_node = node
			}
		}
	}
	return lowest_cost_node
}


func main(){
	// init hardcoded test graph
      // 6      ┌───┐     1
  // ┌─────────►│ a ├───────────┐
  // │          └───┘           ▼
// ┌─┴──┐         ▲           ┌───┐
// │strt│         │           │end│
// └─┬──┘         │           └───┘
  // │          ┌─┴─┐           ▲
  // └─────────►│ b ├───────────┘
      // 2      └───┘     5
	
	graph := map[string]map[string]int{"strt": {"a": 6, "b": 2}, 
	"a": {"end": 1},
	"b": {"a": 3, "end": 5},
	"end": {}}

	parents := map[string]string{"a": "strt", 
	"b": "strt",
	"end": ""}

	costs := map[string]int{"a": 6, 
	"b": 2, 
	"end": inf}

	node := find_lowest_cost_node(costs)
	for node != ""{
		cost := costs[node]
		neighs := graph[node]
		for n := range neighs{
			// unknown costs (inf) becomes current value
			new_cost := cost + neighs[n]
			if new_cost < costs[n]{
				costs[n] = new_cost
				parents[n] = node
			}
		}
		processed[node] = 1
		node = find_lowest_cost_node(costs)
	}
	fmt.Println(costs["end"])
}
