package main

import (
	calc "ceceweek2/src"
	"fmt"
)

func main() {
	//arr := []float64{2, 4, 6, 7, 8, 9}
	//arr = append(arr[:2], arr[3:]...)
	//fmt.Println(arr)

	fmt.Println(calc.Calculate("1+2+3+4", "1-2-3-4", "1*2*3*4", "1/2/3/4"))
	fmt.Println(calc.Addition([]float64{1, 2, 3, 4}))
	fmt.Println(calc.Subtraction([]float64{1, 2, 3, 4}))
	fmt.Println(calc.Multiplication([]float64{1, 2, 3, 4}))
	fmt.Println(calc.Division([]float64{1, 2, 3, 4}))
}
