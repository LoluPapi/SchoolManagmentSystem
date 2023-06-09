package models

import (
	"github.com/jinzhu/gorm"
)

type Course struct {
	gorm.Model
	Title    string    `json:"title"`
	Students []Student `json:"students" gorm:"many2many:student_courses"`
}
type CourseResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
type UpdateCoursesRequest struct {
	CourseIds []int `json:"course_ids" binding:"required"`
}
