package main

import "testing"

func TestGeneratesCorrectNumeral(t *testing.T) {
	testCases := []struct {
		amount   int
		expected string
	}{
		{1, "I"},
		{4, "IV"},
		{10, "X"},
		{101, "CI"},
		{529, "DXXIX"},
		{1200, "MCC"},
		{3999, "MMMCMXCIX"},
	}

	for _, tc := range testCases {
		result := GetNumeral(tc.amount)

		if result != tc.expected {
			t.Errorf("GetNumeral(%d) = %s; expected %s", tc.amount, result, tc.expected)
		}
	}
}
