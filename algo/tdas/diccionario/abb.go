package diccionario

import (
	TDAPila "tdas/pila"
)

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      func(K, K) int
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{cmp: funcion_cmp}
}

//-----------------FUNCIONES AUXILIARES--------------//

//buscar de forma tal que devuelva un **nodoAbb[K, V](doble puntero)//

func (ab *abb[K, V]) buscar(clave K) **nodoAbb[K, V] {
	nodo := &ab.raiz
	for *nodo != nil {
		switch cmp := ab.cmp(clave, (*nodo).clave); {
		case cmp < 0:
			nodo = &(*nodo).izquierdo
		case cmp > 0:
			nodo = &(*nodo).derecho
		default:
			return nodo
		}
	}
	return nodo
}

func (nodo *nodoAbb[K, V]) cantHijos() int {
	if nodo.izquierdo == nil && nodo.derecho == nil {
		return 0
	}
	if nodo.izquierdo != nil && nodo.derecho != nil {
		return 2
	}
	return 1
}

/*FUNCION AUXILIAR*/
/*
Esto podría abstraerse a una función que dado un nodo te devuelva a uno de sus hijos (al izquierdo si no es nil, y si no al derecho). Luego este caso pasa a ser igual que el anterior:

Si el nodo a eliminar tiene 0 hijos: se pisa con nil (que, como no tiene hijos, podría considerarse como su hijo derecho).
Si el nodo a eliminar tiene 1 hijo: se pisa con el hijo que tenga (que puede ser izquierdo o derecho).

*/
func (n *nodoAbb[K, V]) unicoHijo() *nodoAbb[K, V] {
	if n.izquierdo != nil {
		return n.izquierdo
	}
	return n.derecho
}

/*----------------------------------------------------------------------*/

func (ab *abb[K, V]) Cantidad() int {
	return ab.cantidad
}

func (ab *abb[K, V]) Pertenece(clave K) bool {
	nodoPtr := ab.buscar(clave)
	return *nodoPtr != nil
}

func (ab *abb[K, V]) Obtener(clave K) V {
	nodo := ab.buscar(clave)
	if *nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	return (*nodo).dato
}

/*simplificación de la lógica por desarrollo de doble puntero en buscar*/
func (ab *abb[K, V]) Guardar(clave K, dato V) {
	nodoPtr := ab.buscar(clave)
	if *nodoPtr != nil {
		(*nodoPtr).dato = dato
		return
	}
	*nodoPtr = &nodoAbb[K, V]{clave: clave, dato: dato}
	ab.cantidad++
}

//--------------------------------------------------------------------------------/

func (ab *abb[K, V]) Borrar(clave K) V {
	nodoPtr := ab.buscar(clave)
	nodo := *nodoPtr
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}

	datoBorrado := nodo.dato

	switch nodo.cantHijos() {
	case 0:
		*nodoPtr = nil

	case 1:
		*nodoPtr = nodo.unicoHijo()

	case 2:
		sucesor := nodo.derecho
		for sucesor.izquierdo != nil {
			sucesor = sucesor.izquierdo
		}
		claveReemplazo := sucesor.clave
		//borramos recursivamente el reemplazo (llamando a Borrar)
		valorReemplazo := ab.Borrar(claveReemplazo)

		nodo.clave = claveReemplazo
		nodo.dato = valorReemplazo

		ab.cantidad++
	}

	ab.cantidad--
	return datoBorrado
}

//--------------------------------------------------------------------------------/

// CORRECCIÓN SUGERIDA
func (ab *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	ab._IterarRango(ab.raiz, desde, hasta, visitar)
}

func (ab *abb[K, V]) _IterarRango(nodo *nodoAbb[K, V], desde *K, hasta *K, visitar func(K, V) bool) bool {
	if nodo == nil {
		return false
	}

	// visitar subarbol izquierdo si nodo.clave > desde
	if desde == nil || ab.cmp(nodo.clave, *desde) > 0 {
		if ab._IterarRango(nodo.izquierdo, desde, hasta, visitar) {
			return true
		}
	}

	// visitar el nodo si está dentro del rango
	if (desde == nil || ab.cmp(nodo.clave, *desde) >= 0) && (hasta == nil || ab.cmp(nodo.clave, *hasta) <= 0) {
		if !visitar(nodo.clave, nodo.dato) {
			return true
		}
	}

	// visitar subarbol derecho si nodo.clave < hasta
	if hasta == nil || ab.cmp(nodo.clave, *hasta) < 0 {
		if ab._IterarRango(nodo.derecho, desde, hasta, visitar) {
			return true
		}
	}

	return false
}

//--------------------------------------------------------------------------------/

type iterDiccionarioOrdenado[K comparable, V any] struct {
	elementos TDAPila.Pila[*nodoAbb[K, V]]
	desde     *K
	hasta     *K
	cmp       func(K, K) int // Corrección sugerida: enfocar el uso de la función de comparación
}

func (ab *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := &iterDiccionarioOrdenado[K, V]{
		elementos: TDAPila.CrearPilaDinamica[*nodoAbb[K, V]](),
		desde:     desde,
		hasta:     hasta,
		cmp:       ab.cmp,
	}
	apilarDesde(iter, ab.raiz)
	return iter
}

func apilarDesde[K comparable, V any](iter *iterDiccionarioOrdenado[K, V], nodo *nodoAbb[K, V]) {

	if nodo == nil {
		return
	}
	if iter.desde != nil && iter.cmp(nodo.clave, *iter.desde) < 0 {
		apilarDesde(iter, nodo.derecho)
		return
	}
	if iter.hasta != nil && iter.cmp(nodo.clave, *iter.hasta) > 0 {
		apilarDesde(iter, nodo.izquierdo)
		return
	}
	iter.elementos.Apilar(nodo)
	apilarDesde(iter, nodo.izquierdo)
}

func (iter *iterDiccionarioOrdenado[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	nodo := iter.elementos.Desapilar()
	apilarDesde(iter, nodo.derecho)
}

func (iter *iterDiccionarioOrdenado[K, V]) HaySiguiente() bool {
	return !iter.elementos.EstaVacia()
}

func (iter *iterDiccionarioOrdenado[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	nodo := iter.elementos.VerTope()
	return nodo.clave, nodo.dato
}

func (ab *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	ab.IterarRango(nil, nil, visitar)
}

func (ab *abb[K, V]) Iterador() IterDiccionario[K, V] {
	iter := ab.IteradorRango(nil, nil)
	return iter
}
