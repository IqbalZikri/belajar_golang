// Latihan 4
// Buat program yang meminta pengguna memasukkan sebuah huruf. Program kemudian menentukan apakah huruf tersebut adalah vokal atau konsonan.

package main

import "fmt"

func main(){
	var huruf string
	var (
		vokal1 = "a"
		vokal2 = "i"
		vokal3 = "u"
		vokal4 = "e"
		vokal5 = "o"
	)
	fmt.Print("Masukkan huruf : ")
	fmt.Scan(&huruf)
	if huruf != vokal1 && huruf != vokal2 && huruf != vokal3 && huruf != vokal4 && huruf != vokal5 {
		fmt.Println("Anda memasukkan huruf konsonan")
		
	}else{
		fmt.Println("Anda memasukkan huruf vokal")
	}
		
}