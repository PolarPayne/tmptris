package main

import (
	"math/rand"

	"github.com/polarpayne/tmptris/board"
)

type Board struct {
	curPiece board.Piece
	board    [10][20]bool
	gravity  int
	random   rand.Source
}

func NewBoard(seed int64) *Board {
	return &Board{
		gravity: 1,
		random:  rand.NewSource(seed)}
}

func (b *Board) Update(bs buttons) error {
	return nil
}
