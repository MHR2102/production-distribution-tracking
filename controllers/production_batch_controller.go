package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/MHR2102/production-distribution-tracking/database"
	"github.com/MHR2102/production-distribution-tracking/models"
)

func GetProductionBatches(c *gin.Context) {
	var batches []models.ProductionBatch
	database.DB.Find(&batches)
	c.JSON(http.StatusOK, batches)
}

func CreateProductionBatch(c *gin.Context) {
	var input models.ProductionBatch
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	batch := models.ProductionBatch{
		BatchNumber: input.BatchNumber,
		TotalUnits:  input.TotalUnits,
		Status:      "proses",
	}

	if err := database.DB.Create(&batch).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat production batch"})
		return
	}

	c.JSON(http.StatusOK, batch)
}

func UpdateProductionBatch(c *gin.Context) {
	var batch models.ProductionBatch
	id := c.Param("id")

	if err := database.DB.First(&batch, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Production batch tidak ditemukan"})
		return
	}

	var input models.ProductionBatch
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	batch.Status = input.Status

	if err := database.DB.Save(&batch).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal memperbarui production batch"})
		return
	}

	c.JSON(http.StatusOK, batch)
}

func DeleteProductionBatch(c *gin.Context) {
	id := c.Param("id")
	var batch models.ProductionBatch
	if err := database.DB.First(&batch, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Production batch tidak ditemukan"})
		return
	}

	// Hanya memperbolehkan penghapusan jika status "gagal"
	if batch.Status != "gagal" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Hanya batch dengan status 'gagal' yang dapat dihapus"})
		return
	}

	if err := database.DB.Delete(&batch).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus production batch"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Production batch berhasil dihapus"})
}
