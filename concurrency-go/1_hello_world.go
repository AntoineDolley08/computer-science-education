package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Exécution séquentielle ===")
	simple()

	fmt.Println("\n=== Exécution concurrente ===")
	simpleConc()
}

// simple exécute 3 tâches séquentiellement
// Temps total: ~3 secondes
func simple() {
	start := time.Now()

	fmt.Println(time.Now().Format("15:04:05"), "Tâche 0")
	time.Sleep(time.Second)
	fmt.Println(time.Now().Format("15:04:05"), "Tâche 1")
	time.Sleep(time.Second)
	fmt.Println(time.Now().Format("15:04:05"), "Tâche 2")
	time.Sleep(time.Second)

	fmt.Printf("Terminé en %v\n", time.Since(start))
}

// simpleConc exécute 3 tâches en parallèle avec des goroutines
// Temps total: ~1 seconde (toutes les tâches s'exécutent en même temps)
func simpleConc() {
	start := time.Now()

	// Lance 3 goroutines en parallèle
	for i := 0; i < 3; i++ {
		go func(index int) {
			fmt.Println(time.Now().Format("15:04:05"), "Tâche", index)
		}(i) // Important: on passe i en paramètre pour éviter les problèmes de closure
	}

	// On attend que les goroutines se terminent
	time.Sleep(time.Second)
	fmt.Printf("Terminé en %v\n", time.Since(start))
}