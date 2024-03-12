package main

import "testing"

func TestBossBaby(t *testing.T) {
	testCases := []struct {
		s        string
		expected string
	}{
		{"SRSSRRR", "Good boy"},
		{"RSSRR", "Bad boy"},
		{"SSSRRRRS", "Bad boy"},
		{"SSRR", "Good boy"},
		{"SRRSSR", "Bad boy"},
	}

	for _, testCase := range testCases {
		result := bossBaby(testCase.s)
		if result != testCase.expected {
			t.Errorf("bossBaby(%q) = %q; want %q", testCase.s, result, testCase.expected)
		}
	}
}
