// Variadic Function
// Parameter yang berada di posisi terakhir, memiliki kemampuan dijadikan sebuah varargs (variable argumens)
// Varargs artinya datanya bisa menerima lebih dari satu input, atau anggap saja sama seperti array
// Perbedaan parameter biasa dengan array :
//  -> Jika parameter bertipe array, kita wajib membuat array terlbih dahulu sebelum mengirimkan ke function
//  -> Jika parameter menggunakan varargs, kita bisa langsung mengirim datanya, jika lebih dari satu, cukup gunakan tanda koma

package main

import "fmt"


func sumAll(numbers ...int) int{
	total := 0 
	for _, number := range numbers{
		total += number
	}
	return total
}

func sumAllSlice(numbers[]int) int{
	total := 0 
	for _, number := range numbers{
		total += number
	}
	return total
}

// Slice Parameter
// Kadang ada kasus dimana kita menggunakan Variadic Function, namun berupa variable berupa slice
// kita bisa menjadikan slice menjadi Varargs (Variable Argumen)


func main(){
	total := sumAll(10,10,10,10,10,10)
	fmt.Println(total)

	// bisa juga seperti ini
	fmt.Println(sumAll(10,10,10,10))
	fmt.Println(sumAll(10,10,10,10,10,10,10))
	fmt.Println(sumAll(10,10,10,10,10,10,10,10,10,10))

	// Slice
	fmt.Println(sumAllSlice([]int{10,10,10,10,10}))

	// jika kita ingin membuatnya slice maka kita harus membuatnya secara manual
	// Maksudnya manual adalah buat slicenya terlebih dahulu

	// Slice Parameter
	numbers := []int{10,10,10,10,10,10,10,10}
	total = sumAll(numbers...)
	fmt.Println(total)
	// Atau bisa langsung seperti ini
	fmt.Println(sumAll(numbers...))

	// Keuntungan menggunakan Variable Argumens adalah saat kita memanggil functionnya kita bisa kirimkan langsung lebih dari satu data
	// contoh diatas kita memanggil sumAll() disana kita bisa mengirim datanya lebih dari satu, sedangkan saat kita membuat functionnya kira hanya membuat 1 parameter
}
