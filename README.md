# ⚕️ CMedicas Backend — API REST Clínica Médica

Backend de un sistema de gestión para clínica médica, desarrollado en **Go** con arquitectura REST. Gestiona citas médicas, pacientes, médicos, horarios, observaciones y notificaciones. Se conecta con el frontend desarrollado en Angular.

> **Repositorio frontend:** [frontend-clinica-bootstrap](https://github.com/Xchel25/frontend-clinica-bootstrap)

---

## 🧰 Tecnologías utilizadas

- **Go 1.23** — lenguaje principal
- **Gin** — framework web HTTP
- **GORM** — ORM para Go
- **PostgreSQL** — base de datos relacional
- **JWT** — autenticación (dgrijalva/jwt-go)
- **bcrypt** — hash de contraseñas
- **Alembic** — migraciones de base de datos
- **godotenv** — variables de entorno
- **gin-contrib/cors** — configuración CORS

---

## 📁 Estructura del proyecto

```
CMedicas-Backend/
├── controllers/         # Lógica de cada recurso
│   ├── CitaController.go
│   ├── HorariosController.go
│   ├── MedicoController.go
│   ├── NotificacionController.go
│   ├── ObservacionController.go
│   ├── PersonaController.go
│   └── UsuarioController.go
├── models/              # Modelos de base de datos (GORM)
│   ├── cita.go
│   ├── horario.go
│   ├── medico.go
│   ├── notificacion.go
│   ├── observacion.go
│   ├── persona.go
│   └── usuario.go
├── routes/              # Definición de rutas
├── middlewares/         # Middleware JWT
├── initializers/        # Inicialización de DB
├── migrate/             # Migraciones
├── main.go
└── go.mod
```

---

## ⚙️ Instalación y uso

### Requisitos previos
- Go 1.23+
- PostgreSQL

### Pasos

```bash
git clone https://github.com/Xchel25/CMedicas-Backend.git
cd CMedicas-Backend
cp .env.example .env
# Edita .env con tus credenciales de PostgreSQL
go mod tidy
go run main.go
```

El servidor corre por defecto en `http://localhost:8080`

---

## 🔗 Endpoints principales

| Método | Ruta | Descripción |
|--------|------|-------------|
| POST | `/auth/login` | Iniciar sesión |
| GET | `/usuarios` | Listar usuarios |
| GET | `/citas` | Listar citas |
| POST | `/citas` | Crear cita |
| GET | `/medicos` | Listar médicos |
| GET | `/horarios` | Consultar horarios |
| GET | `/notificaciones` | Ver notificaciones |

---

## ✨ Características principales

- API REST con autenticación JWT
- Gestión completa de citas médicas (agendar, ver, cancelar)
- Registro de médicos, pacientes y horarios
- Sistema de notificaciones
- Observaciones clínicas por paciente
- CORS configurado para integración con frontend Angular

---

## 👤 Autor

**Xchel25** — [github.com/Xchel25](https://github.com/Xchel25)
