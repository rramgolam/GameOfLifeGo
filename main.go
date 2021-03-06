package main

import (
	"fmt"
	"sync"
)

type cell struct {
	Alive bool
}


func Main () {
	// Kill all the cells
}

func NewCell() cell {

	return cell{ Alive: true }

}

func (c *cell) Die() {
	c.Alive = false
}

func (c *cell) Live() {
	c.Alive = true
}

func (c *cell) ShouldIBeDead(livingNeighbours int) bool {
	switch livingNeighbours {
	case 0, 1:
		return true
	case 2:
		return !c.Alive
	case 3:
		return false
	case 4:
		return true
	default:
		return true
	}
}

func (c *cell) Update(livingNeighbours int) {
	if c.ShouldIBeDead(livingNeighbours) {
		c.Die()
	} else {
		c.Live()
	}
}

type grid struct {
	X int
	Y int
	Columns [][]cell
}

func NewGrid(x, y int) grid {
	var columns [][]cell
	for i:=0; i < y; i++ {

		var row []cell

		for j:=0; j < x; j++ {
			row = append(row, NewCell())
		}
		columns = append(columns, row)
	}

	return grid { X : x, Y : y, Columns: columns }
}

type coords struct { x, y int}

func (g *grid) getNeighbourCoods(x, y int) []coords {
	minX := x - 1
	if (minX < 0) {
		minX = g.X - 1
	}

	minY := y - 1
	if (minY < 0) {
		minY = g.Y - 1
	}

	maxX := x + 1
	if (maxX >= g.X) {
		maxX = 0;
	}

	maxY := y + 1
	if (maxY >= g.Y) {
		maxY = 0;
	}

	neighbours := []coords {
							{minX, minY},
							{x, minY},
							{maxX, minY},

							{minX, y},
							{maxX, y},

							{minX, maxY},
							{x, maxY},
							{maxX, maxY},
						}
	return neighbours
}

func (g *grid) checkCell(x, y int) (livingNeighbours int) {

	neighbours := g.getNeighbourCoods(x, y)

	for _ , neighbour := range neighbours {
		column := g.Columns[neighbour.x]
		cell := column[neighbour.y]

		if (cell.Alive) {
			livingNeighbours++
		}
	}			
	return 
}

type cellUpdate struct {
	X int
	Y int
	LivingNeighbours int
}

func (g *grid) Update() {

	channel := make(chan cellUpdate, g.X * g.Y)
	var wg sync.WaitGroup

	for y, c := range g.Columns {	
		for x := range c {
			wg.Add(1)
			go func(x, y int, g *grid, channel chan cellUpdate) {
				defer wg.Done()
				livingNeighbours := g.checkCell(x, y)
				channel <- cellUpdate{X: x, Y: y, LivingNeighbours: livingNeighbours}
			}(x, y, g, channel)
		}
	}

	wg.Wait()
	close(channel)

	for cellUpdate := range channel {
		g.Columns[cellUpdate.Y][cellUpdate.X].Update(cellUpdate.LivingNeighbours)
	}
}

func (g *grid) printMinesweeperGrid() {
	for y, c := range g.Columns {
		row := ""
		for x := range c {
			livingNeighbours := g.checkCell(x, y)
			row += fmt.Sprintf(" %d ", livingNeighbours)
		}
		fmt.Println(row)
	}
}

func PrintGrid(g [][]cell) {

	for _, c := range g {
		row := ""
		for _, x := range c {
			if x.Alive {
				row += " * "
			} else {
				row += " - "
			}
		}
		fmt.Println(row)
	}
}