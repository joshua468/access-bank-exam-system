package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshua468/access-bank-exam-system/internal/auth"
	"github.com/joshua468/access-bank-exam-system/internal/repository"
	"github.com/joshua468/access-bank-exam-system/models"
	"github.com/joshua468/access-bank-exam-system/pkg/utils"
	"gorm.io/gorm"
)

// RegisterUser godoc
// @Summary Register a new user
// @Description Register a new user with username and password
// @Tags User
// @Accept  json
// @Produce  json
// @Param   user body models.User true "User"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /register [post]
func RegisterUser(c *gin.Context, db *gorm.DB) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError(err))
		return
	}

	if err := repository.CreateUser(db, &user); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(err))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseSuccess("User registered successfully"))
}

// LoginUser godoc
// @Summary Login a user
// @Description Login a user with username and password
// @Tags User
// @Accept  json
// @Produce  json
// @Param   credentials body models.User true "Credentials"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Failure 401 {object} utils.Response
// @Router /login [post]
func LoginUser(c *gin.Context, db *gorm.DB) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError(err))
		return
	}

	user, err := repository.GetUserByUsername(db, input.Username)
	if err != nil || user.Password != input.Password {
		c.JSON(http.StatusUnauthorized, utils.ResponseError(errors.New("Invalid username or password")))
		return
	}

	token, err := auth.GenerateJWT(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
