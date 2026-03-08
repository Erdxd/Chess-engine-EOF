package movegen

import "ChessEngineEOF/internal/model"

func RookMoveGen(b *model.Board, x, y int) []model.PiecesMove {

	Moves := []model.PiecesMove{}
	if b.White {

		for step := 1; step <= 8; step++ { // up
			Uy := y + step
			if Uy < 0 || Uy > 8 {
				break
			} else if b.Board[x][Uy] > 0 {
				break
			} else if b.Board[x][Uy] != 0 {
				Moves = append(Moves, model.PiecesMove{
					MFX:   x,
					MFY:   y,
					MHX:   x,
					MHY:   Uy,
					Piece: b.Board[x][y],
				})
				break
			} else {
				Moves = append(Moves, model.PiecesMove{
					MFX:   x,
					MFY:   y,
					MHX:   x,
					MHY:   Uy,
					Piece: b.Board[x][y],
				})

			}
		}
		for step := 1; step <= 8; step++ { // up
			RX := x + step
			if RX < 0 || RX > 8 {
				break
			} else if b.Board[RX][y] > 0 {
				break
			} else if b.Board[RX][y] != 0 {
				Moves = append(Moves, model.PiecesMove{
					MFX:   x,
					MFY:   y,
					MHX:   RX,
					MHY:   y,
					Piece: b.Board[x][y],
				})
				break
			} else {
				Moves = append(Moves, model.PiecesMove{
					MFX:   x,
					MFY:   y,
					MHX:   RX,
					MHY:   y,
					Piece: b.Board[x][y],
				})

			}
		}
		for step := 1; step <= 8; step++ { // up
			LX := x - step
			if LX < 0 || LX > 8 {
				break
			} else if b.Board[LX][y] > 0 {
				break
			} else if b.Board[LX][y] != 0 {
				Moves = append(Moves, model.PiecesMove{
					MFX:   x,
					MFY:   y,
					MHX:   LX,
					MHY:   y,
					Piece: b.Board[x][y],
				})
				break
			} else {
				Moves = append(Moves, model.PiecesMove{
					MFX:   x,
					MFY:   y,
					MHX:   LX,
					MHY:   y,
					Piece: b.Board[x][y],
				})

			}
		}
		for step := 1; step <= 8; step++ { // up
			Dy := y - step
			if Dy < 0 || Dy > 8 {
				break
			} else if b.Board[x][Dy] > 0 {
				break
			} else if b.Board[x][Dy] != 0 {
				Moves = append(Moves, model.PiecesMove{
					MFX:   x,
					MFY:   y,
					MHX:   x,
					MHY:   Dy,
					Piece: b.Board[x][y],
				})
				break
			} else {
				Moves = append(Moves, model.PiecesMove{
					MFX:   x,
					MFY:   y,
					MHX:   x,
					MHY:   Dy,
					Piece: b.Board[x][y],
				})
			}

		}

	} else {
		for step := 1; step <= 8; step++ {
			Uy := y + step
			if Uy < 0 || Uy > 8 {
				break
			} else if b.Board[x][Uy] < 0 || b.Board[x][Uy] == 99 {
				break
			} else if b.Board[x][Uy] != 0 {
				Moves = append(Moves, model.PiecesMove{
					MFX:   x,
					MFY:   y,
					MHX:   x,
					MHY:   Uy,
					Piece: b.Board[x][y],
				})
				break
			} else {
				Moves = append(Moves, model.PiecesMove{
					MFX:   x,
					MFY:   y,
					MHX:   x,
					MHY:   Uy,
					Piece: b.Board[x][y],
				})

			}
		}
		for step := 1; step <= 8; step++ { // up
			RX := x + step
			if RX < 0 || RX > 8 {
				break
			} else if b.Board[RX][y] < 0 || b.Board[RX][y] == 99 {
				break
			} else if b.Board[RX][y] != 0 {
				Moves = append(Moves, model.PiecesMove{
					MFX:   x,
					MFY:   y,
					MHX:   RX,
					MHY:   y,
					Piece: b.Board[x][y],
				})
				break
			} else {
				Moves = append(Moves, model.PiecesMove{
					MFX:   x,
					MFY:   y,
					MHX:   RX,
					MHY:   y,
					Piece: b.Board[x][y],
				})

			}
		}
		for step := 1; step <= 8; step++ { // up
			LX := x - step
			if LX < 0 || LX > 8 {
				break
			} else if b.Board[LX][y] < 0 || b.Board[LX][y] == 99 {
				break
			} else if b.Board[LX][y] != 0 {
				Moves = append(Moves, model.PiecesMove{
					MFX:   x,
					MFY:   y,
					MHX:   LX,
					MHY:   y,
					Piece: b.Board[x][y],
				})
				break
			} else {
				Moves = append(Moves, model.PiecesMove{
					MFX:   x,
					MFY:   y,
					MHX:   LX,
					MHY:   y,
					Piece: b.Board[x][y],
				})

			}
		}
		for step := 1; step <= 8; step++ { // up
			Dy := y - step
			if Dy < 0 || Dy > 8 {
				break
			} else if b.Board[x][Dy] < 0 || b.Board[x][Dy] == 99 {
				break
			} else if b.Board[x][Dy] != 0 {
				Moves = append(Moves, model.PiecesMove{
					MFX:   x,
					MFY:   y,
					MHX:   x,
					MHY:   Dy,
					Piece: b.Board[x][y],
				})
				break
			} else {
				Moves = append(Moves, model.PiecesMove{
					MFX:   x,
					MFY:   y,
					MHX:   x,
					MHY:   Dy,
					Piece: b.Board[x][y],
				})
			}

		}
	}
	return Moves
}
