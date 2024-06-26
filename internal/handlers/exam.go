package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshua468/access-bank-exam-system/internal/repository"
	"github.com/joshua468/access-bank-exam-system/models"
	"github.com/joshua468/access-bank-exam-system/pkg/utils"
	"gorm.io/gorm"
)

// CreateExam godoc
// @Summary Create a new exam
// @Description Create a new exam with questions
// @Tags Exam
// @Accept  json
// @Produce  json
// @Param   exam body models.Exam true "Exam"
// @Success 200 {object} utils.Response
// @Failure 400 {object} utils.Response
// @Router /exams [post]
func CreateExam(c *gin.Context, db *gorm.DB) {
	var exam models.Exam
	if err := c.ShouldBindJSON(&exam); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResponseError(err))
		return
	}

	if err := repository.CreateExam(db, &exam); err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(err))
		return
	}

	c.JSON(http.StatusOK, utils.ResponseSuccess("Exam created successfully"))
}

// ListExams godoc
// @Summary List all exams
// @Description Get a list of all exams
// @Tags Exam
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Exam
// @Router /exams [get]
func ListExams(c *gin.Context, db *gorm.DB) {
	exams, err := repository.ListExams(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResponseError(err))
		return
	}

	c.JSON(http.StatusOK, exams)
}
