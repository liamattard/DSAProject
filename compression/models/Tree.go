package compression

import (
	"fmt"
	"sort"
)

// Node : struct representing a node from
// the huffman tree.
type Node struct {
	Left   *Node
	Right  *Node
	Count  int
	Symbol rune
}

// BuildLeaves : Generates the leaves of the
// huffman tree sorted by count.
func BuildLeaves(letters map[rune]int) []Node {

	var nodes []Node
	var counts []int

	for _, values := range letters {

		counts = append(counts, values)

	}

	sort.Ints(counts)

	for _, count := range counts {

		for key, value := range letters {

			if value == count {

				nodes = append(nodes, Node{Count: count, Symbol: key})
				letters[key] = 0
				break

			}

		}
	}

	return nodes
}

// BuildTree : Builds the Huffman tree.
func BuildTree(nodes []Node) Node {

	for len(nodes) > 1 {
		left, right := nodes[0], nodes[1]
		newNode := Node{Count: (nodes[0].Count + nodes[1].Count)}
		newNode.Left = &left
		newNode.Right = &right

		parentPos := 0

		for _, node := range nodes {

			if node.Count < newNode.Count {

				parentPos++

			}

		}

		nodes = append(nodes, Node{})
		copy(nodes[parentPos+1:], nodes[parentPos:])
		nodes[parentPos] = newNode
		nodes = nodes[2:]

	}

	return nodes[0]
}

// PrintTree : Traverses the tree in prefix order and
// generates the new variable sized code
// representing the symbol
func PrintTree(root Node) {

	var prefix func(node Node, code []string)

	prefix = func(node Node, code []string) {

		if node.Left == nil {
			fmt.Println()
			fmt.Printf("'%s': %s ", string(node.Symbol), code)

		} else {
			left := append(code, "0")
			prefix(*node.Left, left)
			right := append(code, "1")
			prefix(*node.Right, right)
		}

	}

	var code []string
	code = append(code, "")
	prefix(root, code)
}
