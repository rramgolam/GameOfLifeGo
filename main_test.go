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

func TestGrid(t *testing.T) {

	t.Run("Grids have a height and width", func(t *testing.T) {

		grid := NewGrid(3, 3)
		assert.Equal(t, 3, grid.X)
		assert.Equal(t, 3, grid.Y)

	})

	t.Run("Grid is full of cells", func(t *testing.T) {
	
		grid := NewGrid(3, 3)
		assert.Equal(t, 3, len(grid.Columns[0]))
		assert.Equal(t, 3, len(grid.Columns[1]))
		assert.Equal(t, 3, len(grid.Columns[2]))
		
		assert.IsType(t, []cell{}, grid.Columns[0])
		assert.IsType(t, []cell{}, grid.Columns[1])
		assert.IsType(t, []cell{}, grid.Columns[2])
		assert.IsType(t, [][]cell{}, grid.Columns)

	})

	t.Run("Grid can check living neighbours for X and Y", func(t *testing.T) {

		grid := NewGrid(3, 3)
		actual := grid.checkCell(1,1)
		assert.Equal(t, 8, actual)

	})

	t.Run("Grid can check living neighbours for cells on the edge", func(t *testing.T) {

		grid := NewGrid(3, 3)
		actual := grid.checkCell(0,0)
		assert.Equal(t, 8, actual)

	})

	t.Run("Grid detects dead cells", func(t *testing.T) {

		grid := NewGrid(3, 3)
		grid.Columns[0][1].Alive = false
		actual := grid.checkCell(1,1)

		assert.Equal(t, 7, actual)

	})

	t.Run("Grid detects all neighbours are dead, oh no", func(t *testing.T) {

		grid := NewGrid(3, 3)
		grid.Columns[0][0].Alive = false
		grid.Columns[0][1].Alive = false
		grid.Columns[0][2].Alive = false

		grid.Columns[1][0].Alive = false
		grid.Columns[1][2].Alive = false
		grid.Columns[1][0].Alive = false

		grid.Columns[2][0].Alive = false
		grid.Columns[2][1].Alive = false
		grid.Columns[2][2].Alive = false

		actual := grid.checkCell(1,1)

		assert.Equal(t, 0, actual)

	})
}

func TestGridUpdate(t *testing.T) {
	t.Run("Grid updates every cell of the board", func(t *testing.T) {

		expectedGrid := NewGrid(6,6)
		expectedGrid.Columns[0][0].Alive = false
		expectedGrid.Columns[0][1].Alive = false
		expectedGrid.Columns[0][2].Alive = false
		expectedGrid.Columns[0][3].Alive = false
		expectedGrid.Columns[0][4].Alive = false
		expectedGrid.Columns[0][5].Alive = false

		expectedGrid.Columns[1][0].Alive = false
		expectedGrid.Columns[1][1].Alive = true
		expectedGrid.Columns[1][2].Alive = true
		expectedGrid.Columns[1][3].Alive = false
		expectedGrid.Columns[1][4].Alive = false
		expectedGrid.Columns[1][5].Alive = false

		expectedGrid.Columns[2][0].Alive = false
		expectedGrid.Columns[2][1].Alive = true
		expectedGrid.Columns[2][2].Alive = false
		expectedGrid.Columns[2][3].Alive = false
		expectedGrid.Columns[2][4].Alive = false
		expectedGrid.Columns[2][5].Alive = false

		expectedGrid.Columns[3][0].Alive = false
		expectedGrid.Columns[3][1].Alive = false
		expectedGrid.Columns[3][2].Alive = false
		expectedGrid.Columns[3][3].Alive = false
		expectedGrid.Columns[3][4].Alive = true
		expectedGrid.Columns[3][5].Alive = false

		expectedGrid.Columns[4][0].Alive = false
		expectedGrid.Columns[4][1].Alive = false
		expectedGrid.Columns[4][2].Alive = false
		expectedGrid.Columns[4][3].Alive = true
		expectedGrid.Columns[4][4].Alive = true
		expectedGrid.Columns[4][5].Alive = false
		
		expectedGrid.Columns[5][0].Alive = false
		expectedGrid.Columns[5][1].Alive = false
		expectedGrid.Columns[5][2].Alive = false
		expectedGrid.Columns[5][3].Alive = false
		expectedGrid.Columns[5][4].Alive = false
		expectedGrid.Columns[5][5].Alive = false

		grid := NewGrid(6,6)
		grid.Columns[0][0].Alive = false
		grid.Columns[0][1].Alive = false
		grid.Columns[0][2].Alive = false
		grid.Columns[0][3].Alive = false
		grid.Columns[0][4].Alive = false
		grid.Columns[0][5].Alive = false

		grid.Columns[1][0].Alive = false
		grid.Columns[1][1].Alive = true
		grid.Columns[1][2].Alive = true
		grid.Columns[1][3].Alive = false
		grid.Columns[1][4].Alive = false
		grid.Columns[1][5].Alive = false

		grid.Columns[2][0].Alive = false
		grid.Columns[2][1].Alive = true
		grid.Columns[2][2].Alive = true
		grid.Columns[2][3].Alive = false
		grid.Columns[2][4].Alive = false
		grid.Columns[2][5].Alive = false

		grid.Columns[3][0].Alive = false
		grid.Columns[3][1].Alive = false
		grid.Columns[3][2].Alive = false
		grid.Columns[3][3].Alive = true
		grid.Columns[3][4].Alive = true
		grid.Columns[3][5].Alive = false

		grid.Columns[4][0].Alive = false
		grid.Columns[4][1].Alive = false
		grid.Columns[4][2].Alive = false
		grid.Columns[4][3].Alive = true
		grid.Columns[4][4].Alive = true
		grid.Columns[4][5].Alive = false
		
		grid.Columns[5][0].Alive = false
		grid.Columns[5][1].Alive = false
		grid.Columns[5][2].Alive = false
		grid.Columns[5][3].Alive = false
		grid.Columns[5][4].Alive = false
		grid.Columns[5][5].Alive = false
		grid.Update()

		assert.Equal(t, expectedGrid, grid)

	})

}
