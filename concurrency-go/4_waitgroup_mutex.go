package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== 1. WaitGroup basique ===")
	basicWaitGroup()

	fmt.Println("\n=== 2. Problème de race condition ===")
	raceCondition()

	fmt.Println("\n=== 3. Mutex pour protéger les données ===")
	mutexProtection()

	fmt.Println("\n=== 4. RWMutex (lecture/écriture) ===")
	rwMutexExample()
}

// basicWaitGroup montre comment attendre que toutes les goroutines terminent
// WaitGroup est meilleur que time.Sleep car on attend exactement le bon temps
func basicWaitGroup() {
	var wg sync.WaitGroup

	// Lancer 5 workers
	for i := 1; i <= 5; i++ {
		wg.Add(1) // Incrémenter le compteur

		go func(id int) {
			defer wg.Done() // Décrémenter quand terminé (toujours utiliser defer!)

			fmt.Printf("Worker %d démarre\n", id)
			time.Sleep(time.Duration(id*100) * time.Millisecond)
			fmt.Printf("Worker %d termine\n", id)
		}(i)
	}

	wg.Wait() // Bloquer jusqu'à ce que le compteur soit à 0
	fmt.Println("Tous les workers ont terminé!")
}

// raceCondition montre un problème de concurrence
// Plusieurs goroutines accèdent à la même variable sans protection
func raceCondition() {
	counter := 0
	var wg sync.WaitGroup

	// Lancer 1000 goroutines qui incrémentent le compteur
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++ // DANGER: race condition!
		}()
	}

	wg.Wait()
	fmt.Printf("Compteur final: %d (devrait être 1000, mais souvent moins!)\n", counter)
	fmt.Println("Exécute avec 'go run -race 4_waitgroup_mutex.go' pour détecter les races")
}

// mutexProtection montre comment utiliser Mutex pour éviter les races
// Mutex = Mutual Exclusion (une seule goroutine à la fois)
func mutexProtection() {
	counter := 0
	var mu sync.Mutex // Mutex pour protéger counter
	var wg sync.WaitGroup

	// Lancer 1000 goroutines qui incrémentent le compteur
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			mu.Lock()   // Verrouiller avant d'accéder à counter
			counter++   // Section critique
			mu.Unlock() // Déverrouiller après

			// Alternative avec defer (plus sûr):
			// mu.Lock()
			// defer mu.Unlock()
			// counter++
		}()
	}

	wg.Wait()
	fmt.Printf("Compteur final: %d (toujours 1000 maintenant!)\n", counter)
}

// rwMutexExample montre RWMutex pour optimiser lecture/écriture
// Permet plusieurs lecteurs simultanés, mais un seul écrivain
func rwMutexExample() {
	type SafeCounter struct {
		mu    sync.RWMutex
		value map[string]int
	}

	counter := SafeCounter{value: make(map[string]int)}
	var wg sync.WaitGroup

	// Fonction pour écrire
	write := func(key string, val int) {
		counter.mu.Lock() // Lock exclusif pour écriture
		defer counter.mu.Unlock()
		counter.value[key] = val
		fmt.Printf("Écrit: %s = %d\n", key, val)
		time.Sleep(100 * time.Millisecond)
	}

	// Fonction pour lire
	read := func(key string) {
		counter.mu.RLock() // RLock pour lecture (plusieurs peuvent lire en même temps)
		defer counter.mu.RUnlock()
		val := counter.value[key]
		fmt.Printf("Lu: %s = %d\n", key, val)
	}

	// Écrire quelques valeurs
	wg.Add(3)
	go func() { defer wg.Done(); write("a", 1) }()
	go func() { defer wg.Done(); write("b", 2) }()
	go func() { defer wg.Done(); write("c", 3) }()
	wg.Wait()

	// Lire en parallèle (tous peuvent lire simultanément)
	wg.Add(6)
	go func() { defer wg.Done(); read("a") }()
	go func() { defer wg.Done(); read("b") }()
	go func() { defer wg.Done(); read("c") }()
	go func() { defer wg.Done(); read("a") }()
	go func() { defer wg.Done(); read("b") }()
	go func() { defer wg.Done(); read("c") }()
	wg.Wait()

	fmt.Println("Toutes les opérations terminées!")
}
