// Latihan 1
// Buatlah program yang meminta dua angka dari pengguna dan menampilkan angka yang lebih besar.

package main

import "fmt"

func perbandingan(a, b int) (nilai1, nilai2 string) {
	nilai1 = fmt.Sprintf("Angka 1 lebih besar dari angka 2 = %d", a)
	nilai2 = fmt.Sprintf("Angka 2 lebih besar dari angka 1 = %d", b)
	return nilai1, nilai2
}

func main() {
	// var angka1 string
	// var angka2 string
	var angka1 int
	var angka2 int
	fmt.Print("Masukkan angka pertama: ")
	fmt.Scan(&angka1)
	fmt.Print("Masukkan angka ke-dua: ")
	fmt.Scan(&angka2)
	// tanpa menggunakan function
	// if angka1 >= angka2 {
	// 	fmt.Printf("Angka 1 lebih besar daripada angka 2 = %s", angka1)
	// } else if angka2 >= angka1 {
	// 	fmt.Printf("Angka 2 lebih besar daripada angka 1 = %s", angka2)
	// }

	// Dengan menggunakan function
	if angka1 >= angka2 {
		nilai1, _ := perbandingan(angka1, angka2)
		fmt.Println(nilai1)
	} else if angka2 >= angka1 {
		_, nilai2 := perbandingan(angka1, angka2)
		fmt.Println(nilai2)
	}
}
