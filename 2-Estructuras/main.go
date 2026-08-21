package main

import "fmt"

func main() {
	//Vectores
	var vector [5]int // sin inicializar

	var vectorDos = [5]int{10, 20, 30, 40, 50} //vector completamente incicializado

	var vectorTres = [...]int{10, 20, 30, 40, 50} //los 3 puntos es para incializarlo pero no sabemos cuantos elementos posee

	vector[0] = 1 //modificar el vector en una posicion especifica

	//Recorrer un vector y modificarlo con un bucle
	for i := 0; i < len(vector); i++ {
		if i == 0 {
			vector[i] = 1
		} else {
			vector[i] = vector[i-1] * 2
		}
	}

	//For con Range
	for index, value := range vectorDos { //si no quiero utilizar el index y solo agarrar el valro en la posicion de index se colocaria "_"
		fmt.Println("indice= %d, Valor= %d\n", index, value)
	}

	fmt.Println(vector, vectorDos, vectorTres)

	//Matriz bidimencional

	var matriz = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	fmt.Println(matriz)

}
