import { NextResponse } from "next/server";

// API ROUTE : endpoint accessible via HTTP
// GET http://localhost:3000/api/hello

export async function GET() {
  // Retourner une réponse JSON
  return NextResponse.json({
    message: "Bonjour depuis l'API Next.js !",
    timestamp: new Date().toISOString(),
    data: {
      framework: "Next.js",
      version: "15",
      feature: "API Routes",
    },
  });
}

// Vous pouvez aussi gérer d'autres méthodes HTTP
export async function POST(request: Request) {
  // Récupérer le body de la requête
  const body = await request.json();

  // Traiter les données
  return NextResponse.json({
    message: "Données reçues !",
    receivedData: body,
  });
}

/*
COMMENT UTILISER CETTE API ROUTE :

1. GET Request :
   - Ouvrez http://localhost:3000/api/hello dans votre navigateur
   - Ou utilisez fetch :
     fetch('/api/hello')
       .then(res => res.json())
       .then(data => console.log(data))

2. POST Request :
   fetch('/api/hello', {
     method: 'POST',
     headers: { 'Content-Type': 'application/json' },
     body: JSON.stringify({ name: 'John' })
   })

STRUCTURE DES API ROUTES :
app/
  api/
    hello/
      route.ts  →  /api/hello
    users/
      route.ts  →  /api/users
    users/
      [id]/
        route.ts  →  /api/users/:id

MÉTHODES HTTP DISPONIBLES :
- GET : récupérer des données
- POST : créer des données
- PUT : mettre à jour des données
- DELETE : supprimer des données
- PATCH : modifier partiellement des données
*/
