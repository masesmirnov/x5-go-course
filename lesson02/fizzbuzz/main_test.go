package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	word := fizzBuzz(1)
	if word != "" {
		t.Errorf("fizzBuzz(1) = %q, want empty", word)
	}
	word = fizzBuzz(3)
	if word != "Fizz" {
		t.Errorf("fizzBuzz(3) = %q, want Fizz", word)
	}
	word = fizzBuzz(5)
	if word != "Buzz" {
		t.Errorf("fizzBuzz(5) = %q, want Buzz", word)
	}
	word = fizzBuzz(15)
	if word != "FizzBuzz" {
		t.Errorf("fizzBuzz(15) = %q, want FizzBuzz", word)
	}
}
