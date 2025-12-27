import Link from "next/link";

// Ceci est un SERVER COMPONENT (par défaut dans Next.js)
// Il est rendu côté serveur, ce qui améliore les performances et le SEO

export default function AboutPage() {
  return (
    <div className="min-h-screen p-8 bg-gray-50 dark:bg-gray-900">
      <div className="max-w-3xl mx-auto">
        <Link
          href="/"
          className="text-blue-600 hover:text-blue-800 dark:text-blue-400 mb-4 inline-block"
        >
          ← Retour à l'accueil
        </Link>

        <div className="bg-white dark:bg-gray-800 rounded-lg shadow-lg p-8 mt-4">
          <h1 className="text-4xl font-bold text-gray-900 dark:text-white mb-6">
            À propos de Next.js
          </h1>

          <div className="space-y-4 text-gray-700 dark:text-gray-300">
            <p>
              Cette page est un exemple de <strong>Server Component</strong> dans Next.js.
            </p>

            <h2 className="text-2xl font-semibold text-gray-900 dark:text-white mt-6 mb-3">
              Qu'est-ce qu'un Server Component ?
            </h2>

            <ul className="list-disc list-inside space-y-2">
              <li>Rendu côté serveur par défaut</li>
              <li>Meilleure performance initiale</li>
              <li>Meilleur pour le SEO</li>
              <li>Peut accéder directement aux bases de données</li>
              <li>Ne peut pas utiliser les hooks React comme useState ou useEffect</li>
            </ul>

            <h2 className="text-2xl font-semibold text-gray-900 dark:text-white mt-6 mb-3">
              Routing dans Next.js
            </h2>

            <p>
              Cette page existe à <code className="bg-gray-100 dark:bg-gray-700 px-2 py-1 rounded">
                app/about/page.tsx
              </code>
            </p>

            <p>
              Next.js utilise le <strong>file-based routing</strong> :
            </p>

            <div className="bg-gray-100 dark:bg-gray-700 p-4 rounded mt-2 font-mono text-sm">
              app/page.tsx → /<br />
              app/about/page.tsx → /about<br />
              app/blog/page.tsx → /blog<br />
              app/blog/[slug]/page.tsx → /blog/:slug
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
