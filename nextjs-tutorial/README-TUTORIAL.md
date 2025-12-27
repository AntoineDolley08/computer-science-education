# Tutoriel Next.js - Guide Complet

Bienvenue dans ce tutoriel Next.js ! Ce guide vous apprendra les concepts fondamentaux à travers des exemples pratiques.

## Table des matières
1. [Introduction à Next.js](#introduction)
2. [Structure du projet](#structure)
3. [Routing (App Router)](#routing)
4. [Server Components vs Client Components](#components)
5. [Data Fetching](#data-fetching)
6. [Exemples pratiques](#exemples)

## Introduction

Next.js est un framework React qui offre :
- **Server-Side Rendering (SSR)** : Rendu côté serveur pour de meilleures performances
- **Static Site Generation (SSG)** : Génération de pages statiques
- **App Router** : Système de routing basé sur les dossiers
- **API Routes** : Création d'APIs facilement
- **Optimisations automatiques** : Images, fonts, etc.

## Structure du projet

```
nextjs-tutorial/
├── app/                    # Dossier principal (App Router)
│   ├── page.tsx           # Page d'accueil (/)
│   ├── layout.tsx         # Layout racine
│   └── globals.css        # Styles globaux
├── public/                # Fichiers statiques
├── node_modules/          # Dépendances
├── package.json           # Configuration npm
├── tsconfig.json          # Configuration TypeScript
└── next.config.ts         # Configuration Next.js
```

## Routing

Next.js utilise le **routing basé sur les fichiers** :

- `app/page.tsx` → `/`
- `app/about/page.tsx` → `/about`
- `app/blog/page.tsx` → `/blog`
- `app/blog/[slug]/page.tsx` → `/blog/:slug` (route dynamique)

## Server Components vs Client Components

**Server Components** (par défaut) :
- Rendus sur le serveur
- Peuvent accéder directement aux bases de données
- Ne peuvent pas utiliser les hooks React (useState, useEffect)
- Meilleurs pour les performances

**Client Components** :
- Marqués avec `"use client"` en haut du fichier
- Peuvent utiliser les hooks React
- Interactivité côté client
- Utilisés pour les formulaires, boutons, etc.

## Data Fetching

Next.js offre plusieurs façons de récupérer des données :

1. **Server Components** : fetch directement dans le composant
2. **API Routes** : Créer des endpoints API
3. **Static Generation** : À la compilation
4. **Server Actions** : Actions serveur pour les formulaires

## Exemples pratiques

J'ai créé plusieurs pages d'exemple dans ce projet :

1. **Page d'accueil** ([app/page.tsx](app/page.tsx)) - Introduction simple
2. **Page À propos** ([app/about/page.tsx](app/about/page.tsx)) - Exemple de routing
3. **Compteur interactif** ([app/counter/page.tsx](app/counter/page.tsx)) - Client Component avec useState
4. **Liste de produits** ([app/products/page.tsx](app/products/page.tsx)) - Data fetching
5. **Route dynamique** ([app/blog/[slug]/page.tsx](app/blog/[slug]/page.tsx)) - Paramètres dynamiques
6. **API Route** ([app/api/hello/route.ts](app/api/hello/route.ts)) - Création d'API

## Démarrer le projet

```bash
# Lancer le serveur de développement
npm run dev

# Ouvrir http://localhost:3000 dans votre navigateur
```

## Commandes utiles

```bash
npm run dev      # Mode développement
npm run build    # Construire pour la production
npm run start    # Démarrer en production
npm run lint     # Vérifier le code
```

## Ressources

- [Documentation officielle Next.js](https://nextjs.org/docs)
- [Apprendre Next.js](https://nextjs.org/learn)
- [Exemples Next.js](https://github.com/vercel/next.js/tree/canary/examples)
