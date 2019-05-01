package main

import (
	"fmt"
	"testing"
)

var charTest = []struct {
	name   string
	input  string
	counts [4]int
}{
	{"all counts populated test", "hello123, world", [4]int{3, 7, 3, 2}},
	{"no digits test", "gltmx", [4]int{0, 5, 0, 0}},
	{"empty string test", "", [4]int{0, 0, 0, 0}},
	{"no letters test", "1234 ,*&", [4]int{0, 0, 4, 4}},
	{"a very long test", "skip-me", [4]int{0, 0, 0, 0}},
	{"valid string", "Mary had a l1tt3e lamb.", [4]int{5, 11, 2, 5}},
}

func TestCountChars(t *testing.T) {
	for _, test := range charTest {

		if testing.Short() && test.name == "a very long test" {
			t.Skip("skipped test as we are in a hurry")
		}

		vowels, consonants, digits, other := CountChars(test.input)
		if vowels != test.counts[0] {
			t.Errorf("%s: Vowel count mismatch: expected %d, got %d", test.name, test.counts[0], vowels)
		}
		if consonants != test.counts[1] {
			t.Errorf("%s: Consonants count mismatch: expected %d, got %d", test.name, test.counts[1], consonants)
		}
		if digits != test.counts[2] {
			t.Errorf("%s: Digits count mismatch: expected %d, got %d", test.name, test.counts[2], digits)
		}
		if other != test.counts[3] {
			t.Errorf("%s: Other count mismatch: expected %d, got %d", test.name, test.counts[3], other)
		}
	}
}

func ExampleCountChars() {
	vowels, consonants, digits, other := CountChars("Mary had a l1tt3e lamb.")

	fmt.Println("", vowels, consonants, digits, other)
	// Output: 5 11 2 5
}

func BenchmarkCountChars(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, _, _ = CountChars("Mary had a l1tt3e lamb.")
		
	}
}
