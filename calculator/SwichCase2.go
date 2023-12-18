// use the defulat in the swich case of the calculator :)

// in this file i learin the simple swich case in the golang :)

package main

import (
	"fmt"
	"math"
)

func main() {

	const (
		pluss      string = "+"
		Minues     string = "-"
		Multiplied string = " '*'"
		sub        string = "/"
		power      string = "^"
	)
	var (
		fir_opeartion int
		sec_operation int
		operation     string
	)
	fmt.Println("enter the first operation ")
	fmt.Scanln(&fir_opeartion)
	fmt.Println("enter the operatort between the first and secend operator :)( + , - , * , / , ^ )  ")
	fmt.Scanln(&operation)
	fmt.Println("enter the secend operation of the calcultation :) ")
	fmt.Scanln(&sec_operation)

	//use teh simple swich case :

	switch operation {
	case "+":
		fmt.Println("output", fir_opeartion+sec_operation)

	case "-":
		fmt.Println("output", fir_opeartion-sec_operation)
	case "*":
		fmt.Println("output", fir_opeartion*sec_operation)
	case "^":
		output := math.Pow(float64(fir_opeartion), float64(sec_operation))
		fmt.Println("output :) ", output)
	case "/":
		output := fir_opeartion / sec_operation
		fmt.Println("output", output)
	default:
		fmt.Println(" hey nuuuuub enter the valid operation or fomr :) Nuuuub sag :)  ")

	}

}
