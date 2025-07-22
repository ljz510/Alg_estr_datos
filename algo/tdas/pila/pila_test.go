package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

// Se pueda crear una Pila vacía, y ésta se comporta como tal.
func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia(), "pila recien creada debe estar vacia")
	// mas pruebas para este caso...

	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() }, "una pila vacia no tiene tope")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() }, "no puedo desapilar una pila vacia")
}

// Se puedan apilar elementos, que al desapilarlos se mantenga el invariante de pila (que esta es LIFO). Probar con elementos diferentes,
// y ver que salgan en el orden deseado.
// (minimo volumen)
func TestApilarDesapilarEnteros(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	pila.Apilar(2)
	pila.Apilar(3)
	pila.Apilar(5)

	pila.Desapilar()
	pila.Desapilar()
	pila.Desapilar()

	require.True(t, pila.EstaVacia(), "pila desapilada totalmente debe estar vacía")

}

func TestApilarDesapilarCadenas(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()

	pila.Apilar("hola")
	pila.Apilar("como")
	pila.Apilar("estan")

	pila.Desapilar()
	pila.Desapilar()
	pila.Desapilar()

	require.True(t, pila.EstaVacia(), "pila desapilada totalmente debería estar vacía")

}

func TestVerTope(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	pila.Apilar(85)
	require.Equal(t, 85, pila.VerTope(), "El tope debería ser 85")
	pila.Apilar(150)
	require.Equal(t, 150, pila.VerTope(), "El tope debería ser 150")
	pila.Apilar(1000)
	require.Equal(t, 1000, pila.VerTope(), "El tope debería ser 1000")

	pila.Desapilar()
	require.Equal(t, 150, pila.VerTope(), "El tope al desapilar debería ser 150")
	pila.Desapilar()
	require.Equal(t, 85, pila.VerTope(), "El tope al desapilar debería ser 85")
	pila.Desapilar()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() }, "una pila vacia no tiene tope")

}

// Prueba de volumen: Se pueden apilar muchos elementos (1.000, 10.000 elementos, o el volumen que corresponda):
//	hacer crecer la pila, y desapilar elementos hasta que esté vacía, comprobando que siempre cumpla el invariante.
// Recordar no apilar siempre lo mismo, validar que se cumpla siempre que el tope de la pila sea el correcto paso a paso, y que
//	el nuevo tope después de cada desapilar también sea el correcto.
//(Gran volumen)

func TestVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	for i := 0; i < 10000; i++ {
		pila.Apilar(i + 1)
	}
	require.Equal(t, 10000, pila.VerTope(), "El tope debería ser 10.000")

	for i := 10000; i > 0; i-- {
		require.Equal(t, i, pila.VerTope(), "El tope no es el esperado antes de desapilar")
		pila.Desapilar()
	}
	require.True(t, pila.EstaVacia(), "La pila debería estar vacía después de desapilar todo")
}
