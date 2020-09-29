package main

import "fmt"

type User struct {
	edad             int
	nombre, apellido string
}

func (usuario User) devuelveUsuario() string {
	return usuario.nombre + " " + usuario.apellido
}

func (this *User) set_name(n string) {
	this.nombre = n
}

func main() {

	uriel := new(User)
	//User{nombre: "Uriel", apellido: "Hernandez"}
	uriel.nombre = "Uriel"
	uriel.set_name("Marcos")

	fmt.Println(uriel.nombre)

}
