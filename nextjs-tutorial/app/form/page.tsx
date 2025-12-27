import Link from "next/link";
import { redirect } from "next/navigation";

// SERVER ACTION : fonction qui s'exécute côté serveur
async function submitForm(formData: FormData) {
  "use server"; // Directive pour indiquer que c'est une Server Action

  // Récupérer les données du formulaire
  const name = formData.get("name") as string;
  const email = formData.get("email") as string;
  const message = formData.get("message") as string;

  // Ici, vous pourriez sauvegarder dans une base de données
  console.log("Formulaire soumis :", { name, email, message });

  // Simuler un délai de traitement
  await new Promise((resolve) => setTimeout(resolve, 1000));

  // Rediriger vers une page de confirmation
  redirect("/form/success");
}

export default function FormPage() {
  return (
    <div className="min-h-screen p-8 bg-gray-50 dark:bg-gray-900">
      <div className="max-w-2xl mx-auto">
        <Link
          href="/"
          className="text-blue-600 hover:text-blue-800 dark:text-blue-400 mb-4 inline-block"
        >
          ← Retour à l'accueil
        </Link>

        <div className="bg-white dark:bg-gray-800 rounded-lg shadow-lg p-8 mt-4">
          <h1 className="text-4xl font-bold text-gray-900 dark:text-white mb-6">
            Formulaire avec Server Actions
          </h1>

          <div className="bg-indigo-50 dark:bg-indigo-900/20 border-l-4 border-indigo-500 p-4 mb-6">
            <p className="text-indigo-800 dark:text-indigo-300">
              <strong>Server Action</strong> : Ce formulaire utilise une Server Action
              pour traiter les données côté serveur sans API Route.
            </p>
          </div>

          <form action={submitForm} className="space-y-6">
            <div>
              <label
                htmlFor="name"
                className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2"
              >
                Nom
              </label>
              <input
                type="text"
                id="name"
                name="name"
                required
                className="w-full px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent dark:bg-gray-700 dark:text-white"
                placeholder="Votre nom"
              />
            </div>

            <div>
              <label
                htmlFor="email"
                className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2"
              >
                Email
              </label>
              <input
                type="email"
                id="email"
                name="email"
                required
                className="w-full px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent dark:bg-gray-700 dark:text-white"
                placeholder="votre@email.com"
              />
            </div>

            <div>
              <label
                htmlFor="message"
                className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-2"
              >
                Message
              </label>
              <textarea
                id="message"
                name="message"
                required
                rows={4}
                className="w-full px-4 py-2 border border-gray-300 dark:border-gray-600 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent dark:bg-gray-700 dark:text-white"
                placeholder="Votre message"
              />
            </div>

            <button
              type="submit"
              className="w-full px-6 py-3 bg-blue-600 hover:bg-blue-700 text-white font-semibold rounded-lg transition-colors"
            >
              Envoyer
            </button>
          </form>

          <div className="mt-8 space-y-4 text-gray-700 dark:text-gray-300">
            <h2 className="text-2xl font-semibold text-gray-900 dark:text-white">
              Server Actions
            </h2>

            <p>
              Les Server Actions sont des fonctions asynchrones qui s'exécutent côté serveur.
            </p>

            <ul className="list-disc list-inside space-y-2">
              <li>Marquées avec la directive <code className="bg-gray-100 dark:bg-gray-700 px-2 py-1 rounded">"use server"</code></li>
              <li>Peuvent être utilisées directement dans les formulaires</li>
              <li>Pas besoin de créer une API Route séparée</li>
              <li>Accès direct aux bases de données et aux ressources serveur</li>
              <li>Gestion automatique de la sérialisation</li>
            </ul>

            <div className="bg-gray-100 dark:bg-gray-700 p-4 rounded mt-4">
              <p className="font-semibold mb-2">Code simplifié :</p>
              <pre className="text-sm overflow-x-auto">
{`async function submitForm(formData: FormData) {
  "use server"; // Server Action

  const name = formData.get("name");

  // Traitement côté serveur
  await saveToDatabase({ name });

  redirect("/success");
}

export default function Page() {
  return (
    <form action={submitForm}>
      <input name="name" />
      <button type="submit">Envoyer</button>
    </form>
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
