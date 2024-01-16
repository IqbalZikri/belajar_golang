// Konversi Tipe Data
// Konversi tipe data digunakan saat kita ingin mengubah satu tipe data ke tipe data lain

package main

import "fmt"

func main(){
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name = "Tubagus Muhammad Iqbal Zikri"
	var ASCII = name[2]
	var ASCIIString = string(ASCII)

	fmt.Println(name)
	fmt.Println(ASCII)
	fmt.Println(ASCIIString)
}