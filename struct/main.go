package main

import "fmt"

func main() {
	type Peserta struct{
		NIM string
		Nama string
		NoHp int
		IPK float64
		Hadir bool
	}

	jamal := Peserta{
		NIM: "123",
		Nama: "Jamaludin Abdul Mamat",
		NoHp: 812474565,
		IPK: 3.5,
		Hadir: true,
	}

	fmt.Println(jamal.Nama)
}