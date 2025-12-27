# Apprendre la Concurrence en Go

Ce dossier contient des exemples progressifs pour apprendre la concurrence en Go.

## 📚 Ordre d'apprentissage

### 1. [1_hello_world.go](1_hello_world.go) - Introduction aux Goroutines
**Concepts:**
- Différence entre exécution séquentielle et concurrente
- Mot-clé `go` pour lancer une goroutine
- Problème de closure dans les boucles

**Exécuter:**
```bash
go run 1_hello_world.go
```

**Points clés:**
- Une goroutine est une fonction qui s'exécute de manière concurrente
- `go func()` lance une nouvelle goroutine
- Il faut passer les variables en paramètre pour éviter les problèmes de closure

---

### 2. [2_channels.go](2_channels.go) - Communication entre Goroutines
**Concepts:**
- Channels (canaux) pour communiquer entre goroutines
- Channels bufferisés vs non-bufferisés
- Channels unidirectionnels (`chan<-` et `<-chan`)
- Fermeture de channels avec `close()`

**Exécuter:**
```bash
go run 2_channels.go
```

**Points clés:**
- `make(chan Type)` crée un channel non-bufferisé (bloquant)
- `make(chan Type, size)` crée un channel bufferisé
- `ch <- valeur` envoie dans le channel
- `valeur := <-ch` reçoit depuis le channel
- `for val := range ch` itère jusqu'à ce que le channel soit fermé

---

### 3. [3_select.go](3_select.go) - Multiplexage de Channels
**Concepts:**
- Statement `select` pour attendre sur plusieurs channels
- Timeouts avec `time.After()`
- Case `default` pour opérations non-bloquantes
- Pattern multiplexing

**Exécuter:**
```bash
go run 3_select.go
```

**Points clés:**
- `select` permet d'attendre sur plusieurs channels simultanément
- Le premier channel prêt est choisi
- `default` s'exécute si aucun channel n'est prêt
- Très utile pour implémenter des timeouts

---

### 4. [4_waitgroup_mutex.go](4_waitgroup_mutex.go) - Synchronisation
**Concepts:**
- `sync.WaitGroup` pour attendre plusieurs goroutines
- Race conditions et pourquoi elles sont dangereuses
- `sync.Mutex` pour protéger les accès concurrents
- `sync.RWMutex` pour optimiser lecture/écriture

**Exécuter:**
```bash
go run 4_waitgroup_mutex.go

# Pour détecter les race conditions:
go run -race 4_waitgroup_mutex.go
```

**Points clés:**
- `wg.Add(1)` avant de lancer une goroutine
- `defer wg.Done()` dans chaque goroutine
- `wg.Wait()` pour attendre que toutes terminent
- Toujours utiliser `Mutex` pour accéder à des données partagées
- `RWMutex` permet plusieurs lecteurs simultanés

---

### 5. [5_patterns.go](5_patterns.go) - Patterns Avancés
**Concepts:**
- Worker Pool: pool de workers qui traitent des jobs
- Pipeline: chaîne de transformations
- Fan-Out/Fan-In: distribuer et combiner le travail
- Cancellation: arrêter proprement des goroutines

**Exécuter:**
```bash
go run 5_patterns.go
```

**Points clés:**
- Worker Pool limite le nombre de goroutines concurrentes
- Pipeline permet de composer des transformations
- Fan-Out distribue le travail, Fan-In combine les résultats
- Utiliser un channel `done` pour signaler l'arrêt

---

## 🎯 Exercices pratiques

### Exercice 1: URL Fetcher
Crée un programme qui télécharge plusieurs URLs en parallèle et affiche le temps de chargement de chacune.

### Exercice 2: Prime Number Finder
Crée un worker pool qui trouve tous les nombres premiers jusqu'à 10000.

### Exercice 3: Rate Limiter
Implémente un rate limiter qui limite les requêtes à N par seconde.

### Exercice 4: Cache Concurrent
Crée un cache thread-safe avec expiration automatique.

---

## 🔧 Commandes utiles

```bash
# Exécuter avec détection de race conditions
go run -race fichier.go

# Formatter le code
go fmt fichier.go

# Build un exécutable
go build fichier.go

# Voir les goroutines en cours (ajoute GODEBUG)
GODEBUG=schedtrace=1000 go run fichier.go
```

---

## 📖 Ressources supplémentaires

- [Go by Example: Goroutines](https://gobyexample.com/goroutines)
- [Go Concurrency Patterns (Google I/O 2012)](https://www.youtube.com/watch?v=f6kdp27TYZs)
- [Advanced Go Concurrency Patterns](https://www.youtube.com/watch?v=QDDwwePbDtw)
- [Effective Go - Concurrency](https://go.dev/doc/effective_go#concurrency)

---

## ⚠️ Règles d'or de la concurrence en Go

1. **"Do not communicate by sharing memory; instead, share memory by communicating."**
   - Utilise des channels plutôt que des variables partagées quand possible

2. **Toujours fermer les channels du côté de l'envoyeur**
   - Jamais du côté du receveur

3. **Utiliser `defer wg.Done()` immédiatement après `wg.Add(1)`**
   - Évite les oublis

4. **Un channel fermé peut toujours être lu**
   - Il retourne la valeur zéro du type

5. **Envoyer sur un channel fermé cause une panique**
   - Vérifier avant d'envoyer si nécessaire

6. **Toujours protéger les accès à des données partagées**
   - Utilise `Mutex` ou des channels

7. **Attention aux deadlocks**
   - Si toutes les goroutines sont bloquées, c'est un deadlock

---

Bon apprentissage! N'hésite pas à modifier et expérimenter avec ces exemples. 🚀
