package main

import "fmt"

func main(){
	var iei = 8

	if iei == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if iei > 5 {
		fmt.Println("lulus")
	} else if iei == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", iei)
	}

	var sc = 6

	switch sc {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}
}