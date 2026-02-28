package rules

import (
	"ChessEngineEOF/internal/model"
	"ChessEngineEOF/internal/pieces"
	"math"
)

func PawnMove(b model.Board, p model.PiecesMove) (bool, bool) {
	CanMove := false
	MoveWhite := false
	if b.White == true {
		if b.Board[p.MHX][p.MHY] > 0 || b.Board[p.MHX][p.MHY] == 99 {
			CanMove = false
		} else {
			if p.MHY-p.MFY < 2 && p.MHY-p.MFY > 0 && pieces.Pawn != b.Board[p.MFX][2] && p.MFX == p.MHX && b.Board[p.MFX][p.MHY] == 0 {
				MoveWhite = true
				CanMove = true

			} else if pieces.Pawn == b.Board[p.MFX][2] && p.MHY-p.MFY > 1 && p.MHY-p.MFY < 3 && b.Board[p.MFX][p.MHY] == 0 && p.MFX == p.MHX && b.Board[p.MFX][p.MHY-1] == 0 {
				MoveWhite = true
				CanMove = true
				b.Move2W = true
			} else if b.Board[p.MHX][p.MHY] < 0 && p.MHY-p.MFY == 1 && math.Abs(float64(p.MHX)-float64(p.MFX)) == 1 {
				MoveWhite = true
				CanMove = true
			} else if b.Move2B == true && b.Board[p.MFX][5] == pieces.Pawn && p.MHY-p.MFY < 2 && math.Abs(float64(p.MHX)-float64(p.MFX)) == 1 && p.MHY-p.MFY > 0 {
				MoveWhite = true
				CanMove = true
				b.Board[p.MHX][p.MHY-1] = pieces.Empty
			} else {

				CanMove = false
			}

		}
		if CanMove {
			b.Move2B = false
		}

	} else {
		if b.Board[p.MHX][p.MHY] < 0 || b.Board[p.MHX][p.MHY] == 99 {
			CanMove = false
		} else {
			if p.MFY-p.MHY < 2 && p.MFY-p.MHY > 0 && pieces.PawnE != b.Board[p.MFX][7] && p.MFX == p.MHX && b.Board[p.MFX][p.MHY] == 0 {
				MoveWhite = false
				CanMove = true
			} else if pieces.PawnE == b.Board[p.MFX][7] && p.MFY-p.MHY > 1 && p.MFY-p.MHY < 3 && b.Board[p.MFX][p.MHY] == 0 && p.MFX == p.MHX && b.Board[p.MFX][p.MHY+1] == 0 {
				MoveWhite = false
				CanMove = true
				b.Move2B = true
			} else if b.Board[p.MHX][p.MHY] > 0 && p.MFY-p.MHY == 1 && math.Abs(float64(p.MFX)-float64(p.MHX)) == 1 {
				MoveWhite = false
				CanMove = true
			} else if b.Move2W == true && b.Board[p.MFX][4] == pieces.PawnE && p.MFY-p.MHY < 2 && math.Abs(float64(p.MFX)-float64(p.MHX)) == 1 && p.MFY-p.MHY > 0 {
				b.Board[p.MHX][p.MHY+1] = pieces.Empty
				MoveWhite = false
				CanMove = true

			} else {
				CanMove = false
			}

		}
		if CanMove {
			b.Move2W = false
		}

	}
	return CanMove, MoveWhite
}
