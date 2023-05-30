package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/panisin000/course/gorm"
	"github.com/panisin000/course/model"
	"github.com/panisin000/course/util"
)

func ListCourses(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		courses, err := db.GetAllCourse()
		if err != nil {
			util.SendError(c, http.StatusInternalServerError, err)
			return
		}
		c.IndentedJSON(http.StatusOK, courses)
	}
}

func GetCourse(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			util.SendError(c, http.StatusBadRequest, err)
			return
		}
		course, err := db.GetCourse(uint(id))
		if err != nil {
			util.SendError(c, http.StatusNotFound, err)
			return
		}
		c.IndentedJSON(http.StatusOK, course)
	}
}

func CreateCourse(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := new(model.Course)
		if err := c.BindJSON(req); err != nil {
			util.SendError(c, http.StatusBadRequest, err)
			return
		}
		if err := db.CreateCourse(req); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.IndentedJSON(http.StatusOK, req)
	}
}
