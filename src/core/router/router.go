package router

import (
	productInfra "API_HEXAGONAL_RECU/src/products/infrastructure"
	userInfra "API_HEXAGONAL_RECU/src/users/infrastructure"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

type Router struct {
	engine *gin.Engine
	db     *gorm.DB
}

func NewRouter(db *gorm.DB) *Router {
	engine := gin.Default()

	// Configuración CORS
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	return &Router{
		engine: engine,
		db:     db,
	}
}

func (r *Router) SetupRoutes() {
	api := r.engine.Group("/api")

	// Setup user routes
	userInfra.SetupUserRoutes(api, r.db)

	// Setup product routes
	productInfra.SetupProductRoutes(api, r.db)
}

func (r *Router) GetEngine() *gin.Engine {
	return r.engine
}
