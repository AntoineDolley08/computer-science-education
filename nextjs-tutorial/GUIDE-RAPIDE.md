# Guide Rapide Next.js

## Démarrage

```bash
# Se déplacer dans le dossier du projet
cd nextjs-tutorial

# Installer les dépendances (déjà fait)
npm install

# Lancer le serveur de développement
npm run dev

# Ouvrir http://localhost:3000 dans votre navigateur
```

## Structure du Projet

```
nextjs-tutorial/
├── app/                      # Dossier principal (App Router)
│   ├── page.tsx             # Page d'accueil (/)
│   ├── layout.tsx           # Layout racine
│   ├── globals.css          # Styles globaux
│   │
│   ├── about/               # Route /about
│   │   └── page.tsx
│   │
│   ├── counter/             # Route /counter
│   │   └── page.tsx
│   │
│   ├── products/            # Route /products
│   │   └── page.tsx
│   │
│   ├── blog/
│   │   └── [slug]/          # Route dynamique /blog/:slug
│   │       └── page.tsx
│   │
│   ├── form/                # Route /form
│   │   ├── page.tsx
│   │   └── success/
│   │       └── page.tsx
│   │
│   └── api/                 # API Routes
│       ├── hello/
│       │   └── route.ts     # GET/POST /api/hello
│       └── users/
│           └── [id]/
│               └── route.ts # GET /api/users/:id
│
├── public/                  # Fichiers statiques
├── package.json
└── next.config.ts
```

## Concepts Fondamentaux

### 1. Routing Basé sur les Fichiers

```
app/page.tsx                 →  /
app/about/page.tsx          →  /about
app/blog/page.tsx           →  /blog
app/blog/[slug]/page.tsx    →  /blog/:slug
```

### 2. Server Components vs Client Components

**Server Components (par défaut)** :
```tsx
// Pas besoin de directive "use client"
export default function Page() {
  // Rendu côté serveur
  return <div>Server Component</div>;
}
```

**Client Components** :
```tsx
"use client"; // Directive requise

import { useState } from "react";

export default function Page() {
  const [count, setCount] = useState(0);
  // Peut utiliser les hooks React
  return <button onClick={() => setCount(count + 1)}>{count}</button>;
}
```

### 3. Data Fetching

**Dans un Server Component** :
```tsx
export default async function Page() {
  // fetch directement dans le composant
  const data = await fetch('https://api.example.com/data');
  const json = await data.json();

  return <div>{json.title}</div>;
}
```

**Options de cache** :
```tsx
// Désactiver le cache
fetch('url', { cache: 'no-store' })

// Revalider toutes les 60 secondes
fetch('url', { next: { revalidate: 60 } })
```

### 4. Routes Dynamiques

**Structure** :
```
app/blog/[slug]/page.tsx
```

**Code** :
```tsx
interface PageProps {
  params: Promise<{ slug: string }>;
}

export default async function Page({ params }: PageProps) {
  const { slug } = await params;
  return <div>Article: {slug}</div>;
}

// Optionnel : générer les routes statiquement
export async function generateStaticParams() {
  return [
    { slug: 'article-1' },
    { slug: 'article-2' },
  ];
}
```

### 5. Server Actions

**Formulaire avec Server Action** :
```tsx
async function submitForm(formData: FormData) {
  "use server"; // Directive Server Action

  const name = formData.get("name");
  // Traitement côté serveur
  await saveToDatabase({ name });
}

export default function Page() {
  return (
    <form action={submitForm}>
      <input name="name" />
      <button type="submit">Envoyer</button>
    </form>
  );
}
```

### 6. API Routes

**Créer une API Route** :
```tsx
// app/api/hello/route.ts
import { NextResponse } from "next/server";

export async function GET() {
  return NextResponse.json({ message: "Hello!" });
}

export async function POST(request: Request) {
  const body = await request.json();
  return NextResponse.json({ received: body });
}
```

**Avec paramètres dynamiques** :
```tsx
// app/api/users/[id]/route.ts
interface RouteContext {
  params: Promise<{ id: string }>;
}

export async function GET(request: Request, context: RouteContext) {
  const { id } = await context.params;
  return NextResponse.json({ userId: id });
}
```

### 7. Layouts

**Layout racine** (app/layout.tsx) :
```tsx
export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="fr">
      <body>{children}</body>
    </html>
  );
}
```

**Layouts imbriqués** :
```
app/
  layout.tsx              # Layout racine
  blog/
    layout.tsx            # Layout pour /blog/*
    page.tsx              # /blog
    [slug]/page.tsx       # /blog/:slug
```

### 8. Navigation

**Liens entre pages** :
```tsx
import Link from "next/link";

<Link href="/about">À propos</Link>
<Link href="/blog/mon-article">Mon article</Link>
```

**Navigation programmatique** :
```tsx
"use client";
import { useRouter } from "next/navigation";

export default function Page() {
  const router = useRouter();

  return (
    <button onClick={() => router.push('/about')}>
      Aller à la page À propos
    </button>
  );
}
```

### 9. Images Optimisées

```tsx
import Image from "next/image";

<Image
  src="/photo.jpg"
  alt="Description"
  width={500}
  height={300}
  priority  // Pour les images au-dessus de la ligne de flottaison
/>
```

### 10. Métadonnées (SEO)

```tsx
import { Metadata } from "next";

export const metadata: Metadata = {
  title: "Titre de la page",
  description: "Description pour le SEO",
};

export default function Page() {
  return <div>Contenu</div>;
}
```

## Commandes Utiles

```bash
# Développement
npm run dev              # Démarrer en mode développement (port 3000)

# Production
npm run build            # Construire pour la production
npm run start            # Démarrer en mode production

# Autres
npm run lint             # Vérifier le code avec ESLint
```

## Bonnes Pratiques

1. **Utilisez Server Components par défaut**
   - Plus rapides, meilleurs pour le SEO
   - Utilisez Client Components uniquement si nécessaire (interactivité, hooks)

2. **Data Fetching**
   - Préférez fetcher dans les Server Components
   - Utilisez les options de cache appropriées

3. **Organisation du code**
   - Un dossier = une route
   - `page.tsx` pour le contenu de la page
   - `layout.tsx` pour les layouts partagés
   - `loading.tsx` pour les états de chargement
   - `error.tsx` pour la gestion des erreurs

4. **Performance**
   - Utilisez `next/image` pour les images
   - Lazy loading avec `dynamic()`
   - Code splitting automatique

## Ressources

- Documentation officielle : https://nextjs.org/docs
- Tutoriel interactif : https://nextjs.org/learn
- Exemples : https://github.com/vercel/next.js/tree/canary/examples

## Prochaines Étapes

1. Explorez les exemples dans ce projet
2. Modifiez les pages existantes
3. Créez vos propres routes
4. Expérimentez avec les API Routes
5. Déployez sur Vercel (gratuit) : https://vercel.com
