package position

import (
	"ChessEngineEOF/internal/model"
	"ChessEngineEOF/internal/pieces"
)

func MateInOne1(b *model.BoardP) *model.BoardP {
	for x := 1; x <= 8; x++ {
		for y := 0; y <= 8; y++ {
			b.Board[x][y] = 0
		}
	}
	b.Board[5][8] = pieces.KingE
	b.Board[5][6] = pieces.King
	b.Board[2][7] = pieces.Queen

	return b

}
