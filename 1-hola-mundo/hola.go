package main

import (
	"fmt"
	"strconv"

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
	var integer16 int16 = 50
	var integer32 int32 = 100
	s := "100"

	i, _ := strconv.Atoi(s) //Convierte un string a entero y devuelve 2 valores un int y un posible error (La "_" es para no almacenar el error)

	n := 42

	s = strconv.Itoa(n)

	var valorFinal int32 = int32(integer16) + integer32
	fmt.Println(valorFinal, i, s)

	//Paquete fmt -------------------------------------------------------------------

	var nameTwo string
	var ageTwo int

	fmt.Print("Ingrese su nombre: ")
	fmt.Scanln(&nameTwo)

	fmt.Print("Ingrese su edad: ")
	fmt.Scanln(&ageTwo)

	fmt.Printf("Hola, me llamo %s y tengo %d años. \n", nameTwo, ageTwo) //"%s" espera un string, %d espera un número entero

	greeting := fmt.Sprintf("Hola, me llamo %s y tengo %d años. \n", nameTwo, ageTwo)

	fmt.Printf(greeting)

	fmt.Printf("El tipo de nameTwo es: %T \n", nameTwo)
	fmt.Printf("El tipo de ageTwo es: %T\n", ageTwo)

}
