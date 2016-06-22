package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    s "strings"
    "sync"
)

func And(inputC1, inputC2, c chan int, wg *sync.WaitGroup) {
	input1 := <-inputC1
	input2 := <-inputC2
	retVal := 0
	if(input1 == 1 && input2 == 1){
		retVal = 1
	}
	
	c <- retVal
	wg.Done()
}

func Or(inputC1, inputC2, c chan int, wg *sync.WaitGroup) {
	input1 := <-inputC1
	input2 := <-inputC2
	retVal := 0 
	if(input1 == 1 || input2 == 1){
		retVal = 1
	}
	//fmt.Println("Finished in Or")
	c <- retVal
	wg.Done()
}

func Not(inputC1, c chan int, wg *sync.WaitGroup) {
	input1 := <-inputC1
	retVal := 0
	if(input1 == 0){
		retVal = 1
	}
	//fmt.Println("Finished in Not")
	c <- retVal
	wg.Done()
}

func Nand(inputC1, inputC2, c chan int, wg *sync.WaitGroup) {
	input1 := <-inputC1
	input2 := <-inputC2
	retVal := 0
	if(input1 == input2){
		retVal = 1
	}
	//fmt.Println("Finished in Nand")
	c <- retVal
	wg.Done()
}

func Nor(inputC1, inputC2, c chan int, wg *sync.WaitGroup) {
	input1 := <-inputC1
	input2 := <-inputC2
	retVal := 0
	if(input1 == 0 && input2 == 0){
		retVal = 1
	}
	//fmt.Println("Finished in Nor")
	c <- retVal
	wg.Done()
}

func Xor(inputC1, inputC2, c chan int, wg *sync.WaitGroup) {
	input1 := <-inputC1
	input2 := <-inputC2
	retVal := 0
	if(input1 != input2){
		retVal = 1
	}
	//fmt.Println("Finished in Xor ", input1)
	c <- retVal
	wg.Done()
}

func intersection(input, output1, output2 chan int, wg *sync.WaitGroup) {
	retVal := <- input
	output1 <- retVal
	output2 <- retVal
	wg.Done()
	//fmt.Println("Finished in intersection")
}

func Output(output1 chan int, outputNum int, wg *sync.WaitGroup){
	fmt.Println("The output for pipe ", outputNum, " is : ", <-output1)
	wg.Done()
	fmt.Println("Finished in Output")
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

//add external input channels
func addChannels(channels []chan int, externals int) []chan int{
	for  i := 0; i < externals; i++{
		channels = append(channels, make(chan int))
	}
	return channels
}

func pipeline(commands []string, channels []chan int, wait *sync.WaitGroup){
	var wg sync.WaitGroup
	for _, lines := range commands {
		split := s.Split(lines, " ")

		switch split[0] {
			case "AND": val1, err:= strconv.Atoi(split[1])
						if err != nil {
							println("err: AND ", err, "split[1]: ", split[1])}
						val2, err:= strconv.Atoi(split[2])
						if err != nil {
							println("err: AND", err, "split[2]: ", split[2])}
						val3, err:= strconv.Atoi(split[3])
						if err != nil {
							println("err AND:", err, "split[3]: ", split[3])}
						wg.Add(1)
						go And(channels[val1], channels[val2], channels[val3], &wg)
			case "OR": 	val1, err:= strconv.Atoi(split[1])
						if err != nil {
							println("err OR:", err, "split[1]: ", split[1])}
						val2, err:= strconv.Atoi(split[2])
						if err != nil {
							println("err OR:", err, "split[2]: ", split[2])}
						val3, err:= strconv.Atoi(split[3])
						if err != nil {
							println("err OR:", err, "split[3]: ", split[3])}
						wg.Add(1)
						go Or(channels[val1], channels[val2], channels[val3], &wg)
			case "NOT": val1, err:= strconv.Atoi(split[1])
						if err != nil {
							println("err: NOT ", err, "split[1]: ", split[1])}
						val2, err:= strconv.Atoi(split[2])
						if err != nil {
							println("err: NOT ", err, "split[2]: ", split[2])}
						wg.Add(1)
						go Not(channels[val1], channels[val2], &wg)
			case "NAND":val1, err:= strconv.Atoi(split[1])
						if err != nil {
							println("err: NAND ", err, "split[1]: ", split[1])}
						val2, err:= strconv.Atoi(split[2])
						if err != nil {
							println("err: NAND ", err, "split[2]: ", split[2])}
						val3, err:= strconv.Atoi(split[3])
						if err != nil {
							println("err: NAND ", err, "split[3]: ", split[3])}
						wg.Add(1)
						go Nand(channels[val1], channels[val2], channels[val3], &wg)
			case "NOR":val1, err:= strconv.Atoi(split[1])
						if err != nil {
							println("err: NOR ", err, "split[1]: ", split[1])}
						val2, err:= strconv.Atoi(split[2])
						if err != nil {
							println("err: NOR ", err, "split[2]: ", split[2])}
						val3, err:= strconv.Atoi(split[3])
						if err != nil {
							println("err: NOR ", err, "split[3]: ", split[3])}
						wg.Add(1)
						go Nor(channels[val1], channels[val2], channels[val3], &wg)
			case "XOR":val1, err:= strconv.Atoi(split[1])
						if err != nil {
							println("err: XOR ", err, "split[1]: ", split[1])}
						val2, err:= strconv.Atoi(split[2])
						if err != nil {
							println("err: XOR ", err, "split[2]: ", split[2])}
						val3, err:= strconv.Atoi(split[3])
						if err != nil {
							println("err: XOR ", err, "split[3]: ", split[3])}
						wg.Add(1)
						go Xor(channels[val1], channels[val2], channels[val3], &wg)
			case "INTERSECT": val1, err:= strconv.Atoi(split[1])
						if err != nil {
							println("err: INTERSECTION ", err, "split[1]: ", split[1])}
						val2, err:= strconv.Atoi(split[2])
						if err != nil {
							println("err: INTERSECTION ", err, "split[2]: ", split[2])}
						val3, err:= strconv.Atoi(split[3])
						if err != nil {
							println("err: INTERSECTION", err, "split[3]: ", split[3])}
						wg.Add(1)
						go intersection(channels[val1], channels[val2], channels[val3], &wg)
			case "OUTPUT": val1, err:= strconv.Atoi(split[1])
						if err != nil {
							fmt.Println("err: OUTPUT ", err, "split[1]: ", split[1])
						}
						wg.Add(1)
						go Output(channels[val1], val1, &wg)
						
			default: fmt.Println("Error in Pipeline", split[0])
		}
	}
	wg.Wait()
	wait.Done()
}


func main() {
	var flipFlopState int
	var clockRequested string
	var clockPulsesPerSec int
	var clockPulsesTot int
	var external int
	var channelNum int
	var numExternals int
	var filename string
	var wg sync.WaitGroup

	fmt.Println("What is the filename with the circuit description?")
	fmt.Scanf("%s\n", &filename)

	fmt.Println("Enter in initial state value for flip-flop")
	fmt.Scanf("%d\n", &flipFlopState )
	
	fmt.Println("How many external values do you have?")
	fmt.Scanf("%d\n", &numExternals)

	commands := readFile(filename)
	ch := createChannels(commands)
	ch = addChannels(ch, numExternals)

	for i := 0; i < numExternals; i++ {
		fmt.Println("What is your external value?")
		fmt.Scanf("%d\n", &external)

		fmt.Println("Which pipe is this external value associated with?")
		fmt.Scanf("%d\n", &channelNum)
		go func() {
			ch[channelNum] <- external
			}()
		
	}
	
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
	wg.Add(1)
	go pipeline(commands, ch, &wg)
	wg.Wait()
	//fmt.Println("ch[15]: ", <-ch[15])
	fmt.Println("Thank you for using this system!")
}