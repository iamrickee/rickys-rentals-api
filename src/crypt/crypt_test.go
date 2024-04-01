package crypt

import (
	"fmt"
	"testing"
)

func TestTwoWayEncryption(t *testing.T) {
	var tests = []struct {
		input string
	}{
		{"TEST_1"},
		{"TEST 2"},
		{"TEST april 1, 2024"},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.input)
		t.Run(testname, func(t *testing.T) {
			enc, err := EncryptPW(tt.input)
			if err != nil {
				t.Errorf("got error: %s", err)
			}
			dec, err := DecryptPW(enc)
			if err != nil {
				t.Errorf("got error: %s", err)
			}
			if dec != tt.input {
				t.Errorf("mismatch, decrypt: %s, input: %s", dec, tt.input)
			}
		})
	}
}
