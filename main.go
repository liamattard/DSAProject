package main

import (
	"DsaUomProject/dpll"
	"fmt"
)

func main() {

	input := dpll.ParseInput()
	sat, proof := dpll.Dpll(input, nil)

	if sat {
		fmt.Println("SAT")
		dpll.PrintProof(proof)
	} else {
		fmt.Println("UNSAT")
	}

}
