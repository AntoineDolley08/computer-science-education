package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// SOLUTIONS DES EXERCICES
// N'ouvre ce fichier qu'après avoir essayé de résoudre les exercices!

func main() {
	// Décommente pour tester les solutions:
	// solution1()
	// solution2()
	// solution3()
	// solution4()
	// solution5()
}

// ============================================================================
// SOLUTION 1: Somme concurrente
// ============================================================================

func solution1() {
	fmt.Println("=== Solution 1: Somme concurrente ===")

	const total = 1000000
	const numWorkers = 4
	const chunkSize = total / numWorkers

	results := make(chan int, numWorkers)
	var wg sync.WaitGroup

	// Lancer les workers
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		start := i*chunkSize + 1
		end := (i + 1) * chunkSize
		if i == numWorkers-1 {
			end = total // S'assurer de couvrir tout
		}

		go func(start, end int) {
			defer wg.Done()
			sum := 0
			for j := start; j <= end; j++ {
				sum += j
			}
			results <- sum
		}(start, end)
	}

	// Attendre et fermer
	go func() {
		wg.Wait()
		close(results)
	}()

	// Combiner les résultats
	totalSum := 0
	for partialSum := range results {
		totalSum += partialSum
	}

	fmt.Printf("Somme: %d\n", totalSum)
	fmt.Printf("Vérification: %d (formule: n*(n+1)/2)\n", total*(total+1)/2)
}

// ============================================================================
// SOLUTION 2: URL Downloader
// ============================================================================

func solution2() {
	fmt.Println("\n=== Solution 2: URL Downloader ===")

	urls := []string{
		"http://example.com/page1",
		"http://example.com/page2",
		"http://example.com/page3",
		"http://example.com/page4",
	}

	type Result struct {
		url      string
		duration time.Duration
		success  bool
	}

	results := make(chan Result, len(urls))

	// Lancer les downloads
	for _, url := range urls {
		go func(u string) {
			start := time.Now()

			// Simuler le download (1-5 secondes)
			downloadTime := time.Duration(rand.Intn(5)+1) * time.Second

			select {
			case <-time.After(downloadTime):
				results <- Result{u, time.Since(start), true}
			case <-time.After(3 * time.Second):
				results <- Result{u, time.Since(start), false}
			}
		}(url)
	}

	// Collecter les résultats
	for i := 0; i < len(urls); i++ {
		result := <-results
		if result.success {
			fmt.Printf("✓ %s téléchargé en %v\n", result.url, result.duration)
		} else {
			fmt.Printf("✗ %s timeout après %v\n", result.url, result.duration)
		}
	}
}

// ============================================================================
// SOLUTION 3: Rate Limiter
// ============================================================================

func solution3() {
	fmt.Println("\n=== Solution 3: Rate Limiter ===")

	const requestsPerSecond = 5
	const numRequests = 20

	// Créer un ticker qui tick N fois par seconde
	limiter := time.Tick(time.Second / requestsPerSecond)

	for i := 1; i <= numRequests; i++ {
		<-limiter // Attendre le prochain tick
		fmt.Printf("%s Requête #%d traitée\n",
			time.Now().Format("15:04:05.000"), i)
	}

	fmt.Println("Toutes les requêtes traitées")
}

// ============================================================================
// SOLUTION 4: Pipeline de traitement d'images
// ============================================================================

func solution4() {
	fmt.Println("\n=== Solution 4: Pipeline de traitement ===")

	// Étape 1: Charger les images
	load := func() <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for i := 1; i <= 10; i++ {
				fmt.Printf("Chargement image %d\n", i)
				out <- i
				time.Sleep(100 * time.Millisecond)
			}
		}()
		return out
	}

	// Étape 2: Appliquer un filtre
	filter := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for img := range in {
				filtered := img * 2
				fmt.Printf("Filtre appliqué: %d -> %d\n", img, filtered)
				out <- filtered
				time.Sleep(50 * time.Millisecond)
			}
		}()
		return out
	}

	// Étape 3: Compresser
	compress := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			for img := range in {
				compressed := img + 100
				fmt.Printf("Compression: %d -> %d\n", img, compressed)
				out <- compressed
				time.Sleep(50 * time.Millisecond)
			}
		}()
		return out
	}

	// Étape 4: Sauvegarder
	save := func(in <-chan int) {
		for img := range in {
			fmt.Printf("✓ Image sauvegardée: %d\n", img)
			time.Sleep(50 * time.Millisecond)
		}
	}

	// Construire et exécuter le pipeline
	images := load()
	filtered := filter(images)
	compressed := compress(filtered)
	save(compressed)

	fmt.Println("Pipeline terminé")
}

// ============================================================================
// SOLUTION 5: Cache concurrent avec expiration
// ============================================================================

type CacheItem5 struct {
	value      interface{}
	expiration time.Time
}

type Cache5 struct {
	mu    sync.RWMutex
	items map[string]CacheItem5
	stop  chan bool
}

func NewCache() *Cache5 {
	c := &Cache5{
		items: make(map[string]CacheItem5),
		stop:  make(chan bool),
	}

	// Lancer le nettoyage automatique
	go c.cleanupLoop()

	return c
}

func (c *Cache5) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = CacheItem5{
		value:      value,
		expiration: time.Now().Add(ttl),
	}
}

func (c *Cache5) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.items[key]
	if !found {
		return nil, false
	}

	// Vérifier l'expiration
	if time.Now().After(item.expiration) {
		return nil, false
	}

	return item.value, true
}

func (c *Cache5) cleanupLoop() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.cleanup()
		case <-c.stop:
			return
		}
	}
}

func (c *Cache5) cleanup() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	for key, item := range c.items {
		if now.After(item.expiration) {
			delete(c.items, key)
			fmt.Printf("🗑️  Clé '%s' expirée et supprimée\n", key)
		}
	}
}

func (c *Cache5) Stop() {
	c.stop <- true
}

func solution5() {
	fmt.Println("\n=== Solution 5: Cache concurrent ===")

	cache := NewCache()
	defer cache.Stop()

	// Test du cache
	cache.Set("user:1", "Alice", 2*time.Second)
	cache.Set("user:2", "Bob", 4*time.Second)
	cache.Set("user:3", "Charlie", 1*time.Second)

	// Lire immédiatement
	if val, ok := cache.Get("user:1"); ok {
		fmt.Printf("Trouvé: user:1 = %v\n", val)
	}

	// Attendre que user:3 expire
	time.Sleep(1500 * time.Millisecond)
	if _, ok := cache.Get("user:3"); !ok {
		fmt.Println("user:3 a expiré ✓")
	}

	// user:1 et user:2 devraient encore être là
	if val, ok := cache.Get("user:1"); ok {
		fmt.Printf("Encore valide: user:1 = %v\n", val)
	}

	// Attendre le nettoyage automatique
	time.Sleep(2 * time.Second)

	fmt.Println("Test du cache terminé")
}

// ============================================================================
// BONUS: Worker Pool avec priorité
// ============================================================================

type Job struct {
	ID       int
	Name     string
	Priority string
}

func bonusSolution() {
	fmt.Println("\n=== BONUS: Worker Pool avec priorité ===")

	highPriority := make(chan Job, 10)
	lowPriority := make(chan Job, 10)
	var wg sync.WaitGroup

	// Worker qui préfère les jobs haute priorité
	worker := func(id int) {
		defer wg.Done()
		for {
			select {
			case job, ok := <-highPriority:
				if !ok {
					return
				}
				fmt.Printf("Worker %d traite job HAUTE priorité: %s\n", id, job.Name)
				time.Sleep(100 * time.Millisecond)

			default:
				select {
				case job, ok := <-highPriority:
					if !ok {
						return
					}
					fmt.Printf("Worker %d traite job HAUTE priorité: %s\n", id, job.Name)
					time.Sleep(100 * time.Millisecond)

				case job, ok := <-lowPriority:
					if !ok {
						return
					}
					fmt.Printf("Worker %d traite job basse priorité: %s\n", id, job.Name)
					time.Sleep(100 * time.Millisecond)
				}
			}
		}
	}

	// Lancer les workers
	const numWorkers = 3
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i)
	}

	// Envoyer des jobs
	go func() {
		for i := 1; i <= 5; i++ {
			lowPriority <- Job{i, fmt.Sprintf("Low-%d", i), "low"}
		}
		for i := 1; i <= 3; i++ {
			highPriority <- Job{i, fmt.Sprintf("High-%d", i), "high"}
		}
		for i := 6; i <= 10; i++ {
			lowPriority <- Job{i, fmt.Sprintf("Low-%d", i), "low"}
		}

		close(highPriority)
		close(lowPriority)
	}()

	wg.Wait()
	fmt.Println("Worker pool terminé")
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
