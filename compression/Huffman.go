package compression

import (
	"DsaUomProject/compression/models"
	"fmt"
	"regexp"
)

// StartFunc : Calls all the required method for
// The Huffman Algorithm.
func StartFunc() {

	text := ReadFile()

	var validID = regexp.MustCompile(`^[A-Za-z0-9]+$+`)
	matched := validID.MatchString(text)

	if !matched {
		fmt.Println("Error, file must contain letters and numbers only")
	}

	frequencies := FrequencyCount(text)

	leaves := compression.BuildLeaves(frequencies)
	fmt.Printf("%v", leaves)

	root := compression.BuildTree(leaves)

	compression.PrintTree(root)
}
