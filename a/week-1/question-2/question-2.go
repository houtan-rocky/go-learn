package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	processText()
}

func shift(message string) string {
	const (
		emoticonStart = 0x1F600
		emoticonEnd   = 0x1F64F
		emoticonCount = emoticonEnd - emoticonStart + 1
	)

	result := ""
	for _, r := range message {
		if unicode.IsLetter(r) {
			shift := rune(r % 5)
			if unicode.IsLower(r) {
				result += string('a' + (r-'a'+shift)%26)
			} else if unicode.IsUpper(r) {
				result += string('A' + (r-'A'+shift)%26)
			}
		} else if r >= emoticonStart && r <= emoticonEnd {
			shift := rune(r % 5)
			result += string(emoticonStart + (r-emoticonStart+shift)%emoticonCount)
		} else {
			result += string(r)
		}
	}
	return result
}

func replaceNumbers(message string, length int) string {
	result := ""

	for _, r := range message {
		if unicode.IsDigit(r) {
			num := int(r)
			transformed := fmt.Sprintf("%d", num*length)
			reversed := reverseLine(transformed)
			result += reversed
		} else {
			result += string(r)
		}
	}

	return result
}

func processSpace(message string) string {
	var result string
	var lastWordLength int

	for _, r := range message {
		if unicode.IsSpace(r) {
			if lastWordLength > 0 {
				result += strings.Repeat("_", lastWordLength)
			}
			// result += string(r)
			lastWordLength = 0
		} else {
			result += string(r)
			lastWordLength++
		}
	}
	return result
}

func reverseLine(message string) string {
	runes := []rune(message)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func processText() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()

	var numOfLines int
	fmt.Sscanf(s.Text(), "%d", &numOfLines)

	textSlice := make([]string, 0, numOfLines)
	for i := 0; i < numOfLines; i++ {
		s.Scan()
		message := s.Text()

		textSlice = append(textSlice, message)
	}

	for _, l := range textSlice {
		var length int
		for range l {
			length++
		}

		spaced := processSpace(l)
		shifted := shift(spaced)
		replaced := replaceNumbers(shifted, length)
		reversed := reverseLine(replaced)
		fmt.Println(reversed)
	}
}
