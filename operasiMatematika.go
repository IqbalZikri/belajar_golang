// Operasi Matematika
// + , - , *, / , % (% = modulo atau sisa hasil bagi)
package main

import "fmt"

func main(){
	var a = 10
	var b = 10
	var c = a + b
	fmt.Println(c)

	// Augmented Assignments
	// a = a + 10 sama saja  dengan a += 10
	// a = a - 10 sama saja  dengan a -= 10
	// a = a * 10 sama saja  dengan a *= 10
	// a = a / 10 sama saja  dengan a /= 10
	// a = a % 10 sama saja  dengan a %= 10

	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)

	// Unary Operator
	// ++ = a + 1 , -- = a - 1 , - = minus , + = plus , ! = Boolean kebalikan

	var d = 1
	d++
	fmt.Println(d)
	d++
	fmt.Println(d)

	d--
	fmt.Println(d)
	d--
	fmt.Println(d)
}