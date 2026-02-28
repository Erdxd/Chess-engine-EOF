package rules

import (
	"ChessEngineEOF/internal/model"
	"math"
)

func KnightMove(b *model.Board, p *model.PiecesMove) (bool, bool) {
	CanMove := false
	MoveWhite := false
	dy := math.Abs(float64(p.MHY) - float64(p.MFY))
	dx := math.Abs(float64(p.MHX) - float64(p.MFX))
	if (dx == 1 && dy == 2) || (dx == 2 && dy == 1) {
		if b.White {
			if b.Board[p.MHX][p.MHY] == 99 || b.Board[p.MHX][p.MHY] > 0 {
				CanMove = false
			} else {
				CanMove = true
				MoveWhite = true
			}

			b.White = false
		} else {
			if b.Board[p.MHX][p.MHY] == 99 || b.Board[p.MHX][p.MHY] < 0 {
				CanMove = false
			} else {
				CanMove = true
				MoveWhite = false
			}

			b.White = true
		}
	}
	return CanMove, MoveWhite
}
