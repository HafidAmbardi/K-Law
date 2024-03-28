package controller

import (
	"OTI-inbound/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type postInput struct {
	UserID       uint   `json:"user_id"`
	Companies    string `json:"companies"`
	Post_Title   string `json:"post_title"`
	Post_Text    string `json:"post_text"`
	CategoriesID uint   `json:"categories_id"`
}

// GetAllPost godoc
// @Summary Get all Post.
// @Description Get a list of Post.
// @Tags Post
// @Produce json
// @Success 200 {object} []models.Post
// @Router /post [get]
func GetAllPost(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var post []models.Post
	db.Preload("Categories").Find(&post)

	c.JSON(http.StatusOK, gin.H{"data": post})
}

// GetPostByText godoc
// @Summary Get Post.
// @Description Get a Post by text.
// @Tags Post
// @Produce json
// @Param text path string true "post text"
// @Success 200 {object} models.Post
// @Router /post/{id} [get]
func GetPostByText(c *gin.Context) {
	var post models.Post

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("post_text = ?", c.Param("text")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
}

// CreatePost godoc
// @Summary Create New Post.
// @Description Creating a new Post.
// @Tags Post
// @Param Body body postInput true "the body to create a new Post"
// @Produce json
// @Success 200 {object} models.Post
// @Router /post [post]
func CreatePost(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input postInput
	var user models.User
	var categories models.Categories
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
	}

	if err := db.Where("id = ?", input.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserID not found!"})
		return
	}

	if err := db.Where("id = ?", input.CategoriesID).First(&categories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "categories not found!"})
		return
	}

	post := models.Post{
		UserID:       input.UserID,
		Companies:    input.Companies,
		Post_Title:   input.Post_Title,
		Post_Text:    input.Post_Text,
		CategoriesID: input.CategoriesID,
	}
	db.Create(&post)

	c.JSON(http.StatusOK, gin.H{"data": post})
}

// UpdatePost godoc
// @Summary Update Post.
// @Description Update Post by id.
// @Tags Post
// @Produce json
// @Param id path string true "post id"
// @Param Body body postInput true "the body to update an Post"
// @Success 200 {object} models.Post
// @Router /post/{id} [put]
func UpdatePost(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var post models.Post

	if err := db.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input postInput
	var categories models.Categories
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("categories_text = ?", input.CategoriesID).First(&categories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ProductID not found!"})
		return
	}

	var updatedInput models.Post
	updatedInput.Companies = input.Companies
	updatedInput.CategoriesID = input.CategoriesID
	updatedInput.Post_Title = input.Post_Title
	updatedInput.Post_Text = input.Post_Text
	updatedInput.UpdatedAt = time.Now()

	db.Model(&post).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": post})
}

// DeletePost godoc
// @Summary Delete one Post.
// @Description Delete a Post by id.
// @Tags Post
// @Produce json
// @Param id path string true "post id"
// @Success 200 {object} map[string]boolean
// @Router /post/{id} [delete]
func DeletePost(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	var post models.Post
	if err := db.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&post)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
