package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/MHR2102/production-distribution-tracking/database"
	"github.com/MHR2102/production-distribution-tracking/models"
)

func GetProductionReport(c *gin.Context) {
	var batches []models.ProductionBatch

	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	query := database.DB

	if startDateStr != "" && endDateStr != "" {
		layout := "2006-01-02"
		startDate, err1 := time.Parse(layout, startDateStr)
		endDate, err2 := time.Parse(layout, endDateStr)

		if err1 != nil || err2 != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
			return
		}

		query = query.Where("created_at BETWEEN ? AND ?", startDate, endDate)
	}

	query.Find(&batches)
	c.JSON(http.StatusOK, batches)
}

func GetDistributionReport(c *gin.Context) {
	var distributions []models.Distribution

	batchID := c.Query("batch_id")
	location := c.Query("location")

	query := database.DB

	if batchID != "" {
		query = query.Where("batch_id = ?", batchID)
	}
	if location != "" {
		query = query.Where("location = ?", location)
	}

	query.Find(&distributions)
	c.JSON(http.StatusOK, distributions)
}
