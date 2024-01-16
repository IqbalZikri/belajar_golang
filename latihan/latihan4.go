// Latihan 4
// Buatlah dua for loops bersarang untuk membuat pola segitiga siku-siku (tinggi = 5)

package main

import "fmt"

func cetakSegitiga(tinggi int){
	for i := 1; i <= tinggi; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println("")
	}
}

func main(){
	cetakSegitiga(5)
}