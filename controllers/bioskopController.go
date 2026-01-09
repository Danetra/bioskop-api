package controllers

import (
	"bioskop-api/config"
	"bioskop-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBioskop(c *gin.Context) {
	var bioskop models.Bioskop

	if err := c.ShouldBindJSON(&bioskop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON",
		})
		return
	}

	if bioskop.Nama == "" || bioskop.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Nama dan Lokasi tidak boleh kosong",
		})
		return
	}

	query := `INSERT INTO bioskop (nama, lokasi, rating) VALUES ($1, $2, $3) RETURNING id`

	insert := config.DB.QueryRow(
		query, bioskop.Nama, bioskop.Lokasi, bioskop.Rating).Scan(&bioskop.ID)

	if insert != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": insert.Error(), // ← tampilkan error asli
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Bioskop berhasil ditambahkan",
		"data":    bioskop,
	})
}
