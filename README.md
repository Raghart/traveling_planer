# Traveling

## Future feature ideas
- SQLite para Caché
- Sincronización y Cola de Tareas en segundo plano con RabbitMQ: Cuando el usuario busque un país, en lugar de hacerlo esperar en la terminal a que se hagan las 5 requests de las APIs una por una, la terminal envía un mensaje a la cola de RabbitMQ y muestra un progreso.
Un proceso de Go independiente (un "worker" o servicio secundario) recibe la tarea, hace las requests a las APIs en paralelo, y guarda el resultado en la base de datos (caché).
La terminal (tu frontend de Bubble Tea) está escuchando otra cola de RabbitMQ. En cuanto el worker termina de guardar los datos, le avisa a la terminal: "¡Oye, ya descargué todo el viaje de Bolivia!", y la terminal actualiza la pantalla al instante.
Esto es exactamente cómo funcionan las arquitecturas de microservicios modernas y distribuidas en empresas gigantes.
- El “Descargador de Paquetes de Viaje con Barra de Progreso” (Caché asíncrona): Cuando el usuario selecciona un país, hay mucha información que descargar. Puedes simular que estás "descargando el paquete offline del país" para que la app funcione sin internet.
La Feature: El usuario presiona "Descargar paquete offline de España". La pantalla muestra una barra de progreso que avanza a medida que se descarga y procesa cada sección (Clima, Festividades, Wiki, Fotos, Moneda).
El Reto Técnico (Concurrencia y UI asíncrona):
Lanzarás tareas asíncronas para cada sección. Cada tarea terminada envía un mensaje de progreso (ej: "Clima descargado (+20%)").
El modelo de Bubble Tea capta estos mensajes individuales en segundo plano y actualiza una barra de progreso animada (bubbles/progress) de forma fluida.
Una vez que la barra llega al 100%, todos los datos se guardan en tu base de datos local y la app pasa a modo "Offline" para ese país.