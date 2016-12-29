package primes

import (
	"testing"
)

func TestPrimeSumFor2000000(t *testing.T) {
	var actSum uint64
	var expSum uint64 = 142913828922
	var n uint64 = 2000000

	actSum, _, _ = PrimeSum(n, false)
	if actSum != expSum {
		t.Errorf("Sum must be %v but is %v", expSum, actSum)
	}
}

func TestPrimeSumAndListFor10(t *testing.T) {
	var actSum uint64
	var actList []uint64
	var expSum uint64 = 17
	var expList []uint64 = []uint64{2, 3, 5, 7}
	var n uint64 = 10

	actSum, actList, _ = PrimeSum(n, true)
	if actSum != expSum {
		t.Errorf("Sum must be %v but is %v", expSum, actSum)
	}
	if !isEquals(actList, expList) {
		t.Errorf("List must be %v but is %v", expList, actList)
	}
}

func isEquals(a []uint64, b []uint64) bool {
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
