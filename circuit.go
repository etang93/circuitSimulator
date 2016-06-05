package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    s "strings"
)

func And(inputC1, inputC2, c chan int) {
	input1 := <-inputC1
	input2 := <-inputC2
	retVal := 0
	if(input1 == 1 && input2 == 1){
		retVal = 1
	}
	c <- retVal
}

func Or(inputC1, inputC2, c chan int) {
	input1 := <-inputC1
	input2 := <-inputC2
	retVal := 0 
	if(input1 == 1 || input2 == 1){
		retVal = 1
	}
	c <- retVal
}

func Not(inputC1, c chan int) {
	input1 := <-inputC1
	retVal := 0
	if(input1 == 0){
		retVal = 1
	}
	c <- retVal
}

func Nand(inputC1, inputC2, c chan int) {
	input1 := <-inputC1
	input2 := <-inputC2
	retVal := 0
	if(input1 == input2){
		retVal = 1
	}
	c <- retVal
}

func Nor(inputC1, inputC2, c chan int) {
	input1 := <-inputC1
	input2 := <-inputC2
	retVal := 0
	if(input1 == 0 && input2 == 0){
		retVal = 1
	}
	c <- retVal
}

func Xor(inputC1, inputC2, c chan int) {
	input1 := <-inputC1
	input2 := <-inputC2
	retVal := 0
	if(input1 != input2){
		retVal = 1
	}
	c <- retVal
}

func intersection(input, output1, output2 chan int) {
	retVal := <- input
	output1 <- retVal
	output2 <- retVal
}

func readFile(filename string) []string{

	var commands []string
	
	file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    	commands = append(commands, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return commands
}

func createChannels(commands []string) []chan int{
	var channels []chan int 
	for _, lines := range commands {
		channels = append(channels,make(chan int))
		if s.Contains(lines, "INTERSECT") {
			channels = append(channels, make(chan int))
		}
	}
	return channels
}

func main() {
	var flipFlopState int
	var clockRequested string
	var clockPulsesPerSec int
	var clockPulsesTot int
	var channels [8]chan int

	filename := "C:/Go/workspace/src/circuit/bin/blah.txt"
	commands := readFile(filename)
	
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
	
	testC1 := make(chan int)
	testC2 := make(chan int)

	go func() {
		testC1 <- 1
		testC2 <- 1
		}()

	go And(testC1, testC2, channels[0])
	valAnd := <-channels[0]

	go func() {
		testC1 <- 1
		testC2 <- 1
		}()

	go Or(testC1, testC2, channels[1])
	valOr := <-channels[1]

	go func() {
		testC1 <- 1
		}()

	go Not(testC1, channels[2])
	valNot := <-channels[2]
	
	go func() {
		testC1 <- 1
		testC2 <- 1
		}()

	go Nand(testC1, testC2, channels[3])
	valNand := <-channels[3]

	go func() {
		testC1 <- 1
		testC2 <- 1
		}()

	go Nor(testC1, testC2, channels[4])
	valNor := <-channels[4]
	
	go func() {
		testC1 <- 1
		testC2 <- 0
		}()

	go Xor(testC1, testC2, channels[5])
	valXor := <-channels[5]

	go func() {
		channels[5] <- 1
		}()

	go intersection(channels[5], channels[6], channels[7])
	intersect1, intersect2 := <-channels[6], <-channels[7]

	fmt.Println("valAnd: ", valAnd)
	fmt.Println("valOr: ", valOr)
	fmt.Println("valNot: ", valNot)	
	fmt.Println("valNand: ", valNand)
	fmt.Println("valNor: ", valNor)
	fmt.Println("valXor: ", valXor)
	fmt.Println("intersect1: ", intersect1, "intersect2: ", intersect2)

	fmt.Println("flipFlopState: ", flipFlopState)
	fmt.Println("clockRequested: ", clockRequested)
	fmt.Println("clockPulsesPerSec: ", clockPulsesPerSec)
	fmt.Println("clockPulsesTot: ", clockPulsesTot)

	for _, r := range commands {
		fmt.Println(r)
	}

	ch := createChannels(commands)

	var count int
	for _, c := range ch {
		count++
		go func() {
			c <- 1
			}()
	}
	fmt.Println("counter: " , count)
}