package main

//posible error con los paquetes, despues arreglar.
import (
	"algo/tp1/infix"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		linea := scanner.Text()
		tokens := infix.Tokenizar(linea)
		posfija := infix.ConvertirInfijaAPosfija(tokens)
		fmt.Fprintln(os.Stdout, strings.Join(posfija, " "))

	}
}
