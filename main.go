package main

import (
	"github.com/gin-gonic/gin"

	"github.com/MHR2102/production-distribution-tracking/controllers"
	"github.com/MHR2102/production-distribution-tracking/database"
	"github.com/MHR2102/production-distribution-tracking/middleware"
)

func main() {
	// Inisialisasi database
	database.ConnectDB()
	database.MigrateDB()

	r := gin.Default()

	// Routing untuk autentikasi
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", controllers.Register)
		authRoutes.POST("/login", controllers.Login)
	}

	// Routing yang membutuhkan autentikasi
	protected := r.Group("/")
	protected.Use(middleware.JWTAuth())

	// Routing untuk Production Batch
	productionRoutes := protected.Group("/production-batches")
	{
		productionRoutes.GET("/", controllers.GetProductionBatches)
		productionRoutes.POST("/", controllers.CreateProductionBatch)
		productionRoutes.PUT("/:id", controllers.UpdateProductionBatch)
		productionRoutes.DELETE("/:id", middleware.Authorize("admin"), controllers.DeleteProductionBatch)
	}

	// Routing untuk Distribution
	distributionRoutes := protected.Group("/distributions")
	{
		distributionRoutes.GET("/", controllers.GetDistributions)
		distributionRoutes.POST("/", controllers.CreateDistribution)
		distributionRoutes.PUT("/:id", controllers.UpdateDistribution)
	}

	// Routing untuk Reports
	reportRoutes := protected.Group("/reports")
	{
		reportRoutes.GET("/production", controllers.GetProductionReport)
		reportRoutes.GET("/distribution", controllers.GetDistributionReport)
	}

	// Menjalankan server
	r.Run(":8080")
}
