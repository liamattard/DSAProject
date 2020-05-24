package dpll

import (
	"DsaUomProject/dpll/models"
	"bufio"
	"fmt"
	"os"
)

var input string

var expression [][]models.Literal

func readInput() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your CNF String: ")
	fmt.Print("-> ")

	input, _ = reader.ReadString('\n')

}

func nextChar() byte {

	if input[0] == '\n' {
		return '\n'
	}
	char := input[0]
	input = input[1:]

	return char
}

func isIn(b byte) bool {

	constants := [4]byte{models.W, models.X, models.Y, models.Z}

	for _, a := range constants {
		if a == b {
			return true
		}
	}
	return false
}

//ParseInput : Read's User Input and performs checks
//to check semantics of the CNF.
func ParseInput() [][]models.Literal {

	readInput()

	validity := true
	a := nextChar()

	for a != '\n' && validity == true {
		var clause []models.Literal

		if a == '(' {

			for a != ')' && validity == true {

				a = nextChar()

				if isIn(a) {

					clause = append(clause, models.Literal{Symbol: a, Validity: true})

				} else if a == '!' {

					a = nextChar()

					if isIn(a) {

						clause = append(clause, models.Literal{Symbol: a, Validity: false})

					}

				} else {

					validity = false

				}

				a = nextChar()

				if a != ',' && a != ')' {

					validity = false

				}

			}

		} else {

			validity = false

		}

		if validity != false {

			a = nextChar()
			expression = append(expression, clause)

		}

	}

	if validity == false {

		fmt.Println("ERROR IN PARSING")
		os.Exit(1)

	} else {
		fmt.Println("Expression: ")
		fmt.Printf("%v", expression)
		fmt.Println("")
	}

	return expression

}
