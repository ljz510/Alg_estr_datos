package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

// Se puedan apilar elementos, que al desapilarlos se mantenga el invariante de la cola (que esta es FIFO). Probar con elementos diferentes,
// y ver que salgan en el orden deseado.
// (minimo volumen)
func TestEncolarDesencolarEnteros(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	cola.Encolar(1)
	require.Equal(t, 1, cola.VerPrimero())

	cola.Encolar(2)
	require.Equal(t, 1, cola.VerPrimero())

	cola.Encolar(3)
	require.Equal(t, 1, cola.VerPrimero())

	require.Equal(t, 1, cola.Desencolar())
	require.Equal(t, 2, cola.VerPrimero())

	require.Equal(t, 2, cola.Desencolar())
	require.Equal(t, 3, cola.VerPrimero())

	require.Equal(t, 3, cola.Desencolar())
	require.True(t, cola.EstaVacia())
}

func TestEncolarDesencolarStrings(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()

	cola.Encolar("hola")
	cola.Encolar("todo bien")

	require.Equal(t, "hola", cola.VerPrimero())
	require.Equal(t, "hola", cola.Desencolar())
	require.Equal(t, "todo bien", cola.Desencolar())
	require.True(t, cola.EstaVacia())
}

func TestVerPrimero(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(10)
	require.Equal(t, 10, cola.VerPrimero())
	cola.Encolar(20)
	require.Equal(t, 10, cola.VerPrimero())
	cola.Desencolar()
	require.Equal(t, 20, cola.VerPrimero())
}

func TestColaVaciaPanics(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
}

// Prueba de volumen: Se pueden encolar muchos elementos y desencolar todos.
// Verifica que se mantenga el invariante.
//(Gran volumen)

func TestVolumen(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	const n = 10000
	for i := 0; i < n; i++ {
		cola.Encolar(i)
	}
	for i := 0; i < n; i++ {
		require.Equal(t, i, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
}
