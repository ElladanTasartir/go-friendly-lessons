package main

import "testing"

func Test_LowestNumber(t *testing.T) {
	cases := []struct {
		testName       string
		number1        int
		number2        int
		number3        int
		number4        int
		number5        int
		expectedResult int
	}{
		{
			testName:       "should return lowest number between positive numbers",
			number1:        1,
			number2:        2,
			number3:        3,
			number4:        4,
			number5:        5,
			expectedResult: 1,
		},
		{
			testName:       "should return lowest number between negative numbers",
			number1:        -1,
			number2:        -2,
			number3:        -3,
			number4:        -4,
			number5:        -5,
			expectedResult: -5,
		},
		{
			testName:       "should return lowest number between natural numbers",
			number1:        -2,
			number2:        -1,
			number3:        0,
			number4:        1,
			number5:        5,
			expectedResult: -2,
		},
	}

	for _, tt := range cases {
		t.Run(tt.testName, func(t *testing.T) {
			result := lowestNumber(tt.number1, tt.number2, tt.number3, tt.number4, tt.number5)
			if result != tt.expectedResult {
				t.Errorf("Test Failed: got = %d, want = %d", result, tt.expectedResult)
			}
		})
	}
}
