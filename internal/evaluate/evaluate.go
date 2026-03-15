package evaluate

import "ChessEngineEOF/internal/model"

func EvaluateBoard(b *model.BoardP) (int, int) {
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
	return White, Black
}
