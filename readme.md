# MyAPI

Este proyecto es una API REST básica construida con Go, Gin y GORM. La API permite gestionar usuarios y utiliza SQLite como base de datos.

## Requisitos previos

- [Go](https://golang.org/) instalado (versión 1.20 o superior).
- [Docker](https://www.docker.com/) y [Docker Compose](https://docs.docker.com/compose/) instalados.

## Configuración

1. Clona este repositorio:
   ```bash
   git clone <URL_DEL_REPOSITORIO>
   cd myapi
   ```

2. Construye y ejecuta el proyecto con Docker Compose:
   ```bash
   docker-compose up --build
   ```

3. La API estará disponible en `http://localhost:8080`.

## Endpoints

### 1. **GET /users**

Obtiene una lista de usuarios con soporte para paginación.

#### Parámetros de consulta (query parameters):
- `page` (opcional): Número de página (por defecto: `1`).
- `limit` (opcional): Número de usuarios por página (por defecto: `10`).

#### Respuesta:
- **Código 200**: Devuelve un JSON con los usuarios y la información de paginación.
  ```json
  {
    "data": [
      {
        "id": 1,
        "name": "John Doe",
        "email": "john.doe@example.com"
      }
    ],
    "total": 100,
    "page": 1,
    "limit": 10
  }
  ```
- **Código 400**: Devuelve un error si los parámetros son inválidos.
  ```json
  {
    "error": "Invalid page parameter"
  }
  ```

#### Ejemplo de uso:
```bash
curl "http://localhost:8080/users?page=2&limit=5"
```

---

### 2. **POST /users**

Crea un nuevo usuario.

#### Cuerpo de la solicitud (JSON):
```json
{
  "name": "John Doe",
  "email": "john.doe@example.com"
}
```

#### Respuesta:
- **Código 201**: Devuelve el usuario creado.
  ```json
  {
    "id": 1,
    "name": "John Doe",
    "email": "john.doe@example.com"
  }
  ```
- **Código 400**: Devuelve un error si el cuerpo de la solicitud es inválido.
  ```json
  {
    "error": "Invalid request body"
  }
  ```

#### Ejemplo de uso:
```bash
curl -X POST "http://localhost:8080/users" \
-H "Content-Type: application/json" \
-d '{
  "name": "John Doe",
  "email": "john.doe@example.com"
}'
```

---

## Pruebas

Para ejecutar las pruebas, utiliza el siguiente comando:
```bash
go test ./...
```

## Contribución

Si deseas contribuir al proyecto, sigue estos pasos:

1. Haz un fork del repositorio.
2. Crea una rama para tu funcionalidad (`git checkout -b feature/nueva-funcionalidad`).
3. Realiza tus cambios y haz un commit (`git commit -m "Descripción de los cambios"`).
4. Haz push a tu rama (`git push origin feature/nueva-funcionalidad`).
5. Abre un Pull Request.

## Licencia

Este proyecto está bajo la licencia [MIT](https://opensource.org/licenses/MIT).