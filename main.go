package main

type cell struct {
	Alive bool
}

func Main () {

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