package controller

import (
	"OTI-inbound/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginIn struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterIn struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type PasswordUpdateIn struct {
	OldPassword     string `json:"old_password" binding:"required"`
	NewPassword     string `json:"new_password" binding:"required"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

// LoginAdmin godoc
// @Summary Login as admin
// @Description Logging in to get jwt token access
// @Tags Auth
// @Param Body body LoginInput true "the body to login a admin"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /loginadmin [post]
func LoginAdmin(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input LoginIn

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.Administrator{}
	u.Username = input.Username
	u.Password = input.Password

	token, err := models.LoginCheckAdmin(u.Username, u.Password, db)

	if err != nil {
		fmt.Println("error")
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
	}

	user := map[string]string{
		"username": u.Username,
	}

	c.JSON(http.StatusOK, gin.H{"message": "login succesful", "user": user, "token": token})

}

// Register godoc
// @Summary Register a admin.
// @Description registering an admin from public access.
// @Tags Auth
// @Param Body body RegisterInput true "the body to register a admin"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /registeradmin [post]
func RegisterAdmin(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input RegisterIn

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.Administrator{}

	u.Username = input.Username
	u.Password = input.Password

	_, err := u.SaveAdmin(db)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := map[string]string{
		"username": input.Username,
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success", "user": user})

}

// UpdatePassword godoc
// @Summary Update admin password.
// @Description Update admin's password.
// @Tags Auth
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "admin id"
// @Param Body body PasswordUpdateInput true "the body to update admin's password"
// @Success 200 {object} map[string]interface{}
// @Router /update-password-admin/{id} [put]
func UpdatePassAdmin(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input PasswordUpdateIn

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.NewPassword != input.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "New password and confirm password do not match"})
		return
	}

	var admin models.Administrator
	if err := db.Where("id = ?", c.Param("id")).First(&admin).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := models.VerifyPassword(input.OldPassword, admin.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Old password is incorrect"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	admin.Password = string(hashedPassword)

	if err := db.Save(&admin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}
