package controller

import (
	"OTI-inbound/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type categoriesInput struct {
	Categories_Text string `json:"categories_text"`
	Categories_Desc string `json:"categories_desc"`
}

// GetAllCategories godoc
// @Summary Get all Categories.
// @Description Get a list of Categories.
// @Tags Categories
// @Produce json
// @Success 200 {object} []models.Categories
// @Router /categories [get]
func GetAllCategories(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var categories []models.Categories
	db.Find(&categories)

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

// CreateCategories godoc
// @Summary Create New Categories.
// @Description Creating a new Categories.
// @Tags Categories
// @Param Body body categoriesInput true "the body to create a new Categories"
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Categories
// @Router /categories [post]
func CreateCategories(c *gin.Context) {
	var input categoriesInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	categories := models.Categories{Categories_Text: input.Categories_Text, Categories_Desc: input.Categories_Desc}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&categories)

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

// UpdateCategories godoc
// @Summary Update Categories.
// @Description Update Categories by id.
// @Tags Categories
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "categories id"
// @Param Body body categoriesInput true "the body to update category"
// @Success 200 {object} models.Categories
// @Router /categories/{id} [put]
func UpdateCategories(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var categories models.Categories
	if err := db.Where("id = ?", c.Param("id")).First(&categories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input categoriesInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Categories
	updatedInput.Categories_Text = input.Categories_Text
	updatedInput.Categories_Desc = input.Categories_Desc

	db.Model(&categories).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

// DeleteCategories godoc
// @Summary Delete one Categories.
// @Description Delete a Categories by id.
// @Tags Categories
// @Produce json
// @Param id path string true "categories id"
// @Success 200 {object} map[string]boolean
// @Router /categories/{id} [delete]
func DeleteCategories(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var categories models.Categories
	if err := db.Where("id = ?", c.Param("id")).First(&categories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "record not found!"})
		return
	}
	db.Delete(&categories)
	c.JSON(http.StatusOK, gin.H{"Data": true})
}
