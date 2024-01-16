// Break and Continue
// Digunakan kata kunci yang bisa digunakan dalam perulangan
// Break digunakan untuk menghentikan seluruh perulangan
// Continue digunakan untuk menghentikan perulangan yang berjalan dan langusung melanjutkan ke perulangan selanjutnya
package main

import "fmt"

func main(){
	count := 10
	for i := 0; i < count; i++ {
		if i == 5 {
			break // akan berhenti disini / angka selanjutnya tidak akan dieksekusi
		}
		fmt.Println("Perulangan ke", i)
	}

	for i := 0; i < count; i++ {
		if i % 2 == 0 {
			continue // akan berhenti dan melanjutkan hitungan 
		}
		fmt.Println("Hitungan ke", i)
	}
}