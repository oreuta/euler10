package primes

import (
	"testing"
)

func TestPrimeSumForValidN(t *testing.T) {
	var actSum int64
	type dataPoint struct {
		N   int64
		Sum int64
	}
	var data []dataPoint = []dataPoint{
		{2000000, 142913828922},
		{2, 2},
		{3, 5},
		{4, 5},
		{5, 10},
		{6, 10},
		{100, 1060},
	}

	for _, exp := range data {
		actSum, _, _ = PrimeSum(exp.N, false, 0)
		if actSum != exp.Sum {
			t.Errorf("For n=%v the sum must be %v but is %v", exp.N, exp.Sum, actSum)
		}
	}
}

func TestPrimeSumAndListFor10(t *testing.T) {
	var actSum int64
	var actList []int64
	var expSum int64 = 17
	var expList []int64 = []int64{2, 3, 5, 7}
	var n int64 = 10

	actSum, actList, _ = PrimeSum(n, true, 0)
	if actSum != expSum {
		t.Errorf("Sum must be %v but is %v", expSum, actSum)
	}
	if !isEquals(actList, expList) {
		t.Errorf("List must be %v but is %v", expList, actList)
	}
}

func TestEmptyRangeErrorFor0(t *testing.T) {
	emptyRangeErrorCheck(t, 0)
}

func TestEmptyRangeErrorFor1(t *testing.T) {
	emptyRangeErrorCheck(t, 1)
}

func emptyRangeErrorCheck(t *testing.T, n int64) {
	_, _, err := PrimeSum(n, true, 0)
	if err != ErrEmptyRange {
		t.Errorf("For n=%v Rrror must be %v but is %v", n, ErrEmptyRange, err)
	}
}

func isEquals(a []int64, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
