package cola_prioridad

const (
	TAM_INI    = 29
	F_REDIM    = 2
	COND_REDIM = 2
)

type colaConPrioridad[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

func CrearHeap[T any](cmp func(T, T) int) ColaPrioridad[T] {
	return &colaConPrioridad[T]{
		datos: make([]T, TAM_INI),
		cant:  0,
		cmp:   cmp,
	}
}

func CrearHeapArr[T any](arreglo []T, cmp func(T, T) int) ColaPrioridad[T] {
	datos := make([]T, len(arreglo)+TAM_INI)
	copy(datos, arreglo)
	cola := &colaConPrioridad[T]{datos: datos, cant: len(arreglo), cmp: cmp}

	// Aplico heapify para construir el heap a partir del arreglo dado
	heapify(cola.datos, cola.cant, cola.cmp)

	/*for i := (cola.cant - 1) / 2; i >= 0; i-- {
		downHeap(cola.datos, i, cola.cant, cola.cmp)
	}*/
	return cola
}

// =================== Primitivas del TDA ===================

func (cola *colaConPrioridad[T]) EstaVacia() bool {
	return cola.Cantidad() == 0
}

func (cola *colaConPrioridad[T]) Cantidad() int {
	return cola.cant
}

func (cola *colaConPrioridad[T]) Encolar(dato T) {

	if cola.Cantidad() == cap(cola.datos) {
		cola.redimensionar(cap(cola.datos) * F_REDIM)
	}
	cola.datos[cola.cant] = dato
	cola.cant++
	upHeap(cola.datos, cola.cant-1, cola.cmp)
}

func (cola *colaConPrioridad[T]) VerMax() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	return cola.datos[0]
}

func (cola *colaConPrioridad[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	dato := cola.datos[0]
	cola.cant--
	if cola.cant > 0 {
		cola.datos[0] = cola.datos[cola.cant]
		downHeap(cola.datos, 0, cola.cant, cola.cmp)
	}
	if cola.cant > 0 && cola.cant <= cap(cola.datos)/COND_REDIM {
		nuevaCapacidad := cap(cola.datos) / F_REDIM
		if nuevaCapacidad >= TAM_INI {
			cola.redimensionar(nuevaCapacidad)
		}
	}
	return dato
}

// =================== Funciones auxiliares ===================

func upHeap[T any](datos []T, pos int, cmp func(T, T) int) {
	for pos > 0 {
		padre := (pos - 1) / 2
		if cmp(datos[pos], datos[padre]) <= 0 {
			return
		}
		datos[pos], datos[padre] = datos[padre], datos[pos]
		pos = padre
	}
}

func downHeap[T any](datos []T, pos, cant int, cmp func(T, T) int) {
	for {
		izq := 2*pos + 1
		der := 2*pos + 2
		mayor := pos

		if izq < cant && cmp(datos[izq], datos[mayor]) > 0 {
			mayor = izq
		}
		if der < cant && cmp(datos[der], datos[mayor]) > 0 {
			mayor = der
		}
		if mayor == pos {
			break
		}
		datos[pos], datos[mayor] = datos[mayor], datos[pos]
		pos = mayor
	}
}

func (cola *colaConPrioridad[T]) redimensionar(nuevaCapacidad int) {
	nuevosDatos := make([]T, nuevaCapacidad)
	copy(nuevosDatos, cola.datos[:cola.cant])
	cola.datos = nuevosDatos
}

// =================== Heapsort ===================

func HeapSort[T any](elementos []T, cmp func(T, T) int) {
	n := len(elementos)
	// Heapify
	heapify(elementos, n, cmp)
	// Sort
	for i := n - 1; i > 0; i-- {
		// Swap:
		elementos[0], elementos[i] = elementos[i], elementos[0]
		downHeap(elementos, 0, i, cmp)
	}
}

// =================== Heapify ===================
func heapify[T any](datos []T, cant int, cmp func(T, T) int) {
	for i := (cant - 1) / 2; i >= 0; i-- {
		downHeap(datos, i, cant, cmp)
	}
}
