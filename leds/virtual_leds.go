package leds

import (
	"github.com/ansoni/termination"
)

type VirtualLEDGrid struct {
	NumRows int
	NumCols int
	Grid    [][]Color
	Term    *termination.Termination
}

var ledShape = termination.Shape{
	"default": []string{""},
	"on":      []string{"*"},
}

func ledMovement(t *termination.Termination, e *termination.Entity, position termination.Position) termination.Position {
	lg := e.Data.(*VirtualLEDGrid)
	if lg.Grid[position.Y][position.X] > 0 {
		e.ShapePath = "on"
	} else {
		e.ShapePath = "default"
	}
	return position
}

func AnimateVirtualGrid(lg *VirtualLEDGrid, framesPerSecond int) {
	lg.Term = termination.New()
	lg.Term.FramesPerSecond = framesPerSecond
	for i := range lg.Grid {
		for j := range lg.Grid[i] {
			ledEntity := lg.Term.NewEntity(termination.Position{j, i, 0})
			ledEntity.Shape = ledShape
			ledEntity.MovementCallback = ledMovement
			ledEntity.Data = lg
		}
	}
	lg.Term.Animate()
}

func StopVirtualGrid(lg *VirtualLEDGrid) {
	lg.Term.Close()
}

func NewVirtualLEDGrid(numRows int, numCols int) *VirtualLEDGrid {
	ledGrid := &VirtualLEDGrid{
		NumRows: numRows,
		NumCols: numCols,
		Grid:    make([][]Color, numRows),
	}
	for i := range ledGrid.Grid {
		ledGrid.Grid[i] = make([]Color, numCols)
	}
	return ledGrid
}

func (lg *VirtualLEDGrid) SetLED(row int, col int, color Color) error {
	if row < 0 || row > lg.NumRows || col < 0 || col > lg.NumCols {
		return ErrLEDOutOfBounds
	}
	lg.Grid[row][col] = color
	return nil
}

func (lg *VirtualLEDGrid) UpdateLEDs() error {
	return nil
}