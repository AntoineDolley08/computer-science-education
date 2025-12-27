import { NextResponse } from "next/server";

// Données fictives
const users = [
  { id: 1, name: "Alice", email: "alice@example.com" },
  { id: 2, name: "Bob", email: "bob@example.com" },
  { id: 3, name: "Charlie", email: "charlie@example.com" },
];

// API Route avec paramètre dynamique [id]
// GET /api/users/1, /api/users/2, etc.

interface RouteContext {
  params: Promise<{
    id: string;
  }>;
}

export async function GET(
  request: Request,
  context: RouteContext
) {
  // Récupérer le paramètre dynamique
  const { id } = await context.params;

  // Trouver l'utilisateur
  const user = users.find((u) => u.id === parseInt(id));

  if (!user) {
    return NextResponse.json(
      { error: "Utilisateur non trouvé" },
      { status: 404 }
    );
  }

  return NextResponse.json(user);
}

/*
EXEMPLES D'UTILISATION :

GET /api/users/1
→ { "id": 1, "name": "Alice", "email": "alice@example.com" }

GET /api/users/2
→ { "id": 2, "name": "Bob", "email": "bob@example.com" }

GET /api/users/999
→ { "error": "Utilisateur non trouvé" } (404)

DANS VOTRE CODE CLIENT :
fetch('/api/users/1')
  .then(res => res.json())
  .then(user => console.log(user))
*/
