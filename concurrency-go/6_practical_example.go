package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Exemple pratique: Simulateur de traitement de commandes e-commerce
// Démontre plusieurs concepts de concurrence ensemble

type Order struct {
	ID       int
	Customer string
	Amount   float64
}

type ProcessedOrder struct {
	Order
	Status      string
	ProcessedAt time.Time
}

func main() {
	fmt.Println("=== Simulateur de traitement de commandes ===\n")

	// Configuration
	const (
		numOrders   = 20
		numWorkers  = 4
		maxRetries  = 3
	)

	// Channels
	orders := make(chan Order, 10)
	processed := make(chan ProcessedOrder, 10)
	failed := make(chan Order, 5)

	// WaitGroups
	var producerWg sync.WaitGroup
	var workerWg sync.WaitGroup
	var collectorWg sync.WaitGroup

	// Statistiques (protégées par mutex)
	var stats struct {
		mu        sync.Mutex
		succeeded int
		failed    int
		totalTime time.Duration
	}

	// 1. PRODUCER: Génère des commandes
	producerWg.Add(1)
	go func() {
		defer producerWg.Done()
		defer close(orders)

		fmt.Println("📦 Génération de", numOrders, "commandes...")
		for i := 1; i <= numOrders; i++ {
			order := Order{
				ID:       i,
				Customer: fmt.Sprintf("Client-%d", i),
				Amount:   rand.Float64() * 1000,
			}
			orders <- order
			fmt.Printf("  ✓ Commande #%d créée (%.2f€)\n", order.ID, order.Amount)
			time.Sleep(50 * time.Millisecond)
		}
		fmt.Println("📦 Toutes les commandes générées!\n")
	}()

	// 2. WORKERS: Traitent les commandes en parallèle
	fmt.Printf("⚙️  Démarrage de %d workers...\n\n", numWorkers)
	for w := 1; w <= numWorkers; w++ {
		workerWg.Add(1)
		go worker(w, orders, processed, failed, &workerWg)
	}

	// 3. COLLECTOR: Collecte les résultats
	collectorWg.Add(1)
	go func() {
		defer collectorWg.Done()
		for result := range processed {
			stats.mu.Lock()
			if result.Status == "success" {
				stats.succeeded++
			}
			stats.totalTime += time.Since(result.ProcessedAt)
			stats.mu.Unlock()

			fmt.Printf("  ✅ Commande #%d [%s] traitée avec succès\n",
				result.ID, result.Status)
		}
	}()

	// 4. FAILURE HANDLER: Gère les commandes échouées
	collectorWg.Add(1)
	go func() {
		defer collectorWg.Done()
		retries := make(map[int]int)

		for order := range failed {
			retries[order.ID]++

			if retries[order.ID] <= maxRetries {
				fmt.Printf("  ⚠️  Commande #%d échouée, tentative %d/%d...\n",
					order.ID, retries[order.ID], maxRetries)
				time.Sleep(200 * time.Millisecond)
				orders <- order // Réessayer
			} else {
				stats.mu.Lock()
				stats.failed++
				stats.mu.Unlock()
				fmt.Printf("  ❌ Commande #%d abandonnée après %d tentatives\n",
					order.ID, maxRetries)
			}
		}
	}()

	// Attendre que tout soit terminé
	producerWg.Wait()    // Attendre que toutes les commandes soient générées
	workerWg.Wait()      // Attendre que tous les workers terminent
	close(processed)     // Fermer le channel de résultats
	close(failed)        // Fermer le channel d'échecs
	collectorWg.Wait()   // Attendre que les collectors terminent

	// Afficher les statistiques finales
	fmt.Println("\n" + "====================================================")
	fmt.Println("📊 STATISTIQUES FINALES")
	fmt.Println("====================================================")
	fmt.Printf("Commandes réussies: %d\n", stats.succeeded)
	fmt.Printf("Commandes échouées: %d\n", stats.failed)
	fmt.Printf("Total:              %d\n", stats.succeeded+stats.failed)
	if stats.succeeded > 0 {
		avgTime := stats.totalTime / time.Duration(stats.succeeded)
		fmt.Printf("Temps moyen:        %v\n", avgTime)
	}
	fmt.Println("====================================================")
}

// worker traite les commandes
func worker(id int, orders <-chan Order, processed chan<- ProcessedOrder,
	failed chan<- Order, wg *sync.WaitGroup) {

	defer wg.Done()

	for order := range orders {
		fmt.Printf("🔧 Worker #%d traite commande #%d...\n", id, order.ID)

		// Simuler le traitement (peut échouer aléatoirement)
		processingTime := time.Duration(100+rand.Intn(400)) * time.Millisecond
		time.Sleep(processingTime)

		// 10% de chance d'échec
		if rand.Float64() < 0.1 {
			failed <- order
			continue
		}

		// Succès
		processed <- ProcessedOrder{
			Order:       order,
			Status:      "success",
			ProcessedAt: time.Now(),
		}
	}
}

// Fonction utilitaire pour initialiser le random
func init() {
	rand.Seed(time.Now().UnixNano())
}
