package config

import (
	"fmt"
	"testing"
)

func TestLoad(t *testing.T) {
	Init(".env.test")
	var tests = []struct {
		name string
		want string
	}{
		{"TEST_1", "test 1 47 2"},
		{"TEST_2", "harf"},
		{"TEST_3", "47000"},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.name)
		t.Run(testname, func(t *testing.T) {
			ans := Load(tt.name)
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
