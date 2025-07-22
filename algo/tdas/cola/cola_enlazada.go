package cola

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

// funcion sugerida por rosita:
// Sería buena idea tener una funcion (privada) de creacion para el nodo (con un nombre adecuado)

func nuevoNodoCola[T any](dato T) *nodoCola[T] {
	return &nodoCola[T]{dato: dato}
}

// inicializa una cola enlazada vacía
func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{}
}

func (c *colaEnlazada[T]) Encolar(dato T) {
	nuevo := nuevoNodoCola(dato)
	if c.EstaVacia() {
		c.primero = nuevo
	} else {
		c.ultimo.prox = nuevo
	}
	// corrección: colocar afuera
	c.ultimo = nuevo
}

func (c *colaEnlazada[T]) Desencolar() T {
	if c.EstaVacia() {
		// corrección: llamar a panic de forma directa
		panic("La cola esta vacia")
	}
	dato := c.primero.dato
	c.primero = c.primero.prox
	if c.primero == nil {
		c.ultimo = nil
	}
	return dato
}

func (c *colaEnlazada[T]) EstaVacia() bool {
	return c.primero == nil
}

func (c *colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		paniquear()
	}
	return c.primero.dato
}

func paniquear() {
	panic("La cola esta vacia")
}
