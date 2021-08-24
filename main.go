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