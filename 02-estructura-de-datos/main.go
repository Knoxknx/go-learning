package main

import (
	"bufio"
	"fmt"
	"os"
)

type Tarea struct {
	nombre     string
	desc       string
	completado bool
}

type ListaTareas struct {
	tareas []Tarea
}

// Metodo para agregar tareas
func (l *ListaTareas) agregarTarea(t Tarea) {
	l.tareas = append(l.tareas, t)
}

// Metodo para marcar completado la tarea
func (l *ListaTareas) marcarCompletado(index int) {
	l.tareas[index].completado = true
}

// Metodo para editar tarea
func (l *ListaTareas) editarTarea(index int, t Tarea) {
	l.tareas[index] = t
}

// Metodo para eliminar tarea
func (l *ListaTareas) eliminarTarea(index int) {
	l.tareas = append(l.tareas[:index], l.tareas[:index+1]...)
}

func main() {
	// Crear instancia de lista de tareas
	lista := ListaTareas{}

	leer := bufio.NewReader(os.Stdin)

	for {
		var opcion int
		fmt.Println("Seleccione una opción: \n",
			"1. Agregar tarea \n",
			"2. Marcar tarea completada \n",
			"3. Editar tarea \n",
			"4. Eliminar tarea \n",
			"5. Salir")

		fmt.Print("Ingrese la opción: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			var t Tarea
			fmt.Print("Ingrese nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Print("Ingrese una descripción de la tarea: ")
			t.desc, _ = leer.ReadString('\n')
			lista.agregarTarea(t)
			fmt.Println("Tarea agregada correctamente")
		case 2:
			var index int
			fmt.Println("Ingrese el indice de la tarea que desea marcar como completado: ")
			fmt.Scanln(&index)
			lista.marcarCompletado((index))
			fmt.Println("Tarea completada: ")
		case 3:
			var index int
			var t Tarea
			fmt.Print("Ingrese el indice de la tarea que desea actualizar:")
			fmt.Scanln(&index)
			fmt.Print("Ingrese nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Print("Ingrese una descripción de la tarea: ")
			t.desc, _ = leer.ReadString('\n')
			lista.editarTarea(index, t)
			fmt.Print("Tarea actualizada: ")
		case 4:
			var index int
			fmt.Print("Ingrese el indice de la tarea que desea eliminar:")
			fmt.Scanln(&index)
			lista.eliminarTarea(index)
			fmt.Print("Tarea eliminada:")
		case 5:
			fmt.Println("Saliendo del programa")
			return
		default:
			fmt.Println("Opción invalida")
		}
		// Listar todas las tareas
		fmt.Println("Lista de tareas")
		fmt.Println("************************************************************************")
		for i, t := range lista.tareas {
			fmt.Printf("%d. %s - %s - Completado: %t\n", i, t.nombre, t.desc, t.completado)
		}
		fmt.Println("************************************************************************")
	}

}
