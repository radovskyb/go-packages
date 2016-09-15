package main

import (
	"fmt"
	"unicode"
)

func main() {
	// SimpleFold iterates over Unicode code points equivalent under
	// the Unicode-defined simple case folding.  Among the code points
	// equivalent to rune (including rune itself), SimpleFold returns the
	// smallest rune > r if one exists, or else the smallest rune >= 0.
	//
	// For example:
	//	SimpleFold('A') = 'a'
	//	SimpleFold('a') = 'A'
	//
	//	SimpleFold('K') = 'k'
	//	SimpleFold('k') = '\u212A' (Kelvin symbol, K)
	//	SimpleFold('\u212A') = 'K'
	//
	//	SimpleFold('1') = '1'
	sr := unicode.SimpleFold('\u212A')
	fmt.Println(sr) // wrong way to print!

	fmt.Printf("%+q\n", sr) // simpleFold will return K (Kelvin symbol)

	src := unicode.SimpleFold('你')
	fmt.Printf("%q\n", src)  // simpleFold will return the same character 你
	fmt.Printf("%+q\n", src) // simpleFold will return the unicode codepoint

	sra := unicode.SimpleFold('A')
	fmt.Printf("%+q\n", sra) // simpleFold will return the lower case 'a'

	fmt.Println()
	fmt.Printf("%#U\n", unicode.SimpleFold('A'))      // 'a'
	fmt.Printf("%#U\n", unicode.SimpleFold('a'))      // 'A'
	fmt.Printf("%#U\n", unicode.SimpleFold('K'))      // 'k'
	fmt.Printf("%#U\n", unicode.SimpleFold('k'))      // '\u212A' (Kelvin symbol, K)
	fmt.Printf("%#U\n", unicode.SimpleFold('\u212A')) // 'K'
	fmt.Printf("%#U\n", unicode.SimpleFold('1'))      // '1
}
