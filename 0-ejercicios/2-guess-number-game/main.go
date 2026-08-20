package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	play()
}

// funciones
func play() {
	numAleatorio := rand.IntN(100)
	var numIngresado int
	var intentos int
	const maxIntentos = 10

	for intentos < maxIntentos {
		intentos++
		fmt.Printf("Ingresa un numero: (intentos restantes %d): ", maxIntentos-intentos+1)
		fmt.Scanln(&numIngresado)

		if numIngresado == numAleatorio {
			fmt.Println("Felicitaciones adivinaste")
			playAgain()
			return
		} else if numIngresado > numAleatorio {
			fmt.Println("El número aleatorio es menor")
		} else if numIngresado < numAleatorio {
			fmt.Println("El número aleatorio es mayor")
		}

	}

	fmt.Println("Se acabaron los intentos, el número era:", numAleatorio)

	playAgain()
}

func playAgain() {
	var selec string
	fmt.Println("Quieres jugar nuevamente? (s/n)")
	fmt.Scanln(&selec)

	switch selec {
	case "s":
		play()

	case "n":
		fmt.Println("Gracias por jugar!")

	default:
		fmt.Println("Opción invalida")
		playAgain()
	}
}
