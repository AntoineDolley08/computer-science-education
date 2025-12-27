import Link from "next/link";

// Interface pour les paramètres de la route
interface PageProps {
  params: Promise<{
    slug: string;
  }>;
}

// Données fictives pour les articles de blog
const blogPosts: Record<string, { title: string; content: string; date: string }> = {
  "mon-premier-article": {
    title: "Mon Premier Article",
    content: "Bienvenue sur mon blog Next.js ! Ceci est un exemple de route dynamique.",
    date: "2025-01-15",
  },
  "apprendre-nextjs": {
    title: "Apprendre Next.js",
    content: "Next.js est un framework React incroyable pour créer des applications web modernes.",
    date: "2025-01-16",
  },
  "server-components": {
    title: "Les Server Components",
    content: "Les Server Components révolutionnent la façon dont nous construisons des applications React.",
    date: "2025-01-17",
  },
};

// ROUTE DYNAMIQUE : [slug] devient un paramètre
export default async function BlogPostPage({ params }: PageProps) {
  // Dans Next.js 15, params est une Promise
  const { slug } = await params;

  // Récupérer l'article correspondant au slug
  const post = blogPosts[slug];

  // Si l'article n'existe pas, afficher un message
  if (!post) {
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
            <h1 className="text-4xl font-bold text-red-600 mb-4">Article non trouvé</h1>
            <p className="text-gray-700 dark:text-gray-300">
              L'article avec le slug "{slug}" n'existe pas.
            </p>
          </div>
        </div>
      </div>
    );
  }

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
          <div className="mb-6">
            <span className="text-sm text-gray-500 dark:text-gray-400">
              {post.date}
            </span>
            <h1 className="text-4xl font-bold text-gray-900 dark:text-white mt-2">
              {post.title}
            </h1>
          </div>

          <div className="bg-purple-50 dark:bg-purple-900/20 border-l-4 border-purple-500 p-4 mb-6">
            <p className="text-purple-800 dark:text-purple-300">
              <strong>Route Dynamique</strong> : Cette page utilise le paramètre{" "}
              <code className="bg-purple-100 dark:bg-purple-800 px-2 py-1 rounded">
                [slug]
              </code>{" "}
              pour afficher différents articles. Slug actuel : <strong>{slug}</strong>
            </p>
          </div>

          <div className="prose dark:prose-invert max-w-none mb-8">
            <p className="text-gray-700 dark:text-gray-300 text-lg">
              {post.content}
            </p>
          </div>

          <div className="border-t border-gray-200 dark:border-gray-700 pt-6">
            <h2 className="text-2xl font-semibold text-gray-900 dark:text-white mb-4">
              Autres articles disponibles
            </h2>
            <div className="space-y-2">
              {Object.entries(blogPosts).map(([key, article]) => (
                <Link
                  key={key}
                  href={`/blog/${key}`}
                  className={`block p-3 rounded hover:bg-gray-100 dark:hover:bg-gray-700 transition-colors ${
                    key === slug ? "bg-gray-100 dark:bg-gray-700" : ""
                  }`}
                >
                  <div className="font-semibold text-gray-900 dark:text-white">
                    {article.title}
                  </div>
                  <div className="text-sm text-gray-500 dark:text-gray-400">
                    {article.date}
                  </div>
                </Link>
              ))}
            </div>
          </div>

          <div className="mt-8 space-y-4 text-gray-700 dark:text-gray-300">
            <h2 className="text-2xl font-semibold text-gray-900 dark:text-white">
              Routes Dynamiques
            </h2>

            <p>
              Les routes dynamiques utilisent des crochets <code className="bg-gray-100 dark:bg-gray-700 px-2 py-1 rounded">[param]</code> dans le nom du dossier.
            </p>

            <div className="bg-gray-100 dark:bg-gray-700 p-4 rounded">
              <p className="font-semibold mb-2">Structure de fichiers :</p>
              <pre className="text-sm">
{`app/
  blog/
    [slug]/
      page.tsx  →  /blog/:slug

Exemples d'URLs :
/blog/mon-premier-article
/blog/apprendre-nextjs
/blog/n-importe-quel-slug`}
              </pre>
            </div>

            <div className="bg-gray-100 dark:bg-gray-700 p-4 rounded mt-4">
              <p className="font-semibold mb-2">Accéder au paramètre :</p>
              <pre className="text-sm overflow-x-auto">
{`interface PageProps {
  params: Promise<{ slug: string }>;
}

export default async function BlogPostPage({ params }: PageProps) {
  const { slug } = await params;

  // Utiliser le slug pour récupérer les données
  return <div>Article: {slug}</div>;
}`}
              </pre>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

// Génération statique : définir les slugs possibles à la compilation
// Ceci est optionnel mais améliore les performances
export async function generateStaticParams() {
  return Object.keys(blogPosts).map((slug) => ({
    slug,
  }));
}
