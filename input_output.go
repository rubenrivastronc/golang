package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*var edad int
	var nombre string
	fmt.Println("Indica tu edad: ")
	fmt.Scanf("%d\n", &edad)
	fmt.Println("Indica tu nombre: ")
	fmt.Scanf("%s\n", &nombre)
	fmt.Printf(nombre+" su edad es %d\n ", edad)*/
	buff()
}

func buff() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingresa tu nombre: ")
	nombre, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hola " + nombre)
	}
}
