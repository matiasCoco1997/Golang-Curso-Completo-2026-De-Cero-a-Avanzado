package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hola mundo")
	fmt.Println(quote.Hello())

	//Declaración de variables

	//Como declarar una variable nueva con el simbolo ":="
	// firstName := "Matías"

	var firstName, lastName string
	var age int = 28
	var a byte = 'a'
	var r rune = '❤'

	firstName = "Matías"
	lastName = "Coco"

	fmt.Println(firstName, lastName, age, r, a)

	//Valores default -------------------------------------------------------------------
	var (
		defaultInt    int
		defauiltUint  uint
		defaulFloat   float32
		defaultBool   bool
		defaultString string
	)
	fmt.Println("\nValores default")
	fmt.Println(defaultInt, defauiltUint, defaulFloat, defaultBool, defaultString) //String es un espacio vacio por default

	//Conversiones de tipos (son explicitas) -------------------------------------------------------------------

}
