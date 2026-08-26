package main

import (
	"errors"
	"fmt"
	"strconv"
)

func divide(dividendo, divisor int) (int, error) {

	if divisor == 0 {
		return 0, errors.New("No es posible dividir por 0")
	}
	return dividendo / divisor, nil
}

func main() {

	str := "123"
	numero, error := strconv.Atoi(str)

	if error != nil {
		fmt.Println("Error:", error)
		return
	} else {
		fmt.Println("El número es:", numero)
	}

	resultado, err := divide(1, 0)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("El resultado es:", resultado)
	}
}
