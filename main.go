package main

import (
	_ "SchoolMgtmt/docs"
	"SchoolMgtmt/src/controller"
	"SchoolMgtmt/src/setup"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           School management system
// @version         1.0
// @description     This is a School management system
// @termsOfService  http://swagger.io/terms/

// @contact.name   Emmanuel Adeleke
// @contact.url    github.com/LoluPapi
// @contact.email  aemmanlolu@gmail.com

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	setup.Connection()
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "welcome to deez administration",
		})
	})
	r.POST("/student", controller.CreateStudent)
	r.GET("/student/:id", controller.GetStudentById)
	r.PUT("/student/:id", controller.UpdateStudent)
	r.DELETE("/student/:id", controller.DeleteStudent)
	r.GET("/list/student", controller.ListStudents)
	r.GET("/list/course", controller.ListCourses)
	r.GET("/list/student/:course_name", controller.ListStudentsByCourse)
	r.PUT("/course/:id/student", controller.UpdateStudentCourseById)
	r.Run()
}
