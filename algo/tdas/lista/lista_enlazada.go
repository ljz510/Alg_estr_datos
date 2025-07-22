package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

// Estructura implementada para la interfaz Lista[T].
type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

// Función privada que había sugerido el bot en el tda cola, ahora en el tda lista.
func nuevoNodoLista[T any](dato T) *nodoLista[T] {
	return &nodoLista[T]{dato: dato}
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
}

func (l *listaEnlazada[T]) EstaVacia() bool {
	return l.primero == nil
}

func (l *listaEnlazada[T]) InsertarPrimero(elemento T) {
	nuevo := nuevoNodoLista(elemento)
	if l.EstaVacia() {
		l.ultimo = nuevo
	} else {
		nuevo.siguiente = l.primero
	}
	l.primero = nuevo
	l.largo++
}

func (l *listaEnlazada[T]) InsertarUltimo(elemento T) {
	nuevo := nuevoNodoLista(elemento)
	if l.EstaVacia() {
		l.primero = nuevo
	} else {
		l.ultimo.siguiente = nuevo
	}
	l.ultimo = nuevo
	l.largo++
}

func (l *listaEnlazada[T]) BorrarPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	dato := l.primero.dato
	l.primero = l.primero.siguiente
	if l.primero == nil {
		l.ultimo = nil
	}
	l.largo--
	return dato
}

func (l *listaEnlazada[T]) VerPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.primero.dato
}

func (l *listaEnlazada[T]) VerUltimo() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.ultimo.dato
}

func (l *listaEnlazada[T]) Largo() int {
	return l.largo
}

// ITERADOR INTERNO ----------
func (l *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := l.primero
	for actual != nil {
		if !visitar(actual.dato) {
			break
		}
		actual = actual.siguiente
	}
}

// ITERADOR EXTERNO ----------

func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iteradorListaEnlazada[T]{lista: l, actual: l.primero, anterior: nil}
}

type iteradorListaEnlazada[T any] struct {
	lista    *listaEnlazada[T]
	actual   *nodoLista[T]
	anterior *nodoLista[T]
}

func (i *iteradorListaEnlazada[T]) VerActual() T {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return i.actual.dato
}

func (i *iteradorListaEnlazada[T]) HaySiguiente() bool {
	return i.actual != nil
}

func (i *iteradorListaEnlazada[T]) Siguiente() {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	i.anterior = i.actual
	i.actual = i.actual.siguiente
}

func (i *iteradorListaEnlazada[T]) Insertar(elemento T) {
	// corrección: SIMPLIFICACIÓN DE LA LÓGICA
	nuevo := &nodoLista[T]{dato: elemento, siguiente: i.actual}

	if i.anterior == nil {
		i.lista.primero = nuevo
	} else {
		i.anterior.siguiente = nuevo
	}

	if i.actual == nil {
		i.lista.ultimo = nuevo
	}

	i.actual = nuevo
	i.lista.largo++
}

func (i *iteradorListaEnlazada[T]) Borrar() T {
	// corrección: SIMPLIFICACIÓN DE LA LÓGICA
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	dato := i.actual.dato

	if i.anterior == nil {
		i.lista.primero = i.actual.siguiente
	} else {
		i.anterior.siguiente = i.actual.siguiente
	}

	if i.actual == i.lista.ultimo {
		i.lista.ultimo = i.anterior
	}

	i.actual = i.actual.siguiente
	i.lista.largo--

	return dato
}
