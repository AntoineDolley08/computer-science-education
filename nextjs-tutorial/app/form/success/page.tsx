import Link from "next/link";

export default function SuccessPage() {
  return (
    <div className="min-h-screen p-8 bg-gray-50 dark:bg-gray-900 flex items-center justify-center">
      <div className="max-w-md w-full">
        <div className="bg-white dark:bg-gray-800 rounded-lg shadow-lg p-8 text-center">
          <div className="w-16 h-16 bg-green-100 dark:bg-green-900 rounded-full flex items-center justify-center mx-auto mb-4">
            <svg
              className="w-8 h-8 text-green-600 dark:text-green-400"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M5 13l4 4L19 7"
              />
            </svg>
          </div>

          <h1 className="text-3xl font-bold text-gray-900 dark:text-white mb-4">
            Formulaire envoyé avec succès !
          </h1>

          <p className="text-gray-700 dark:text-gray-300 mb-6">
            Votre message a été traité par la Server Action.
          </p>

          <div className="space-y-3">
            <Link
              href="/form"
              className="block px-6 py-3 bg-blue-600 hover:bg-blue-700 text-white font-semibold rounded-lg transition-colors"
            >
              Retour au formulaire
            </Link>

            <Link
              href="/"
              className="block px-6 py-3 bg-gray-200 hover:bg-gray-300 dark:bg-gray-700 dark:hover:bg-gray-600 text-gray-900 dark:text-white font-semibold rounded-lg transition-colors"
            >
              Retour à l'accueil
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
}
