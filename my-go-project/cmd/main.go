package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	ifcond()
	checkResults()
	checkStock()
	chooseTransport()
	weekday()
	oddEvenDay()
	forLoop()
}

func ifcond() {
	if 20 > 18 {
		fmt.Println("20 is greater than 18")
	}

	x := 12
	y := 34
	if x > y {
		fmt.Println("x is greater than y")
	} else {
		fmt.Println("y is greater than x")
	}
}

func checkResults() {
	marks := 85

	if marks >= 90 {
		fmt.Println("Grade: A")
	} else if marks >= 80 {
		fmt.Println("Grade: B")
	} else if marks >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: D")
	}
}

func checkStock() {
	marketcap := 200
	if marketcap > 100 {
		fmt.Println("Stock is larger cap")
	} else if marketcap > 50 {
		fmt.Println("Stock is mid cap")
	} else {
		fmt.Println("Stock is small cap")
	}
}

func chooseTransport() {
	var distance = 2
	switch distance {
	case 0:
		fmt.Println("Walking")
	case 1:
		fmt.Println("Bicycling")
	case 2:
		fmt.Println("Public Transport")
	default:
		fmt.Println("Driving")
	}
}

func weekday() {
	var day = 4
	switch day {
	case 0:
		fmt.Println("Sunday")
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	default:
		fmt.Println("Invalid day")
	}
}

func oddEvenDay() {
	var day = 6
	switch {
	case day%2 == 0:
		fmt.Println("Even Day")
	default:
		fmt.Println("Odd Day")
	}
}

func forLoop() {
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}
}
