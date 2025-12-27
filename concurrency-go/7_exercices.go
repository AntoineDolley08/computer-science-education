package main

import (
	"fmt"
	"time"
)

// Ce fichier contient des exercices à compléter
// Décommente chaque fonction dans main() pour tester ta solution

func main() {
	// Décommente les exercices un par un:

	// exercice1()
	// exercice2()
	// exercice3()
	// exercice4()
	// exercice5()
}

// ============================================================================
// EXERCICE 1: Somme concurrente
// ============================================================================
// Objectif: Calculer la somme de 1 à 1000000 en utilisant plusieurs goroutines
// Concepts: goroutines, channels, WaitGroup
//
// Instructions:
// 1. Diviser le travail en N parties (ex: 4 goroutines)
// 2. Chaque goroutine calcule la somme de sa partie
// 3. Combiner les résultats pour obtenir la somme totale
//
// Résultat attendu: 500000500000

func exercice1() {
	fmt.Println("=== Exercice 1: Somme concurrente ===")

	const total = 1000000
	const numWorkers = 4

	// TON CODE ICI
	// ...

	// Affiche le résultat
	// fmt.Printf("Somme: %d\n", sum)
}

// ============================================================================
// EXERCICE 2: URL Downloader
// ============================================================================
// Objectif: Télécharger plusieurs URLs en parallèle et mesurer le temps
// Concepts: goroutines, channels, select avec timeout
//
// Instructions:
// 1. Crée une fonction qui simule le téléchargement (time.Sleep aléatoire)
// 2. Lance plusieurs goroutines pour télécharger en parallèle
// 3. Collecte les résultats avec un timeout de 3 secondes
// 4. Affiche quelles URLs ont réussi et lesquelles ont timeout

func exercice2() {
	fmt.Println("\n=== Exercice 2: URL Downloader ===")

	urls := []string{
		"http://example.com/page1",
		"http://example.com/page2",
		"http://example.com/page3",
		"http://example.com/page4",
	}

	// TON CODE ICI
	// Simule un download avec:
	// time.Sleep(time.Duration(rand.Intn(5)) * time.Second)

	_ = urls // Utilise cette variable
}

// ============================================================================
// EXERCICE 3: Rate Limiter
// ============================================================================
// Objectif: Implémenter un rate limiter qui limite à N opérations par seconde
// Concepts: channels, time.Ticker, goroutines
//
// Instructions:
// 1. Crée un rate limiter qui permet 5 requêtes par seconde
// 2. Traite 20 requêtes en respectant cette limite
// 3. Affiche l'heure de traitement de chaque requête

func exercice3() {
	fmt.Println("\n=== Exercice 3: Rate Limiter ===")

	const requestsPerSecond = 5
	const numRequests = 20

	// TON CODE ICI
	// Indice: utilise time.Tick() ou time.NewTicker()

	fmt.Println("Toutes les requêtes traitées")
}

// ============================================================================
// EXERCICE 4: Pipeline de traitement d'images
// ============================================================================
// Objectif: Créer un pipeline pour traiter des "images"
// Concepts: pipeline pattern, channels
//
// Instructions:
// 1. Étape 1: Charge les images (1-10)
// 2. Étape 2: Applique un filtre (multiplie par 2)
// 3. Étape 3: Compresse (ajoute 100)
// 4. Étape 4: Sauvegarde (affiche le résultat)
// Chaque étape doit être une fonction séparée qui communique par channels

func exercice4() {
	fmt.Println("\n=== Exercice 4: Pipeline de traitement ===")

	// TON CODE ICI
	// Crée des fonctions: load(), filter(), compress(), save()
	// Connecte-les avec des channels

	fmt.Println("Pipeline terminé")
}

// ============================================================================
// EXERCICE 5: Cache concurrent avec expiration
// ============================================================================
// Objectif: Implémenter un cache thread-safe avec expiration
// Concepts: sync.RWMutex, time.After, goroutines
//
// Instructions:
// 1. Crée une structure Cache avec un map et un RWMutex
// 2. Implémente Set(key, value, ttl) qui stocke avec expiration
// 3. Implémente Get(key) qui retourne la valeur si pas expirée
// 4. Implémente un goroutine de nettoyage qui supprime les entrées expirées

type CacheItem struct {
	value      interface{}
	expiration time.Time
}

type Cache struct {
	// TON CODE ICI
}

func exercice5() {
	fmt.Println("\n=== Exercice 5: Cache concurrent ===")

	// TON CODE ICI
	// cache := NewCache()
	// cache.Set("key1", "value1", 2*time.Second)
	// ...

	fmt.Println("Test du cache terminé")
}

// ============================================================================
// BONUS: Worker Pool avec priorité
// ============================================================================
// Objectif: Créer un worker pool qui traite les jobs selon leur priorité
// Concepts: worker pool, select, channels multiples
//
// Instructions:
// 1. Crée deux channels: highPriority et lowPriority
// 2. Les workers traitent d'abord les jobs haute priorité
// 3. Utilise select avec priorité (un select dans un select)

type Job struct {
	ID   int
	Name string
}

func bonusExercice() {
	fmt.Println("\n=== BONUS: Worker Pool avec priorité ===")

	// TON CODE ICI

	fmt.Println("Worker pool terminé")
}

// ============================================================================
// SOLUTIONS
// ============================================================================
// Les solutions sont dans un fichier séparé: 8_solutions.go
// Essaie de résoudre les exercices par toi-même avant de regarder!
