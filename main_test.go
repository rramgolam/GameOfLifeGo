package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCell(t *testing.T) {
	t.Run("Cells can be alive", func(t *testing.T) {

		cell := NewCell()
		assert.NotEqual(t, "", cell.Alive)
	})

	t.Run("Cells are alive when created", func(t *testing.T) {

		cell := NewCell()
		assert.True(t, cell.Alive)

	})

	t.Run("A cell can die", func(t *testing.T) {
		
		cell := NewCell()
		cell.Die()
		assert.False(t, cell.Alive)

	})

	t.Run("A cell can come back to life", func(t *testing.T) {

		cell := NewCell()
		cell.Alive = false
		cell.Live()
		assert.True(t, cell.Alive)

	})

	t.Run("When we have two cells, one does not kill the other", func(t *testing.T) {

		cellOne := NewCell()
		cellTwo := NewCell()

		cellOne.Die()
		assert.True(t, cellTwo.Alive)

	})
}

func TestCellRules(t *testing.T) {
	t.Run("cells with fewer than 2 living neighbours dies", func(t *testing.T) {

		cellOne := NewCell()
		cellTwo := NewCell()
		cellOne.Alive = true
		cellTwo.Alive = true

		cellOne.Update(0)
		cellTwo.Update(1)

		assert.False(t, cellOne.Alive)
		assert.False(t, cellTwo.Alive)
	})

	t.Run("cells with 2 or 3 living neighbours survive", func(t *testing.T) {

		cellOne := NewCell()
		cellTwo := NewCell()
		cellOne.Alive = true
		cellTwo.Alive = true

		cellOne.Update(2)
		cellTwo.Update(3)

		assert.True(t, cellOne.Alive)
		assert.True(t, cellTwo.Alive)

	})

	t.Run("cells with more than 3 neighbours die", func(t *testing.T) {

		cellOne := NewCell()
		cellTwo := NewCell()
		cellOne.Alive = true
		cellTwo.Alive = true

		cellOne.Update(4)
		cellTwo.Update(5)

		assert.False(t, cellOne.Alive)
		assert.False(t, cellTwo.Alive)

	})

	t.Run("dead cells with exactly 3 neighbours come back to life", func(t *testing.T) {

		cell := NewCell()
		cell.Alive = false

		cell.Update(3)

		assert.True(t, cell.Alive)
	})

	t.Run("dead cells with 2 neighbours do not come back to life", func(t *testing.T) {

		cell := NewCell()
		cell.Alive = false

		cell.Update(2)

		assert.False(t, cell.Alive)
	})
}