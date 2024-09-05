package main

import (
	"fmt"
	"math/cmplx"
	"unsafe"
)

func main() {
	// Enteros con signo
	var i int = 42
	var i8 int8 = -128
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = -9223372036854775808

	// Enteros sin signo
	var u uint = 42
	var u8 uint8 = 255
	var u16 uint16 = 65535
	var u32 uint32 = 4294967295
	var u64 uint64 = 18446744073709551615
	var uptr uintptr = uintptr(unsafe.Pointer(&i))

	// Números de punto flotante
	var f32 float32 = 3.14
	var f64 float64 = 2.718

	// Números complejos
	var c64 complex64 = 1 + 2i
	var c128 complex128 = cmplx.Sqrt(-5 + 12i)

	// Alias de tipos
	var b byte = 255
	var r rune = '♥' // Unicode code point for the heart symbol

	// Booleanos
	var flag bool = true

	// Cadenas de caracteres
	var s string = "Hello, Go!"

	// Imprimir los valores
	fmt.Println("int:", i)
	fmt.Println("int8:", i8)
	fmt.Println("int16:", i16)
	fmt.Println("int32:", i32)
	fmt.Println("int64:", i64)
	fmt.Println("uint:", u)
	fmt.Println("uint8:", u8)
	fmt.Println("uint16:", u16)
	fmt.Println("uint32:", u32)
	fmt.Println("uint64:", u64)
	fmt.Println("uintptr:", uptr)
	fmt.Println("float32:", f32)
	fmt.Println("float64:", f64)
	fmt.Println("complex64:", c64)
	fmt.Println("complex128:", c128)
	fmt.Println("byte:", b)
	fmt.Println("rune:", r)
	fmt.Println("bool:", flag)
	fmt.Println("string:", s)
}