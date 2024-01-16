// Laithan
/* Buatlah sebuah fungsi yang menerima dua parameter, x dan y. Fungsi ini harus mengembalikan nilai
// berupa string "x lebih  besar dari y" jika x lebih besar dari y, atau "y lebih besar dari x" jika
// y lebih besar dari x. */
package main

import "fmt"

func XY(X, Y int)(hasil1, hasil2 string){
	hasil1 = "Nilai X lebih besar dari nilai Y"
	hasil2 = "Nilai Y lebih besar dari nilai X"
	return hasil1, hasil2
}

func main(){
	angka1 := 1
	angka2 := 2
	if angka1 > angka2 {
		hasil1, _ := XY(angka1,angka2)
		fmt.Println(hasil1)
	}else if angka2 > angka1 {
		_, hasil2 := XY(angka1,angka2)
		fmt.Println(hasil2)
	}

}

