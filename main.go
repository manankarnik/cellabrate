package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

var grid = [][]int{}
var quit = make(chan bool)
var cursor = [2]int{0, 0}
var simulate = false

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

	width, height := screen.Size()
	if width == 0 || height == 0 {
		return
	}

	go poll(screen, width, height)

	for col := range width {
		grid = append(grid, []int{})
		for range height * 2 {
			grid[col] = append(grid[col], 0)

		}
	}
	grid[width/2][height] = 1
	grid[width/2][height+1] = 1
	grid[width/2][height+2] = 1
	cursor[0] = width / 2
	cursor[1] = height + 1
	screen.Show()

	go func() {
		for {
			if simulate {
				update()
				time.Sleep(time.Millisecond * 200)
			}
			draw(screen, width, height)
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

func poll(screen tcell.Screen, width, height int) {
	for {
		event := screen.PollEvent()
		switch event := event.(type) {
		case *tcell.EventKey:
			switch event.Key() {
			case tcell.KeyEscape:
				close(quit)
				return
			case tcell.KeyUp:
				if cursor[1] > 0 {
					cursor[1] -= 1
				}
			case tcell.KeyLeft:
				if cursor[0] > 0 {
					cursor[0] -= 1
				}
			case tcell.KeyDown:
				if cursor[1] < height*2-1 {
					cursor[1] += 1
				}
			case tcell.KeyRight:
				if cursor[0] < width-1 {
					cursor[0] += 1
				}
			}
			switch event.Rune() {
			case 'q':
				close(quit)
				return
			case 'w':
				if cursor[1] > 0 {
					cursor[1] -= 1
				}
			case 'a':
				if cursor[0] > 0 {
					cursor[0] -= 1
				}
			case 's':
				if cursor[1] < height*2-1 {
					cursor[1] += 1
				}
			case 'd':
				if cursor[0] < width-1 {
					cursor[0] += 1
				}
			case ' ':
				grid[cursor[0]][cursor[1]] ^= 1
			case 'f':
				simulate = !simulate
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
	if !simulate {
		if cursor[1]%2 == 0 {
			style = style.Foreground(tcell.ColorNone)
			style = style.Background(tcell.ColorRed)
		} else {
			style = style.Foreground(tcell.ColorRed)
			style = style.Background(tcell.ColorNone)
		}
		screen.SetContent(cursor[0], cursor[1]/2, ch, nil, style)
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
