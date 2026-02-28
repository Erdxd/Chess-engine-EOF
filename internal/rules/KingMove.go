package rules

import (
	"ChessEngineEOF/internal/model"
	"math"
)

func KingMove(b model.Board, p model.PiecesMove) (bool, bool) {
	CanMove := false
	MoveWhite := true
	dx := math.Abs(float64(p.MFX) - float64(p.MHX))
	dy := math.Abs(float64(p.MFY) - float64(p.MHY))
	if b.White == true {
		if b.Board[p.MHX][p.MHY] == 99 || b.Board[p.MHX][p.MHY] > 0 {
			CanMove = false
			MoveWhite = true
		} else if (dx == 1 || dy == 1) && (dx < 2 && dy < 2) {
			CanMove = true
			MoveWhite = false
		}
	} else {
		if b.Board[p.MHX][p.MHY] == 99 || b.Board[p.MHX][p.MHY] < 0 {
			CanMove = false
			MoveWhite = true
		} else if (dx == 1 || dy == 1) && (dx < 2 && dy < 2) {
			CanMove = true
			MoveWhite = false
		}
	}
	return CanMove, MoveWhite
}
