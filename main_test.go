package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCell(t *testing.T) {
	t.Run("Cells can be alive", func(t *testing.T) {

		cell := NewCell()
		assert.NotEqual(t, nil, cell.Alive)
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