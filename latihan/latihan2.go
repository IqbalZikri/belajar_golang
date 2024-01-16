// Latihan 2
/* Buatlah sebuah fungsi yang menerima satu parameter berupa angka bulat positif n. Fungsi ini
harus menghitung dan mengembalikan hasil dari penjumlahan semua bilangan ganjil dari 1 hingga
n */

package main

import "fmt"

func angkaBulatPositif(N int)int{
	return N
}

func main(){
	nilaiN := angkaBulatPositif(20)
	for i := 1; i < nilaiN; i++ {
		if i % 2 == 1 {
			fmt.Println(i + i) // angka ganjil ditambah angka ganjil
		}
	}
}