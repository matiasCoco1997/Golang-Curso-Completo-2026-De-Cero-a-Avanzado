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

	//Slices (Rebanadas)
	var a []int
	diasSemana := []string{"Domingo", "Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado"}

	diasRebanada := diasSemana[0:5] //desde el indice 0 al 5

	diasRebanada = append(diasRebanada, "Viernes", "Sabado", "Otro dia")

	diasRebanada = append(diasRebanada[0:2], diasRebanada[3:]...) // Eliminamos al indice 2 Martes

	fmt.Println(a)
	fmt.Println(diasSemana)
	fmt.Println(diasRebanada)
	fmt.Println(len(diasRebanada)) //Longitud del vector
	fmt.Println(cap(diasRebanada)) //Cuánto puede crecer el slice desde su posición inicial hasta el final del array subyacente

	//Funcion make
	nombres := make([]string, 5, 10) // Tipo de dato - longitud - Capacidad
	nombres[3] = "Matias"            //Agrego un elemento
	fmt.Println(nombres)

	//Funcion Copy
	rebanadaUno := []int{1, 2, 3, 4, 5}
	rebanadaDos := make([]int, 5)
	copy(rebanadaDos, rebanadaUno)
	fmt.Println(rebanadaDos, rebanadaUno) //El de la derecha copia los elementos del que esta a la izquierda

	//Mapas
	colors := map[string]string{
		"rojo":  "#FF0000",
		"verde": "#00FF00",
		"azul":  "#0000FF",
	}

	//Agregar un elemento al mapa
	colors["negro"] = "#000000"

	valor, verificacionExistencia := colors["blanco"] // Si no existe el valor te devuelve un string vacio

	fmt.Println(colors)
	fmt.Println(colors["rojo"])

	if verificacionExistencia {
		fmt.Println(valor)
	} else {
		fmt.Println("No existe la clave ingresada")
	}

	//eliminar un elemento del mapa
	delete(colors, "rojo")
	fmt.Println(colors)

}
