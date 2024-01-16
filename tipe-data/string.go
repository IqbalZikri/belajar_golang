// String
// String adalah Kumpulan karakter
// Jumlah karakter bisa dari nol sampe tak terhingga
// dipresentasikan dengan kata kunci string
// Nilai data String diawali dengan tanda petik 2 (" ")

package main

import "fmt"

func main(){
	fmt.Println("Tubagus")
	fmt.Println("Muhammad Iqbal")
	fmt.Println("Zikri")

	fmt.Println(len("Tubagus"))
	// len("") digunakan untuk menghitung jumlah string
	fmt.Println("Muhammad Iqbal"[0])
	// [0] digunakan untuk mengambil karakter pada string (sesuai indeks) data yang keluar akan berupa kode ASCII
	fmt.Println("Zikri"[3])
}