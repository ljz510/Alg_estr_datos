package diccionario_test

import (
	"fmt"
	"strings"
	"tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

func comparacionInt(a, b int) int {
	return a - b
}

func comparacionString(a, b string) int {
	return strings.Compare(a, b)
}

func Test(t *testing.T) {
	abb := diccionario.CrearABB[int, int](comparacionInt)
	abb.Guardar(1, 1)
	abb.Guardar(2, 2)
	abb.Guardar(3, 3)
	abb.Guardar(4, 4)
	abb.Guardar(6, 6)
	abb.Guardar(5, 5)
	abb.Guardar(7, 7)

	desde := 2
	hasta := 5

	iterador := abb.IteradorRango(&desde, &hasta)

	esperado := []int{2, 3, 4, 5}
	indice := 0

	for iterador.HaySiguiente() {
		claveActual, _ := iterador.VerActual()
		fmt.Print(claveActual)
		require.Equal(t, esperado[indice], claveActual)
		indice++
		iterador.Siguiente()
	}
}

func TestABBVacio(t *testing.T) {
	// Test de un ABB vacío
	dic := diccionario.CrearABB[string, string](comparacionString)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })
}

func TestABBClaveDefault(t *testing.T) {
	// Test de un ABB con clave por defecto

	dic := diccionario.CrearABB[string, string](comparacionString)
	require.False(t, dic.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("") })

	dicNum := diccionario.CrearABB[int, string](comparacionInt)
	require.False(t, dicNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Borrar(0) })
}

func TestUnElemento(t *testing.T) {

	dic := diccionario.CrearABB[string, int](comparacionString)
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestABBAgregar(t *testing.T) {

	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := diccionario.CrearABB[string, string](comparacionString)
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	require.False(t, dic.Pertenece(claves[1]))
	dic.Guardar(claves[1], valores[1])
	require.EqualValues(t, 2, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))

	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[2], valores[2])
	require.EqualValues(t, 3, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], dic.Obtener(claves[2]))
}

func TestReemplazoD(t *testing.T) {
	// Test de reemplazo de un elemento en el ABB
	clave := "Gato"
	clave2 := "Perro"
	dic := diccionario.CrearABB[string, string](comparacionString)
	dic.Guardar(clave, "miau")
	dic.Guardar(clave2, "guau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, "miau", dic.Obtener(clave))
	require.EqualValues(t, "guau", dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave, "miu")
	dic.Guardar(clave2, "baubau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "miu", dic.Obtener(clave))
	require.EqualValues(t, "baubau", dic.Obtener(clave2))
}

func TestABBBorrar(t *testing.T) {

	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := diccionario.CrearABB[string, string](comparacionString)

	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])

	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], dic.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[2]) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[2]))

	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[0]) })
	require.EqualValues(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[0]) })

	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dic.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[1]) })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[1]) })
}

func TestConClavesNum(t *testing.T) {

	dic := diccionario.CrearABB[int, string](comparacionInt)
	clave := 10
	valor := "numero"

	dic.Guardar(clave, valor)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, valor, dic.Obtener(clave))
	require.EqualValues(t, valor, dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(clave) })
}

func TestPruebasDeVolumen(t *testing.T) {
	// Pruebas de volumen con un ABB con claves numéricas
	dic := diccionario.CrearABB[int, int](comparacionInt)
	cantidad := 10000
	for i := 0; i < cantidad; i++ {
		dic.Guardar(i, i)
	}
	require.EqualValues(t, cantidad, dic.Cantidad())
	for i := 0; i < cantidad; i++ {
		require.True(t, dic.Pertenece(i))
		require.EqualValues(t, i, dic.Obtener(i))
	}
	for i := 0; i < cantidad; i++ {
		require.EqualValues(t, i, dic.Borrar(i))
	}
	require.EqualValues(t, 0, dic.Cantidad())
}

func TestPruebasDeVolumenConClavesNumericas(t *testing.T) {
	// Pruebas de volumen con un ABB con claves numéricas
	dic := diccionario.CrearABB[int, int](comparacionInt)
	cantidad := 10000
	for i := 0; i < cantidad; i++ {
		dic.Guardar(i, i)
	}
	require.EqualValues(t, cantidad, dic.Cantidad())
	for i := 0; i < cantidad; i++ {
		require.True(t, dic.Pertenece(i))
		require.EqualValues(t, i, dic.Obtener(i))
	}
	for i := 0; i < cantidad; i++ {
		require.EqualValues(t, i, dic.Borrar(i))
	}
	require.EqualValues(t, 0, dic.Cantidad())
}

func TestPruebasDeVolumenConClavesNumericasConBorrados(t *testing.T) {
	// Pruebas de volumen con un ABB con claves numéricas con borrados
	dic := diccionario.CrearABB[int, int](comparacionInt)
	cantidad := 10000
	for i := 0; i < cantidad; i++ {
		dic.Guardar(i, i)
	}
	require.EqualValues(t, cantidad, dic.Cantidad())
	for i := 0; i < cantidad; i++ {
		require.True(t, dic.Pertenece(i))
		require.EqualValues(t, i, dic.Obtener(i))
	}
	for i := 0; i < cantidad; i++ {
		require.EqualValues(t, i, dic.Borrar(i))
	}
	require.EqualValues(t, 0, dic.Cantidad())
	for i := 0; i < cantidad; i++ {
		require.False(t, dic.Pertenece(i))
		require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(i) })
		require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(i) })
	}
}

func TestPruebasDeVolumenConClavesNumericasConBorradosIntermedios(t *testing.T) {
	// Pruebas de volumen con un ABB con claves numéricas con borrados intermedios
	dic := diccionario.CrearABB[int, int](comparacionInt)
	cantidad := 10000
	for i := 0; i < cantidad; i++ {
		dic.Guardar(i, i)
	}
	require.EqualValues(t, cantidad, dic.Cantidad())
	for i := 0; i < cantidad; i += 2 {
		require.EqualValues(t, i, dic.Borrar(i))
	}
	require.EqualValues(t, cantidad/2, dic.Cantidad())
	for i := 0; i < cantidad; i += 2 {
		require.False(t, dic.Pertenece(i))
		require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(i) })
		require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(i) })
	}
	for i := 1; i < cantidad; i += 2 {
		require.True(t, dic.Pertenece(i))
		require.EqualValues(t, i, dic.Obtener(i))
	}
}
