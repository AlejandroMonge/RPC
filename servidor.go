package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	alumno "./alumno"
)

type Servidor struct{}
var estudiantes = make(map[string]map[string]float64)

func (this *Servidor) AgregarCalifDeAlumno(alumn alumno.Alumno, respuesta *bool) error {

	asignatura := make(map[string]float64)
	asignatura[alumn.Asignatura] = alumn.Calificacion

	existe := false
	for v := range estudiantes {
		if v == alumn.Nombre {
			existe = true
			break
		}
	}
	//tamanio := len(estudiantes)
	if existe {
		estudiantes[alumn.Nombre][alumn.Asignatura] = alumn.Calificacion
	} else {
		estudiantes[alumn.Nombre] = asignatura
	}
	*respuesta = true
	return nil
}

func (this *Servidor) ObtenerPromedioAlumno(nombreAlumno string, respuesta *float64) error {

	existe := false
	var alumn string
	if len(estudiantes) > 0 {
		for v := range estudiantes {
			if v == nombreAlumno {
				alumn = v
				existe = true
				break
			}
		}
		if existe {
			var sum float64
			for _, calificacion := range estudiantes[alumn] {
				sum = sum + calificacion
			}
			promedio := sum / float64(len(estudiantes[alumn]))
			*respuesta = promedio
		} else {
			return errors.New("No se pudo encontrar al alumno ")
		}
	} else {
		return errors.New("Aun no hay registros")
	}
	return nil
}

func (this *Servidor) ObtenerPromedioAlumnos(inutil string, respuesta *float64) error {

	if len(estudiantes) > 0 {
		var promedio float64
		var promedioGeneral float64
		numAlumnos := len(estudiantes)

		for v := range estudiantes {
			var sum float64
			for _, calificacion := range estudiantes[v] {
				sum = sum + calificacion
			}
			promedio = sum / float64(len(estudiantes[v]))
			promedioGeneral = promedioGeneral + promedio
		}
		*respuesta = promedioGeneral / float64(numAlumnos)
	} else {
		return errors.New("Aun no hay registros")
	}

	return nil 
}

func (this *Servidor) ObtenerPromedioMateria(materia string, respuesta *float64) error {

	if len(estudiantes) > 0 {
		var sum float64
		count := 0
		for v := range estudiantes {
			for mat, calificacion := range estudiantes[v] {
				if mat == materia {
					sum = sum + calificacion
					count++
				}
			}
		}
		promedio := sum / float64(count)
		*respuesta = promedio

	} else {
		return errors.New("No hay elementos registrados")
	}

	return nil
}

func servidor() {
	rpc.Register(new(Servidor))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}

	for {
		c, err := ln.Accept() 
		if err != nil {
			fmt.Println(err)
			continue
		}

		go rpc.ServeConn(c)
	}
}

func main() {
	go servidor()

	var input string
	fmt.Scanln(&input)
}