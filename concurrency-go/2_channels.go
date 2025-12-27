package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== 1. Channel basique ===")
	basicChannel()

	fmt.Println("\n=== 2. Channel bufferisé ===")
	bufferedChannel()

	fmt.Println("\n=== 3. Channel unidirectionnel ===")
	unidirectionalChannel()

	fmt.Println("\n=== 4. Fermeture de channel ===")
	channelClosure()
}

// basicChannel démontre un channel simple (non-bufferisé)
// Les channels permettent la communication entre goroutines
func basicChannel() {
	// Créer un channel qui transporte des strings
	messages := make(chan string)

	// Envoyer un message dans une goroutine
	go func() {
		time.Sleep(500 * time.Millisecond)
		messages <- "Hello from goroutine!" // Envoi dans le channel
	}()

	// Recevoir le message (bloque jusqu'à réception)
	msg := <-messages // Réception depuis le channel
	fmt.Println("Message reçu:", msg)
}

// bufferedChannel montre un channel avec buffer
// Un channel bufferisé peut contenir plusieurs valeurs sans bloquer
func bufferedChannel() {
	// Channel qui peut contenir 2 valeurs
	messages := make(chan string, 2)

	// On peut envoyer 2 messages sans bloquer
	messages <- "Message 1"
	messages <- "Message 2"
	// Un 3ème message bloquerait car le buffer est plein

	fmt.Println("Reçu:", <-messages)
	fmt.Println("Reçu:", <-messages)
}

// unidirectionalChannel montre les channels avec direction
// On peut restreindre un channel à envoi-seulement ou réception-seulement
func unidirectionalChannel() {
	messages := make(chan string)

	// Lance une goroutine qui envoie
	go sender(messages)

	// Lance une goroutine qui reçoit
	go receiver(messages)

	time.Sleep(time.Second)
}

// sender accepte seulement un channel d'envoi (chan<-)
func sender(ch chan<- string) {
	ch <- "Message depuis sender"
}

// receiver accepte seulement un channel de réception (<-chan)
func receiver(ch <-chan string) {
	msg := <-ch
	fmt.Println("Reçu:", msg)
}

// channelClosure montre comment fermer un channel
// Fermer un channel signale qu'aucune autre valeur ne sera envoyée
func channelClosure() {
	jobs := make(chan int, 5)

	// Envoyer des jobs
	go func() {
		for i := 1; i <= 3; i++ {
			jobs <- i
		}
		close(jobs) // Important: fermer le channel quand on a fini
	}()

	// Recevoir jusqu'à ce que le channel soit fermé
	for job := range jobs {
		fmt.Println("Job reçu:", job)
	}
	fmt.Println("Tous les jobs traités!")
}
