// Tipe data Array
// Tipe data yang berisi kumpulan data dengan tipe yang sama
// daya tampung array tidak bibsa ditambah setelah array dibuat

package main

import "fmt"

func main(){
	var names[3] string
	names[1] = "tb"
	names[0] = "Muhammad"
	names[2] = "Iqbal"

	fmt.Println(names[1])
	fmt.Println(names[0])

	// membuat array secara langsung
	var values = [3]int{
		90,
		85,
		80,
	}

	fmt.Println(values)

	// function array
	// len(array) = untuk mendapatkan panjang array
	// array[index] = untuk mendapatkan nilai array berdasarkan posisi index
	// array[index] = value == mengubah data di posisi index 

	var nilai = [...] int {
		10,
		20,
		30,
		40,
		50,
	}

	fmt.Println(nilai)
	fmt.Println(len(nilai))
	fmt.Println(nilai[3])
	nilai[4] = 100 // data nilai index ke 4 akan berubah yang tadinya 50 menjadi 100
	fmt.Println(nilai)
}