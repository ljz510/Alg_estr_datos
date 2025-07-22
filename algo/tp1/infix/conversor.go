package infix

import (
	"strconv"
	"strings"
	"tdas/pila"
)

var precedencia = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
	"^": 3,
}

func esAsociativoDerecha(op string) bool {
	return op == "^"
}

func esOperador(s string) bool {
	_, ok := precedencia[s]
	return ok
}

func esNumero(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

// simplifica la separacion de las cadenas en tokens, teniendo en cuenta los operadores y los parentesis
func Tokenizar(linea string) []string {
	var tokens []string
	actual := ""

	for _, r := range linea {
		switch {
		case r == ' ':
			if actual != "" {
				tokens = append(tokens, actual)
				actual = ""
			}
		case strings.ContainsRune("()+-*/^", r):
			if actual != "" {
				tokens = append(tokens, actual)
				actual = ""
			}
			tokens = append(tokens, string(r))
		default:
			actual += string(r)
		}
	}
	if actual != "" {
		tokens = append(tokens, actual)
	}
	return tokens
}

// basado en el algoritmo shunting yard
func ConvertirInfijaAPosfija(tokens []string) []string {
	var salida []string
	p := pila.CrearPilaDinamica[string]()

	for _, token := range tokens {
		switch {
		case esNumero(token):
			salida = append(salida, token)

		case esOperador(token):
			for !p.EstaVacia() {
				tope := p.VerTope()
				if !esOperador(tope) {
					break
				}
				if (precedencia[token] < precedencia[tope]) ||
					(precedencia[token] == precedencia[tope] && !esAsociativoDerecha(token)) {
					salida = append(salida, p.Desapilar())
				} else {
					break
				}
			}
			p.Apilar(token)

		case token == "(":
			p.Apilar(token)

		case token == ")":
			for !p.EstaVacia() && p.VerTope() != "(" {
				salida = append(salida, p.Desapilar())
			}
			if !p.EstaVacia() && p.VerTope() == "(" {
				p.Desapilar()
			}
		}
	}

	for !p.EstaVacia() {
		salida = append(salida, p.Desapilar())
	}

	return salida
}
