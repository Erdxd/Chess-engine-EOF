package search

import (
	movegen "ChessEngineEOF/internal/MoveGen"
	"ChessEngineEOF/internal/evaluate"
	"ChessEngineEOF/internal/model"
	"ChessEngineEOF/internal/pieces"
	"ChessEngineEOF/internal/rules"
)

func MiniMax(deep int, b *model.BoardP, Max bool) int {
	var best int

	var CanMove bool
	if deep == 0 {
		val := evaluate.EvaluateBoard(b)

		return val
	}
	King, check := rules.IsCheck(b)
	Moves := movegen.GenerateMoves(b)
	if b.White {

		if len(Moves) == 0 && rules.IsMate(b) {
			return pieces.Mate

		} else if !King {
			return pieces.Mate
		} else if len(Moves) == 0 && !check {
			return 0
		}
	} else {
		if len(Moves) == 0 && rules.IsMate(b) {
			return -pieces.Mate

		} else if !King {
			return -pieces.Mate
		} else if len(Moves) == 0 && !check {
			return 0
		}
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
