package main

import (
	// "fmt"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/panisin000/course/gorm"
	"github.com/panisin000/course/model"
	"github.com/panisin000/course/server"
)

type mockDB struct {
}

func (m *mockDB) GetAllCourse() ([]model.Course, error) {
	return nil, errors.New("This is error")
}
func main() {
	gorm, err := gorm.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	// if err := db.Reset(); err != nil {
	// 	log.Fatal(err)
	// }
	if err := gorm.AutoMigrate(); err != nil {
		log.Fatal(err)
	}

	r := server.Init(db, gorm)
	// r.Use(cors.Default())
	// // mdb := mockDB{}
	// r.GET("/courses", handler.ListCourses(db))
	// r.GET("/courses/:id", handler.GetCourse(db))
	// r.POST("/courses", handler.CreateCourse(db))
	// r.POST("/classes", handler.CreateClasses(db))
	// r.POST("/enrollments", middleware.RequireUser(db), handler.EnrollClass(db))
	// r.POST("/register", handler.Register(db))
	// r.POST("/login", handler.Login(db))
	r.Run(":8080")
}

func Error(c *gin.Context, status int, err error) {
	log.Println(err)
	c.JSON(status, gin.H{
		"message": err.Error(),
	})
}
