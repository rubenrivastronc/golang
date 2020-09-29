package main

import "fmt"

func main() {

	/* 1. Puntero = dirección de memoria
	   2. En lugar del valor, tenemos la direccion en la que está el valor
	   3. Modificar X,Y => apuntan al mismo dato (as123d) => 5
	   4. X => 6as123d -> 6
		El asterisco de acceso al valor en ese espacio de memoria.
	*/

	var x, y *int
	entero := 5

	x = &entero
	y = &entero

	*x = 6

	fmt.Println(*x)
	fmt.Println(y)
}
