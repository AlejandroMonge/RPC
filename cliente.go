package main

import (
	"fmt"
	"net/rpc"
	alumno "./alumno"
)

func main() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var opc int64
	for {
		fmt.Println("\n\nAgregar la calificación de un alumno por materia . [1]")
		fmt.Println("Obtener el promedio del alumno ....................[2]")
		fmt.Println("Obtener el promedio de todos los alumnos ..........[3]")
		fmt.Println("Obtener el promedio por materia ...................[4]")
		fmt.Println("Salir .............................................[0]")
		fmt.Print("Elige una opcion: ")
		fmt.Scanln(&opc)

		switch opc {
		case 1:
			var nombre string
			var materia string
			var calificacion float64
			fmt.Println("\n::::: Agregar la calificación de un alumno por materia :::::")
			fmt.Print("Nombre alumno: ")
			fmt.Scanln(&nombre)
			fmt.Print("Materia: ")
			fmt.Scanln(&materia)
			fmt.Print("Calificación: ")
			fmt.Scanln(&calificacion)
			alumn := alumno.Alumno{nombre, materia, calificacion}
			resultado := false
			err = c.Call("Servidor.AgregarCalifDeAlumno", alumn, &resultado)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Alumno y calificacion agregados")
			}
		case 2:
			var nombre string
			var resultado float64
			fmt.Println("\n::::: Obtener el promedio del alumno :::::")
			fmt.Print("Nombre alumno: ")
			fmt.Scanln(&nombre)
			err = c.Call("Servidor.ObtenerPromedioAlumno", nombre, &resultado)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Promedio de", nombre, ":", resultado)
			}
		case 3:
			var resultado float64
			fmt.Println("\n::::: Obtener el promedio de todos los alumnos :::::")
			err = c.Call("Servidor.ObtenerPromedioAlumnos", "Chale", &resultado)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("El promedio general es:", resultado)
			}
		case 4:
			var materia string
			var resultado float64
			fmt.Println("\n::::: Obtener el promedio por materia :::::")
			fmt.Print("Materia: ")
			fmt.Scanln(&materia)
			err = c.Call("Servidor.ObtenerPromedioMateria", materia, &resultado)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("El promedio de la materia:", resultado)
			}
		case 0:
			fmt.Println("\nBYE\n")
			return
		default:
			fmt.Println("Error\n")
		}
	}
}