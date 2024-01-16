// Switch
// Sama seperti if else
// Lebih

package main

import (
	"fmt"

	"golang.org/x/text/cases"
)

func main(){

	var name string
	fmt.Print("Masukkan nama :")
	fmt.Scanln(&name)

	// switch name {
	// case "TB","tb":
	// 	fmt.Println("TB")
	// case "Bang parjo":
	// 	fmt.Println("Bang Parjo")
	// case "Bang Billy":
	// 	fmt.Println("Bang Billy")
	// default : 
	// 	fmt.Println("Kamu Siapaa",name)
	// }	

	// switch dengan short statement
	switch length := len(name); length >= 5 {
	case true :
		fmt.Println("Nama anda kepanjangan")
	case false :
		fmt.Println("Nama diterima")
	}

	// switch tanpa kondisi
	length := len(name)
	switch  {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default :
		fmt.Println("Nama sudah benar")
	}

	
    // var inputString string

    // fmt.Print("Masukkan teks (termasuk spasi): ")
    // fmt.Scanln(&inputString)

    // fmt.Println("Anda memasukkan:", inputString)
}

