package diccionario

import (
	"fmt"

	"hash/fnv"
)

const (
	VACIO   = 0
	OCUPADO = 1
	BORRADO = 2

	FACTOR_REDIMENSION = 2
	TAMAÑO_INICIAL     = 23
	FACTOR_CARGA       = 0.7
)

type celdaHash[K comparable, V any] struct {
	clave  K
	dato   V
	estado int
}

type hashCerrado[K comparable, V any] struct {
	tabla    []celdaHash[K, V]
	cantidad int
	tam      int
	borrados int
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	hash := &hashCerrado[K, V]{}
	hash.tabla = hash.crearTabla(TAMAÑO_INICIAL)
	hash.tam = TAMAÑO_INICIAL
	return hash
}

func (hash *hashCerrado[K, V]) crearTabla(tam int) []celdaHash[K, V] {
	return make([]celdaHash[K, V], tam)
}

//INICIO DE FUNCIONES AUXILIARES:

// Transforma un tipo de dato genérico a un array de bytes
func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func hashFNV(data []byte) uint64 {
	hasher := fnv.New64a()
	hasher.Write(data)
	return hasher.Sum64()
}

func (hash *hashCerrado[K, V]) obtPos(clave K, tam int) int {
	return int(hashFNV(convertirABytes(clave)) % uint64(tam))
}

// correcion de la funcion buscar
func (hash *hashCerrado[K, V]) buscar(clave K) int {
	//simplificacion de funcion buscar para que devuleva la posicion
	pos := hash.obtPos(clave, hash.tam)
	primeraLibre := -1
	for {
		if hash.tabla[pos].estado == OCUPADO {
			if hash.tabla[pos].clave == clave {
				return pos
			}
		} else {
			if primeraLibre == -1 {
				primeraLibre = pos
			}
			// si la celda está vacía y no hemos encontrado la clave,
			// devolvemos la posición
			if hash.tabla[pos].estado == VACIO {
				return primeraLibre
			}
		}
		pos = (pos + 1) % hash.tam
	}
}

func (hash *hashCerrado[K, V]) factorDeCarga() float64 {
	return (float64(hash.cantidad + hash.borrados)) / float64(hash.tam)
}

func esPrimo(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Si haces el redimensionado con Guardar(), esta función pasa a ser únicamente usada en Guardar. Luego buscar la posición de nuevo es innecesario,
// ya que ya la buscaste en el Guardar.
/*func (hash *hashCerrado[K, V]) insertarCelda(celda celdaHash[K, V], tabla []celdaHash[K, V], tam int) {
	pos := hash.obtPos(celda.clave, tam)
	for tabla[pos].estado == OCUPADO {
		pos = (pos + 1) % tam
	}
	tabla[pos] = celda
}*/
//funcion pasada a comentario, no es necesaria

func (hash *hashCerrado[K, V]) redimensionar() {

	n := hash.tam*FACTOR_REDIMENSION + 1
	for !esPrimo(n) {
		n++
	}
	nuevoTam := n
	viejaTabla := hash.tabla

	hash.tabla = hash.crearTabla(nuevoTam)
	hash.tam = nuevoTam
	hash.cantidad = 0
	hash.borrados = 0

	for _, celda := range viejaTabla {
		if celda.estado == OCUPADO {
			//Acá queres guardar los datos de tu tabla de hash vieja en la nueva... no deberías poder usar hash.Guardar()?
			hash.Guardar(celda.clave, celda.dato)
		}
	}
}

// FIN DE FUNCIONES AUXILIARES

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {
	if hash.factorDeCarga() >= FACTOR_CARGA {
		hash.redimensionar()
	}
	pos := hash.buscar(clave)
	if hash.tabla[pos].estado == OCUPADO {
		hash.tabla[pos].dato = dato
		return
	}
	hash.tabla[pos].clave = clave
	hash.tabla[pos].dato = dato
	hash.tabla[pos].estado = OCUPADO
	hash.cantidad++
}

func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {
	//De nuevo, si sacas lo del encontrado de tu función buscar, el funcionamiento se simplifica a este:
	return hash.tabla[hash.buscar(clave)].estado == OCUPADO
}

// aplique el mismo concepto de la funcion buscar
func (hash *hashCerrado[K, V]) Obtener(clave K) V {
	pos := hash.buscar(clave)
	if hash.Pertenece(clave) {
		return hash.tabla[pos].dato
	}
	panic("La clave no pertenece al diccionario")
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	pos := hash.buscar(clave)
	if !hash.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
	valor := hash.tabla[pos].dato
	hash.tabla[pos].estado = BORRADO
	hash.cantidad--
	hash.borrados++
	if hash.factorDeCarga() >= FACTOR_CARGA {
		hash.redimensionar()
	}
	return valor
}

func (hash *hashCerrado[K, V]) Cantidad() int {
	return hash.cantidad
}

// ITERADOR INTERNO
func (hash *hashCerrado[K, V]) Iterar(funcion func(clave K, dato V) bool) {
	for _, celda := range hash.tabla {
		if celda.estado == OCUPADO {
			if !funcion(celda.clave, celda.dato) {
				return
			}
		}
	}
}

// ITERADOR EXTERNO
type iterHashCerrado[K comparable, V any] struct {
	hash      *hashCerrado[K, V]
	posActual int
}

func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	for i := 0; i < hash.tam; i++ {
		if hash.tabla[i].estado == OCUPADO {
			return &iterHashCerrado[K, V]{hash: hash, posActual: i}
		}
	}
	return &iterHashCerrado[K, V]{hash: hash, posActual: hash.tam}
}

func (iter *iterHashCerrado[K, V]) HaySiguiente() bool {
	// verifico si pos actual es igual al largo de la tabla
	//Para asegurar buen funcionamiento, es
	//un poco más correcto verificar estar "dentro" del límite,
	// y no directamente "en" el límite.
	return iter.posActual < iter.hash.tam
}

func (iter *iterHashCerrado[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.hash.tabla[iter.posActual].clave, iter.hash.tabla[iter.posActual].dato
}

func (iter *iterHashCerrado[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iter.posActual++
	//corrección del iterador:
	for iter.posActual < iter.hash.tam && iter.hash.tabla[iter.posActual].estado != OCUPADO {
		iter.posActual++
	}
	if iter.posActual == iter.hash.tam {
		return
	}
}
