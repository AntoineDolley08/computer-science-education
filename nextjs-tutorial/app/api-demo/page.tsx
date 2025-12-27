"use client";

import Link from "next/link";
import { useState } from "react";

export default function ApiDemoPage() {
  const [apiResponse, setApiResponse] = useState<any>(null);
  const [loading, setLoading] = useState(false);
  const [userId, setUserId] = useState("1");

  const fetchHelloApi = async () => {
    setLoading(true);
    try {
      const response = await fetch("/api/hello");
      const data = await response.json();
      setApiResponse(data);
    } catch (error) {
      setApiResponse({ error: "Erreur lors de la requête" });
    } finally {
      setLoading(false);
    }
  };

  const fetchUserApi = async () => {
    setLoading(true);
    try {
      const response = await fetch(`/api/users/${userId}`);
      const data = await response.json();
      setApiResponse(data);
    } catch (error) {
      setApiResponse({ error: "Erreur lors de la requête" });
    } finally {
      setLoading(false);
    }
  };

  const postToApi = async () => {
    setLoading(true);
    try {
      const response = await fetch("/api/hello", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          name: "Utilisateur Next.js",
          message: "Test de POST request",
        }),
      });
      const data = await response.json();
      setApiResponse(data);
    } catch (error) {
      setApiResponse({ error: "Erreur lors de la requête" });
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen p-8 bg-gray-50 dark:bg-gray-900">
      <div className="max-w-4xl mx-auto">
        <Link
          href="/"
          className="text-blue-600 hover:text-blue-800 dark:text-blue-400 mb-4 inline-block"
        >
          ← Retour à l'accueil
        </Link>

        <div className="bg-white dark:bg-gray-800 rounded-lg shadow-lg p-8 mt-4">
          <h1 className="text-4xl font-bold text-gray-900 dark:text-white mb-6">
            Démonstration des API Routes
          </h1>

          <div className="bg-orange-50 dark:bg-orange-900/20 border-l-4 border-orange-500 p-4 mb-6">
            <p className="text-orange-800 dark:text-orange-300">
              <strong>Client Component</strong> : Cette page utilise fetch pour
              appeler les API Routes créées dans ce projet.
            </p>
          </div>

          <div className="space-y-6">
            <div className="border border-gray-200 dark:border-gray-700 rounded-lg p-6">
              <h2 className="text-xl font-semibold text-gray-900 dark:text-white mb-4">
                GET /api/hello
              </h2>
              <button
                onClick={fetchHelloApi}
                disabled={loading}
                className="px-6 py-3 bg-blue-600 hover:bg-blue-700 disabled:bg-gray-400 text-white font-semibold rounded-lg transition-colors"
              >
                {loading ? "Chargement..." : "Appeler l'API"}
              </button>
            </div>

            <div className="border border-gray-200 dark:border-gray-700 rounded-lg p-6">
              <h2 className="text-xl font-semibold text-gray-900 dark:text-white mb-4">
                POST /api/hello
              </h2>
              <button
                onClick={postToApi}
                disabled={loading}
                className="px-6 py-3 bg-green-600 hover:bg-green-700 disabled:bg-gray-400 text-white font-semibold rounded-lg transition-colors"
              >
                {loading ? "Envoi..." : "Envoyer des données (POST)"}
              </button>
            </div>

            <div className="border border-gray-200 dark:border-gray-700 rounded-lg p-6">
              <h2 className="text-xl font-semibold text-gray-900 dark:text-white mb-4">
                GET /api/users/[id]
              </h2>
              <div className="flex gap-4">
                <input
                  type="number"
                  value={userId}
                  onChange={(e) => setUserId(e.target.value)}
                  min="1"
                  max="3"
                  className="px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-lg dark:bg-gray-700 dark:text-white"
                  placeholder="User ID (1-3)"
                />
                <button
                  onClick={fetchUserApi}
                  disabled={loading}
                  className="px-6 py-3 bg-purple-600 hover:bg-purple-700 disabled:bg-gray-400 text-white font-semibold rounded-lg transition-colors"
                >
                  {loading ? "Chargement..." : "Récupérer l'utilisateur"}
                </button>
              </div>
            </div>

            {apiResponse && (
              <div className="border border-gray-200 dark:border-gray-700 rounded-lg p-6 bg-gray-50 dark:bg-gray-900">
                <h3 className="text-lg font-semibold text-gray-900 dark:text-white mb-3">
                  Réponse de l'API :
                </h3>
                <pre className="bg-gray-800 text-green-400 p-4 rounded overflow-x-auto">
                  {JSON.stringify(apiResponse, null, 2)}
                </pre>
              </div>
            )}
          </div>

          <div className="mt-8 space-y-4 text-gray-700 dark:text-gray-300">
            <h2 className="text-2xl font-semibold text-gray-900 dark:text-white">
              API Routes dans Next.js
            </h2>

            <p>
              Les API Routes permettent de créer des endpoints API directement dans
              votre application Next.js.
            </p>

            <ul className="list-disc list-inside space-y-2">
              <li>Créées dans le dossier <code className="bg-gray-100 dark:bg-gray-700 px-2 py-1 rounded">app/api/</code></li>
              <li>Supportent toutes les méthodes HTTP (GET, POST, PUT, DELETE, etc.)</li>
              <li>Peuvent utiliser des routes dynamiques avec [param]</li>
              <li>Accès complet à Node.js et aux bases de données</li>
              <li>Parfait pour créer des backends sans serveur séparé</li>
            </ul>

            <div className="bg-gray-100 dark:bg-gray-700 p-4 rounded mt-4">
              <p className="font-semibold mb-2">Structure des fichiers :</p>
              <pre className="text-sm">
{`app/
  api/
    hello/
      route.ts      →  /api/hello
    users/
      [id]/
        route.ts    →  /api/users/:id`}
              </pre>
            </div>

            <div className="bg-gray-100 dark:bg-gray-700 p-4 rounded mt-4">
              <p className="font-semibold mb-2">Exemple d'API Route :</p>
              <pre className="text-sm overflow-x-auto">
{`import { NextResponse } from "next/server";

export async function GET() {
  return NextResponse.json({ message: "Hello!" });
}

export async function POST(request: Request) {
  const body = await request.json();
  return NextResponse.json({ received: body });
}`}
              </pre>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
