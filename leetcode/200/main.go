package main

import (
	"fmt"
)

func numIslands(input [][]string) int {
	width := len(input[0])
	height := len(input)
	
	return 0
}

func main(){
	inputs :=[][][]string {
		{
			{"1","1","1","1","0"},
			{"1","1","0","1","0"},
			{"1","1","0","0","0"},
			{"0","0","0","0","0"},
		},
		{
			{"1","1","0","0","0"},
			{"1","1","0","0","0"},
			{"0","0","1","0","0"},
			{"0","0","0","1","1"},
		},
	}
	outputs := []int{1,3}
	_ = inputs
	_ = outputs

	for i:=0; i < 2 ;i++{
		fmt.Printf("%d: result %d | expected: %d\n", i, numIslands(inputs[i]), outputs[i])
	}
}
