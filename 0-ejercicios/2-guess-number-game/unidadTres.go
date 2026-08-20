package main

import (
	"fmt"
	"runtime"
	"time"
)

func mainTwo() {

	t := time.Now().Hour()

	if t < 12 {
		fmt.Println("Mañana")
	} else if t < 17 {
		fmt.Println("tarde")
	} else {
		fmt.Println("noche")
	}

	//Se puede declarar una variable adentro del if
	if t := time.Now().Hour(); t < 12 {
		fmt.Println("Mañana")
	} else if t < 17 {
		fmt.Println("tarde")
	} else {
		fmt.Println("noche")
	}

	os := runtime.GOOS

	switch os {
	case "windows":
		fmt.Println("Windows")

	case "linux":
		fmt.Println("Linux")

	case "darwin":
		fmt.Println("MAC")

	default:
		fmt.Println("Otro OS")
	}

	//inicializando una variable dentro del switch
	switch os2 := runtime.GOOS; os2 {
	case "windows":
		fmt.Println("Windows")

	case "linux":
		fmt.Println("Linux")

	case "darwin":
		fmt.Println("MAC")

	default:
		fmt.Println("Otro OS")
	}

	// Bucle for ----------------------------------

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		if i == 5 {
			break
		}
	}

	fmt.Println(hello("matias"))

	fmt.Println(suma(1, 2))

	//asigno los valores a variables
	sum, prod := calc(4, 5)
	fmt.Println("La suma es:", sum)
	fmt.Println("El producto es:", prod)

}

//funciones

func hello(name string) string {
	return "Hola, " + name
}

func suma(a, b int) int {
	return a + b
}

// se pueden retornar mas de un valor
func calc(a, b int) (sum, prod int) {
	sum = a + b
	prod = a * b
	return
}
