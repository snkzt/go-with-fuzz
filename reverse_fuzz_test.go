package main

import (
  "testing"
  "unicode/utf8"
)

// Fuzzing comes up with inputs for the code, and may identify
// edge cases that the original test cases didn't cover. 
func FuzzReverse(f *testing.F) {
  testCases := []string{"Hello, world", " ", "!12345"}
  for _, tc := range testCases {
    // https://go.dev/security/fuzz/
    f.Add(tc)  // Use f.Add to provide a seed corpus
  }
  f.Fuzz(func(t *testing.T, original string) {
    // The fuzz test will run until it encounters a failing input
    // unless it passes the -fuzztime flag.
    // The default is to run forever if no failures occur. (Ctrl+C to stop)
    // go test -fuzz=Fuzz  
    reverse, err1 := Reverse(original)
    if err1 != nil {
      return
    }
    doubleReverse, err2 := Reverse(reverse)
    if err2 != nil {
      return
    }
    // In fuzzing, you cannot predict expected output,
    // since no control over inputs but can check the existing
    // facts as below.
    if original != doubleReverse {
      t.Errorf("Before: %q, after: %q", original, doubleReverse)
    }
    if utf8.ValidString(original) && !utf8.ValidString(reverse) {
      t.Errorf("Reverse produced invalid UTF-8 string %q", reverse)
    }
  })
}

