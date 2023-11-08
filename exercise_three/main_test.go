package main

import "testing"

func Test_HasRepeatingValue(t *testing.T) {
	cases := []struct {
		testName       string
		expectedResult bool
		numbers        []int
	}{
		{
			testName:       "should return true if has repeating value",
			expectedResult: true,
			numbers:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1},
		},
		{
			testName:       "should return false if has no repeating value",
			expectedResult: false,
			numbers:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
		},
	}

	for _, tt := range cases {
		t.Run(tt.testName, func(t *testing.T) {
			result := hasRepeatingValue(tt.numbers)
			if result != tt.expectedResult {
				t.Errorf("Test Failed: got = %t, want = %t", result, tt.expectedResult)
			}
		})
	}
}
