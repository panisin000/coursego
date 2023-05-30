package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/panisin000/course/gorm"
	"github.com/panisin000/course/handler"
	"github.com/panisin000/course/middleware"
)

func Init(db db.DB, gorm *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	// mdb := mockDBgorm
	r.GET("/courses", handler.ListCourses(db))
	r.GET("/courses/:id", handler.GetCourse(gorm))
	r.POST("/courses", handler.CreateCourse(gorm))
	r.POST("/classes", handler.CreateClasses(gorm))
	r.POST("/enrollments", middleware.RequireUser(gorm), handler.EnrollClass(gorm))
	r.POST("/register", handler.Register(gorm))
	r.POST("/login", handler.Login(gorm))
	return r
}
