package position

import (
	"ChessEngineEOF/internal/model"
	"ChessEngineEOF/internal/pieces"
)

func Fork1(b *model.BoardP) *model.BoardP {
	for x := 0; x <= 8; x++ {
		for y := 0; y <= 8; y++ {
			b.Board[x][y] = 0
		}
	}
	b.Board[2][3] = pieces.KnightE
	b.Board[3][6] = pieces.Queen
	b.Board[5][6] = pieces.King
	b.Board[1][1] = pieces.KingE

	return b

}
func Fork2(b *model.BoardP) *model.BoardP {
	for x := 0; x <= 8; x++ {
		for y := 0; y <= 8; y++ {
			b.Board[x][y] = 0
		}
	}
	b.Board[2][3] = pieces.KnightE
	b.Board[2][7] = pieces.Queen
	b.Board[4][7] = pieces.King

	return b

}
func Fork3(b *model.BoardP) *model.BoardP {
	for x := 0; x <= 8; x++ {
		for y := 0; y <= 8; y++ {
			b.Board[x][y] = 0
		}
	}
	b.Board[2][3] = pieces.KnightE
	b.Board[1][5] = pieces.Queen
	b.Board[3][5] = pieces.King

	return b

}
