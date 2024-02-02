package problems

func IsWinner(player1 []int, player2 []int) int {
	var score1 int
	var score2 int
	x2p1, x2p2 := 1, 1
	strike1, strike2 := 0, 0
	for k, v := range player1 {
		if strike1+1 == k {
			x2p1 = 1
		}
		if strike2+1 == k {
			x2p2 = 1
		}

		score1 += (x2p1 * v)
		score2 += (x2p2 * player2[k])
		if v == 10 {
			strike1 = k + 2
			x2p1 = 2
		}
		if player2[k] == 10 {
			strike2 = k + 2
			x2p2 = 2
		}

	}
	if score1 > score2 {
		return 1
	} else if score1 < score2 {
		return 2
	} else {
		return 0
	}
}
