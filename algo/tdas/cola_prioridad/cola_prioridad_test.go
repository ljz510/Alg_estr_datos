package cola_prioridad_test

import (
	cola_prioridad "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

// Comparación para un Heap de Máximo
func CompararMaximo(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

// Comparación para un Heap de Mínimo
func CompararMinimo(a, b int) int {
	if a < b {
		return 1
	} else if a > b {
		return -1
	}
	return 0
}

// Heap de Máximos
func TestHeapMaximo(t *testing.T) {
	heap := cola_prioridad.CrearHeap(CompararMaximo)

	// Encolar elementos
	heap.Encolar(10)
	heap.Encolar(20)
	heap.Encolar(5)

	// Desencolar y verificar el orden
	if dato := heap.Desencolar(); dato != 20 {
		t.Errorf("Esperado 20, obtenido %d", dato)
	}
	if dato := heap.Desencolar(); dato != 10 {
		t.Errorf("Esperado 10, obtenido %d", dato)
	}
	if dato := heap.Desencolar(); dato != 5 {
		t.Errorf("Esperado 5, obtenido %d", dato)
	}
}

// Heap de Mínimos
func TestHeapMinimo(t *testing.T) {
	heap := cola_prioridad.CrearHeap(CompararMinimo)

	// Encolar elementos
	heap.Encolar(10)
	heap.Encolar(20)
	heap.Encolar(5)

	// Desencolar y verificar el orden
	if dato := heap.Desencolar(); dato != 5 {
		t.Errorf("Esperado 5, obtenido %d", dato)
	}
	if dato := heap.Desencolar(); dato != 10 {
		t.Errorf("Esperado 10, obtenido %d", dato)
	}
	if dato := heap.Desencolar(); dato != 20 {
		t.Errorf("Esperado 20, obtenido %d", dato)
	}
}

// Prueba de volumen: Se pueden encolar muchos elementos y desencolar todos.
func TestVolumen(t *testing.T) {
	heap := cola_prioridad.CrearHeap(CompararMaximo)
	const n = 100000
	for i := 0; i < n; i++ {
		heap.Encolar(i)
	}
	for i := n - 1; i >= 0; i-- {
		if dato := heap.Desencolar(); dato != i {
			t.Errorf("Esperado %d, obtenido %d", i, dato)
		}
	}
}

func TestHeapVacioPanics(t *testing.T) {
	heap := cola_prioridad.CrearHeap(CompararMaximo)
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
}

func TestVerPrimeroHeap(t *testing.T) {
	heap := cola_prioridad.CrearHeap(CompararMaximo)
	heap.Encolar(10)
	require.Equal(t, 10, heap.VerMax())
	heap.Encolar(20)
	require.Equal(t, 20, heap.VerMax())
	heap.Desencolar()
	require.Equal(t, 10, heap.VerMax())
}
