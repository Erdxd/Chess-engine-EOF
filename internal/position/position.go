package position

import (
	"ChessEngineEOF/internal/model"
	"ChessEngineEOF/internal/pieces"
)

func Pos1NoPawn(b *model.BoardP) *model.BoardP {
	b.Board[1][7] = 0
	return b
}
func PosGetPawnByKnight(b *model.BoardP) *model.BoardP {
	for x := 0; x <= 8; x++ {
		for y := 0; y <= 8; y++ {
			b.Board[x][y] = 0
		}
	}
	b.Board[4][4] = pieces.KnightE
	b.Board[3][6] = pieces.Pawn

	return b

}
