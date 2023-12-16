package races

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRace_PredictDistance(t *testing.T) {
	race := Race{
		Duration:       7,
		RecordDistance: 9,
	}

	distance := race.PredictDistance(5)

	assert.Equal(t, 10, distance, "distance did not match up")
}

func testRaceHoldMS(t *testing.T, duration, distance, expectedWinners int) {
	race := Race{
		Duration:       duration,
		RecordDistance: distance,
	}

	actualWinners := race.WinnerHoldsMS()

	assert.Equalf(t, expectedWinners, actualWinners, "there should be this many winning numbers. \n winners are: %v", actualWinners)
}

func TestRace_WinnerHoldsMS(t *testing.T) {
	testRaceHoldMS(t, 7, 9, 4)
	testRaceHoldMS(t, 15, 40, 8)
	testRaceHoldMS(t, 30, 200, 9)
}
