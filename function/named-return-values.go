// Named Return Values
// Biasanya saat kita membuat return values kita langsung memberikan tipe data return value pada functionnya
// namun kita juga bisa membuat variable secara langsung di tipe data return functionnya

package main

import "fmt"

func getCompleteName()(firstName, middleName, lastName string){
// bisa juga seperti ini
// func getCompleteName()(firstName string, middleName string, lastName string){
	firstName = "Tubagus"
	middleName = "Muhammad"
	lastName = "Iqbal"
	// bisa juga dengan menghapus salah satu variable di atas maka akan menjadi string kosong dan tidak akan terjadi error
	return firstName, middleName, lastName
}

func main(){
	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName)
}