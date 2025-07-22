package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	TDAAeropuerto "tp2/sistema/aeropuerto"
	"tp2/sistema/comandos"
)

func main() {
	algueiza := TDAAeropuerto.CrearAeropuerto()
	entrada := bufio.NewScanner(os.Stdin)

	for entrada.Scan() {
		ingresado := strings.Split(entrada.Text(), " ")
		comando := ingresado[0]

		var err error
		var res string

		switch comando {
		case "agregar_archivo":
			err = comandos.ComandoAgregarArchivo(algueiza, ingresado)
		case "info_vuelo":
			res, err = comandos.ComandoInfoVuelo(algueiza, ingresado)
		case "prioridad_vuelos":
			err = comandos.ComandoPrioridadVuelos(algueiza, ingresado)
		case "siguiente_vuelo":
			res, err = comandos.ComandoSiguienteVuelo(algueiza, ingresado)
		case "ver_tablero":
			err = comandos.ComandoVerTablero(algueiza, ingresado)
		case "borrar":
			err = comandos.ComandoBorrar(algueiza, ingresado)
		default:
			continue
		}

		// eliminacion de lógica repetida...
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		if res != "" {
			fmt.Println(res)
		}
		fmt.Println("OK")
	}
}
