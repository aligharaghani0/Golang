// swith case without any swich :)

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

	switch {
	case operation == "+":
		fmt.Println("output", fir_opeartion+sec_operation)

	case operation == "-":
		fmt.Println("output", fir_opeartion-sec_operation)
	case "*" == operation:
		fmt.Println("output", fir_opeartion*sec_operation)
	case operation == "^":
		output := math.Pow(float64(fir_opeartion), float64(sec_operation))
		fmt.Println("output :) ", output)
	case "/" == operation:
		output := fir_opeartion / sec_operation
		fmt.Println("output", output)
	}

}
