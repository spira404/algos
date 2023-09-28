package main

import(
	"fmt"
)

// using this insted of math.Inf bc it"s int
var inf = 10000000
var processed = map[string]int{}

func bellmanford(){
	pass
}


func main(){
	// init test graph by hands bc it's about algo
	// not about methods of creating graphs
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

	}
