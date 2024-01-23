// Latihan 3
// Buat program yang meminta pengguna memasukkan angka dan menentukan apakah angka tersebut genap atau ganjil.

package main

import "fmt"

func main(){
	var angka int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)
	if angka % 2 == 0 {
		fmt.Println("Angka anda termasuk bilangan genap")
	}else {
		fmt.Println("Bilangan ini termasuk bilangan ganjil")
	}
}