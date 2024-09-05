package main

func main() {
	// Usando var
	// var age int = 15
	// var name = "Alice"
	// var a, b, c int = 1, 2, 3
	// var x, y = 1, "Hello"

	// var isReady bool // false
	// var count int // 0

	// var a, b, c int // 0

	// Usando :=
	// age := 25
	// name := "Alice"

	// a, b, c := 1, 2, 3
	// x, y := 1, "hello"

}

// Ámbito de Uso:

// var: Puede usarse tanto dentro como fuera de funciones.
// :=: Solo puede usarse dentro de funciones.
// Flexibilidad:

// var: Permite declarar variables sin inicialización (asignando el valor cero por defecto del tipo).
// :=: Siempre requiere una inicialización explícita.
// Redundancia:

// var: Puede ser redundante en casos donde el tipo es evidente.
// :=: Más conciso y elimina la redundancia, ya que infiere el tipo automáticamente.