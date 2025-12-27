package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== 1. Select basique ===")
	basicSelect()

	fmt.Println("\n=== 2. Timeout avec select ===")
	selectTimeout()

	fmt.Println("\n=== 3. Default case (non-bloquant) ===")
	selectDefault()

	fmt.Println("\n=== 4. Pattern multiplexing ===")
	selectMultiplex()
}

// basicSelect montre comment select permet d'attendre sur plusieurs channels
// Select bloque jusqu'à ce qu'un des channels soit prêt
func basicSelect() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message du channel 1"
	}()

	go func() {
		time.Sleep(500 * time.Millisecond)
		ch2 <- "Message du channel 2"
	}()

	// Select attend le premier channel prêt
	select {
	case msg1 := <-ch1:
		fmt.Println("Reçu:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Reçu:", msg2)
	}
}

// selectTimeout montre comment implémenter un timeout
// Très utile pour éviter de bloquer indéfiniment
func selectTimeout() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second) // Trop lent!
		ch <- "Message tardif"
	}()

	select {
	case msg := <-ch:
		fmt.Println("Reçu:", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout! Aucun message reçu en 1 seconde")
	}
}

// selectDefault utilise le case default pour non-bloquant
// Le case default s'exécute immédiatement si aucun channel n'est prêt
func selectDefault() {
	ch := make(chan string)

	// Pas de goroutine qui envoie, donc le channel ne sera jamais prêt

	select {
	case msg := <-ch:
		fmt.Println("Reçu:", msg)
	default:
		fmt.Println("Aucun message disponible")
	}
}

// selectMultiplex montre comment traiter plusieurs sources
// Pattern courant: boucle + select pour gérer plusieurs channels
func selectMultiplex() {
	tick := time.Tick(200 * time.Millisecond)  // Émet toutes les 200ms
	boom := time.After(1 * time.Second)         // Émet après 1 seconde

	for {
		select {
		case <-tick:
			fmt.Println("Tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return // Sortir de la fonction
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
