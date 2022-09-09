package main

import (
	"testing"
)

func TestValid(t *testing.T) {
	tests := []struct {
		u    string
		p    string
		want bool
		name string
	}{
		{"", "", false, "nil,nil"},
		{"", "pass", false, "nil,pass"},
		{"user", "", false, "use,nil"},
		{"user", "pass", true, "user,pass"},
	}
	for _, tc := range tests {
		ft := func(t *testing.T) {
			t.Parallel()
			user := NewUser(tc.u, tc.p)
			got := user.Valid()
			if got != tc.want {
				t.Fatalf("got: %v, want: %v", got, tc.want)
			}
		}
		t.Run(tc.name, ft)
	}
}
