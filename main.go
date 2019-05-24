/*
*	This is an introduction to Golang (GO!) I've put together for myself.
*	I will be adding onto this as it also serves as my practice with git.
*
*	run on Terminal using: go run main.go
*	
*	BY: Sebastian Medina	May 23, 2019
*/

// <- Usual comment

// Each Go application must contain a package
package main

// "library" imports
import (
	"fmt"
	"errors"
)

// Struct example
type dog struct {
	breed string
	name string
	age int
}

// main function / where program begins
func main(){
	// Variable declaraction
	fmt.Println("variables:")
	var x int = 5				//Regular
	y := 10						//Quick
	fmt.Println(y, x, "\n")

	// Slices (Arrays which can be added to)
	// To get fixed arrays add a value between brackets []
	fmt.Println("Array (Slice):")
	a := []int{3, 4}
	fmt.Println(a)

	// Adding to array (slices)
	a = append(a, 4)
	fmt.Println(a, "\n")

	// MAPS -- make(map[first_parameter]second_parameter)
	// Like the Dictionary ADT created in java, or Dictionaries in python
	fmt.Println("Maps:")
	firstMap := make(map[string]int)

	firstMap["str1"] = 1
	firstMap["str2"] = 2
	firstMap["str3"] = 3

	fmt.Println(firstMap)

	fmt.Println("Find particular key: ", firstMap["str2"])

	// Delete from map
	delete(firstMap, "str2")
	fmt.Println(firstMap, "\n")

	// LOOPS (the only kind are for)
	fmt.Println("for Loop(check code):")
	for i := 0; i < 4; i++ {
		fmt.Println(i)
	}
	fmt.Println("\n")

	// For acting as while
	fmt.Println("for acting as while Loop:")
	count := 0
	for count < 4{
		fmt.Println(count)
		count++
	}
	fmt.Println("\n")

	// For loop with RANGE (similar to "for each" loop in java)
	fmt.Println("For loop with range")
	arr := []string{"one", "two", "three"}
	for index, value := range arr {
		fmt.Println("Index: ", index, "Value", value)
	}
	fmt.Println("\n")

	// For loop w/ range of MAP
	secondMap := make(map[string]string)
	secondMap["first"] = "1"
	secondMap["two"] = "2"
	secondMap["three"] = "3"

	for key, keyVal := range secondMap {
		fmt.Println("Key: ", key, "Value: ", keyVal)
	}
	fmt.Println("\n")

	// Separate Functions
	fmt.Println("Function calling:\nsqrt(5)")
	result, err := sqrt(5)
	
	if err == nil{
		fmt.Println(result)
	} else{
		fmt.Println(err)	
	}
	fmt.Println("\n")

	// STRUCTS USAGE
	fmt.Println("Struct Usage: ")
	pupOne := dog{breed: "Russell Terrier", name: "Bentley", age: 10}
	fmt.Println(pupOne)
	fmt.Println("Getting individual variables from struct:", pupOne.name, "\n")

	// Pointers
	fmt.Println("Brief pointer walkthrough:")
	m := 5
	increment_fail(m)
	fmt.Println("Without pointer the value doesn't change:", m)

	increment_pass(&m)
	fmt.Println("With address/pointer:", m)
}

// func Name(Parameters) return_Type1, return_Type2,...{}
// We use error since Go has no form of exceptions like Java
func sqrt(a int) (int, error){
	if a < 0{
		return 0, errors.New("Undefined for negative square roots")
	}

	return (a*a), nil
}

func increment_fail(a int) {
	a++
}

func increment_pass(a *int) {
	*a++
}

