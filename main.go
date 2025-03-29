package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

var grid = [][]int{}
var quit = make(chan bool)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}

	screen.Clear()

	if err := screen.Init(); err != nil {
		panic(err)
	}
	defer screen.Fini()

	go poll(screen)

	width, height := screen.Size()
	if width == 0 || height == 0 {
		return
	}
	for col := range width {
		grid = append(grid, []int{})
		for range height * 2 {
			grid[col] = append(grid[col], 0)

		}
	}
	grid[width/2][height] = 1
	grid[width/2][height+1] = 1
	grid[width/2][height+2] = 1
	screen.Show()

	go func() {
		for {
			update()
			draw(screen, width, height)
			time.Sleep(time.Millisecond * 200)
		}
	}()

	<-quit
}

func getNeighbors(row, col int) int {
	neighbors := 0
	if row > 0 {
		neighbors += grid[row-1][col]
		if col > 0 {
			neighbors += grid[row-1][col-1]
		}
		if col < len(grid[0])-1 {
			neighbors += grid[row-1][col+1]
		}
	}
	if row < len(grid)-1 {
		neighbors += grid[row+1][col]
		if col > 0 {
			neighbors += grid[row+1][col-1]
		}
		if col < len(grid[0])-1 {
			neighbors += grid[row+1][col+1]
		}
	}
	if col > 0 {
		neighbors += grid[row][col-1]
	}
	if col < len(grid[0])-1 {
		neighbors += grid[row][col+1]
	}
	return neighbors
}

func poll(screen tcell.Screen) {
	for {
		event := screen.PollEvent()
		switch event := event.(type) {
		case *tcell.EventKey:
			switch event.Key() {
			case tcell.KeyEscape:
				close(quit)
				return
			}
			switch event.Rune() {
			case 'q':
				close(quit)
				return
			}
		}
	}
}

func draw(screen tcell.Screen, width, height int) {
	const ch = '▄'
	style := tcell.StyleDefault
	for col := range width {
		for row := range height {
			if grid[col][row*2] == 0 {
				style = style.Background(tcell.ColorBlack)
			} else {
				style = style.Background(tcell.ColorWhite)
			}
			if grid[col][row*2+1] == 0 {
				style = style.Foreground(tcell.ColorBlack)
			} else {
				style = style.Foreground(tcell.ColorWhite)
			}
			screen.SetContent(col, row, ch, nil, style)
		}
	}
	screen.Show()
}

func update() {
	neighbors := [][]int{}
	for row := range grid {
		neighbors = append(neighbors, []int{})
		for col := range grid[row] {
			neighbors[row] = append(neighbors[row], getNeighbors(row, col))
		}
	}
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == 1 {
				if neighbors[row][col] < 2 || neighbors[row][col] > 3 {
					grid[row][col] = 0
				}
			} else {
				if neighbors[row][col] == 3 {
					grid[row][col] = 1
				}
			}
		}
	}
}
