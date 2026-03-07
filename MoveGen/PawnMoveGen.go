package movegen

import (
	"ChessEngineEOF/internal/model"
)

func PawnMoveGenerate(b *model.Board, x, y int) []model.PiecesMove {
	Moves := []model.PiecesMove{}

	if b.White {
		if b.Board[x][y+1] == 0 && b.Board[x][y+1] != 99 {
			Moves = append(Moves, model.PiecesMove{

				MFX:   x,
				MFY:   y,
				MHX:   x,
				MHY:   y + 1,
				Piece: b.Board[x][y],
			})

		}
		if b.Board[x][y+2] == 0 && y == 2 {
			Moves = append(Moves, model.PiecesMove{

				MFX:   x,
				MFY:   y,
				MHX:   x,
				MHY:   y + 1,
				Piece: b.Board[x][y],
			})
		}
		if b.Board[x-1][y] < 0 && b.Move2B == true && b.Board[x-1][y+1] != 99 {
			Moves = append(Moves, model.PiecesMove{

				MFX:   x,
				MFY:   y,
				MHX:   x - 1,
				MHY:   y + 1,
				Piece: b.Board[x][y],
			})

		}
		if b.Board[x+1][y] < 0 && b.Move2B == true && b.Board[x+1][y+1] != 99 {
			Moves = append(Moves, model.PiecesMove{

				MFX:   x,
				MFY:   y,
				MHX:   x + 1,
				MHY:   y + 1,
				Piece: b.Board[x][y],
			})
		}
		if b.Board[x+1][y+1] < 0 && b.Board[x+1][y+1] != 99 {
			Moves = append(Moves, model.PiecesMove{

				MFX:   x,
				MFY:   y,
				MHX:   x + 1,
				MHY:   y + 1,
				Piece: b.Board[x][y],
			})
		}
		if b.Board[x-1][y+1] < 0 && b.Board[x-1][y+1] != 99 {
			Moves = append(Moves, model.PiecesMove{

				MFX:   x,
				MFY:   y,
				MHX:   x - 1,
				MHY:   y + 1,
				Piece: b.Board[x][y],
			})

		}

	}
	return Moves
}
