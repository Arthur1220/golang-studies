package routes

import (
	"gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/", controllers.ReturnSignal)
	r.GET("/introduce/:name", controllers.Introduce)

	r.GET("/students", controllers.GetStudents)
	r.GET("/students/:id", controllers.GetStudentByID)
	r.POST("/students", controllers.AddStudent)
	r.DELETE("/students/:id", controllers.DeleteStudent)
	r.PUT("/students/:id", controllers.UpdateStudent)
	r.GET("/students/name/:name", controllers.GetStudentByName)

	r.Run(":8080")
}
