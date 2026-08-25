package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tarea struct {
	nombre      string
	descripcion string
	completado  bool
}

type ListaTareas struct {
	tareas []Tarea
}

func (l *ListaTareas) agregarTarea(t Tarea) {
	l.tareas = append(l.tareas, t)
}

func (l *ListaTareas) marcarCompletado(index int) {
	l.tareas[index].completado = true
}

func (l *ListaTareas) editarTarea(index int, t Tarea) {
	l.tareas[index] = t
}

func (l *ListaTareas) eliminarTarea(index int) {
	l.tareas = append(l.tareas[0:index], l.tareas[index+1:]...) // Elimina la tarea ubicada en el índice indicado, uniendo las tareas anteriores y posteriores a ese índice.
}

func main() {

	lista := ListaTareas{}
	leer := bufio.NewReader(os.Stdin) // Este paquete se usa para recibir muchos caracteres (fmt no se banca tantos caracteres)

	for {
		var opcion int
		fmt.Println(
			"Seleccione una opción: \n",
			"1. Agregar tarea\n",
			"2. Marcar tarea como completada\n",
			"3. Editar tarea\n",
			"4. Eliminar tarea\n",
			"5. Salir",
		)

		fmt.Print("Ingrese una opcion: ")
		fmt.Scanln(&opcion)

		switch opcion {

		case 1: // Agregar tarea
			var t Tarea

			fmt.Println("Ingrese el nombre de la tarea")
			t.nombre, _ = leer.ReadString('\n')

			fmt.Println("Ingrese la descripción de la tarea")
			t.descripcion, _ = leer.ReadString('\n')

			lista.agregarTarea(t)
			fmt.Println("Tarea agregada correctamente")

		case 2: //Marcar tarea como completada
			var index int

			fmt.Print("Ingrese el índice de la tarea que desea marcar como completada: ")
			fmt.Scanln(&index)

			lista.marcarCompletado(index)
			fmt.Println("Tarea N° %d completada correctamente", index)

		case 3: //Editar tarea

			var t Tarea
			var index int

			fmt.Print("Ingrese el índice de la tarea que desea actualizar/modificar: ")
			fmt.Scanln(&index)

			fmt.Println("Ingrese el nombre de la tarea")
			t.nombre, _ = leer.ReadString('\n')

			fmt.Println("Ingrese la descripción de la tarea")
			t.descripcion, _ = leer.ReadString('\n')

			lista.editarTarea(index, t)

			fmt.Println("Tarea N° %d se ha modificado correctamente", index)

		case 4: //Eliminar tarea
			var index int

			fmt.Print("Ingrese el índice de la tarea que desea eliminar: ")
			fmt.Scanln(&index)

			lista.eliminarTarea(index)
			fmt.Println("Tarea N° %d se ha eliminado correctamente", index)

		case 5: //Salir
			fmt.Print("Saliendo del programa")
			return

		default:
			fmt.Print("Opción invalida")
		}

		fmt.Println("Lista de tareas")
		fmt.Println("===========================================")
		for indice, tarea := range lista.tareas {
			fmt.Printf("%d. %s - %s - Completado: %t\n", indice, tarea.nombre, tarea.descripcion, tarea.completado)
		}
	}

}
