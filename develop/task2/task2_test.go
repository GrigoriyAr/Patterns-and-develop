package main

import "testing"

func TestSeq(t *testing.T) {
	tests := []struct {
		test        string
		expectation string
	}{
		{"a4", "aaaa"},
		{"", ""},
		{"a2b2c2", "aabbcc"},
		{"aaa3bbb2", "aaaaabbbb"},
	}

	for _, test := range tests {
		unpack, _ := Unpack(test.test)

		if unpack != test.expectation {
			t.Error("Expected result not equal: ", unpack, "!=", test.expectation)
		}
	}
}
