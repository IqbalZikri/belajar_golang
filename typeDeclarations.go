// Type Declarations
// digunakan untuk membuat alias atau nama lain terhadap tipe data yang sudah ada , dengan tujuan agar lebuh mudah dimengerti
// gunakan kata kunci type untuk membuat alias

package main

import "fmt"

func main(){
	type noKTP string

	var ktplur noKTP = "1111"
	fmt.Println(ktplur)

	// noKTP menjadi tipe data string
}