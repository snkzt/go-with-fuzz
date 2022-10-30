package main

import "testing"

func TestReverse(t *testing.T) {
  testCases := []struct {
    in, want string
  }{
    {"Hello, world", "dlrow ,olleH"},
    {"", ""},
    {"!12345", "54321!"},
  }
  for _, tc := range testCases {
    reverse := Reverse(tc.in)
    if reverse != tc.want {
      t.Errorf("Reverse: %q, want %q", reverse, tc.want)
    }
  }
}
