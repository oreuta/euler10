package primes

import (
	"testing"
)

func TestPrimeSumFor2000000(t *testing.T) {
	var actSum int64
	var expSum int64 = 142913828922
	var n int64 = 2000000

	actSum, _, _ = PrimeSum(n, false, 0)
	if actSum != expSum {
		t.Errorf("Sum must be %v but is %v", expSum, actSum)
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
