package evaluate

import (
	"ChessEngineEOF/internal/model"
)

func EvaluateBoard(b *model.BoardP) int {
	White := 0
	Black := 0
	for _, row := range b.Board {
		for _, cell := range row {
			if cell > 0 && cell != 99 {
				White += cell
			} else if cell < 0 {
				Black += cell
			}

		}
	}
	if b.White {
		return White + Black
	} else {
		return -1 * (White + Black)
	}
}
