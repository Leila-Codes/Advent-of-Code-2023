package races

type Race struct {
	Duration, RecordDistance int
}

func (r Race) PredictDistance(heldMS int) int {
	return heldMS * (r.Duration - heldMS)
}

func (r Race) WinnerHoldsMS() (winningNumbers int) {
	for i := 1; i < r.Duration; i++ {
		if r.PredictDistance(i) > r.RecordDistance {
			winningNumbers++
		}
	}

	return winningNumbers
}
