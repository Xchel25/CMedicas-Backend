package main

import (
	"github.com/Ilimm9/CMedicas/initializers"
	"github.com/Ilimm9/CMedicas/migrate"
	"github.com/Ilimm9/CMedicas/routes"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"  // <-- Agregamos este
	"time"                          // <-- Agregamos este también
)

func init(){
	initializers.LoadEnv()
	initializers.ConnectDB()
	migrate.Migrations()
}

func main() {
	r := gin.Default()

	// Configuración de CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	// Rutas
	routes.AdminRutas(r)

	r.Run()
}
