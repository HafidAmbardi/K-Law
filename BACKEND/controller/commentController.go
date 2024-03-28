package controller

import (
	"OTI-inbound/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type commentInput struct {
	PostID       uint   `json:"post_id"`
	Comment_Text string `json:"comment_text"`
}

// GetAllComment godoc
// @Summary Get all Comment.
// @Description Get a list of Comment.
// @Tags Comment
// @Produce json
// @Success 200 {object} []models.Comment
// @Router /comment [get]
func GetAllComment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var comment []models.Comment
	db.Preload("Post").Find(&comment)

	c.JSON(http.StatusOK, gin.H{"data": comment})
}

// CreateComment godoc
// @Summary Create New Comment.
// @Description Creating a new Comment.
// @Tags Comment
// @Param Body body commentInput true "the body to create a new comment"
// @Produce json
// @Success 200 {object} models.Comment
// @Router /comment [post]
func CreateComment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input commentInput
	var post models.Post
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.PostID).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PostID not found!"})
		return
	}

	comment := models.Comment{PostID: input.PostID, Comment_Text: input.Comment_Text}
	db.Create(&comment)

	c.JSON(http.StatusOK, gin.H{"data": comment})
}

// UpdateComment godoc
// @Summary Update Comment.
// @Description Update Comment by id.
// @Tags Comment
// @Produce json
// @Param id path string true "comment id"
// @Param Body body commentInput true "the body to update an comment"
// @Success 200 {object} models.Comment
// @Router /comment/{id} [put]
func UpdateComment(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var comment models.Comment

	if err := db.Where("id = ?", c.Param("id")).First(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var post models.Post
	var input commentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.PostID).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ProductID not found!"})
		return
	}

	var updatedInput models.Comment
	updatedInput.PostID = input.PostID
	updatedInput.Comment_Text = input.Comment_Text
	updatedInput.UpdatedAt = time.Now()

	db.Model(&comment).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": comment})
}

// DeleteComment godoc
// @Summary Delete one Comment.
// @Description Delete a Comment by id.
// @Tags Comment
// @Produce json
// @Param id path string true "Comment id"
// @Success 200 {object} map[string]boolean
// @Router /comment/{id} [delete]
func DeleteComment(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	var comment models.Comment
	if err := db.Where("id = ?", c.Param("id")).First(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&comment)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
