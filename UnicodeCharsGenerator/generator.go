package main

import (
	"flag"
	"os"
)

// Unicode has 143859 ("active"|"usable") characters
// Unicode`s character codes go from (0) U+0000 to U+10FFFF (1114111)
// These code points can be decoded by UTF-8, UTF-16 and UTF-32.
// Most people use UTF-8 encoding daily that covers all necessary characters and many "exotic" ones )
// That means that in order to "get" all the unicode characters we need to loop over all
// codes (from 0 to 1114111) and convert each code into rune

// In order to read these runes correctly, we can use UTF-8, UTF-16 and other encodings
// this is done by a text editor that you`re using.

// This is new for me, so I can be wrong about something, I`m sorry if that`s the case

// Sources
// https://en.wikipedia.org/wiki/Unicode
// https://en.wikipedia.org/wiki/List_of_Unicode_characters
// https://www.unicode.org/faq/utf_bom.html
// https://www.thoughtco.com/what-is-unicode-2034272
// https://golangdocs.com/rune-in-golang

func generateUnicodeChars(runesInline int, limit int32) {
	const UNICODEMAX int = 1114111 // ALL codes

	// create output file
	outputFile, err := os.Create("Unicode.txt")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// process given arguments
	if runesInline <= 0 {
		// all characters will be put in one line
		runesInline = UNICODEMAX
	}

	if limit <= 0 {
		// no limit
		limit = int32(UNICODEMAX)
	}

	// looping through all codes in unicode
	var runeCounter uint64 = 0
	for i := 0; i < UNICODEMAX; i++ {

		if runeCounter == uint64(runesInline) {
			// inject a new line in file and reset the counter
			outputFile.Write([]byte("\n"))
			runeCounter = 0
		}
		outputFile.Write([]byte(string(rune(i))))
		runeCounter++

		// the current Unicode code point is the last one
		if i == int(limit) {
			break
		}
	}
}

func main() {
	var RUNESINLINE *int = flag.Int("inline", 0, "How many characters will be placed in one line in the output file before the \\n")
	var LIMIT *int = flag.Int("limit", 0, "Set the limit")
	flag.Parse()

	generateUnicodeChars(*RUNESINLINE, int32(*LIMIT))

}
