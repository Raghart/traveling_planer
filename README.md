# Traveling

## Project Ideas
--> popular sights, best months to visit
--> where to buy sim card, how to take taxis, what time does restaurants close, which applications to download, how to take public transportation

## Future feature ideas
- SQLite para Caché
- XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX Sincronización y Cola de Tareas en segundo plano con RabbitMQ: Cuando el usuario busque un país, en lugar de hacerlo esperar en la terminal a que se hagan las 5 requests de las APIs una por una, la terminal envía un mensaje a la cola de RabbitMQ y muestra un progreso.
Un proceso de Go independiente (un "worker" o servicio secundario) recibe la tarea, hace las requests a las APIs en paralelo, y guarda el resultado en la base de datos (caché).
La terminal (tu frontend de Bubble Tea) está escuchando otra cola de RabbitMQ. En cuanto el worker termina de guardar los datos, le avisa a la terminal: "¡Oye, ya descargué todo el viaje de Bolivia!", y la terminal actualiza la pantalla al instante.
Esto es exactamente cómo funcionan las arquitecturas de microservicios modernas y distribuidas en empresas gigantes.

- El “Optimizador de Ruta de Viaje Multiciudad” (Algoritmo de Grafos / Traveling Salesperson)
Si un usuario quiere visitar varias ciudades en su viaje (ej: Tokio, Kyoto, Osaka y Hiroshima), el orden en que las visite puede hacerle ahorrar cientos de dólares y horas de viaje.

La Utilidad: El usuario ingresa una lista de ciudades que quiere visitar. La app calcula el orden óptimo para visitarlas todas haciendo la ruta más corta y eficiente posible.
Tus Programming Skills (Demostración de talento):
Implementarás el famoso problema de ciencias de la computación TSP (Traveling Salesperson Problem) usando estructuras de datos de Grafos en Go.
Escribirás un algoritmo (como Dijkstra o un algoritmo voraz de búsqueda heurística) para encontrar el camino más corto entre nodos (ciudades).
Manejarás matrices de adyacencia o listas de adyacencia en memoria para representar las distancias de transporte.
Por qué impresiona: Demuestra que dominas algoritmos complejos y teoría de grafos, resolviendo un problema real de optimización logística.
- https://en.wikipedia.org/w/api.php?action=query&list=geosearch&gscoord=35.6894|139.6917&gsradius=10000&gslimit=5&format=json