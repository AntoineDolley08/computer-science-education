# 🚀 Démarrage Rapide

## Lancer le projet

```bash
# 1. Aller dans le dossier du projet
cd nextjs-tutorial

# 2. Lancer le serveur de développement
npm run dev

# 3. Ouvrir votre navigateur
# http://localhost:3000
```

Le serveur démarre sur **http://localhost:3000**

## 📂 Structure des Fichiers Créés

```
nextjs-tutorial/
│
├── 📄 DEMARRAGE.md         ← Vous êtes ici !
├── 📄 README-TUTORIAL.md   ← Vue d'ensemble de Next.js
├── 📄 GUIDE-RAPIDE.md      ← Référence rapide des concepts
├── 📄 EXEMPLES.md          ← Description détaillée des exemples
│
├── app/                    ← Dossier principal de l'application
│   │
│   ├── page.tsx            ← 🏠 Page d'accueil (/)
│   ├── layout.tsx          ← Layout racine
│   ├── globals.css         ← Styles globaux
│   │
│   ├── about/
│   │   └── page.tsx        ← 📄 Page à propos (/about)
│   │
│   ├── counter/
│   │   └── page.tsx        ← 🔢 Compteur interactif (/counter)
│   │
│   ├── products/
│   │   └── page.tsx        ← 🛍️ Liste de produits (/products)
│   │
│   ├── blog/
│   │   └── [slug]/
│   │       └── page.tsx    ← 📝 Articles de blog (/blog/:slug)
│   │
│   ├── form/
│   │   ├── page.tsx        ← 📋 Formulaire (/form)
│   │   └── success/
│   │       └── page.tsx    ← ✅ Page de confirmation
│   │
│   ├── api-demo/
│   │   └── page.tsx        ← 🔌 Démo des API Routes (/api-demo)
│   │
│   └── api/                ← API Routes
│       ├── hello/
│       │   └── route.ts    ← API: /api/hello
│       └── users/
│           └── [id]/
│               └── route.ts ← API: /api/users/:id
│
├── public/                 ← Fichiers statiques (images, etc.)
├── package.json            ← Dépendances du projet
└── tsconfig.json           ← Configuration TypeScript
```

## 🎯 Parcours d'Apprentissage

Suivez cet ordre pour apprendre progressivement :

### Niveau 1 : Les Bases
1. **Page d'accueil** → http://localhost:3000
   - Comprenez la structure
   - Voyez tous les exemples disponibles

2. **Page À propos** → http://localhost:3000/about
   - Routing simple
   - Server Component

### Niveau 2 : Interactivité
3. **Compteur** → http://localhost:3000/counter
   - Client Component
   - useState, événements

### Niveau 3 : Données
4. **Produits** → http://localhost:3000/products
   - Data fetching
   - Appels API

### Niveau 4 : Routes Avancées
5. **Blog** → http://localhost:3000/blog/mon-premier-article
   - Routes dynamiques
   - Paramètres d'URL

### Niveau 5 : Formulaires et API
6. **Formulaire** → http://localhost:3000/form
   - Server Actions
   - Traitement de formulaires

7. **API Demo** → http://localhost:3000/api-demo
   - Création d'API
   - Consommation d'API

## 📚 Documentation

### Fichiers à lire dans l'ordre :

1. **README-TUTORIAL.md** - Commencez ici pour la vue d'ensemble
2. **EXEMPLES.md** - Descriptions détaillées de chaque exemple
3. **GUIDE-RAPIDE.md** - Référence rapide quand vous codez

## 💻 Commandes Essentielles

```bash
# Développement
npm run dev              # Démarrer le serveur (port 3000)

# Production
npm run build            # Construire pour la production
npm run start            # Lancer en mode production

# Maintenance
npm run lint             # Vérifier le code
```

## 🎨 Modifier le Code

Tous les fichiers sont dans le dossier `app/`. Pour modifier :

1. Ouvrez un fichier `.tsx` dans votre éditeur
2. Faites vos modifications
3. Sauvegardez (Ctrl+S / Cmd+S)
4. Le navigateur se met à jour automatiquement !

### Exemple : Modifier la page d'accueil

```bash
# Ouvrir dans votre éditeur
code app/page.tsx

# ou
nano app/page.tsx
```

## 🔑 Concepts Clés à Retenir

### Server Component (par défaut)
```tsx
// app/ma-page/page.tsx
export default function MaPage() {
  return <div>Hello</div>;
}
```
- Rendu côté serveur
- Meilleur SEO
- Pas d'interactivité

### Client Component (avec "use client")
```tsx
// app/ma-page/page.tsx
"use client";

import { useState } from "react";

export default function MaPage() {
  const [count, setCount] = useState(0);
  return <button onClick={() => setCount(count + 1)}>{count}</button>;
}
```
- Rendu côté client
- Interactivité possible
- Hooks React disponibles

### Data Fetching
```tsx
export default async function MaPage() {
  const data = await fetch('https://api.example.com/data');
  const json = await data.json();
  return <div>{json.title}</div>;
}
```

### Routes Dynamiques
```
app/blog/[slug]/page.tsx  →  /blog/n-importe-quoi
```

## ⚡ Astuces

1. **Hot Reload** : Les changements apparaissent instantanément
2. **Console** : Ouvrez F12 dans le navigateur pour voir les logs
3. **Erreurs** : Lisez attentivement les messages d'erreur
4. **TypeScript** : Profitez de l'auto-complétion dans VS Code

## 🐛 Problèmes Courants

**Le serveur ne démarre pas ?**
```bash
# Vérifiez que vous êtes dans le bon dossier
cd nextjs-tutorial

# Réinstallez les dépendances
npm install

# Relancez
npm run dev
```

**Page blanche ?**
- Vérifiez la console du navigateur (F12)
- Vérifiez les erreurs dans le terminal

**Changements non visibles ?**
- Rafraîchissez le navigateur (Ctrl+R / Cmd+R)
- Vérifiez que vous avez sauvegardé le fichier

## 📖 Prochaines Étapes

Une fois que vous maîtrisez les exemples :

1. Modifiez les pages existantes
2. Créez vos propres pages
3. Ajoutez de nouvelles fonctionnalités
4. Expérimentez avec les styles
5. Déployez sur Vercel (gratuit)

## 🌐 Ressources Externes

- [Next.js Docs](https://nextjs.org/docs) - Documentation officielle
- [Next.js Learn](https://nextjs.org/learn) - Tutoriel interactif
- [React Docs](https://react.dev) - Documentation React

## ✨ Amusez-vous bien !

N'hésitez pas à expérimenter et à casser des choses. C'est comme ça qu'on apprend ! 🚀

---

**Besoin d'aide ?** Lisez les commentaires dans le code, chaque fichier est documenté.
