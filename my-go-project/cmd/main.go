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
	countToHundredByTens()
	continueExample()
	breakStatement()
	rangeStatement()
	useLenToIterate()
	printiIFlessThan()
	callMe()
	myFunc()
	fmt.Println("Sum of 5 and 10 is:", add(5, 10))
	familyName("Smith", 41)
	fmt.Println("Factorial of 5 is:", factorial(5))
	fmt.Println("Subtracting 10 from 20 gives:", subtract(20, 10))
	firstName, lastName := returnName()
	fmt.Println("First Name:", firstName, "Last Name:", lastName)
	structExample()
	var per1 Person
	per1.firstName = "Anil"
	per1.lastName = "Kumar"
	per1.age = 40
	passStruct(per1)
	mapExample()
	housePrices()
	createMapInTwoWays()
	addAndDeleteElements()
	checkForElementInMap()

	rect := Rectangle{length: 10, width: 5}
	fmt.Println("Area of rectangle:", rect.area())
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

func countToHundredByTens() {
	for i := 0; i <= 100; i += 10 {
		fmt.Println("Count ", i)
	}
}

func continueExample() {
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Println("Iteration:", i)
	}
}

func breakStatement() {
	for i := 0; i < 5; i++ {
		if i == 2 {
			break
		}
		fmt.Println("breakStatement:", i)
	}
}

func rangeStatement() {
	numbers := []int{1, 2, 3, 4, 5}
	for i, v := range numbers {
		fmt.Println("Index:", i, "Value:", v)
	}

	marks := []int{85, 90, 78, 92, 88}
	for i, v := range marks {
		fmt.Println("Index:", i, "Value:", v)
	}

	for _, v := range marks {
		fmt.Println("Value:", v)
	}

	var person = map[string]string{"name": "Alice", "city": "Wonderland"}
	for key, value := range person {
		fmt.Println("Key:", key, "Value:", value)
	}

	//only print keys
	for key := range marks {
		fmt.Println("Key:", key)
	}
}

func useLenToIterate() {
	var fruits = []string{"Apple", "Banana", "Cherry"}
	for i := 0; i < len(fruits); i++ {
		fmt.Println("Index:", i, "Value:", fruits[i])
	}
}

func printiIFlessThan() {
	for i := 0; i < 6; i++ {
		fmt.Println("i is less than 6:", i)
	}
}

func callMe() {
	fmt.Println("Hello from callMe!")
}

func myFunc() {
	fmt.Println("Hello from myFunc!")
}

func add(a int, b int) int {
	return a + b
}

func familyName(name string, age int) {
	fmt.Println("Hello Family Name:", name, "Age:", age)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// named return values
func subtract(a int, b int) (diff int) {
	diff = a - b
	return diff
}

func returnName() (firstName string, lastName string) {
	firstName = "Alice"
	lastName = "Wonderland"
	return
}

type Person struct {
	firstName string
	lastName  string
	age       int
}

func structExample() {
	var per1 Person
	per1.firstName = "John"
	per1.lastName = "Doe"
	per1.age = 30

	fmt.Println("Person:", per1.firstName, per1.lastName, "Age:", per1.age)
}

func passStruct(p Person) {
	fmt.Println("Person:", p.firstName, p.lastName, "Age:", p.age)
}

func mapExample() {
	var marks = map[string]int{
		"Math":    90,
		"Science": 85,
		"English": 88,
	}
	for subject, mark := range marks {
		fmt.Println("Subject:", subject, "Mark:", mark)
	}
}

func housePrices() {
	var housePrices = make(map[string]int)
	housePrices["House A"] = 300000
	housePrices["House B"] = 450000
	housePrices["House C"] = 500000

	for house, price := range housePrices {
		fmt.Println("House:", house, "Price:", price)
	}
}

func createMapInTwoWays() {
	var myMap = make(map[string]int)
	fmt.Println("ismyMap nil", myMap == nil)

	var myMap2 map[string]int
	fmt.Println("ismyMap2 nil", myMap2 == nil)

}

func addAndDeleteElements() {
	var myMap = make(map[string]int)
	myMap["A"] = 1
	myMap["B"] = 2
	myMap["C"] = 3

	fmt.Println("Before deletion:", myMap)

	delete(myMap, "B")
	fmt.Println("After deletion:", myMap)

	myMap["D"] = 4
	fmt.Println("After addition:", myMap)
}

func checkForElementInMap() {
	var myMap = make(map[string]int)
	myMap["A"] = 1
	myMap["B"] = 2
	myMap["C"] = 3

	value, exists := myMap["B"]

	fmt.Println("Value for key 'B':", value)
	fmt.Println("Key 'B' exists:", exists)
}

type Rectangle struct {
	length int
	width  int
}

func (r Rectangle) area() int {
	return r.length * r.width
}
