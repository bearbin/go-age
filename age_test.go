package age

import (
	"testing"
	"time"
)

type AgeTestCandidate struct {
	BirthDate    time.Time
	CheckingTime time.Time
	ExpectedAge  int
}

var AgeTestCandidates = []AgeTestCandidate{
	{time.Date(2000, 3, 14, 0, 0, 0, 0, time.UTC), time.Date(2010, 3, 14, 0, 0, 0, 0, time.UTC), 10},
	{time.Date(2001, 3, 14, 0, 0, 0, 0, time.UTC), time.Date(2009, 3, 14, 0, 0, 0, 0, time.UTC), 8},
	{time.Date(2004, 6, 18, 0, 0, 0, 0, time.UTC), time.Date(2005, 5, 12, 0, 0, 0, 0, time.UTC), 0},
}

func TestAgeAt(t *testing.T) {
	for _, candidate := range AgeTestCandidates {
		gotAge := AgeAt(candidate.BirthDate, candidate.CheckingTime)
		if gotAge != candidate.ExpectedAge {
			t.Error(
				"For", candidate.BirthDate,
				"Expected", candidate.ExpectedAge,
				"Got", gotAge,
			)
		}
	}
}
