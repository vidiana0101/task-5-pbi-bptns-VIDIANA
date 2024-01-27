package usercontrollers

import (
	"encoding/json"
	"net/http"

	"github.com/vidiana0101/task-5-pbi-bptns-VIDIANA/models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	var users []models.User

	models.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"users": users})

}

func Show(c *gin.Context) {
	var users models.User
	id := c.Param("id")

	if err := models.DB.First(&users, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"user": users})
}

func Create(c *gin.Context) {

	var users models.User

	if err := c.ShouldBindJSON(&users); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&users)
	c.JSON(http.StatusOK, gin.H{"user": users})
}

func Update(c *gin.Context) {
	var users models.User
	id := c.Param("id")

	if err := c.ShouldBindJSON(&users); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&users).Where("id = ?", id).Updates(&users).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})

}

func Delete(c *gin.Context) {

	var users models.User

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&users, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
