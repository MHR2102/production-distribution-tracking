package main

import (
	"github.com/gin-gonic/gin"

	"github.com/MHR2102/production-distribution-tracking/controllers"
	"github.com/MHR2102/production-distribution-tracking/database"
	"github.com/MHR2102/production-distribution-tracking/middleware"
)

func main() {
	// Initialize the database
	database.ConnectDB()
	database.MigrateDB()

	r := gin.Default()

	// Auth routes
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/register", controllers.Register)
		authRoutes.POST("/login", controllers.Login)
	}

	// Protected routes
	protected := r.Group("/")
	protected.Use(middleware.JWTAuth())

	// Production Batch routes
	productionRoutes := protected.Group("/production-batches")
	{
		productionRoutes.GET("/", controllers.GetProductionBatches)
		productionRoutes.POST("/", controllers.CreateProductionBatch)
		productionRoutes.PUT("/:id", controllers.UpdateProductionBatch)
		productionRoutes.DELETE("/:id", middleware.Authorize("admin"), controllers.DeleteProductionBatch)
	}

	// Distribution routes
	distributionRoutes := protected.Group("/distributions")
	{
		distributionRoutes.GET("/", controllers.GetDistributions)
		distributionRoutes.POST("/", controllers.CreateDistribution)
		distributionRoutes.PUT("/:id", controllers.UpdateDistribution)
	}

	// Report routes
	reportRoutes := protected.Group("/reports")
	{
		reportRoutes.GET("/production", controllers.GetProductionReport)
		reportRoutes.GET("/distribution", controllers.GetDistributionReport)
	}

	// Run the server
	r.Run(":8080")
}
