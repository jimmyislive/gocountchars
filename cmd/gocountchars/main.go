package main

import (
	"flag"
	"fmt"
)

// CountChars counts the number of characters in a string
func CountChars(str string) (int, int, int, int) {
	vcount := 0
	ccount := 0
	dcount := 0
	ocount := 0

	for i := 0; i < len(str); i++ {
		switch str[i] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vcount++
		case 'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n', 'p', 'q', 'r', 's', 't', 'v', 'w', 'x', 'y', 'z',
			'B', 'C', 'D', 'F', 'G', 'H', 'J', 'K', 'L', 'M', 'N', 'P', 'Q', 'R', 'S', 'T', 'V', 'W', 'X', 'Y', 'Z':
			ccount++
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			dcount++
		default:
			ocount++
		}

	}

	return vcount, ccount, dcount, ocount
}

func main() {
	inputStr := flag.String("inputStr", "", "The input string")

	flag.Parse()

	vowels, consonants, digits, other := CountChars(*inputStr)
	fmt.Println("Number of vowel chars: ", vowels)
	fmt.Println("Number of consonant chars: ", consonants)
	fmt.Println("Number of digit chars: ", digits)
	fmt.Println("Number of other chars: ", other)
}
