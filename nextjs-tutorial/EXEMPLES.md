# Exemples Créés dans ce Projet

Ce projet contient 6 exemples pratiques pour apprendre Next.js.

## 📋 Liste des Exemples

### 1. Page d'Accueil - Introduction
**Fichier** : [app/page.tsx](app/page.tsx)
**Route** : `/`
**Concepts** :
- Server Component
- Composants React
- Tailwind CSS
- Link component de Next.js

### 2. Page À Propos - Routing Simple
**Fichier** : [app/about/page.tsx](app/about/page.tsx)
**Route** : `/about`
**Concepts** :
- Routing basé sur les fichiers
- Server Component
- Navigation avec Link

**Ce que vous apprenez** :
- Comment créer une nouvelle route (créer un dossier + page.tsx)
- La différence entre Server et Client Components
- Comment la structure des dossiers définit les URLs

### 3. Compteur Interactif - Client Component
**Fichier** : [app/counter/page.tsx](app/counter/page.tsx)
**Route** : `/counter`
**Concepts** :
- Client Component avec `"use client"`
- useState hook
- Gestion d'événements (onClick)
- Interactivité côté client

**Ce que vous apprenez** :
- Quand et comment utiliser `"use client"`
- Utiliser les hooks React (useState)
- Créer des composants interactifs
- Différence entre rendu serveur et client

### 4. Liste de Produits - Data Fetching
**Fichier** : [app/products/page.tsx](app/products/page.tsx)
**Route** : `/products`
**Concepts** :
- Server Component asynchrone
- Data fetching avec fetch
- async/await
- Appel API externe
- TypeScript interfaces

**Ce que vous apprenez** :
- Récupérer des données depuis une API
- Utiliser async/await dans les composants
- Afficher des données dynamiques
- Options de cache de Next.js

### 5. Blog - Routes Dynamiques
**Fichier** : [app/blog/[slug]/page.tsx](app/blog/[slug]/page.tsx)
**Routes** : `/blog/mon-premier-article`, `/blog/apprendre-nextjs`, etc.
**Concepts** :
- Routes dynamiques avec [param]
- Paramètres d'URL
- generateStaticParams pour la génération statique
- Gestion de 404 (article non trouvé)

**Ce que vous apprenez** :
- Créer des routes dynamiques avec []
- Accéder aux paramètres d'URL
- Générer des pages statiques à la compilation
- Gérer les cas où les données n'existent pas

**Articles disponibles** :
- `/blog/mon-premier-article`
- `/blog/apprendre-nextjs`
- `/blog/server-components`

### 6. Formulaire - Server Actions
**Fichier** : [app/form/page.tsx](app/form/page.tsx)
**Route** : `/form`
**Concepts** :
- Server Actions avec `"use server"`
- Gestion de formulaires
- formData API
- Redirection avec redirect()

**Ce que vous apprenez** :
- Créer une Server Action
- Traiter des formulaires sans API Route
- Récupérer les données avec formData
- Rediriger après soumission

### 7. API Routes - Création d'API
**Fichiers** :
- [app/api/hello/route.ts](app/api/hello/route.ts)
- [app/api/users/[id]/route.ts](app/api/users/[id]/route.ts)

**Routes** :
- `GET /api/hello`
- `POST /api/hello`
- `GET /api/users/:id`

**Concepts** :
- API Routes
- Méthodes HTTP (GET, POST)
- NextResponse
- Routes API dynamiques

**Ce que vous apprenez** :
- Créer des endpoints API
- Gérer différentes méthodes HTTP
- Retourner du JSON
- Paramètres dynamiques dans les API

### 8. Démo API - Consommation d'API
**Fichier** : [app/api-demo/page.tsx](app/api-demo/page.tsx)
**Route** : `/api-demo`
**Concepts** :
- Client Component
- Fetch API
- useState pour gérer les réponses
- Appel aux API Routes locales

**Ce que vous apprenez** :
- Appeler vos propres API Routes
- Gérer les états de chargement
- Afficher les réponses API
- Tester différentes méthodes HTTP

## 🎯 Parcours d'Apprentissage Recommandé

1. **Commencez ici** : Page d'Accueil (/)
   - Comprenez la structure générale

2. **Routing** : Page À propos (/about)
   - Voyez comment les routes fonctionnent

3. **Interactivité** : Compteur (/counter)
   - Découvrez les Client Components

4. **Données** : Liste de Produits (/products)
   - Apprenez le data fetching

5. **Dynamique** : Blog (/blog/mon-premier-article)
   - Maîtrisez les routes dynamiques

6. **Formulaires** : Formulaire (/form)
   - Utilisez les Server Actions

7. **API** : Démo API (/api-demo)
   - Créez et consommez des API

## 🔧 Comment Modifier ces Exemples

### Ajouter une nouvelle page

1. Créez un dossier dans `app/`
2. Ajoutez un fichier `page.tsx`
3. Exportez un composant par défaut

```bash
mkdir app/ma-page
```

```tsx
// app/ma-page/page.tsx
export default function MaPage() {
  return <div>Ma nouvelle page</div>;
}
```

Votre page sera accessible à `/ma-page`

### Modifier un exemple existant

1. Ouvrez le fichier correspondant
2. Modifiez le code
3. Sauvegardez
4. Le navigateur se rafraîchit automatiquement (Hot Reload)

### Expérimenter

Essayez de :
- Changer les styles Tailwind
- Ajouter de nouveaux boutons
- Modifier les données affichées
- Créer vos propres composants
- Ajouter de nouvelles routes

## 📚 Fichiers de Documentation

- [README-TUTORIAL.md](README-TUTORIAL.md) - Vue d'ensemble du tutoriel
- [GUIDE-RAPIDE.md](GUIDE-RAPIDE.md) - Référence rapide des concepts
- [EXEMPLES.md](EXEMPLES.md) - Ce fichier, liste des exemples

## 🚀 Démarrer

```bash
# Lancer le serveur de développement
npm run dev

# Ouvrir http://localhost:3000
```

## 💡 Conseils

1. **Lisez les commentaires** : Chaque fichier contient des commentaires expliquant le code
2. **Expérimentez** : Modifiez le code et voyez ce qui se passe
3. **Console du navigateur** : Ouvrez F12 pour voir les logs et erreurs
4. **Hot Reload** : Les changements apparaissent automatiquement
5. **TypeScript** : Les erreurs de type s'affichent dans VS Code

## ❓ Questions Fréquentes

**Q : Quelle est la différence entre Server et Client Components ?**
R : Server Components sont rendus sur le serveur (par défaut), Client Components avec `"use client"` peuvent utiliser les hooks React et l'interactivité.

**Q : Quand utiliser une API Route vs une Server Action ?**
R : Server Actions pour les formulaires simples, API Routes pour des endpoints réutilisables ou appelés depuis l'extérieur.

**Q : Comment ajouter du CSS personnalisé ?**
R : Utilisez Tailwind (classes), ou ajoutez des styles dans `app/globals.css`, ou créez des modules CSS.

**Q : Les données sont-elles en temps réel ?**
R : Par défaut, Next.js cache les fetch. Utilisez `{ cache: 'no-store' }` pour des données fraîches ou `{ next: { revalidate: 60 } }` pour revalider.

**Q : Puis-je utiliser des bibliothèques npm ?**
R : Oui ! Installez avec `npm install nom-du-package` et importez-les.

## 🎓 Aller Plus Loin

Une fois ces exemples maîtrisés :
- Ajoutez une vraie base de données (Prisma + PostgreSQL)
- Implémentez l'authentification (NextAuth.js)
- Déployez sur Vercel
- Ajoutez des tests (Jest, React Testing Library)
- Utilisez des bibliothèques UI (shadcn/ui, Radix UI)
