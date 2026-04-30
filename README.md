# Epic: The Musical Api

API RESTful para gestionar sagas y canciones, construida en Go con SQLite.

## Características

- **Tecnología**: Go + SQLite (github.com/mattn/go-sqlite3)
- **Puerto**: `8088`
- **Base de datos**: Se crea automáticamente desde `data/ddl.sql`
- **Semilla**: Los datos iniciales se cargan desde `data/dml.sql` si la base está vacía

---

## Endpoints

### Sagas

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| `GET` | `/sagas` | Obtener todas las sagas |
| `GET` | `/sagas/{id}` | Obtener una saga por ID (incluye sus canciones) |
| `POST` | `/sagas` | Crear una nueva saga |
| `PUT` | `/sagas/{id}` | Actualizar una saga |
| `DELETE` | `/sagas/{id}` | Eliminar una saga |

### Songs (Canciones)

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| `GET` | `/songs` | Obtener todas las canciones |
| `GET` | `/songs/{id}` | Obtener una canción por ID |
| `POST` | `/songs` | Crear una nueva canción |
| `PUT` | `/songs/{id}` | Actualizar una canción |
| `DELETE` | `/songs/{id}` | Eliminar una canción |

---

## Estructura de Datos

### Saga

```json
{
  "id": 1,
  "title": "Título de la saga",
  "description": "Descripción de la saga",
  "image_url": "https://ejemplo.com/imagen.jpg",
  "songs": [...]
}
```

### Song

```json
{
  "id": 1,
  "saga_id": 1,
  "title": "Título de la canción",
  "description": "Descripción de la canción",
  "singers": "Cantantes",
  "image_url": "https://ejemplo.com/imagen.jpg"
}
```

---

## Docker

```bash
# Construir y ejecutar
docker-compose up --build

# Detener
docker-compose down
```

---

## Desarrollo Local

```bash
# Instalar dependencias
go mod tidy

# Ejecutar
go run main.go
```

El servidor estará disponible en: `http://localhost:8088`