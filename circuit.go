package main

import "fmt"

func andGate(input1, input2 int) int {
	retVal := 0
	if(input1 == 1 && input2 == 1){
		retVal = 1
	}
	fmt.Printf("Process in AndGate finished\n")
	return retVal
}

func orGate(input1, input2 int) int {
	retVal := 0 
	if(input1 == 1 || input2 == 1){
		retVal = 1
	}
	return retVal
}

func notGate(input1 int) int {
	retVal := 0
	if(input1 == 0){
		retVal = 1
	}
	return retVal
}

func nandGate(input1, input2 int) int {
	retVal := 0
	if(input1 == input2){
		retVal = 1
	}
	return retVal
}

func norGate(input1, input2 int) int {
	retVal := 0
	if(input1 == 0 && input2 == 0){
		retVal = 1
	}
	return retVal
}

func xorGate(input1, input2 int) int {
	retVal := 0
	if(input1 != input2){
		retVal = 1
	}
	return retVal
}

func main() {
	valAnd := andGate(0, 0)
	valOr := orGate(0, 1)
	valNot := notGate(0)
	valNand := nandGate(0, 0)
	valNor := norGate(1, 1)
	valXor := xorGate(0,1)
	fmt.Print("valAnd: ", valAnd, "\n")
	fmt.Print("valOr: ", valOr, "\n")
	fmt.Print("valNot: ", valNot, "\n")	
	fmt.Print("valNand: ", valNand, "\n")
	fmt.Print("valNor: ", valNor, "\n")
	fmt.Print("valXor: ", valXor, "\n")
}
