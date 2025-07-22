package comandos

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDAAeropuerto "tp2/sistema/aeropuerto"
)

// La función LeerArchivos abre un único archivo, no lo lee

func LeerArchivos(ruta string) (*os.File, error) {
	return os.Open(ruta)
}

func IngresarVuelos(archivo *os.File, aeropuerto TDAAeropuerto.Aeropuerto) {
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		linea := scanner.Text()
		lista := strings.Split(linea, ",")

		if len(lista) != TDAAeropuerto.CANT_INFORMACION {
			continue
		}

		var info [TDAAeropuerto.CANT_INFORMACION]string
		copy(info[:], lista)

		vuelo := TDAAeropuerto.CrearVuelo(info)
		aeropuerto.Agregar_vuelo(vuelo)
	}
}

func ComandoAgregarArchivo(aeropuerto TDAAeropuerto.Aeropuerto, ingresado []string) error {
	if len(ingresado) != 2 {
		return fmt.Errorf("Error en comando agregar_archivo")
	}
	ruta := ingresado[1]
	archivo, err := LeerArchivos(ruta)
	if err != nil {
		return fmt.Errorf("Error en comando agregar_archivo")
	}
	defer archivo.Close()
	IngresarVuelos(archivo, aeropuerto)
	return nil
}

func ComandoInfoVuelo(aeropuerto TDAAeropuerto.Aeropuerto, ingresado []string) (string, error) {
	if len(ingresado) != 2 {
		return "", fmt.Errorf("Error en comando info_vuelo")
	}
	codigo := ingresado[1]
	info, err := aeropuerto.Info_vuelo(codigo)
	if err != nil {
		return "", fmt.Errorf("Error en comando info_vuelo")
	}
	return info, nil
}

func ComandoPrioridadVuelos(aeropuerto TDAAeropuerto.Aeropuerto, ingresado []string) error {
	if len(ingresado) != 2 {
		return fmt.Errorf("Error en comando prioridad_vuelos")
	}
	k, err := strconv.Atoi(ingresado[1])
	if err != nil || k < 0 {
		return fmt.Errorf("Error en comando prioridad_vuelos")
	}
	aeropuerto.Mostrar_prioridad(k)
	return nil
}

func ComandoSiguienteVuelo(aeropuerto TDAAeropuerto.Aeropuerto, ingresado []string) (string, error) {
	if len(ingresado) != 4 {
		return "", fmt.Errorf("Error en comando siguiente_vuelo")
	}
	origen, destino, fecha := ingresado[1], ingresado[2], ingresado[3]
	vuelo := aeropuerto.Siguiente_vuelo(origen, destino, fecha)
	if vuelo == nil {
		return "", nil
	}
	return vuelo.ObtenerInfo(), nil
}

func ComandoVerTablero(aeropuerto TDAAeropuerto.Aeropuerto, ingresado []string) error {
	if len(ingresado) != 5 {
		return fmt.Errorf("Error en comando ver_tablero")
	}
	k, err := strconv.Atoi(ingresado[1])
	modo := ingresado[2]
	desde, hasta := ingresado[3], ingresado[4]

	if err != nil || k <= 0 || (modo != "asc" && modo != "desc") {
		return fmt.Errorf("Error en comando ver_tablero")
	}
	if strings.Compare(desde, hasta) > 0 {
		return fmt.Errorf("Error en comando ver_tablero")
	}

	if modo == "asc" {
		aeropuerto.Ver_tablero_ASC(k, desde, hasta)
	} else {
		aeropuerto.Ver_tablero_DESC(k, desde, hasta)
	}
	return nil
}

func ComandoBorrar(aeropuerto TDAAeropuerto.Aeropuerto, ingresado []string) error {
	if len(ingresado) != 3 {
		return fmt.Errorf("Error en comando borrar")
	}
	desde, hasta := ingresado[1], ingresado[2]
	if strings.Compare(desde, hasta) > 0 {
		return fmt.Errorf("Error en comando borrar")
	}
	aeropuerto.Borrar(desde, hasta)
	return nil
}
