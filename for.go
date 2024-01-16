// For loops
package main

import "fmt"

func main(){
	// cara 1
	count := 100
	for i := 0; i <= count; i++ {
	// init statement, yaity statement sebelum for di eksekusi ()
	// post statement, yaitu statement yang akan selalu dieksekusi di akhir perulangan 
		fmt.Println("angka ke ", i)
	}

	// cara 2
	counter := 1
	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}
	fmt.Println("Selesai")

	// for bisa digunakan untuk melalukan iterasi terhadap semua data collection
	// data collection seperti Array, Slice dan Map

	// Slice
	nama := []string{"iqbal", "Azki", "Joko", "Anies" , "Bambang"}
	for i := 0; i < len(nama); i++ {
		fmt.Println(nama[i])
	}

	// for range
	name := []string{"Iqbal", "Ball", "TB"}
	for index, name := range name {
		fmt.Println("index", index, "=", name)
	}

	// Array
	kampus := [...]string{"UI", "ITB", "UNNES"}
	for index, kampus := range kampus{
		fmt.Println("index", index, "=", kampus)
	}

	for _, kampus := range kampus{
		fmt.Println(kampus)
	}
	
}