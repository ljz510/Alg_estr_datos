package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

const (
	CAPACIDAD_INICIAL   int = 1
	CONST_REDIMENSION_1 int = 2
	CONST_REDIMENSION_2 int = 4
)

func CrearPilaDinamica[T any]() Pila[T] {

	return &pilaDinamica[T]{datos: make([]T, CAPACIDAD_INICIAL)}

}

// Apilar agrega un nuevo elemento a la pila.
func (pila *pilaDinamica[T]) Apilar(elemento T) {

	if pila.cantidad == len(pila.datos) {
		pila.redimensionar((len(pila.datos)) * CONST_REDIMENSION_1)
	}

	pila.datos[pila.cantidad] = elemento
	pila.cantidad++

}

// Desapilar saca el elemento tope de la pila. Si la pila tiene elementos, se quita el tope de la pila, y
// se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La pila esta vacia".
func (pila *pilaDinamica[T]) Desapilar() T {

	pila.paniquear()

	if pila.cantidad*CONST_REDIMENSION_2 <= len(pila.datos) {
		pila.redimensionar(len(pila.datos) / CONST_REDIMENSION_1)
	}

	pila.cantidad--
	return pila.datos[pila.cantidad]

}

// VerTope obtiene el valor del tope de la pila. Si la pila tiene elementos se devuelve el valor del tope.
// Si está vacía, entra en pánico con un mensaje "La pila esta vacia".
func (pila *pilaDinamica[T]) VerTope() T {
	pila.paniquear()
	return pila.datos[pila.cantidad-1]
}

// EstaVacia devuelve verdadero si la pila no tiene elementos apilados, false en caso contrario.
func (pila *pilaDinamica[T]) EstaVacia() bool {

	return pila.cantidad == 0

}

// Redimensiona la pila de manera implicita (no ve barbara esta funcion)
func (pila *pilaDinamica[T]) redimensionar(nuevaCap int) {
	nuevosDatos := make([]T, nuevaCap)
	copy(nuevosDatos, pila.datos)
	pila.datos = nuevosDatos
}

// Funcion auxiliar sugerida para el control de pila vacia
func (pila *pilaDinamica[T]) paniquear() {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}

}
