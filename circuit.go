package main

import "fmt"

func andGate(input1, input2 int, c chan int) {
	retVal := 0
	if(input1 == 1 && input2 == 1){
		retVal = 1
	}
	c <- retVal
}

func orGate(input1, input2 int, c chan int) {
	retVal := 0 
	if(input1 == 1 || input2 == 1){
		retVal = 1
	}
	c <- retVal
}

func notGate(input1 int, c chan int) {
	retVal := 0
	if(input1 == 0){
		retVal = 1
	}
	c <- retVal
}

func nandGate(input1, input2 int, c chan int) {
	retVal := 0
	if(input1 == input2){
		retVal = 1
	}
	c <- retVal
}

func norGate(input1, input2 int, c chan int) {
	retVal := 0
	if(input1 == 0 && input2 == 0){
		retVal = 1
	}
	c <- retVal
}

func xorGate(input1, input2 int, c chan int) {
	retVal := 0
	if(input1 != input2){
		retVal = 1
	}
	c <- retVal
}

func main() {
	var flipFlopState int
	var clockRequested string
	var clockPulsesPerSec int
	var clockPulsesTot int
	var channels [6]chan int
	for i := range channels {
		channels[i] = make(chan int)
	}
	fmt.Println("Enter in initial state value for flip-flop")
	fmt.Scanf("%d\n", &flipFlopState )
	
	
	for {
		fmt.Println("Do you want a clock? (y/n)")
		
		fmt.Scanf("%s\n", &clockRequested)
		if (clockRequested == "y" || clockRequested == "n") {
			break
		} else {
			fmt.Println("Invalid choice")
		}
	}
	
	if(clockRequested == "y"){
		fmt.Println("How many clock pulses do you want per second?")
		fmt.Scanf("%d\n", &clockPulsesPerSec)
		fmt.Println("How many clock pulses do you want in the simulation?")
		fmt.Scanf("%d\n", &clockPulsesTot)
	}
	
	go andGate(1, 1, channels[0])
	valAnd := <-channels[0]
	go orGate(0, 1, channels[1])
	valOr := <-channels[1]
	go notGate(0, channels[2])
	valNot := <- channels[2]
	go nandGate(0, 0, channels[3])
	valNand := <- channels[3]
	go norGate(1, 1, channels[4])
	valNor := <- channels[4]
	go xorGate(0,1, channels[5])
	valXor := <- channels[5]
	fmt.Print("valAnd: ", valAnd, "\n")
	fmt.Print("valOr: ", valOr, "\n")
	fmt.Print("valNot: ", valNot, "\n")	
	fmt.Print("valNand: ", valNand, "\n")
	fmt.Print("valNor: ", valNor, "\n")
	fmt.Print("valXor: ", valXor, "\n")
	fmt.Print("flipFlopState: ", flipFlopState, "\n")
	fmt.Println("clockRequested: ", clockRequested)
	fmt.Println("clockPulsesPerSec: ", clockPulsesPerSec)
	fmt.Println("clockPulsesTot: ", clockPulsesTot)
}