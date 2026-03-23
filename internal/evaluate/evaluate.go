package evaluate

import (
	"ChessEngineEOF/internal/model"
	"ChessEngineEOF/internal/pieces"
)

func EvaluateBoard(b *model.BoardP) int {
	White := 0
	Black := 0
	for x := 1; x <= 8; x++ {
		for y := 1; y <= 8; y++ {
			if b.Board[x][y] > 0 && b.Board[x][y] != 99 {
				if b.Board[x][y] == pieces.Pawn {
					White += PawnTableWhite[x][y]

				}
				White += b.Board[x][y]
			}
			if b.Board[x][y] < 0 {
				Black += PawnTableBlack[x][y]
			}
			Black += b.Board[x][y]
		}

	}
	return White + Black
}
