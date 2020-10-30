package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/haithngn/go-crud/model"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type QuestionContoller struct {
	DB *gorm.DB
}

func (controller *QuestionContoller) CreatingNewQuestion(context *gin.Context) {
	title := context.PostForm("title")
	content := context.PostForm("content")
	question := model.Question{
		Title:   title,
		Content: content,
	}
	err := controller.DB.Model(&model.Question{}).Create(&question).Error
	if err != nil {
		context.AbortWithStatusJSON(http.StatusConflict, err)
		return
	}
	context.JSON(http.StatusOK, question)
}

func (controller *QuestionContoller) GetQuestion(context *gin.Context) {
	questionId := context.Param("id")
	var question model.Question
	err := controller.DB.Model(model.Question{}).Where("id = ?", questionId).Take(&question).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		context.AbortWithStatusJSON(http.StatusNoContent, err)
		return
	}
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnprocessableEntity, err)
		return
	}

	context.JSON(http.StatusOK, question)
}

func (controller *QuestionContoller) UpdateQuestion(context *gin.Context) {
	questionId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}
	title := context.PostForm("title")
	content := context.PostForm("content")
	question := model.Question{
		ID:      questionId,
		Title:   title,
		Content: content,
	}
	err = controller.DB.Model(model.Question{}).Where("id = ?", questionId).Updates(&question).Error
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnprocessableEntity, err)
		return
	}
	context.JSON(http.StatusOK, question)
}

func (controller *QuestionContoller) RemoveQuestion(context *gin.Context) {
	questionId, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err = controller.DB.Delete(model.Question{}, questionId).Error
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnprocessableEntity, err)
		return
	}
	context.String(http.StatusOK, "The question has been removed.")
}
