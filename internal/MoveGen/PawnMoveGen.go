package movegen

import (
	"ChessEngineEOF/internal/model"
	"log"
)

func PawnMoveGenerate(b *model.BoardP, x, y int) []model.PiecesMove {
	Moves := []model.PiecesMove{}
	log.Println(Moves)
	log.Println(x)
	log.Println(y)
	log.Println(b.Board[x][y])
	log.Println(b.Board[7][y])
	log.Println(b.Board[7][5])

	if b.White {
		if y < 8 && b.Board[x][y+1] == 0 && b.Board[x][y+1] != 99 {
			Moves = append(Moves, model.PiecesMove{

				MFX:   x,
				MFY:   y,
				MHX:   x,
				MHY:   y + 1,
				Piece: b.Board[x][y],
			})

		}
		if y < 8 && b.Board[x][y+2] == 0 && y == 2 && b.Board[x][y+1] == 0 {
			Moves = append(Moves, model.PiecesMove{

				MFX:   x,
				MFY:   y,
				MHX:   x,
				MHY:   y + 2,
				Piece: b.Board[x][y],
			})
		}
		if y < 8 && x > 1 && b.Board[x-1][y] < 0 && b.Move2B == true && b.Board[x-1][y+1] != 99 {
			Moves = append(Moves, model.PiecesMove{

				MFX:   x,
				MFY:   y,
				MHX:   x - 1,
				MHY:   y + 1,
				Piece: b.Board[x][y],
			})

		}
		if y < 8 && x < 8 && b.Board[x+1][y] < 0 && b.Move2B == true && b.Board[x+1][y+1] != 99 {
			Moves = append(Moves, model.PiecesMove{

				MFX:   x,
				MFY:   y,
				MHX:   x + 1,
				MHY:   y + 1,
				Piece: b.Board[x][y],
			})
		}
		if x < 8 && y < 8 && b.Board[x+1][y+1] < 0 && b.Board[x+1][y+1] != 99 {
			Moves = append(Moves, model.PiecesMove{

				MFX:   x,
				MFY:   y,
				MHX:   x + 1,
				MHY:   y + 1,
				Piece: b.Board[x][y],
			})
		}
		if x > 1 && y < 8 && b.Board[x-1][y+1] < 0 && b.Board[x-1][y+1] != 99 {
			Moves = append(Moves, model.PiecesMove{

				MFX:   x,
				MFY:   y,
				MHX:   x - 1,
				MHY:   y + 1,
				Piece: b.Board[x][y],
			})

		}

	} else {
		if y > 1 && b.Board[x][y-1] == 0 && b.Board[x][y-1] != 99 {
			Moves = append(Moves, model.PiecesMove{

				MFX:   x,
				MFY:   y,
				MHX:   x,
				MHY:   y - 1,
				Piece: b.Board[x][y],
			})

		}
		if y > 2 && b.Board[x][y-2] == 0 && y == 7 && b.Board[x][y-1] == 0 {
			Moves = append(Moves, model.PiecesMove{

				MFX:   x,
				MFY:   y,
				MHX:   x,
				MHY:   y - 2,
				Piece: b.Board[x][y],
			})
		}
		if y > 1 && x < 8 && b.Board[x+1][y] > 0 && b.Move2W == true && b.Board[x+1][y-1] != 99 {
			Moves = append(Moves, model.PiecesMove{

				MFX:   x,
				MFY:   y,
				MHX:   x + 1,
				MHY:   y - 1,
				Piece: b.Board[x][y],
			})

		}
		if x > 1 && b.Board[x-1][y] > 0 && b.Move2W == true && b.Board[x-1][y-1] != 99 {
			Moves = append(Moves, model.PiecesMove{

				MFX:   x,
				MFY:   y,
				MHX:   x - 1,
				MHY:   y - 1,
				Piece: b.Board[x][y],
			})
		}
		if x > 1 && b.Board[x-1][y-1] > 0 && b.Board[x-1][y-1] != 99 && y > 1 {
			Moves = append(Moves, model.PiecesMove{

				MFX:   x,
				MFY:   y,
				MHX:   x - 1,
				MHY:   y - 1,
				Piece: b.Board[x][y],
			})
		}
		if x < 8 && y > 1 && b.Board[x+1][y-1] > 0 && b.Board[x+1][y-1] != 99 {
			Moves = append(Moves, model.PiecesMove{

				MFX:   x,
				MFY:   y,
				MHX:   x + 1,
				MHY:   y - 1,
				Piece: b.Board[x][y],
			})

		}
	}
	log.Println(Moves)
	return Moves
}
