// Function Return Value
// function bisa mengembalikan data
// Untuk memberi tahu function mengembalikan data, kita harus menuliskan tipe data kembali dari function tersebut
// Jika function tersebut di deklarasikan dengan tipe data pengembalian, maka wajib di dalam functionnya kita harus mengembalikan data
// Untuk mengembalikan data dari function, kita bisa menggunakan kata kunci return, diikuti dengan datanya

package main

import "fmt"

func getHello(names string)string{
	return "Hello" + names
}

// bisa juga seperti ini
func getVHello(nama string)string{
	hello := "Hello" + nama
	return hello
}

func main(){
	result := getHello("Iqbal")
	fmt.Println(result)

	fmt.Println(getHello("Budi"))
	fmt.Println(getHello("Joko"))
}