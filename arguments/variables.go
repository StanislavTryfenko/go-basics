package arguments

import "fmt"

func Variables() {

	fmt.Printf("Se vedi questo messaggio non c'é nulla di attivo in variables.go \n")

	// Principali interpolazioni
	// %v per valori normali
	// %s per stringhe
	// %d per numeri interi
	// %f per numeri con parte decimale
	// %t per valori booleani
	// %q per stringhe tra virgolette

	// fmt.Printf("Lui è %v \n", "Mario") // comune e vale  un po' per tutto
	// fmt.Printf("Lui è %s \n", "Mario") // esclusivo per stringhe
	// fmt.Printf("Lui é %q \n", "Mario") // mette le virgolette e funziona solo per stringhe

	// Tipizzazione base delle variabili in go
	// bool (true/false)
	// string ("Hello, World!")
	// int(64 di base) int8 int16 int32 int64 (numeri interi sia positivi che negativi con 8, 16, 32, 64 bit es. -128 +127, -32768 +32767, -2147483648 +2147483647, -9223372036854775808 +9223372036854775807)
	// uint(64 di base) uint8 uint16 uint32 uint64  (numeri positivi sia positivi che negativi con 8, 16, 32, 64 bit es. 0, 255, 65535, 4294967295, 18446744073709551615)
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64 ( numeri con parte decimale es. 3.14)
	// complex64 complex128 ( numeri complessi es. 1+2i)

	// In questo modo dichiaro variabili con una precisa tipizzazione statica (estremamente efficente)
	// var smsSendingLimit int
	// var costPerSms float32
	// var hasPemission bool
	// var username string

	// Le stampo in console con valori non assegnati: 0 0 false ""
	// fmt.Printf("smsSendingLimit: %d, costPerSms: %f, hasPemission: %t, username: %q \n", smsSendingLimit, costPerSms, hasPemission, username)

	// Dichiaro variabili in modalitá tipizzata dinamicamente (go riconosce questa come sting in automatico)
	// int
	// euroPerText := 1
	// fmt.Printf("Il valore di euroPerText é %T \n", euroPerText) // Uso %T per stampareil tipo della variabile

	// float64
	// euroPerText := 1.0
	// fmt.Printf("Il valore di euroPerText è %T \n", euroPerText)

	// Dichiarazione di piú variabili assieme
	// smsSendingLimit, costPerSms, hasPemission, username := 100, 0.05, true, "mario"
	// fmt.Printf("smsSendingLimit: %d, costPerSms: %f, hasPemission: %t, username: %q \n", smsSendingLimit, costPerSms, hasPemission, username)

	// Cambio tipo a una variabile
	// euroPerText := 1.2
	// euroPerTextInt := int(euroPerText) // In questo caso il valore verrá troncato

	// fmt.Printf("Il valore di euroPerTextInt è %T \n", euroPerTextInt)
	// fmt.Printf("Il valore di euroPerTextInt è %v \n", euroPerTextInt)

	// Le costanti in Go vengono registrate durante la compilazione e non durante l'esecuzione

	/* // Non potrá mai funzionare
	costPerSms, numbersOfSms := 1, 100
	const allSmsCost = costPerSms * numbersOfSms
	fmt.Printf("AllSmsCost: %d \n", allSmsCost) */

	// const costPerSms, numbersOfSms = 1, 100
	// const allSmsCost = costPerSms * numbersOfSms
	// fmt.Printf("AllSmsCost: %d \n", allSmsCost)

}
