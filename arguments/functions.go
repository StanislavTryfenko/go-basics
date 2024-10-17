package arguments

import (
	"fmt"
)

func Functions() {
	// fmt.Println("Se vedi questo messaggio non c'é nulla di attivo in functions.go \n")

    func sub(x int, y int) int {
        return x - y
    }

	fmt.Println(sub(10, 5))
}

