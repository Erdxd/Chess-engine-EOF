package search

import (
	movegen "ChessEngineEOF/internal/MoveGen"
	"ChessEngineEOF/internal/model"
	"ChessEngineEOF/internal/rules"
)

func TheBestMove(b *model.BoardP, deep int) model.PiecesMove {
	BestMove := model.PiecesMove{}
	Moves := movegen.GenerateMoves(b)
	if b.White {

		bestscore := -999999

		for _, move := range Moves {
			b, wasmove, waspieces, Move := rules.ApplyMove(true, b.White, b, &move)
			if !wasmove {
				continue
			}
			score := MiniMax(deep-1, b, false)

			*b = rules.UndoMove(b, waspieces, &Move)
			if bestscore < score {
				bestscore = score
				BestMove = move
			}
		}
		return BestMove
	} else {
		bestScore := 9999999 // Ищем самое маленькое число
		for _, move := range Moves {
			b, wasMove, wasPiece, pStr := rules.ApplyMove(true, b.White, b, &move)
			if !wasMove {
				continue
			}

			score := MiniMax(deep-1, b, true) // Следующий ход белых (Max)
			*b = rules.UndoMove(b, wasPiece, &pStr)

			if score < bestScore {
				bestScore = score
				BestMove = move
			}
		}

	}
	return BestMove
}
