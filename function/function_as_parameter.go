// Function as parameter
// function tidak hanya bisa kita simpan di dalam variable sebagai value
// namun kita juga bisa gunakan sebagai parameter untuk function lain

package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string){

	// Bisa juga seperti ini
	filteredName := filter(name)
	fmt.Println("Hello",filteredName)
}

func spamFilter(name string)string{
	if name == "Anjing"{
		return "..."
	}else{
		return name
	}
}

// Function Type Declarations
// Kadang jika function terlalu panjang, agak sulit untuk menuliskannya dalam parameter
// Type Declarations juga bisa digunakan untuk membuat alias function, sehingga akan mempermudah kita menggunakan function sebagai parameter
// Contoh Function Type Declarations

type Filter func(string)string
// bisa juga ditambahkan seperti ini 
// type Filter func(string, int, bool)string

// func sayHelloWithFilter(name string, filter Filter){
// 	fmt.Println("Hello", filter(name))
// }

func main(){
	sayHelloWithFilter("Iqbal", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}