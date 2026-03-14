package movegen

import (
	"ChessEngineEOF/internal/model"
)

func BishopMoveGen(b *model.Board, x, y int) []model.PiecesMove {
	Moves := []model.PiecesMove{}
	if b.White {

		forplusplus(&Moves, x, y, b, 8)
		forminusplus(&Moves, x, y, b, 8)
		forminusminus(&Moves, x, y, b, 8)
		forplusminus(&Moves, x, y, b, 8)

	} else {
		forplusplusB(&Moves, x, y, b, 8)
		forminusplusB(&Moves, x, y, b, 8)
		forminusminusB(&Moves, x, y, b, 8)
		forplusminusB(&Moves, x, y, b, 8)

	}

	return Moves
}
func forplusplus(Moves *[]model.PiecesMove, x, y int, b *model.Board, maxstep int) {
	for step := 1; step <= maxstep; step++ {
		RUX := step + x //right up x
		RUY := step + y //right up y

		if RUX > 8 || RUY > 8 || RUX < 0 || RUY < 0 {
			break
		} else if b.Board[RUX][RUY] == 99 || b.Board[RUX][RUY] > 0 {
			break
		} else if b.Board[RUX][RUY] != 0 {
			*Moves = append(*Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   RUX,
				MHY:   RUY,
				Piece: b.Board[x][y],
			})
			break
		} else {
			*Moves = append(*Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   RUX,
				MHY:   RUY,
				Piece: b.Board[x][y],
			})
		}

	}
}
func forminusplus(Moves *[]model.PiecesMove, x, y int, b *model.Board, maxstep int) {
	for stepLU := 1; stepLU <= maxstep; stepLU++ {
		LUX := x - stepLU //left up x
		LUY := stepLU + y //left up y
		if LUX > 8 || LUY > 8 || LUX < 0 || LUY < 0 {
			break
		} else if b.Board[LUX][LUY] == 99 || b.Board[LUX][LUY] > 0 {
			break
		} else if b.Board[LUX][LUY] != 0 {
			*Moves = append(*Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   LUX,
				MHY:   LUY,
				Piece: b.Board[x][y],
			})
			break
		} else {
			*Moves = append(*Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   LUX,
				MHY:   LUY,
				Piece: b.Board[x][y],
			})
		}

	}

}
func forminusminus(Moves *[]model.PiecesMove, x, y int, b *model.Board, maxstep int) {
	for stepLD := 1; stepLD <= maxstep; stepLD++ {
		LDX := x - stepLD //left up x
		LDY := y - stepLD //left up y
		if LDX > 8 || LDY > 8 || LDX < 0 || LDY < 0 {
			break
		} else if b.Board[LDX][LDY] == 99 || b.Board[LDX][LDY] > 0 {
			break
		} else if b.Board[LDX][LDY] != 0 {
			*Moves = append(*Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   LDX,
				MHY:   LDY,
				Piece: b.Board[x][y],
			})

			break
		} else {
			*Moves = append(*Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   LDX,
				MHY:   LDY,
				Piece: b.Board[x][y],
			})
		}
	}
}
func forplusminus(Moves *[]model.PiecesMove, x, y int, b *model.Board, maxstep int) {
	for stepRD := 1; stepRD <= maxstep; stepRD++ {
		RDX := x + stepRD //left up x
		RDY := y - stepRD //left up y
		if RDX > 8 || RDY > 8 || RDX < 0 || RDY < 0 {
			break
		} else if b.Board[RDX][RDY] == 99 || b.Board[RDX][RDY] > 0 {
			break
		} else if b.Board[RDX][RDY] != 0 {
			*Moves = append(*Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   RDX,
				MHY:   RDY,
				Piece: b.Board[x][y],
			})
			break
		} else {
			*Moves = append(*Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   RDX,
				MHY:   RDY,
				Piece: b.Board[x][y],
			})
		}
	}
}
func forplusplusB(Moves *[]model.PiecesMove, x, y int, b *model.Board, maxstep int) {
	for step := 1; step <= maxstep; step++ {
		RUX := step + x //right up x
		RUY := step + y //right up y

		if RUX > 8 || RUY > 8 || RUX < 0 || RUY < 0 {
			break
		} else if b.Board[RUX][RUY] == 99 || b.Board[RUX][RUY] < 0 {
			break
		} else if b.Board[RUX][RUY] != 0 {
			*Moves = append(*Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   RUX,
				MHY:   RUY,
				Piece: b.Board[x][y],
			})
			break
		} else {
			*Moves = append(*Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   RUX,
				MHY:   RUY,
				Piece: b.Board[x][y],
			})
		}

	}
}
func forminusplusB(Moves *[]model.PiecesMove, x, y int, b *model.Board, maxstep int) {
	for stepLU := 1; stepLU <= maxstep; stepLU++ {
		LUX := x - stepLU //left up x
		LUY := stepLU + y //left up y
		if LUX > 8 || LUY > 8 || LUX < 0 || LUY < 0 {
			break
		} else if b.Board[LUX][LUY] == 99 || b.Board[LUX][LUY] < 0 {
			break
		} else if b.Board[LUX][LUY] != 0 {
			*Moves = append(*Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   LUX,
				MHY:   LUY,
				Piece: b.Board[x][y],
			})
			break
		} else {
			*Moves = append(*Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   LUX,
				MHY:   LUY,
				Piece: b.Board[x][y],
			})
		}
	}
}
func forminusminusB(Moves *[]model.PiecesMove, x, y int, b *model.Board, maxstep int) {
	for stepLD := 1; stepLD <= maxstep; stepLD++ {
		LDX := x - stepLD //left up x
		LDY := y - stepLD //left up y
		if LDX > 8 || LDY > 8 || LDX < 0 || LDY < 0 {
			break
		} else if b.Board[LDX][LDY] == 99 || b.Board[LDX][LDY] < 0 {
			break
		} else if b.Board[LDX][LDY] != 0 {
			*Moves = append(*Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   LDX,
				MHY:   LDY,
				Piece: b.Board[x][y],
			})

			break
		} else {
			*Moves = append(*Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   LDX,
				MHY:   LDY,
				Piece: b.Board[x][y],
			})
		}
	}
}
func forplusminusB(Moves *[]model.PiecesMove, x, y int, b *model.Board, maxstep int) {
	for stepRD := 1; stepRD <= maxstep; stepRD++ {
		RDX := x + stepRD //left up x
		RDY := y - stepRD //left up y
		if RDX > 8 || RDY > 8 || RDX < 0 || RDY < 0 {
			break
		} else if b.Board[RDX][RDY] == 99 || b.Board[RDX][RDY] < 0 {
			break
		} else if b.Board[RDX][RDY] != 0 {
			*Moves = append(*Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   RDX,
				MHY:   RDY,
				Piece: b.Board[x][y],
			})
			break
		} else {
			*Moves = append(*Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   RDX,
				MHY:   RDY,
				Piece: b.Board[x][y],
			})
		}
	}
}
