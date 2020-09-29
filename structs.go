package main

import "fmt"

type User struct {
	edad             int
	nombre, apellido string
}

func main() {

	uriel := new(User)
	//User{nombre: "Uriel", apellido: "Hernandez"}
	uriel.nombre = "Uriel"

	fmt.Println(uriel.nombre)

}
