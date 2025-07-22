package lista

type Lista[T any] interface {
	// EstaVacia devuelve verdadero si la lista no tiene elementos, false en caso contrario.
	EstaVacia() bool
	// InsertarPrimero agrega un nuevo elemento al principio de la lista.
	InsertarPrimero(T)
	// InsertarUltimo agrega un nuevo elemento al final de la lista.
	InsertarUltimo(T)
	// BorrarPrimero saca el primer elemento de la lista. En caso de no tener elementos deberia entrar en panico.
	BorrarPrimero() T
	// BorrarUltimo saca el último elemento de la lista. En caso de no tener elementos deberia entrar en panico.
	VerPrimero() T
	// VerUltimo devuelve el último elemento de la lista. En caso de no tener elementos deberia entrar en panico.
	VerUltimo() T
	// Largo devuelve la cantidad de elementos que tiene la lista.
	Largo() int
	// Iterador interno, itera de manera transparente al usuario
	Iterar(visitar func(T) bool)
	// Iterador externo, devuelve un iterador(estructura) para recorrer la lista
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	// VerActual devuelve el elemento actual del iterador. Si ya se iteró todo, entra en pánico con un mensaje
	VerActual() T
	// HaySiguiente devuelve verdadero si el iterador tiene un siguiente elemento, false en caso contrario.
	HaySiguiente() bool
	// Siguiente avanza el iterador al siguiente elemento. Si ya se iteró todo, entra en pánico con un mensaje
	Siguiente()
	// Insertar inserta un nuevo elemento en la posición actual del iterador.
	Insertar(T)
	// Borrar borra el elemento actual del iterador. Si ya se iteró todo, entra en pánico con un mensaje
	Borrar() T
}
