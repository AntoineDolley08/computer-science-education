package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== 1. Worker Pool Pattern ===")
	workerPool()

	fmt.Println("\n=== 2. Pipeline Pattern ===")
	pipeline()

	fmt.Println("\n=== 3. Fan-Out / Fan-In Pattern ===")
	fanOutFanIn()

	fmt.Println("\n=== 4. Context-like Cancellation ===")
	cancellation()
}

// workerPool montre le pattern Worker Pool
// Un nombre fixe de workers traite des jobs depuis un channel
func workerPool() {
	const numWorkers = 3
	const numJobs = 10

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Créer le worker pool
	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Envoyer les jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Important: fermer pour signaler qu'il n'y a plus de jobs

	// Attendre que tous les workers terminent
	wg.Wait()
	close(results)

	// Afficher les résultats
	fmt.Println("\nRésultats:")
	for result := range results {
		fmt.Printf("  %d\n", result)
	}
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d traite job %d\n", id, job)
		time.Sleep(100 * time.Millisecond) // Simule du travail
		results <- job * 2
	}
}

// pipeline montre le pattern Pipeline
// Chaque étape transforme les données et les passe à l'étape suivante
func pipeline() {
	// Étape 1: Génère des nombres
	gen := func(nums ...int) <-chan int {
		out := make(chan int)
		go func() {
			for _, n := range nums {
				out <- n
			}
			close(out)
		}()
		return out
	}

	// Étape 2: Carré des nombres
	sq := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			for n := range in {
				out <- n * n
			}
			close(out)
		}()
		return out
	}

	// Étape 3: Additionne 10
	addTen := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			for n := range in {
				out <- n + 10
			}
			close(out)
		}()
		return out
	}

	// Construire et exécuter le pipeline
	// gen -> sq -> addTen
	numbers := gen(1, 2, 3, 4, 5)
	squared := sq(numbers)
	final := addTen(squared)

	// Afficher les résultats
	for result := range final {
		fmt.Printf("Résultat: %d\n", result)
	}
}

// fanOutFanIn montre le pattern Fan-Out/Fan-In
// Fan-Out: Distribuer le travail à plusieurs goroutines
// Fan-In: Combiner les résultats de plusieurs goroutines
func fanOutFanIn() {
	// Source de données
	source := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			source <- i
		}
		close(source)
	}()

	// Fan-Out: Créer plusieurs workers
	numWorkers := 3
	workers := make([]<-chan int, numWorkers)
	for i := 0; i < numWorkers; i++ {
		workers[i] = fanOutWorker(source)
	}

	// Fan-In: Combiner les résultats
	results := fanIn(workers...)

	// Afficher les résultats
	for result := range results {
		fmt.Printf("Résultat: %d\n", result)
	}
}

func fanOutWorker(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			time.Sleep(50 * time.Millisecond) // Simule du travail
			out <- n * n
		}
		close(out)
	}()
	return out
}

func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	// Pour chaque channel d'entrée, créer une goroutine qui copie vers out
	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for n := range c {
				out <- n
			}
		}(ch)
	}

	// Fermer out quand tous les channels sont terminés
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// cancellation montre comment implémenter une annulation
// Pattern utilisé pour arrêter des goroutines proprement
func cancellation() {
	// Channel pour signaler l'annulation
	done := make(chan bool)

	// Goroutine qui fait un travail continu
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Goroutine: Reçu signal d'arrêt, je termine!")
				return
			default:
				fmt.Println("Goroutine: Travail en cours...")
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	// Laisser la goroutine travailler un peu
	time.Sleep(800 * time.Millisecond)

	// Envoyer le signal d'arrêt
	fmt.Println("Main: Envoi du signal d'arrêt")
	done <- true

	// Attendre un peu pour voir le message de fin
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Main: Programme terminé")
}
