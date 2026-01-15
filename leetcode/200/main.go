package main

import (
	"fmt"
	"slices"
)

type VisitedList [][2]int

func (v *VisitedList) has(row, column int) bool {
	return slices.Contains(*v, [2]int{row, column})
}

func (v *VisitedList) push(row, column int) {
	if !v.has(row, column) {
		*v = append(*v, [2]int{row, column})
	}
}

func bfs(grid [][]string, row int, col int, visited *VisitedList) {
	gridNextPositions := [][2]int{
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
		{-1, 0}, // up
	}
	q := [][2]int{{row, col}}
	dc := 0

	for len(q) > 0 {
		current := q[0]
		q = q[1:] // jump first element in the queue -- mark it as consumed
		for _, pos := range gridNextPositions {
			nextRow := current[0] + pos[0]
			nextCol := current[1] + pos[1]

			if !visited.has(nextRow, nextCol) && nextRow >= 0 && nextCol >= 0 && nextRow < len(grid) && nextCol < len(grid[0]) {
				nextItemValue := grid[nextRow][nextCol]
				if nextItemValue == "1" {
					grid[nextRow][nextCol] = "0"
					dc++
					q = append(q, [2]int{nextRow, nextCol})
				}
				visited.push(current[0], current[1])
			}

		}

	}
}

func numIslands(grid [][]string) int {
	rows := len(grid)
	cols := len(grid[0])

	visited := &VisitedList{}
	counter := 0

	for row := range rows {
		for col := range cols {
			current := grid[row][col]
			if current == "1" && !visited.has(row, col) {
				counter++
				bfs(grid, row, col, visited)
			}
		}
	}
	return counter
}

func main() {
	inputs := [][][]string{
		{
			{"1", "1", "1", "1", "0"},
			{"1", "1", "0", "1", "0"},
			{"1", "1", "0", "0", "0"},
			{"0", "0", "0", "0", "0"},
		},
		{
			{"1", "1", "0", "0", "0"},
			{"1", "1", "0", "0", "0"},
			{"0", "0", "1", "0", "0"},
			{"0", "0", "0", "1", "1"},
		},
	}
	outputs := []int{
		1,
		3,
	}
	_ = inputs
	_ = outputs

	for i := range len(inputs) {
		fmt.Printf("%d: result %d | expected: %d\n", i, numIslands(inputs[i]), outputs[i])
	}
}
