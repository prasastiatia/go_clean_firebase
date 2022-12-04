package v1Student

import (
	studentController "remake_crud_golang/app/interfaces/controllers/v1/student"

	"github.com/gin-gonic/gin"
)

func MountStudent(r *gin.Engine) {
	student := r.Group("student")
	{
		student.POST("/create", studentController.CreateStudent())
		student.GET("/:id", studentController.DetailStudent())
		student.GET("/list", studentController.ListStudent())
		student.PUT("/:id", studentController.Update())
		student.DELETE("/:id", studentController.DeleteStudent())
	}
}
