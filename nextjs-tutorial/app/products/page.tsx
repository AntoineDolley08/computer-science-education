import Link from "next/link";

// Interface TypeScript pour typer nos données
interface Product {
  id: number;
  title: string;
  price: number;
  description: string;
  category: string;
  image: string;
}

// SERVER COMPONENT avec Data Fetching
// async/await fonctionne directement dans les Server Components
export default async function ProductsPage() {
  // Récupération des données depuis une API publique
  // Dans Next.js 15, fetch est mis en cache par défaut
  const response = await fetch("https://fakestoreapi.com/products?limit=6");
  const products: Product[] = await response.json();

  return (
    <div className="min-h-screen p-8 bg-gray-50 dark:bg-gray-900">
      <div className="max-w-6xl mx-auto">
        <Link
          href="/"
          className="text-blue-600 hover:text-blue-800 dark:text-blue-400 mb-4 inline-block"
        >
          ← Retour à l'accueil
        </Link>

        <div className="bg-white dark:bg-gray-800 rounded-lg shadow-lg p-8 mt-4">
          <h1 className="text-4xl font-bold text-gray-900 dark:text-white mb-6">
            Liste de Produits
          </h1>

          <div className="bg-green-50 dark:bg-green-900/20 border-l-4 border-green-500 p-4 mb-6">
            <p className="text-green-800 dark:text-green-300">
              <strong>Server Component avec Data Fetching</strong> : Cette page récupère
              des données directement côté serveur avec async/await.
            </p>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-8">
            {products.map((product) => (
              <div
                key={product.id}
                className="border border-gray-200 dark:border-gray-700 rounded-lg p-4 hover:shadow-lg transition-shadow"
              >
                <img
                  src={product.image}
                  alt={product.title}
                  className="w-full h-48 object-contain mb-4"
                />
                <h3 className="font-semibold text-gray-900 dark:text-white mb-2 line-clamp-2">
                  {product.title}
                </h3>
                <p className="text-sm text-gray-600 dark:text-gray-400 mb-2 line-clamp-2">
                  {product.description}
                </p>
                <div className="flex justify-between items-center mt-4">
                  <span className="text-lg font-bold text-green-600 dark:text-green-400">
                    ${product.price}
                  </span>
                  <span className="text-sm text-gray-500 dark:text-gray-400">
                    {product.category}
                  </span>
                </div>
              </div>
            ))}
          </div>

          <div className="space-y-4 text-gray-700 dark:text-gray-300">
            <h2 className="text-2xl font-semibold text-gray-900 dark:text-white">
              Data Fetching dans Next.js
            </h2>

            <ul className="list-disc list-inside space-y-2">
              <li>
                Les Server Components peuvent utiliser <code className="bg-gray-100 dark:bg-gray-700 px-2 py-1 rounded">async/await</code> directement
              </li>
              <li>
                Les données sont récupérées côté serveur (meilleur SEO)
              </li>
              <li>
                Par défaut, fetch est mis en cache pour améliorer les performances
              </li>
              <li>
                Pas besoin de useEffect ou de gestion d'état complexe
              </li>
            </ul>

            <div className="bg-gray-100 dark:bg-gray-700 p-4 rounded mt-4">
              <p className="font-semibold mb-2">Code simplifié :</p>
              <pre className="text-sm overflow-x-auto">
{`export default async function ProductsPage() {
  // Récupération des données côté serveur
  const response = await fetch("https://api.example.com/products");
  const products = await response.json();

  return (
    <div>
      {products.map(product => (
        <div key={product.id}>{product.title}</div>
      ))}
    </div>
  );
}`}
              </pre>
            </div>

            <div className="bg-yellow-50 dark:bg-yellow-900/20 border-l-4 border-yellow-500 p-4 mt-4">
              <p className="text-yellow-800 dark:text-yellow-300">
                <strong>Options de fetch :</strong> Vous pouvez contrôler le cache avec{" "}
                <code className="bg-yellow-100 dark:bg-yellow-800 px-2 py-1 rounded">
                  {`{ cache: 'no-store' }`}
                </code>{" "}
                pour désactiver le cache ou{" "}
                <code className="bg-yellow-100 dark:bg-yellow-800 px-2 py-1 rounded">
                  {`{ next: { revalidate: 60 } }`}
                </code>{" "}
                pour revalider toutes les 60 secondes.
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
