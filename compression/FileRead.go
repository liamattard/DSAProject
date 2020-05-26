package compression

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// ReadFile : Reads the ascii file from the 'asciiFiles'
// folder
func ReadFile() string {

	data, err := ioutil.ReadFile("compression/asciiFiles/file.ascii")

	if err != nil {
		fmt.Println("File reading error", err)
	}

	return string(data)
}

// FrequencyCount : Returns a map of each symbol
// in the file by it's count
func FrequencyCount(text string) map[rune]int {

	frequencies := make(map[rune]int)

	for _, letter := range text {

		frequencies[letter] = strings.Count(text, string(letter))

	}

	return frequencies
}
