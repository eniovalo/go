package prime

// Return prime factors in increasing order

import (
	"reflect"
	"testing"
)

const testVersion = 1

// Retired testVersions
// (none) 1c5d360b98fc3a9f59c1e5e2f3a668ff8438419d

var tests = []struct {
	input    int64
	expected []int64
}{
	{1, []int64{}},
	{2, []int64{2}},
	{3, []int64{3}},
	{4, []int64{2, 2}},
	{6, []int64{2, 3}},
	{8, []int64{2, 2, 2}},
	{9, []int64{3, 3}},
	{27, []int64{3, 3, 3}},
	{625, []int64{5, 5, 5, 5}},
	{901255, []int64{5, 17, 23, 461}},
	{93819012551, []int64{11, 9539, 894119}},
}

func TestPrimeFactors(t *testing.T) {
	if TestVersion != testVersion {
		t.Fatalf("Found TestVersion = %v, want %v", TestVersion, testVersion)
	}
	for _, test := range tests {
		actual := Factors(test.input)
		if !reflect.DeepEqual(actual, test.expected) {
			t.Errorf("prime.Factors(%d) = %v; expected %v",
				test.input, actual, test.expected)
		}
	}
}

func BenchmarkPrimeFactors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			Factors(test.input)
		}
	}
}
