package sheets

import (
	"testing"
)

func TestGetSheetID(t *testing.T) {
	testCases := []struct {
		url      string
		expected string
	}{
		{
			url:      "https://docs.google.com/spreadsheets/d/123456abcdefg/edit#gid=0",
			expected: "123456abcdefg",
		},
		{
			url:      "https://docs.google.com/spreadsheets/d/987654zyxwvu/edit#gid=0",
			expected: "987654zyxwvu",
		},
		{
			url:      "https://docs.google.com/spreadsheets/d/1z2y3x4w5v6u7t/edit#gid=0",
			expected: "1z2y3x4w5v6u7t",
		},
		{
			url:      "https://docs.google.com/spreadsheets/d/abcdef123456/edit#gid=0",
			expected: "abcdef123456",
		},
	}

	for _, tc := range testCases {
		actual, err := getSheetID(tc.url)
		if err != nil {
			t.Errorf("Unexpected error for url %s: %v", tc.url, err)
		}

		if actual != tc.expected {
			t.Errorf("Expected sheet ID %s for url %s, but got %s", tc.expected, tc.url, actual)
		}
	}
}
