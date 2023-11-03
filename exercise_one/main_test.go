package main

import "testing"

func Test_ReviewExercises(t *testing.T) {
	cases := []struct {
		testName       string
		studentName    string
		grade1         float64
		grade2         float64
		grade3         float64
		grade4         float64
		expectedResult string
	}{
		{
			testName:       "should return student is reproved",
			studentName:    "Student 1",
			grade1:         0,
			grade2:         8,
			grade3:         5,
			grade4:         2,
			expectedResult: "Student 1: REPROVADO",
		},
		{
			testName:       "should return student is approved",
			studentName:    "Student 1",
			grade1:         7,
			grade2:         8,
			grade3:         6,
			grade4:         7,
			expectedResult: "Student 1: APROVADO",
		},
		{
			testName:       "should return student is congratulated",
			studentName:    "Student 1",
			grade1:         10,
			grade2:         9,
			grade3:         9.5,
			grade4:         10,
			expectedResult: "Student 1: PARABÉNS",
		},
	}

	for _, tt := range cases {
		t.Run(tt.testName, func(t *testing.T) {
			result := reviewStudents(tt.studentName, tt.grade1, tt.grade2, tt.grade3, tt.grade4)
			if result != tt.expectedResult {
				t.Errorf("Test Failed: got = %s, want = %s", result, tt.expectedResult)
			}
		})
	}
}
