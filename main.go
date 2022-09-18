package main

import "fmt"

func main() {
  // Initialise the input string
  input := "The quick brown fox jumped over the lazy dog"
  // Reverse input
  reverse := Reverse(input)
  doubleReverse := Reverse(reverse)
  
  fmt.Printf("original: %q\n", input)
  fmt.Printf("reverse: %q\n", reverse)
  fmt.Printf("reversed again %q\n", doubleReverse)
}


// Reverse function returns reversed string.
func Reverse(s string) string {
    // Convert string into byte and put into a slice to loop over it.
    rune := []rune(s)
    for i, j := 0, len(rune)-1; i < len(rune)/2; i, j = i+1, j-1 {
      // Swap the word from the front (i) and the back (j)
      rune[i], rune[j] = rune[j], rune[i] 
    }
    return string(rune)
}


