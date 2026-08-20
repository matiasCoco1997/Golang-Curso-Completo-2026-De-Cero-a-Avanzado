package main

import (
	"fmt"
	"time"
)

func main() {

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

}
