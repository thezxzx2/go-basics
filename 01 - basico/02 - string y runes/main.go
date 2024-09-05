package main

import "fmt"

func main() {
	// // Declaracion de una cadena literal
	// var str1 string = "Hello, Go!"

	// // Declaracion con inferencia de tipo
	// str2 := "Go es genial"

	// // Declaracion de una cadena multilinea usando backticks
	// str3 := `esto es
	// una cadena
	// multilinea`

	// fmt.Println(str1)
	// fmt.Println(str2)
	// fmt.Println(str3)

	// ___________________________

	// Operaciones con Strings
	// Concatenacion
	// str1 := "Hello, "
	// str2 := "world!"
	// str3 := str1 + str2
	// fmt.Println(str3)

	// // Longitud
	// length := len(str3)
	// fmt.Println(length)

	// // Acceso a bytes individuales
	// char := str3[0]
	// fmt.Printf("El primer byte de la cadena es: %c\n", char)

	// for i := 0; i < len(str3); i++ {
	// 	fmt.Printf("%c ", str3[i])
	// }

	// fmt.Println()

	// ____________________________

	// Runes
	// Declaracion de una runa literal
	// var r1 rune = 'a'

	// // Declaracion con inferencia de tipo
	// r2 := 'b'

	// // Declaracion de una runa Unicode
	// r3 := '♥'

	// fmt.Printf("r1: %c, r2: %c, r3: %c\n", r1, r2, r3)

	// Operaciones con runas

	// Conversion de string a runas
	str := "Hola, mundo"
	runes := []rune(str)
	fmt.Printf("Runas: %v\n", runes)

	// Iteracion sobre rnas
	for _, r := range runes {
		fmt.Printf("%c ", r)
	}
	
	fmt.Println()
}