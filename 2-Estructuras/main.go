package main

import "fmt"

// Definición de una estructura (struct).
type Persona struct {
	nombre string
	edad   int
	correo string
}

func (p *Persona) diHola() { // El método pertenece al tipo Persona y recibe un puntero a una Persona.
	fmt.Println("Hola mi nobmre es ", p.nombre)
}

func main() {
	//Vectores***************************************************************************************************************************************
	var vector [5]int // sin inicializar

	var vectorDos = [5]int{10, 20, 30, 40, 50} //vector completamente incicializado

	var vectorTres = [...]int{10, 20, 30, 40, 50} // [...] permite que Go determine automáticamente la cantidad de elementos.

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

	//Matriz bidimencional***************************************************************************************************************************************

	var matriz = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	fmt.Println(matriz)

	//Slices (Rebanadas)***************************************************************************************************************************************
	var a []int
	diasSemana := []string{"Domingo", "Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado"}

	diasRebanada := diasSemana[0:5] //desde el indice 0 al 5

	diasRebanada = append(diasRebanada, "Viernes", "Sabado", "Otro dia")

	diasRebanada = append(diasRebanada[0:2], diasRebanada[3:]...) // Eliminamos al indice 2 Martes (reconstruimos el slice sin "martes")

	fmt.Println(a)
	fmt.Println(diasSemana)
	fmt.Println(diasRebanada)
	fmt.Println(len(diasRebanada)) //Longitud del vector
	fmt.Println(cap(diasRebanada)) //Cuánto puede crecer el slice desde su posición inicial hasta el final del array subyacente

	//Funcion make***************************************************************************************************************************************
	nombres := make([]string, 5, 10) // Tipo de dato - longitud - Capacidad
	nombres[3] = "Matias"            //Agrego un elemento
	fmt.Println(nombres)

	//Funcion Copy***************************************************************************************************************************************
	rebanadaUno := []int{1, 2, 3, 4, 5}
	rebanadaDos := make([]int, 5)
	copy(rebanadaDos, rebanadaUno)
	fmt.Println(rebanadaDos, rebanadaUno) //copy(destino, origen) - Copia los elementos de rebanadaUno (origen) en rebanadaDos (destino).

	//Mapas***************************************************************************************************************************************
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

	//Estructuras***************************************************************************************************************************************
	var p Persona
	p.nombre = "Matias"
	p.edad = 28
	p.correo = "matias@gmail.com"

	personaDos := Persona{"Matias", 28, "matias@gmail.com"}

	fmt.Println(p)
	fmt.Println(personaDos)

	//Punteros y metodos***************************************************************************************************************************************
	var x int = 10
	//var puntero *int = &x // &x obtiene la dirección de memoria de x. Un puntero guarda esa dirección y permite acceder/modificar el valor original.

	fmt.Println(x)
	editar(&x) // &x obtiene la dirección de memoria de x y se la pasa a editar.
	fmt.Println(x)

	//METODO CON PUNTEROS
	personaTres := Persona{"Matias", 28, "matias@gmail.com"}
	personaTres.diHola()
}

// Recibe un puntero a un entero y modifica directamente el valor original.
func editar(x *int) {
	*x = 20
}
