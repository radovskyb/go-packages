package main

import (
	"fmt"
	"unicode"
)

// constant with mixed type runes
const mixed = "\b5Ὂg̀9! ℃ᾭG"

func main() {
	for _, c := range mixed {
		fmt.Printf("For %q:\n", c)

		// IsControl reports whether the rune is a control character.
		// The C (Other) Unicode category includes more code points
		// such as surrogates; use Is(C, r) to test for them.
		if unicode.IsControl(c) {
			fmt.Println("\tis control rune")
		}

		// IsDigit reports whether the rune is a decimal digit.
		if unicode.IsDigit(c) {
			fmt.Println("\tis digit rune")
		}

		// IsGraphic reports whether the rune is defined as a Graphic by Unicode.
		// Such characters include letters, marks, numbers, punctuation, symbols, and
		// spaces, from categories L, M, N, P, S, Zs.
		if unicode.IsGraphic(c) {
			fmt.Println("\tis graphic rune")
		}

		// IsLetter reports whether the rune is a letter (category L).
		if unicode.IsLetter(c) {
			fmt.Println("\tis letter rune")
		}

		// IsLower reports whether the rune is a lower case letter.
		if unicode.IsLower(c) {
			fmt.Println("\tis lower case rune")
		}

		// IsMark reports whether the rune is a mark character (category M).
		if unicode.IsMark(c) {
			fmt.Println("\tis mark rune")
		}

		// IsNumber reports whether the rune is a number (category N).
		if unicode.IsNumber(c) {
			fmt.Println("\tis number rune")
		}

		// IsPrint reports whether the rune is defined as printable by Go. Such
		// characters include letters, marks, numbers, punctuation, symbols, and the
		// ASCII space character, from categories L, M, N, P, S and the ASCII space
		// character.  This categorization is the same as IsGraphic except that the
		// only spacing character is ASCII space, U+0020.
		if unicode.IsPrint(c) {
			fmt.Println("\tis printable rune")
		}

		// IsPrint reports whether the rune is defined as printable by Go. Such
		// characters include letters, marks, numbers, punctuation, symbols, and the
		// ASCII space character, from categories L, M, N, P, S and the ASCII space
		// character.  This categorization is the same as IsGraphic except that the
		// only spacing character is ASCII space, U+0020.
		if !unicode.IsPrint(c) {
			fmt.Println("\tis not printable rune")
		}

		// IsPunct reports whether the rune is a Unicode punctuation character
		// (category P).
		if unicode.IsPunct(c) {
			fmt.Println("\tis punct rune")
		}

		// IsSpace reports whether the rune is a space character as defined
		// by Unicode's White Space property; in the Latin-1 space
		// this is
		//	'\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP).
		// Other definitions of spacing characters are set by category
		// Z and property Pattern_White_Space.
		if unicode.IsSpace(c) {
			fmt.Println("\tis space rune")
		}

		// IsSymbol reports whether the rune is a symbolic character.
		if unicode.IsSymbol(c) {
			fmt.Println("\tis symbol rune")
		}

		// IsTitle reports whether the rune is a title case letter.
		if unicode.IsTitle(c) {
			fmt.Println("\tis title case rune")
		}

		// IsUpper reports whether the rune is an upper case letter.
		if unicode.IsUpper(c) {
			fmt.Println("\tis upper case rune")
		}
	}

}
