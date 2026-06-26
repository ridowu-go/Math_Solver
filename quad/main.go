package main

import (
	"fmt"
	"strings"
	"os"
	"strconv"
	"math"
)

func main() {
	var a, b, c int
	Problem :=  os.Args[1]
	CleenProblem := strings.ToLower(strings.ReplaceAll(Problem, " ", ""))
	CleenProblem = strings.ToLower(CleenProblem)
	CleanProblem := strings.ReplaceAll(CleenProblem, "²", "")
	NumX := strings.Count(CleanProblem, "x")
	FirstLetter := string(CleanProblem[0])
	ValueOfB := string(CleanProblem[2])
	SecondValueOfB := string(CleanProblem[3])
	// // validate that a is not equals to 0 and that the first part of the equation exists)
	if NumX != 2 || FirstLetter == "0" {
		fmt.Println("Error: a must not be 0, only variable 'x' is allowed")
		return
	}
	// // Locate the value of a by checking wether x is the first character or has coefficient
	if FirstLetter == "x" && ValueOfB == "x" {
		a = 1
		b = 1
	} else if FirstLetter != "x" && SecondValueOfB == "x" {
		a, _ = strconv.Atoi(FirstLetter)
		b = 1
	} else if FirstLetter != "x" && SecondValueOfB != "x" {
		a, _ = strconv.Atoi(FirstLetter)
		b, _ = strconv.Atoi(SecondValueOfB)
	} else if FirstLetter == "x" && ValueOfB != "x" {
		a = 1
		b, _ = strconv.Atoi(ValueOfB)
	}
	//DETERMINE THE VALUE OF C
	//when a and b == 1, position of c is fifth  or last position i.e CleanProblem[4]
	if FirstLetter == "x" && ValueOfB == "x" {
		ValueOfC := string(CleanProblem[4])
		c, _ = strconv.Atoi(ValueOfC)
	
	//if a has coeficient but b == 1 and viceversa
	} else if (FirstLetter != "x" && SecondValueOfB == "x") || (FirstLetter == "x" && ValueOfB != "x") { 
		c, _ = strconv.Atoi(string(CleanProblem[5]))
	} else if FirstLetter != "x" && SecondValueOfB != "x" {
		c, _ = strconv.Atoi(string(CleanProblem[6]))
	}
	var Sq_Root int
	Sq_Root = (b * b) - (4 * a * c)
	PositiveValue := float64(-5) + math.Sqrt(float64(Sq_Root))
	PositiveValue /= float64(2 * a)
	NegativeValue := float64(-5) - math.Sqrt(float64(Sq_Root))
	NegativeValue /= float64(2 * a)
	// fmt.Printf("Given Equation: %s\n", CleanProblem)
	// fmt.Println(a, b, c, PositiveValue, NegativeValue)
	fmt.Printf("Given Question: %s, value of a = %d, b = %d, c = %d\n", CleenProblem, a, b, c)
	fmt.Printf("x = %v or %v\n", PositiveValue, NegativeValue)
}