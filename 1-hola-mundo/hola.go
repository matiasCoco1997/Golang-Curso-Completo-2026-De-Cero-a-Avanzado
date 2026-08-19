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

	firstName = "Matías"
	lastName = "Coco"

	fmt.Println(firstName, lastName, age)
}
