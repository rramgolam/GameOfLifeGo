package main

type cell struct {
	Alive bool
	LivingNeighbours int
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

func (c *cell) Update(livingNeighbours int) {
	switch livingNeighbours {
	case 0, 1:
		c.Die()
	case 2:
		break
	case 3:
		c.Live()
	case 4:
		c.Die()
	default:
		c.Die()
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

func (g *grid) checkCell(x, y int) (livingNeighbours int) {

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
							{minX, maxY},
						}

	for _ , neighbour := range neighbours {
		column := g.Columns[neighbour.y]
		cell := column[neighbour.x]

		if (cell.Alive) {
			livingNeighbours++
		}
	}			
	return 
}

func (g *grid) Update() {
	
}