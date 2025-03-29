package main

import "fmt"

var board = [][]int{
	{0, 1, 0},
	{0, 1, 0},
	{0, 1, 0},
}

func main() {
	printBoard()
	neighbors := [][]int{}
	for row := range board {
		neighbors = append(neighbors, []int{})
		for col := range board[row] {
			neighbors[row] = append(neighbors[row], getNeighbors(row, col))
		}
	}
	for row := range board {
		for col := range board[row] {
			if board[row][col] == 1 {
				if neighbors[row][col] < 2 || neighbors[row][col] > 3 {
					board[row][col] = 0
				}
			} else {
				if neighbors[row][col] == 3 {
					board[row][col] = 1
				}
			}
		}
	}
	fmt.Println()
	printBoard()
}

func printBoard() {
	for row := range board {
		fmt.Println(board[row])
	}
}

func getNeighbors(row, col int) int {
	neighbors := 0
	if row > 0 {
		neighbors += board[row-1][col]
		if col > 0 {
			neighbors += board[row-1][col-1]
		}
		if col < len(board[0])-1 {
			neighbors += board[row-1][col+1]
		}
	}
	if row < len(board)-1 {
		neighbors += board[row+1][col]
		if col > 0 {
			neighbors += board[row+1][col-1]
		}
		if col < len(board[0])-1 {
			neighbors += board[row+1][col+1]
		}
	}
	if col > 0 {
		neighbors += board[row][col-1]
	}
	if col < len(board[0])-1 {
		neighbors += board[row][col+1]
	}
	return neighbors
}
