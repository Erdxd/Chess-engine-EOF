package movegen

import "ChessEngineEOF/internal/model"

func KnightMoveGen(b *model.BoardP, x, y int) []model.PiecesMove {
	Moves := []model.PiecesMove{}
	if b.White {
		if x < 8 && y < 8 && b.Board[x+2][y+1] <= 0 {
			Moves = append(Moves, model.PiecesMove{

				MFX:   x,
				MFY:   y,
				MHX:   x + 2,
				MHY:   y + 1,
				Piece: b.Board[x][y],
			})

		}
		if x < 8 && y < 8 && b.Board[x+1][y+2] <= 0 {
			Moves = append(Moves, model.PiecesMove{

				MFX:   x,
				MFY:   y,
				MHX:   x + 1,
				MHY:   y + 2,
				Piece: b.Board[x][y],
			})

		}
		if x > 1 && y < 8 && b.Board[x-1][y+2] <= 0 {
			Moves = append(Moves, model.PiecesMove{

				MFX:   x,
				MFY:   y,
				MHX:   x - 1,
				MHY:   y + 2,
				Piece: b.Board[x][y],
			})

		}
		if x < 8 && y > 1 && b.Board[x+2][y-1] <= 0 {
			Moves = append(Moves, model.PiecesMove{

				MFX:   x,
				MFY:   y,
				MHX:   x + 2,
				MHY:   y - 1,
				Piece: b.Board[x][y],
			})

		}
		if x < 8 && y > 2 && b.Board[x+1][y-2] <= 0 {
			Moves = append(Moves, model.PiecesMove{

				MFX:   x,
				MFY:   y,
				MHX:   x + 1,
				MHY:   y - 2,
				Piece: b.Board[x][y],
			})

		}
		if x > 1 && y > 2 && b.Board[x-1][y-2] <= 0 {
			Moves = append(Moves, model.PiecesMove{

				MFX:   x,
				MFY:   y,
				MHX:   x - 1,
				MHY:   y - 2,
				Piece: b.Board[x][y],
			})

		}
		if x > 2 && y > 2 && b.Board[x-2][y+1] <= 0 {
			Moves = append(Moves, model.PiecesMove{

				MFX:   x,
				MFY:   y,
				MHX:   x - 2,
				MHY:   y + 1,
				Piece: b.Board[x][y],
			})

		}
		if x > 2 && y > 1 && b.Board[x-2][y-1] <= 0 {
			Moves = append(Moves, model.PiecesMove{

				MFX:   x,
				MFY:   y,
				MHX:   x - 2,
				MHY:   y - 1,
				Piece: b.Board[x][y],
			})

		}

	} else {
		if x < 8 && y < 8 && b.Board[x+2][y+1] >= 0 && b.Board[x+2][y+1] != 99 {
			Moves = append(Moves, model.PiecesMove{

				MFX:   x,
				MFY:   y,
				MHX:   x + 2,
				MHY:   y + 1,
				Piece: b.Board[x][y],
			})

		}
		if x < 8 && y < 8 && b.Board[x+1][y+2] >= 0 && b.Board[x+1][y+2] != 99 {
			Moves = append(Moves, model.PiecesMove{

				MFX:   x,
				MFY:   y,
				MHX:   x + 1,
				MHY:   y + 2,
				Piece: b.Board[x][y],
			})

		}
		if x > 1 && y < 8 && b.Board[x-1][y+2] >= 0 && b.Board[x-1][y+2] != 99 {
			Moves = append(Moves, model.PiecesMove{

				MFX:   x,
				MFY:   y,
				MHX:   x - 1,
				MHY:   y + 2,
				Piece: b.Board[x][y],
			})

		}
		if x < 8 && y > 1 && b.Board[x+2][y-1] >= 0 && b.Board[x+2][y-1] != 99 {
			Moves = append(Moves, model.PiecesMove{

				MFX:   x,
				MFY:   y,
				MHX:   x + 2,
				MHY:   y - 1,
				Piece: b.Board[x][y],
			})

		}
		if x < 8 && y > 2 && b.Board[x+1][y-2] >= 0 && b.Board[x+1][y-2] != 99 {
			Moves = append(Moves, model.PiecesMove{

				MFX:   x,
				MFY:   y,
				MHX:   x + 1,
				MHY:   y - 2,
				Piece: b.Board[x][y],
			})

		}
		if x > 1 && y > 2 && b.Board[x-1][y-2] >= 0 && b.Board[x-1][y-2] != 99 {
			Moves = append(Moves, model.PiecesMove{

				MFX:   x,
				MFY:   y,
				MHX:   x - 1,
				MHY:   y - 2,
				Piece: b.Board[x][y],
			})

		}
		if x > 2 && y < 8 && b.Board[x-2][y+1] >= 0 && b.Board[x-2][y+1] != 99 {
			Moves = append(Moves, model.PiecesMove{

				MFX:   x,
				MFY:   y,
				MHX:   x - 2,
				MHY:   y + 1,
				Piece: b.Board[x][y],
			})

		}
		if x > 2 && y > 1 && b.Board[x-2][y-1] >= 0 && b.Board[x-2][y-1] != 99 {
			Moves = append(Moves, model.PiecesMove{

				MFX:   x,
				MFY:   y,
				MHX:   x - 2,
				MHY:   y - 1,
				Piece: b.Board[x][y],
			})

		}
	}

	return Moves
}
