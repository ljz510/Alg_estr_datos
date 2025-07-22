package aeropuerto

import (
	"strconv"
	"strings"
)

type vueloImplementacion struct {
	informacion [CANT_INFORMACION]string
}

func CrearVuelo(informacion [CANT_INFORMACION]string) Vuelo {
	vuelo := new(vueloImplementacion)
	vuelo.informacion = informacion
	return vuelo
}

func (vuelo *vueloImplementacion) ObtenerPrioridad() string {
	return vuelo.informacion[PRIORITY]
}

func (vuelo *vueloImplementacion) ObtenerFecha() string {
	return vuelo.informacion[DATE]
}

func (vuelo *vueloImplementacion) ObtenerCodigoVuelo() string {
	return vuelo.informacion[FLIGHT_NUMBER]
}

func (vuelo *vueloImplementacion) ObtenerOrigenYdestino() (string, string) {
	return vuelo.informacion[ORIGIN_AIRPORT], vuelo.informacion[DESTINATION_AIRPORT]
}

func (vuelo *vueloImplementacion) ObtenerInfo() string {
	campos := vuelo.informacion

	campos[5] = strconv.Itoa(funcionaParse(campos[5])) // PRIORITY
	campos[7] = strconv.Itoa(funcionaParse(campos[7])) // DEPARTURE_DELAY
	campos[8] = strconv.Itoa(funcionaParse(campos[8])) // AIR_TIME
	campos[9] = strconv.Itoa(funcionaParse(campos[9])) // CANCELLED

	return strings.Join(campos[:], " ")
}

func funcionaParse(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
