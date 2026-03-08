package movegen

import (
	"ChessEngineEOF/internal/model"
)

func BishopMoveGen(b *model.Board, x, y int) []model.PiecesMove {
	Moves := []model.PiecesMove{}

	for step := 1; step <= 8; step++ {
		RUX := step + x //right up x
		RUY := step + y //right up y

		if RUX > 8 || RUY > 8 || RUX < 0 || RUY < 0 {
			break
		} else if b.Board[RUX][RUY] == 99 || b.Board[RUX][RUY] > 0 {
			break
		} else if b.Board[RUX][RUY] != 0 {
			Moves = append(Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   RUX,
				MHY:   RUY,
				Piece: b.Board[x][y],
			})
			break
		} else {
			Moves = append(Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   RUX,
				MHY:   RUY,
				Piece: b.Board[x][y],
			})
		}

	}
	for stepLU := 1; stepLU <= 8; stepLU++ {
		LUX := x - stepLU //left up x
		LUY := stepLU + y //left up y
		if LUX > 8 || LUY > 8 || LUX < 0 || LUY < 0 {
			break
		} else if b.Board[LUX][LUY] == 99 || b.Board[LUX][LUY] > 0 {
			break
		} else if b.Board[LUX][LUY] != 0 {
			Moves = append(Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   LUX,
				MHY:   LUY,
				Piece: b.Board[x][y],
			})
			break
		} else {
			Moves = append(Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   LUX,
				MHY:   LUY,
				Piece: b.Board[x][y],
			})
		}
	}
	for stepLD := 1; stepLD <= 8; stepLD++ {
		LDX := x - stepLD //left up x
		LDY := y - stepLD //left up y
		if LDX > 8 || LDY > 8 || LDX < 0 || LDY < 0 {
			break
		} else if b.Board[LDX][LDY] == 99 || b.Board[LDX][LDY] > 0 {
			break
		} else if b.Board[LDX][LDY] != 0 {
			Moves = append(Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   LDX,
				MHY:   LDY,
				Piece: b.Board[x][y],
			})

			break
		} else {
			Moves = append(Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   LDX,
				MHY:   LDY,
				Piece: b.Board[x][y],
			})
		}
	}
	for stepRD := 1; stepRD <= 8; stepRD++ {
		RDX := x + stepRD //left up x
		RDY := y - stepRD //left up y
		if RDX > 8 || RDY > 8 || RDX < 0 || RDY < 0 {
			break
		} else if b.Board[RDX][RDY] == 99 || b.Board[RDX][RDY] > 0 {
			break
		} else if b.Board[RDX][RDY] != 0 {
			Moves = append(Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   RDX,
				MHY:   RDY,
				Piece: b.Board[x][y],
			})
			break
		} else {
			Moves = append(Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   RDX,
				MHY:   RDY,
				Piece: b.Board[x][y],
			})
		}
	}

	return Moves
}
