import Link from "next/link";

export default function Home() {
  return (
    <div className="min-h-screen p-8 bg-gradient-to-br from-blue-50 to-indigo-100 dark:from-gray-900 dark:to-gray-800">
      <main className="max-w-4xl mx-auto">
        <h1 className="text-5xl font-bold text-gray-900 dark:text-white mb-4">
          Bienvenue dans le Tutoriel Next.js
        </h1>

        <p className="text-xl text-gray-700 dark:text-gray-300 mb-8">
          Explorez les exemples ci-dessous pour apprendre Next.js pas à pas.
        </p>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-6 mb-8">
          <ExampleCard
            href="/about"
            title="1. À propos"
            description="Exemple de routing simple avec une page statique"
            icon="📄"
          />

          <ExampleCard
            href="/counter"
            title="2. Compteur interactif"
            description="Client Component avec useState (interactivité)"
            icon="🔢"
          />

          <ExampleCard
            href="/products"
            title="3. Liste de produits"
            description="Server Component avec data fetching"
            icon="🛍️"
          />

          <ExampleCard
            href="/blog/mon-premier-article"
            title="4. Blog (route dynamique)"
            description="Routes dynamiques avec paramètres"
            icon="📝"
          />

          <ExampleCard
            href="/form"
            title="5. Formulaire"
            description="Gestion de formulaire avec Server Actions"
            icon="📋"
          />

          <ExampleCard
            href="/api-demo"
            title="6. API Routes Demo"
            description="Tester les API Routes de manière interactive"
            icon="🔌"
          />
        </div>

        <div className="bg-white dark:bg-gray-800 rounded-lg p-6 shadow-lg">
          <h2 className="text-2xl font-semibold text-gray-900 dark:text-white mb-3">
            Concepts clés
          </h2>
          <ul className="space-y-2 text-gray-700 dark:text-gray-300">
            <li>✅ <strong>App Router</strong> : Routing basé sur les dossiers</li>
            <li>✅ <strong>Server Components</strong> : Rendus côté serveur par défaut</li>
            <li>✅ <strong>Client Components</strong> : Avec "use client" pour l'interactivité</li>
            <li>✅ <strong>Data Fetching</strong> : Récupération de données simplifiée</li>
            <li>✅ <strong>API Routes</strong> : Créer des endpoints facilement</li>
          </ul>
        </div>
      </main>
    </div>
  );
}

function ExampleCard({ href, title, description, icon }: {
  href: string;
  title: string;
  description: string;
  icon: string;
}) {
  return (
    <Link
      href={href}
      className="block p-6 bg-white dark:bg-gray-800 rounded-lg shadow-md hover:shadow-xl transition-shadow border-2 border-transparent hover:border-indigo-500"
    >
      <div className="text-4xl mb-3">{icon}</div>
      <h3 className="text-xl font-semibold text-gray-900 dark:text-white mb-2">
        {title}
      </h3>
      <p className="text-gray-600 dark:text-gray-400">
        {description}
      </p>
    </Link>
  );
}
