package main

import (
  "fmt"
  "errors"
  "unicode/utf8"
)
func main() {
  // Initialise the input string
  input := "The quick brown fox jumped over the lazy dog"
  // Reverse input
  reverse, reverseErr := Reverse(input)
  doubleReverse, doubleReverseErr := Reverse(reverse)
  
  fmt.Printf("original: %q\n", input)
  fmt.Printf("reverse: %q, err: %v\n", reverse, reverseErr)
  fmt.Printf("reversed again %q err:%v\n", doubleReverse, doubleReverseErr)
}


// Reverse function returns reversed string.
func Reverse(s string) (string, error) {
  if !utf8.ValidString(s) {
    return s, errors.New("input s is not valid UTF-8")
  }
  fmt.Printf("input: %q\n", s)
  // Convert string into byte and put into a slice to loop over it.
  rune := []rune(s)
  fmt.Printf("input: %q\n", rune)
  for i, j := 0, len(rune)-1; i < len(rune)/2; i, j = i+1, j-1 {
    // Swap the word from the front (i) and the back (j)
    rune[i], rune[j] = rune[j], rune[i] 
  }
  return string(rune), nil
}

