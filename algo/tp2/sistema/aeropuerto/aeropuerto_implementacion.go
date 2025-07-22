package aeropuerto

import (
	"fmt"
	"strconv"
	"strings"
	TDAHeap "tdas/cola_prioridad"
	TDADiccionario "tdas/diccionario"
	TDALista "tdas/lista"
	TDAPila "tdas/pila"
)

type aeropuertoImplementacion struct {
	vuelos     TDADiccionario.Diccionario[string, Vuelo]
	conexiones TDADiccionario.Diccionario[Conexion, TDADiccionario.DiccionarioOrdenado[FechaYCodigo, Vuelo]]
	//modificacion sugerida por el profesor, la utilización de punteros reduce la cantidad de informacion duplicada en memoria
	tableroAsc  TDADiccionario.DiccionarioOrdenado[string, *TDALista.Lista[string]]
	tableroDesc TDADiccionario.DiccionarioOrdenado[string, *TDALista.Lista[string]]
}

func CrearAeropuerto() Aeropuerto {
	return &aeropuertoImplementacion{
		vuelos:     TDADiccionario.CrearHash[string, Vuelo](),
		conexiones: TDADiccionario.CrearHash[Conexion, TDADiccionario.DiccionarioOrdenado[FechaYCodigo, Vuelo]](),
		tableroAsc: TDADiccionario.CrearABB[string, *TDALista.Lista[string]](strings.Compare),
		tableroDesc: TDADiccionario.CrearABB[string, *TDALista.Lista[string]](
			func(a, b string) int { return strings.Compare(b, a) }),
	}
}

func compararVuelos(a, b Vuelo) int {
	prior1, _ := strconv.Atoi(a.ObtenerPrioridad())
	prior2, _ := strconv.Atoi(b.ObtenerPrioridad())
	if res := prior1 - prior2; res != 0 {
		return res
	}
	return strings.Compare(a.ObtenerCodigoVuelo(), b.ObtenerCodigoVuelo()) * -1
}

func CompararFechaYCodigo(a, b FechaYCodigo) int {
	if res := strings.Compare(a.fecha, b.fecha); res != 0 {
		return res
	}
	return strings.Compare(a.codigo, b.codigo)
}

//Construir una nueva lista y copiar todos los elementos para insertar un elemento estropea la complejidad. Simplemente se debería insertar en la lista existente.

func insertarOrdenado(lista TDALista.Lista[string], codigo string) {
	it := lista.Iterador()
	for it.HaySiguiente() {
		if strings.Compare(codigo, it.VerActual()) < 0 {
			it.Insertar(codigo)
			return
		}
		it.Siguiente()
	}
	lista.InsertarUltimo(codigo)
}

func (a *aeropuertoImplementacion) Agregar_vuelo(vuelo Vuelo) {
	codigo, fecha := vuelo.ObtenerCodigoVuelo(), vuelo.ObtenerFecha()
	origen, destino := vuelo.ObtenerOrigenYdestino()
	conexion := Conexion{origen, destino}
	clave := FechaYCodigo{fecha, codigo}

	if a.vuelos.Pertenece(codigo) {
		ant := a.vuelos.Obtener(codigo)
		fechaAnt := ant.ObtenerFecha()
		claveAnt := FechaYCodigo{fechaAnt, codigo}
		origAnt, destAnt := ant.ObtenerOrigenYdestino()
		conxAnt := Conexion{origAnt, destAnt}

		if a.conexiones.Pertenece(conxAnt) {
			dicc := a.conexiones.Obtener(conxAnt)
			dicc.Borrar(claveAnt)
			if dicc.Cantidad() == 0 {
				a.conexiones.Borrar(conxAnt)
			}
		}

		if a.tableroAsc.Pertenece(fechaAnt) {
			lista := a.tableroAsc.Obtener(fechaAnt)
			it := (*lista).Iterador()
			for it.HaySiguiente() {
				if it.VerActual() == codigo {
					it.Borrar()
					break
				}
				it.Siguiente()
			}
			if (*lista).Largo() == 0 {
				a.tableroAsc.Borrar(fechaAnt)
				a.tableroDesc.Borrar(fechaAnt)
			}
		}
	}

	if !a.conexiones.Pertenece(conexion) {
		dicc := TDADiccionario.CrearABB[FechaYCodigo, Vuelo](CompararFechaYCodigo)
		dicc.Guardar(clave, vuelo)
		a.conexiones.Guardar(conexion, dicc)
	} else {
		a.conexiones.Obtener(conexion).Guardar(clave, vuelo)
	}

	var lista *TDALista.Lista[string]
	if !a.tableroAsc.Pertenece(fecha) {
		nueva := TDALista.CrearListaEnlazada[string]()
		lista = &nueva
		a.tableroAsc.Guardar(fecha, lista)
		a.tableroDesc.Guardar(fecha, lista)
	} else {
		lista = a.tableroAsc.Obtener(fecha)
	}
	insertarOrdenado(*lista, codigo)

	a.vuelos.Guardar(codigo, vuelo)
}

func (a *aeropuertoImplementacion) Info_vuelo(codigo string) (string, error) {
	if !a.vuelos.Pertenece(codigo) {
		return "", fmt.Errorf("Error en comando info_vuelo")
	}
	return a.vuelos.Obtener(codigo).ObtenerInfo(), nil
}

func (a *aeropuertoImplementacion) Mostrar_prioridad(k int) {
	iter := a.vuelos.Iterador()
	var vuelos []Vuelo
	for iter.HaySiguiente() {
		_, v := iter.VerActual()
		vuelos = append(vuelos, v)
		iter.Siguiente()
	}
	heap := TDAHeap.CrearHeapArr(vuelos, compararVuelos)
	for i := 0; i < k && !heap.EstaVacia(); i++ {
		v := heap.Desencolar()
		fmt.Printf("%s - %s\n", v.ObtenerPrioridad(), v.ObtenerCodigoVuelo())
	}
}

func (a *aeropuertoImplementacion) Siguiente_vuelo(origen, destino, fecha string) Vuelo {
	conx := Conexion{origen, destino}
	if !a.conexiones.Pertenece(conx) {
		fmt.Printf("No hay vuelo registrado desde %s hacia %s desde %s\n", origen, destino, fecha)
		return nil
	}
	desde := FechaYCodigo{fecha, "-0000"}
	iter := a.conexiones.Obtener(conx).IteradorRango(&desde, nil)
	if iter.HaySiguiente() {
		_, v := iter.VerActual()
		return v
	}
	fmt.Printf("No hay vuelo registrado desde %s hacia %s desde %s\n", origen, destino, fecha)
	return nil
}

func (a *aeropuertoImplementacion) Ver_tablero_ASC(k int, desde, hasta string) {
	iter := a.tableroAsc.IteradorRango(&desde, &hasta)
	i := 0
	for iter.HaySiguiente() && i < k {
		fecha, lista := iter.VerActual()
		it := (*lista).Iterador()
		for it.HaySiguiente() && i < k {
			fmt.Printf("%s - %s\n", fecha, it.VerActual())

			it.Siguiente()
			i++
		}
		iter.Siguiente()
	}
}

func (a *aeropuertoImplementacion) Ver_tablero_DESC(k int, desde, hasta string) {
	iter := a.tableroDesc.IteradorRango(&hasta, &desde)
	i := 0
	for iter.HaySiguiente() && i < k {
		fecha, lista := iter.VerActual()
		pila := TDAPila.CrearPilaDinamica[string]()
		it := (*lista).Iterador()
		for it.HaySiguiente() {
			pila.Apilar(it.VerActual())
			it.Siguiente()
		}
		for !pila.EstaVacia() && i < k {
			fmt.Printf("%s - %s\n", fecha, pila.Desapilar())
			i++
		}
		iter.Siguiente()
	}
}

func (a *aeropuertoImplementacion) Borrar(desde, hasta string) {
	iter := a.tableroAsc.IteradorRango(&desde, &hasta)
	var fechasABorrar []string

	for iter.HaySiguiente() {
		fecha, lista := iter.VerActual()
		it := (*lista).Iterador()

		for it.HaySiguiente() {
			codigo := it.VerActual()

			if a.vuelos.Pertenece(codigo) {
				vuelo := a.vuelos.Obtener(codigo)
				origen, destino := vuelo.ObtenerOrigenYdestino()
				conexion := Conexion{origen, destino}
				clave := FechaYCodigo{vuelo.ObtenerFecha(), codigo}

				a.vuelos.Borrar(codigo)

				if a.conexiones.Pertenece(conexion) {
					conx := a.conexiones.Obtener(conexion)
					conx.Borrar(clave)
					if conx.Cantidad() == 0 {
						a.conexiones.Borrar(conexion)
					}
				}

				fmt.Println(vuelo.ObtenerInfo())
			}

			it.Borrar() //borra directamente el vuelo de la lista
		}

		if (*lista).Largo() == 0 {
			fechasABorrar = append(fechasABorrar, fecha)
		}

		iter.Siguiente()
	}

	for _, fecha := range fechasABorrar {
		a.tableroAsc.Borrar(fecha)
		a.tableroDesc.Borrar(fecha)
	}
}
