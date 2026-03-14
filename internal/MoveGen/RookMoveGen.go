package movegen

import "ChessEngineEOF/internal/model"

func RookMoveGen(b *model.BoardP, x, y int) []model.PiecesMove {

	Moves := []model.PiecesMove{}
	if b.White {
		forUp(&Moves, x, y, b, 8)
		forRight(&Moves, x, y, b, 8)
		forLeft(&Moves, x, y, b, 8)
		forDown(&Moves, x, y, b, 8)
	} else {
		forUpB(&Moves, x, y, b, 8)
		forRightB(&Moves, x, y, b, 8)
		forLeftB(&Moves, x, y, b, 8)
		forDownB(&Moves, x, y, b, 8)
	}
	return Moves
}
func forUp(Moves *[]model.PiecesMove, x, y int, b *model.BoardP, maxstep int) {
	for step := 1; step <= maxstep; step++ { // up
		Uy := y + step
		if Uy < 0 || Uy > 8 {
			break
		} else if b.Board[x][Uy] > 0 {
			break
		} else if b.Board[x][Uy] != 0 {
			*Moves = append(*Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   x,
				MHY:   Uy,
				Piece: b.Board[x][y],
			})
			break
		} else {
			*Moves = append(*Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   x,
				MHY:   Uy,
				Piece: b.Board[x][y],
			})

		}
	}
}
func forRight(Moves *[]model.PiecesMove, x, y int, b *model.BoardP, maxstep int) {
	for step := 1; step <= maxstep; step++ { // up
		RX := x + step
		if RX < 0 || RX > 8 {
			break
		} else if b.Board[RX][y] > 0 {
			break
		} else if b.Board[RX][y] != 0 {
			*Moves = append(*Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   RX,
				MHY:   y,
				Piece: b.Board[x][y],
			})
			break
		} else {
			*Moves = append(*Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   RX,
				MHY:   y,
				Piece: b.Board[x][y],
			})

		}
	}
}
func forLeft(Moves *[]model.PiecesMove, x, y int, b *model.BoardP, maxstep int) {
	for step := 1; step <= maxstep; step++ { // up
		LX := x - step
		if LX < 0 || LX > 8 {
			break
		} else if b.Board[LX][y] > 0 {
			break
		} else if b.Board[LX][y] != 0 {
			*Moves = append(*Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   LX,
				MHY:   y,
				Piece: b.Board[x][y],
			})
			break
		} else {
			*Moves = append(*Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   LX,
				MHY:   y,
				Piece: b.Board[x][y],
			})

		}
	}
}
func forDown(Moves *[]model.PiecesMove, x, y int, b *model.BoardP, maxstep int) {
	for step := 1; step <= maxstep; step++ { // up
		Dy := y - step
		if Dy < 0 || Dy > 8 {
			break
		} else if b.Board[x][Dy] > 0 {
			break
		} else if b.Board[x][Dy] != 0 {
			*Moves = append(*Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   x,
				MHY:   Dy,
				Piece: b.Board[x][y],
			})
			break
		} else {
			*Moves = append(*Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   x,
				MHY:   Dy,
				Piece: b.Board[x][y],
			})
		}

	}
}
func forUpB(Moves *[]model.PiecesMove, x, y int, b *model.BoardP, maxstep int) {
	for step := 1; step <= maxstep; step++ {
		Uy := y + step
		if Uy < 0 || Uy > 8 {
			break
		} else if b.Board[x][Uy] < 0 || b.Board[x][Uy] == 99 {
			break
		} else if b.Board[x][Uy] != 0 {
			*Moves = append(*Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   x,
				MHY:   Uy,
				Piece: b.Board[x][y],
			})
			break
		} else if b.Board[x][Uy] != 99 {
			*Moves = append(*Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   x,
				MHY:   Uy,
				Piece: b.Board[x][y],
			})

		}
	}
}
func forRightB(Moves *[]model.PiecesMove, x, y int, b *model.BoardP, maxstep int) {
	for step := 1; step <= maxstep; step++ { // up
		RX := x + step
		if RX < 0 || RX > 8 {
			break
		} else if b.Board[RX][y] < 0 || b.Board[RX][y] == 99 {
			break
		} else if b.Board[RX][y] != 0 {
			*Moves = append(*Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   RX,
				MHY:   y,
				Piece: b.Board[x][y],
			})
			break
		} else if b.Board[RX][y] != 99 {
			*Moves = append(*Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   RX,
				MHY:   y,
				Piece: b.Board[x][y],
			})

		}
	}
}
func forLeftB(Moves *[]model.PiecesMove, x, y int, b *model.BoardP, maxstep int) {
	for step := 1; step <= maxstep; step++ { // up
		LX := x - step
		if LX < 0 || LX > 8 {
			break
		} else if b.Board[LX][y] < 0 || b.Board[LX][y] == 99 {
			break
		} else if b.Board[LX][y] != 0 {
			*Moves = append(*Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   LX,
				MHY:   y,
				Piece: b.Board[x][y],
			})
			break
		} else if b.Board[LX][y] != 99 {
			*Moves = append(*Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   LX,
				MHY:   y,
				Piece: b.Board[x][y],
			})

		}
	}
}
func forDownB(Moves *[]model.PiecesMove, x, y int, b *model.BoardP, maxstep int) {
	for step := 1; step <= maxstep; step++ { // up
		Dy := y - step
		if Dy < 0 || Dy > 8 {
			break
		} else if b.Board[x][Dy] < 0 || b.Board[x][Dy] == 99 {
			break
		} else if b.Board[x][Dy] != 0 {
			*Moves = append(*Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   x,
				MHY:   Dy,
				Piece: b.Board[x][y],
			})
			break
		} else if b.Board[x][Dy] != 99 {
			*Moves = append(*Moves, model.PiecesMove{
				MFX:   x,
				MFY:   y,
				MHX:   x,
				MHY:   Dy,
				Piece: b.Board[x][y],
			})
		}

	}
}
