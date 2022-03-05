package src

import (
	"strconv"
	"strings"
)

func Calculate(s ...string) []float64 {

	var finalResult []float64

	// ADDITION CALCULATION IN THE VARIADIC FUNCTION
	for i := 0; i < len(s); i++ {
		if strings.Contains(s[i], "+") {

			addArray := strings.Split(s[i], "+")
			var arr = changeStringSliceToFloatSlice(addArray)
			finalResult = append(finalResult, Addition(arr))
		}

		//SUBTRACTION CALCULATION IN THE VARIADIC FUNCTION
		if strings.Contains(s[i], "-") {
			subArray := strings.Split(s[i], "-")
			var arr = changeStringSliceToFloatSlice(subArray)
			finalResult = append(finalResult, Subtraction(arr))
		}

		//MULTIPLICATION CALCULATION IN THE VARIADIC FUNCTION
		if strings.Contains(s[i], "*") {
			multArray := strings.Split(s[i], "*")
			var arr = changeStringSliceToFloatSlice(multArray)
			finalResult = append(finalResult, Multiplication(arr))
		}

		//DIVISION CALCULATION IN THE VARIADIC FUNCTION
		if strings.Contains(s[i], "/") {
			divArray := strings.Split(s[i], "/")
			var arr = changeStringSliceToFloatSlice(divArray)
			finalResult = append(finalResult, Division(arr))
		}
	}
	return finalResult
}

// Addition  function
func Addition(a []float64) float64 {
	var result float64

	for _, v := range a {
		result += v
	}
	return result
}

// Subtraction  function
func Subtraction(b []float64) float64 {
	var result float64

	for i, v := range b {
		if i == 0 {
			result = v
			continue
		}
		result -= v
	}
	return result
}

// Multiplication  function
func Multiplication(c []float64) float64 {

	var result float64 = 1
	for _, v := range c {
		result *= v
	}
	return result
}

// Division  function
func Division(d []float64) float64 {
	var result float64

	for i, v := range d {
		if i == 0 {
			result = v
			continue
		}
		if i > 0 {
			if v == 0 {
				panic("can't divide by zero, input a real number")
			}
		}
		result /= v
	}
	return result
}

func changeStringSliceToFloatSlice(s []string) []float64 {
	var arr []float64
	for j := 0; j < len(s); j++ {
		subConvert, _ := strconv.ParseFloat(s[j], 64)
		arr = append(arr, subConvert)
	}
	return arr
}
