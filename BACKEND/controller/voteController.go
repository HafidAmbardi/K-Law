package controller

import (
	"OTI-inbound/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type voteInput struct {
	PostID    uint `json:"post_id"`
	UserID    uint `json : "user_id"`
	Vote_Type bool `json:"vote_type"`
}

// GetAllVote godoc
// @Summary Get all Vote.
// @Description Get a list of Vote.
// @Tags Vote
// @Produce json
// @Success 200 {object} []models.Votes
// @Router /vote [get]
func GetAllVote(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var vote []models.Votes
	db.Find(&vote)

	c.JSON(http.StatusOK, gin.H{"data": vote})
}

// CountVote godoc
// @Summary Get Count Vote.
// @Description Get a count of Vote.
// @Tags Vote
// @Produce json
// @Success 200 {object} []models.Votes
// @Router /countvote [get]
func CountVote(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var vote []models.Votes
	db.Find(&vote)

	upCount := 0
	downCount := 0
	for _, votes := range vote {
		if votes.Vote_Type {
			upCount++
		} else {
			downCount++
		}
	}

	result := map[string]int{
		"up_count":   upCount,
		"down_count": downCount,
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

// CreateVote godoc
// @Summary Create New Vote.
// @Description Creating a new Vote.
// @Tags Vote
// @Param Body body voteInput true "the body to create a new vote"
// @Produce json
// @Success 200 {object} models.Votes
// @Router /vote [post]
func CreateVote(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input voteInput
	var user models.User
	var post models.Post
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("id = ?", input.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserID not found!"})
		return
	}

	if err := db.Where("id = ?", input.PostID).First(&post).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "PostID not found!"})
		return
	}

	vote := models.Votes{UserID: input.UserID, PostID: input.PostID, Vote_Type: input.Vote_Type}
	db.Create(&vote)

	c.JSON(http.StatusOK, gin.H{"data": vote})
}

// DeleteVote godoc
// @Summary Delete one Vote.
// @Description Delete a Vote by id.
// @Tags Vote
// @Produce json
// @Param id path string true "Vote id"
// @Success 200 {object} map[string]boolean
// @Router /vote/{id} [delete]
func DeleteVote(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	var vote models.Votes
	if err := db.Where("id = ?", c.Param("id")).First(&vote).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&vote)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
