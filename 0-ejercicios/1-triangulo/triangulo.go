package main

import (
	"fmt"
	"math"
)

func main() {

	var (
		ladoUno   float64
		ladoDos   float64
		ladoTres  float64
		area      float64
		perimetro float64
	)

	const decimales = 2

	fmt.Println("Ingrese el valor del primer lado del triángulo: ")
	fmt.Scanln(&ladoUno)

	fmt.Println("Ingrese el valor del segundo lado del triángulo: ")
	fmt.Scanln(&ladoDos)

	ladoTres = math.Sqrt(math.Pow(float64(ladoUno), 2) + math.Pow(float64(ladoDos), 2))

	perimetro = ladoUno + ladoDos + ladoTres
	area = (ladoUno * ladoDos) / 2

	fmt.Printf("El valor de los lados es: %f - %f - %f \n", ladoUno, ladoDos, ladoTres)
	fmt.Printf("El valor del perímetro es: %.*f \n", decimales, perimetro)
	fmt.Printf("El valor del área es: %.*f \n", decimales, area)

}
