package db

import (

	"github.com/panisin000/course/model"

)


type DB struct {
	GetAllCourse() ([]model.Course,error)
}