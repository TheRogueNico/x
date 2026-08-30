package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Grid represents a Game of Life board of fixed size.
type Grid struct {
	cells [][]bool
	w, h  int
}

// NewGrid creates a gird with random live cells.
func NewGrid(w, h int) *Grid {
	cells := make([][]bool, h)
	for y := range cells {
		cells[y] = make([]bool, w)
		for x := range cells[y] {
			cells[y][x] = rand.Intn(4) == 0
		}
	}
	return &Grid{cells: cells, w: w, h: h}
}

// liveNeighbors returns the number of live cells surrounding (x, y).
func (g *Grid) liveNeighbors(x, y int) int {
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := x+dx, y+dy
			if nx < 0 || nx >= g.w || ny < 0 || ny >= g.h {
				continue
			}
			if g.cells[ny][nx] {
				count++
			}
		}
	}
	return count
}

// Next computes the next generation frame.
func (g *Grid) Next() *Grid {
	next := NewGrid(g.w, g.h)
	for y := 0; y < g.h; y++ {
		for x := 0; x < g.w; x++ {
			n := g.liveNeighbors(x, y)
			next.cells[y][x] = n == 3 || (g.cells[y][x] && n == 2)
		}
	}
	return next
}

// String renders the grid as a block of text.
func (g *Grid) String() string {
	var b strings.Builder
	for _, row := range g.cells {
		for _, alive := range row {
			if alive {
				b.WriteRune('█')
			} else {
				b.WriteRune(' ')
			}
		}
		b.WriteRune('\n')
	}
	return b.String()
}

func main() {
	const width, height = 60, 30

	grid := NewGrid(width, height)
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Print("\033[H\033[2J") // clear the terminal screen
		fmt.Print(grid)
		grid = grid.Next()
	}
}
