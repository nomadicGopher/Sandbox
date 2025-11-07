package main

import (
	"fmt"
	"reflect"
)

func checkEmptyNumber[T int | float64](v *T) string {
	if v == nil {
		return ""
	}

	vType := reflect.TypeOf(*v).Kind() // can also use .String() to compare against == "string"

	switch vType {
	case reflect.Int:
		return fmt.Sprintf("%d", *v)
	case reflect.Float64:
		return fmt.Sprintf("%.2f", *v)
	}

	return "Logic Error"
}

func main() {
	input1 := 1
	input1Ptr := &input1
	result1 := checkEmptyNumber(input1Ptr)
	fmt.Printf("Int value: %s\n", result1)

	input2 := 1.00
	input2Ptr := &input2
	result2 := checkEmptyNumber(input2Ptr)
	fmt.Printf("Float value: %s\n", result2)

	var input3Ptr *int
	result3 := checkEmptyNumber(input3Ptr)
	fmt.Printf("Nil value: %s\n", result3)
}
