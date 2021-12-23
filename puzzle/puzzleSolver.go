package puzzle

import log "github.com/sirupsen/logrus"

type Question20Solver struct{}

func (q Question20Solver) Solve(board Board) (Boards, Boards, error) {
	boards := make(Boards, 0, 4)
	results := make(Boards, 0, 1)

	log.Printf("Current board is %d\n", board.Count)
	if board.Count >= 200 {
		results = append(results, board)
		return nil, results, nil
	}

	clone := board
	clone.Count += 1

	boards = append(boards, clone)

	return boards, nil, nil
}
