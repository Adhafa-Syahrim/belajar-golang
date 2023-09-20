package main

import (
	"strconv"
	"fmt"
)

func main() {
	//string to integer
	id := "123"
	idInt, _ := strconv.Atoi(id)
	fmt.Printf("%T %v\n", idInt, idInt)

	//integer to string
	uang := 600
	uangstr := strconv.Itoa(uang)
	fmt.Printf("%T %v\n", uangstr, uangstr)

	//float to string
	nilaifloat := 3.14
	stringfloat := strconv.FormatFloat(nilaifloat, 'f', -1, 64)
	stringfloat2 := fmt.Sprintf("%v", nilaifloat)
	fmt.Printf("valuenya adalah %v tipe datanya adalha %T\n", stringfloat, stringfloat)
	fmt.Printf("valuenya adalah %v tipe datanya adalha %T\n", stringfloat2, stringfloat2)

	//string float
	stringfloat3:="6.12"
	nilaifloat3, _ := strconv.ParseFloat(stringfloat3, 32)
	fmt.Printf("valuenya adalah %v tipe datanya adalha %T\n", float32(nilaifloat3), float32(nilaifloat3))
}