# Tutoriel Next.js Complet

Projet d'apprentissage Next.js avec des exemples pratiques et commentés.

## Démarrage Rapide

```bash
# Lancer le serveur de développement
npm run dev

# Ouvrir http://localhost:3000
```

## Ce que vous allez apprendre

Ce projet contient **8 exemples pratiques** qui couvrent tous les concepts fondamentaux de Next.js :

1. **Page d'accueil** - Introduction et navigation
2. **Routing simple** - Page À propos
3. **Client Components** - Compteur interactif
4. **Data Fetching** - Liste de produits depuis une API
5. **Routes dynamiques** - Blog avec paramètres
6. **Server Actions** - Formulaire
7. **API Routes** - Création d'endpoints
8. **Consommation d'API** - Appels API interactifs

## Documentation

Lisez les fichiers de documentation dans cet ordre :

1. **[DEMARRAGE.md](DEMARRAGE.md)** - Commencez ici
2. **[README-TUTORIAL.md](README-TUTORIAL.md)** - Vue d'ensemble de Next.js
3. **[EXEMPLES.md](EXEMPLES.md)** - Description détaillée des exemples
4. **[GUIDE-RAPIDE.md](GUIDE-RAPIDE.md)** - Référence rapide des concepts

## Structure du Projet

```
app/
├── page.tsx              # Page d'accueil (/)
├── about/                # /about
├── counter/              # /counter - Client Component
├── products/             # /products - Data Fetching
├── blog/[slug]/          # /blog/:slug - Route dynamique
├── form/                 # /form - Server Actions
├── api-demo/             # /api-demo - Test des API
└── api/                  # API Routes
    ├── hello/            # /api/hello
    └── users/[id]/       # /api/users/:id
```

## Exemples Disponibles

Visitez ces URLs après avoir lancé le serveur :

- http://localhost:3000 - Page d'accueil
- http://localhost:3000/about - À propos
- http://localhost:3000/counter - Compteur interactif
- http://localhost:3000/products - Liste de produits
- http://localhost:3000/blog/mon-premier-article - Blog
- http://localhost:3000/form - Formulaire
- http://localhost:3000/api-demo - Test des API
- http://localhost:3000/api/hello - API endpoint

## Concepts Couverts

- App Router et routing basé sur les fichiers
- Server Components vs Client Components
- Data Fetching avec async/await
- Routes dynamiques avec [param]
- Server Actions pour les formulaires
- API Routes (GET, POST)
- Navigation avec Link
- TypeScript
- Tailwind CSS

## Commandes

```bash
npm run dev      # Développement (port 3000)
npm run build    # Build production
npm run start    # Démarrer en production
npm run lint     # Vérifier le code
```

## Ressources

- [Next.js Documentation](https://nextjs.org/docs)
- [Learn Next.js](https://nextjs.org/learn)
- [React Documentation](https://react.dev)

## Notes

Chaque fichier contient des commentaires détaillés expliquant le code. N'hésitez pas à modifier et expérimenter !
