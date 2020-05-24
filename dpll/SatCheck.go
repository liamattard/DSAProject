package dpll

import (
	"DsaUomProject/dpll/models"
	// "fmt"
)

// Dpll : uses the dpll algorithm to check
// satisfiability of the given expression
func Dpll(expression [][]models.Literal, proof []models.Literal) (bool, []models.Literal) {

	if len(expression) == 0 {

		return true, proof

	}

	for _, clause := range expression {
		if len(clause) == 0 {
			return false, proof
		}
	}

	selectedLiteral := expression[0][0]
	negatedLiteral := models.Literal{Symbol: selectedLiteral.Symbol, Validity: !selectedLiteral.Validity}

	var newExpression [][]models.Literal
	var negatedExpression [][]models.Literal

	// For each unit clause u in Clauses:
	for _, clause := range expression {

		var literalInClause bool = false
		var negLiteralInClause bool = false

		var newClause []models.Literal
		var negatedClause []models.Literal

		for _, literal := range clause {

			// One Literal Rule
			if selectedLiteral == literal {
				literalInClause = true
			}

			if negatedLiteral == literal {
				negLiteralInClause = true
			}

			// Pure Literal Rule
			if !((selectedLiteral.Symbol == literal.Symbol) && (selectedLiteral.Validity != literal.Validity)) {
				newClause = append(newClause, literal)
			}

			if !((negatedLiteral.Symbol == literal.Symbol) && (negatedLiteral.Validity != literal.Validity)) {
				negatedClause = append(negatedClause, literal)
			}

		}

		if !literalInClause {
			newExpression = append(newExpression, newClause)
		}

		if !negLiteralInClause {
			negatedExpression = append(negatedExpression, negatedClause)
		}

		// trivially unsat
		if literalInClause && negLiteralInClause {
			return false, nil
		}

	}

	// Left Split
	sat, proof := Dpll(newExpression, append(proof, selectedLiteral))
	if sat {
		return sat, proof
	}

	// Right Split
	sat, proof = Dpll(negatedExpression, append(proof, negatedLiteral))
	if sat {
		return sat, proof
	}

	// Recursion with -ve literal
	// fmt.Println("")
	// fmt.Printf("new Left expression :%v", newExpression)
	// fmt.Println("")
	// fmt.Printf("new Right expression :%v", negatedExpression)

	return false, nil
}
