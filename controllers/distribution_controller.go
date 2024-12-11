package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/MHR2102/production-distribution-tracking/database"
	"github.com/MHR2102/production-distribution-tracking/models"
)

func GetDistributions(c *gin.Context) {
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

func CreateDistribution(c *gin.Context) {
	var input models.Distribution
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Memeriksa apakah batch ada
	var batch models.ProductionBatch
	if err := database.DB.First(&batch, input.BatchID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Production batch tidak ditemukan"})
		return
	}

	distribution := models.Distribution{
		BatchID:      input.BatchID,
		SerialNumber: input.SerialNumber,
		Location:     input.Location,
		Status:       "dikirim",
	}

	if err := database.DB.Create(&distribution).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat distribution"})
		return
	}

	c.JSON(http.StatusOK, distribution)
}

func UpdateDistribution(c *gin.Context) {
	var distribution models.Distribution
	id := c.Param("id")

	if err := database.DB.First(&distribution, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Distribution tidak ditemukan"})
		return
	}

	var input models.Distribution
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	distribution.Location = input.Location
	distribution.Status = input.Status

	if err := database.DB.Save(&distribution).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui distribution"})
		return
	}

	c.JSON(http.StatusOK, distribution)
}
