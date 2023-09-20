package main

import ("fmt" 
		"os")

func tambah(i float64, j float64) float64{
	return i + j;
}

func kurang(i float64, j float64) float64{
	return i - j;
}

func kali(i float64, j float64) float64{
	return i * j;
}

func bagi(i float64, j float64) float64{
	return i / j;
}

func main() {
	var i, j float64
	var jenis int

	fmt.Println("Program Kalkulator")
	fmt.Println("Masukkan angka pertama")
	fmt.Scanln(&i)
	fmt.Println("Masukkan angka Kedua")
	fmt.Scanln(&j)

	fmt.Println(" ")
	fmt.Println("Pilih jenis perhitungan")
	fmt.Println("1. Tambah")
	fmt.Println("2. kurang")
	fmt.Println("3. kali")
	fmt.Println("4. bagi")

	fmt.Scanln(&jenis)

	switch jenis {
	case 1:
		fmt.Println(" ")
		fmt.Println("Hasil perhitungan = ", i+j)
		fmt.Println("hasil fungsi return: ", tambah(i,j))
	case 2:
		fmt.Println(" ")
		fmt.Println("Hasil perhitungan = ", i-j)
		fmt.Println("hasil fungsi return: ", kurang(i,j))
	case 3:
		fmt.Println(" ")
		fmt.Println("Hasil perhitungan = ", i*j)
		fmt.Println("hasil fungsi return: ", kali(i,j))
	case 4:
		fmt.Println(" ")
		fmt.Println("Hasil perhitungan = ", i/j)
		fmt.Println("hasil fungsi return: ", bagi(i,j))
	default:
		fmt.Fprintln(os.Stderr, "pilih yang benar, opsi yang anda pilih salah")
		os.Exit(1)	
	}

}
