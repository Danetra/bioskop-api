package controllers

import (
	"bioskop-api/config"
	"bioskop-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBioskopAll(c *gin.Context) {
	rows, err := config.DB.Query("SELECT * FROM bioskop")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer rows.Close()

	var bioskops []models.Bioskop

	for rows.Next() {
		var bioskop models.Bioskop
		rows.Scan(&bioskop.ID, &bioskop.Nama, &bioskop.Lokasi, &bioskop.Rating)
		bioskops = append(bioskops, bioskop)
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Data Bioskop Successed Load", "data": bioskops})
}

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

func GetBioskopByID(c *gin.Context) {
	id := c.Param("id")

	var b models.Bioskop
	err := config.DB.
		QueryRow(`SELECT * FROM bioskop WHERE id=$1`, id).
		Scan(&b.ID, &b.Nama, &b.Lokasi, &b.Rating)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Data id: " + id + " is exist", "data": b})
}

func UpdateBioskop(c *gin.Context) {
	id := c.Param("id")
	var b models.Bioskop

	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if b.Nama == "" || b.Lokasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi wajib diisi"})
		return
	}

	result, err := config.DB.Exec(
		`UPDATE bioskop SET nama=$1, lokasi=$2, rating=$3 WHERE id=$4`,
		b.Nama, b.Lokasi, b.Rating, id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Data berhasil diupdate"})
}

func DeleteBioskop(c *gin.Context) {
	id := c.Param("id")

	result, err := config.DB.Exec(`DELETE FROM bioskop WHERE id=$1`, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data tidak ditemukan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Data berhasil dihapus"})
}
