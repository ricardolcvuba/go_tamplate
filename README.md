## Estructura del Proyecto

La plantilla sigue una estructura organizada en carpetas para separar las responsabilidades de cada componente. A continuación, se detalla la estructura principal:

```
cmd/
└── api/             # Punto de entrada principal de la aplicación

internal/
├── builders/        # Constructores para inicializar entidades o configuraciones
├── config/          # Configuración de la aplicación
├── dto/             # Objetos de transferencia de datos (Data Transfer Objects)
├── handlers/        # Controladores para manejar las solicitudes HTTP
├── middleware/      # Middlewares para la gestión de CORS, errores, etc.
├── model/           # Modelos de datos
├── repository/      # Interfaces y repositorios para acceso a datos
├── server/          # Configuración del router y servidor HTTP
├── services/        # Lógica de negocio y servicios
└── utils/           # Utilidades y helpers
    └── apperror/    # Gestión de errores personalizados

test/
├── end_to_end/      # Pruebas de extremo a extremo
├── integration/     # Pruebas de integración
└── shared/          # Helpers compartidos para pruebas
```

---

## Configuración del Entorno

### Requisitos Previos

1. **Go**: Asegúrate de tener Go instalado en tu sistema. Puedes descargarlo desde [golang.org](https://golang.org/).
2. **Docker**: Se utiliza Docker para la configuración del entorno de desarrollo. Instálalo desde [docker.com](https://www.docker.com/).
3. **Docker Compose**: Asegúrate de tener Docker Compose instalado para orquestar los servicios.

### Instalación

1. Clona este repositorio:
   ```bash
   git clone https://github.com/ricardolcvuba/go_tamplate.git
   cd go_tamplate
   ```

2. Instala las dependencias de Go:
   ```bash
   go mod tidy
   ```

3. Configura las variables de entorno necesarias. Puedes usar el archivo `.env` para definir las configuraciones específicas de tu entorno.

4. Levanta los servicios con Docker Compose:
   ```bash
   docker-compose up --build
   ```

---

## Descripción de Carpetas Clave

### `cmd/`
Contiene el punto de entrada principal de la aplicación. En este caso, el archivo `main.go` inicializa el servidor y las configuraciones necesarias.

### `internal/`
Esta carpeta contiene la lógica principal de la aplicación, dividida en subcarpetas:
- **`builders/`**: Constructores para inicializar entidades o configuraciones.
- **`config/`**: Configuración de la aplicación.
- **`dto/`**: Objetos de transferencia de datos.
- **`handlers/`**: Controladores para manejar las solicitudes HTTP.
- **`middleware/`**: Middlewares para la gestión de CORS, errores, etc.
- **`model/`**: Modelos de datos.
- **`repository/`**: Interfaces y repositorios para acceso a datos.
- **`server/`**: Configuración del router y servidor HTTP.
- **`services/`**: Lógica de negocio y servicios.
- **`utils/`**: Utilidades y helpers.

### `test/`
Contiene las pruebas de la aplicación, organizadas en:
- **`end_to_end/`**: Pruebas de extremo a extremo.
- **`integration/`**: Pruebas de integración.
- **`shared/`**: Helpers compartidos para pruebas.

---

## Uso

1. Ejecuta la aplicación:
   ```bash
   go run cmd/api/main.go
   ```

2. Accede a la API en `http://localhost:3000` (o el puerto configurado).

3. Para ejecutar las pruebas:
   ```bash
   go test ./...
   ```

---