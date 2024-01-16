// 											Tipe Data Number
// 										Integer dan FLoating Point

// 											-----Integer-----
// int8, int16, int32, int64
// uint8, uint16, uint32, uint64 (tidak ada nilai minus)

// Tipe Data Floating Point (Nilai Desimal)
// float32, float64, complex64, complex128

//											-----alias-----
// byte = uint8 , rune = int32 , int = minimal int32 , uint = minimal uint32 
package main

import "fmt"

func main(){
	fmt.Println("Satu =", 1)
	fmt.Println("Dua =", 2)
	fmt.Println("Tiga Koma Lima =", 3.5)


	barang := 120
	fmt.Println(float64(barang))
}