"use client"; // Cette directive transforme ce composant en CLIENT COMPONENT

import Link from "next/link";
import { useState } from "react";

// CLIENT COMPONENT : peut utiliser les hooks React et l'interactivité
export default function CounterPage() {
  // useState est un hook React qui permet de gérer l'état local
  const [count, setCount] = useState(0);

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
            Compteur Interactif
          </h1>

          <div className="bg-blue-50 dark:bg-blue-900/20 border-l-4 border-blue-500 p-4 mb-6">
            <p className="text-blue-800 dark:text-blue-300">
              <strong>Client Component</strong> : Cette page utilise{" "}
              <code className="bg-blue-100 dark:bg-blue-800 px-2 py-1 rounded">
                "use client"
              </code>{" "}
              pour activer l'interactivité côté client.
            </p>
          </div>

          <div className="text-center space-y-6">
            <div className="text-6xl font-bold text-gray-900 dark:text-white">
              {count}
            </div>

            <div className="flex gap-4 justify-center">
              <button
                onClick={() => setCount(count - 1)}
                className="px-6 py-3 bg-red-500 hover:bg-red-600 text-white font-semibold rounded-lg transition-colors"
              >
                - Décrémenter
              </button>

              <button
                onClick={() => setCount(0)}
                className="px-6 py-3 bg-gray-500 hover:bg-gray-600 text-white font-semibold rounded-lg transition-colors"
              >
                Reset
              </button>

              <button
                onClick={() => setCount(count + 1)}
                className="px-6 py-3 bg-green-500 hover:bg-green-600 text-white font-semibold rounded-lg transition-colors"
              >
                + Incrémenter
              </button>
            </div>
          </div>

          <div className="mt-8 space-y-4 text-gray-700 dark:text-gray-300">
            <h2 className="text-2xl font-semibold text-gray-900 dark:text-white">
              Pourquoi un Client Component ?
            </h2>

            <ul className="list-disc list-inside space-y-2">
              <li>Utilise <code className="bg-gray-100 dark:bg-gray-700 px-2 py-1 rounded">useState</code> pour gérer l'état</li>
              <li>Gère les événements onClick</li>
              <li>Nécessite l'interactivité côté client</li>
              <li>Le code JavaScript est envoyé au navigateur</li>
            </ul>

            <div className="bg-gray-100 dark:bg-gray-700 p-4 rounded mt-4">
              <p className="font-semibold mb-2">Code de cette page :</p>
              <pre className="text-sm overflow-x-auto">
{`"use client"; // Active le mode client

import { useState } from "react";

export default function CounterPage() {
  const [count, setCount] = useState(0);

  return (
    <button onClick={() => setCount(count + 1)}>
      Compteur: {count}
    </button>
  );
}`}
              </pre>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
