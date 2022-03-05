package src

import (
	"fmt"
	"testing"
)

//TESTING CALCULATE VARIADIC FUNCTION
func TestCalculate(t *testing.T) {

	var calc = []struct {
		input  []string
		output []float64
	}{
		{[]string{"1+2+3+4", "1-2-3-4", "1*2*3*4", "4/2"},
			[]float64{10, -8, 24, 2}},
		{[]string{"6+7+8+9", "6-7-8-9", "6*7*8*9", "6/2"},
			[]float64{30, -18, 3024, 3}},
		{[]string{"32+45+56", "789-67-45", "56*7*45", "65/5"},
			[]float64{133, 677, 17640, 13}},
	}

	for _, val := range calc {
		result := Calculate(val.input...)
		for i := 0; i < len(val.output); i++ {
			if result[i] != val.output[i] {
				t.Errorf("Expected output %v received %v", val.output, result)
				fmt.Println(result[i])
			}
		}
	}
}

//TESTING ADDITION
func TestAddition(t *testing.T) {

	var nums = []struct {
		input  []float64
		output float64
	}{
		{[]float64{1, 3, 5, 6, 7, 9}, 31},
		{[]float64{23, 34, 45, 56, 67, 78, 89}, 392},
		{[]float64{567, 234, 345, 789, 000}, 1935},
	}

	for _, val := range nums {
		result := Addition(val.input)
		if result != val.output {
			t.Errorf("Addition : Expected output %v : Output %v", val.output, result)
		}
	}
}

//TESTING SUBTRACTION
func TestSubtraction(t *testing.T) {
	var nums = []struct {
		input  []float64
		output float64
	}{
		{[]float64{100, 20, 30, 5}, 45},
		{[]float64{200, 90, 10, 20}, 80},
	}

	for _, val := range nums {
		result := Subtraction(val.input)
		if result != val.output {
			t.Errorf("Subtract : Expected output %v : Output %v", val.output, result)
		}
	}
}

//TESTING DIVISION
func TestDivision(t *testing.T) {
	var nums = []struct {
		input  []float64
		output float64
	}{
		{[]float64{3000, 5, 2, 2}, 150},
		{[]float64{20000, 5, 4, 2}, 500},
	}
	for _, val := range nums {
		result := Division(val.input)
		if result != val.output {
			t.Errorf("Division: Expected output %v : output %v", val.output, result)
		}
	}
}

func TestMultiplication(t *testing.T) {
	var nums = []struct {
		input  []float64
		output float64
	}{
		{[]float64{3, 2, 5, 6}, 180},
		{[]float64{4, 2, 6, 7}, 336},
	}

	for _, val := range nums {
		result := Multiplication(val.input)
		if result != val.output {
			t.Errorf("Multiply: Expected output %v: output %v", val.output, result)
		}
	}
}
