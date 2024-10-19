package arguments

import (
	"errors"
	"fmt"
)

func Functions() {

	//fmt.Println("Se vedi questo messaggio non c'è nulla di attivo in functions.go")

	fmt.Println(sub(10, 5))

	x := 10
	x = increment(x)
	fmt.Println(x)

	// Go non permette di avere vareabele non utilizzate in una funzione
	// Uso l'_ per evitare di usare la variabile in una funzione e prendnrn solo cio che mi interessa
	name, _ := getNames()
	fmt.Println("Ciao", name)

	x, y := getCoords()
	fmt.Println(x, y)

	fmt.Println(divide(10, 0))

}

// In go non si puo' dichiarare una funzione dentro una funzione
/* func sub(x int, y int) int {
	return x - y
} */

// sugar version (puoi ragruppare i parametri per tipo cosi da essere meno verboso)
func sub(x, y int) int {
	return x - y
}

/* func increment(x int) {
	// senza return usandola in un altra funzione non si avrá il valore incrementato perché go funziona a value e non a pointer
	x++
} */

func increment(x int) int {
	// col return invece si avrá la modifica
	return x + 1
}

func getNames() (string, string) {
	return "Mario", "Rossi"
}

// Queste due fenzioni sono uguali nell'atto pratico,
/* func getCoords() (int, int) {
	var x, y int
	return x, y
} */
/* func getCoords() (x, y int) {
	return
} */
// Peró secondo molti (me compreso) per funzioni piú lunghe conviene sempre scrivere il return per una miglior leggibilitá
func getCoords() (x, y int) {
	return x, y
}

func divide(divided int, divisor int) (int, error) {

	// Guard clause per gestire l'errore
	if divisor == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	// Nil rappresenta un valore nullo,
	// É spesso utilizzato per rappresentare l'assenza di un valore, come ad esempio nel caso di un errore
	return divided / divisor, nil
}
