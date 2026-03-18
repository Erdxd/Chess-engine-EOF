package search

import (
	movegen "ChessEngineEOF/internal/MoveGen"
	"ChessEngineEOF/internal/evaluate"
	"ChessEngineEOF/internal/model"
	"ChessEngineEOF/internal/rules"
)

func MiniMax(deep int, b *model.BoardP, Max bool) int {
	var best int
	var CanMove bool
	if deep == 0 {
		val := evaluate.EvaluateBoard(b)

		return val
	}
	Moves := movegen.GenerateMoves(b)
	if len(Moves) == 0 && rules.IsMate(b) {
		return evaluate.EvaluateBoard(b)
	} else if len(Moves) == 0 {
		return evaluate.EvaluateBoard(b)
	}
	if len(Moves) != 0 {
		CanMove = true
	}
	if Max {
		score := 0
		best = -999999
		for _, move := range Moves {
			b, wasMove, wasPiece, Pstr := rules.ApplyMove(CanMove, b.White, b, &move)
			if wasMove {
				score = MiniMax(deep-1, b, !Max)
				*b = rules.UndoMove(b, wasPiece, &Pstr)
			} else {
				continue
			}

			if score > best {
				best = score
			}

		}
	} else {
		score := 0
		best = 999999
		for _, move := range Moves {
			b, wasMove, wasPiece, Pstr := rules.ApplyMove(CanMove, b.White, b, &move)
			if wasMove {
				score = MiniMax(deep-1, b, !Max)
				*b = rules.UndoMove(b, wasPiece, &Pstr)
			} else {
				continue
			}
			if score < best {
				best = score
			}

		}

	}
	return best

}
