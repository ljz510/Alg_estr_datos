package lista_test

import (
	"testing"

	"tdas/lista"

	"github.com/stretchr/testify/require"
)

/*
Pruebas:
-Al insertar un elemento en la posición en la que se crea el iterador, efectivamente se inserta al principio.
-Insertar un elemento cuando el iterador está al final efectivamente es equivalente a insertar al final.
-Insertar un elemento en el medio se hace en la posición correcta.
-Al remover el elemento cuando se crea el iterador, cambia el primer elemento de la lista.
-Remover el último elemento con el iterador cambia el último de la lista.
-Verificar que al remover un elemento del medio, este no está.
-Otros casos borde que pueden encontrarse al utilizar el iterador externo.
-Casos del iterador interno, incluyendo casos con corte (la función visitar devuelve false eventualmente).
*/

func TestListaVacia(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	require.True(t, l.EstaVacia())
	require.Equal(t, 0, l.Largo())

	require.PanicsWithValue(t, "La lista esta vacia", func() { l.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { l.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { l.BorrarPrimero() })
}

func TestInsertarYVer(t *testing.T) {
	l := lista.CrearListaEnlazada[string]()
	l.InsertarPrimero("a")
	require.False(t, l.EstaVacia())
	require.Equal(t, 1, l.Largo())
	require.Equal(t, "a", l.VerPrimero())
	require.Equal(t, "a", l.VerUltimo())

	l.InsertarUltimo("b")
	require.Equal(t, 2, l.Largo())
	require.Equal(t, "a", l.VerPrimero())
	require.Equal(t, "b", l.VerUltimo())
}

func TestBorrar(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarPrimero(1)
	l.InsertarUltimo(2)
	l.InsertarUltimo(3)

	require.Equal(t, 3, l.Largo())
	require.Equal(t, 1, l.BorrarPrimero())
	require.Equal(t, 2, l.Largo())
	require.Equal(t, 2, l.VerPrimero())
	require.Equal(t, 3, l.VerUltimo())

	require.Equal(t, 2, l.BorrarPrimero())
	require.Equal(t, 3, l.BorrarPrimero())

	require.True(t, l.EstaVacia())
}

//// ------------------- PRUEBAS DEL ITERADOR EXTERNO -------------------
//Se agregaron más pruebas para el iterador externo, incluyendo casos de borde.

//CORRECCIÓN: AGREGAR CASOS BORDE PARA EL ITERADOR EXTERNO

func TestIteradorInsertarPrincipio(t *testing.T) {
	// Al insertar un elemento en la posición en la que se crea el iterador, efectivamente se inserta al principio.
	l := lista.CrearListaEnlazada[int]()
	l.InsertarUltimo(2)
	l.InsertarUltimo(3)

	iter := l.Iterador()
	iter.Insertar(1)

	require.Equal(t, 3, l.Largo())
	require.Equal(t, 1, l.VerPrimero())
}

func TestIteradorInsertarFinal(t *testing.T) {
	// Insertar un elemento cuando el iterador está al final efectivamente es equivalente a insertar al final.
	l := lista.CrearListaEnlazada[int]()
	l.InsertarUltimo(1)
	l.InsertarUltimo(2)

	iter := l.Iterador()
	for iter.HaySiguiente() {
		iter.Siguiente()
	}
	iter.Insertar(3)

	require.Equal(t, 3, l.Largo())
	require.Equal(t, 3, l.VerUltimo())
}

func TestIteradorInsertarMedio(t *testing.T) {
	// Insertar un elemento en el medio se hace en la posición correcta.
	l := lista.CrearListaEnlazada[int]()
	l.InsertarUltimo(1)
	l.InsertarUltimo(3)

	iter := l.Iterador()
	iter.Siguiente()
	iter.Insertar(2)

	// Esperado: 1,2,3
	valores := []int{}
	l.Iterar(func(x int) bool {
		valores = append(valores, x)
		return true
	})

	require.Equal(t, []int{1, 2, 3}, valores)
}

func TestIteradorRemoverPrimero(t *testing.T) {
	// Al remover el elemento cuando se crea el iterador, cambia el primer elemento de la lista.
	l := lista.CrearListaEnlazada[int]()
	l.InsertarUltimo(1)
	l.InsertarUltimo(2)
	l.InsertarUltimo(3)

	iter := l.Iterador()
	require.Equal(t, 1, iter.Borrar())

	require.Equal(t, 2, l.VerPrimero())
}

func TestIteradorRemoverUltimo(t *testing.T) {
	// Remover el último elemento con el iterador cambia el último de la lista.
	l := lista.CrearListaEnlazada[int]()
	l.InsertarUltimo(1)
	l.InsertarUltimo(2)
	l.InsertarUltimo(3)

	iter := l.Iterador()
	for iter.HaySiguiente() {
		iter.Siguiente()
	}
	iter = l.Iterador()
	iter.Siguiente()
	iter.Siguiente()
	require.Equal(t, 3, iter.Borrar())

	require.Equal(t, 2, l.VerUltimo())
}

func TestIteradorRemoverMedio(t *testing.T) {
	// Verificar que al remover un elemento del medio, este no está.
	l := lista.CrearListaEnlazada[int]()
	l.InsertarUltimo(1)
	l.InsertarUltimo(2)
	l.InsertarUltimo(3)

	iter := l.Iterador()
	iter.Siguiente()
	require.Equal(t, 2, iter.VerActual())
	iter.Borrar()

	// Esperado: 1,3
	valores := []int{}
	l.Iterar(func(x int) bool {
		valores = append(valores, x)
		return true
	})

	require.Equal(t, []int{1, 3}, valores)
}

// AGREGACIONES DE TESTS: CASOS BORDE.
func TestCasoBordeIteradorInsertarEnListaVacia(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	iter := l.Iterador()
	iter.Insertar(42)

	require.Equal(t, 1, l.Largo())
	require.Equal(t, 42, l.VerPrimero())
	require.Equal(t, 42, l.VerUltimo())
}
func TestCasoBordeIteradorBorrarUnicoElemento(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarUltimo(100)

	iter := l.Iterador()
	require.Equal(t, 100, iter.Borrar())

	require.True(t, l.EstaVacia())
}
func TestCasoBordeIteradorVerActualEnFinPaniquear(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	iter := l.Iterador()

	require.Panics(t, func() {
		iter.VerActual()
	})
}

func TestCasoBordeIteradorInsertarDespuesDeBorrar(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarUltimo(1)
	l.InsertarUltimo(3)

	iter := l.Iterador()
	require.Equal(t, 1, iter.VerActual())
	iter.Borrar()
	iter.Insertar(2)

	valores := []int{}
	l.Iterar(func(x int) bool {
		valores = append(valores, x)
		return true
	})

	require.Equal(t, []int{2, 3}, valores)
}

//// ------------------- PRUEBAS DEL ITERADOR INTERNO -------------------

func TestIteradorInternoCompleto(t *testing.T) {
	// Probar el caso de iteración sin condición de corte (iterar toda la lista).
	l := lista.CrearListaEnlazada[int]()
	l.InsertarUltimo(1)
	l.InsertarUltimo(2)
	l.InsertarUltimo(3)

	suma := 0
	l.Iterar(func(x int) bool {
		suma += x
		return true
	})

	require.Equal(t, 6, suma)
}

func TestIteradorInternoConCorte(t *testing.T) {
	// Probar iteración con condición de corte (que en un momento determinado la función visitar dé false).
	l := lista.CrearListaEnlazada[int]()
	l.InsertarUltimo(1)
	l.InsertarUltimo(2)
	l.InsertarUltimo(3)
	l.InsertarUltimo(4)

	suma := 0
	l.Iterar(func(x int) bool {
		suma += x
		return x != 2 // Cortar cuando x == 2
	})

	require.Equal(t, 3, suma) // 1+2
}
