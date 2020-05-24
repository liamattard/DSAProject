package main

import (
	"DsaUomProject/dpll"
	"fmt"
)

func main() {

	input := dpll.ParseInput()
	sat, proof := dpll.Dpll(input, nil)

	if sat {
		fmt.Println("SAT: ")
		fmt.Printf("%v", proof)
	} else {
		fmt.Println("UNSAT")
	}

}
