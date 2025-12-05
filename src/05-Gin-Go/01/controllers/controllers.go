package controllers

import (
	"gin/database"
	"gin/models"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func ReturnSignal(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

func Introduce(c *gin.Context) {
	name := c.Param("name")
	c.JSON(200, gin.H{
		"message": "Hello, " + name + "! Welcome to Gin framework.",
	})
}

func GetStudents(c *gin.Context) {
	database.DB.Find(&[]models.Student{})
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(200, students)
}

func GetStudentByID(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	result := database.DB.First(&student, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(200, student)
}

func AddStudent(c *gin.Context) {
	var newStudent models.Student

	if err := c.ShouldBindJSON(&newStudent); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&newStudent)
	c.JSON(201, newStudent)
}

func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student
	result := database.DB.First(&student, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Student not found"})
		return
	}

	database.DB.Delete(&student)
	c.JSON(200, gin.H{"message": "Student deleted successfully"})
}

func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student

	if err := database.DB.First(&student, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Student not found"})
		return
	}

	var updatedData models.Student
	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	student.Name = updatedData.Name
	student.Age = updatedData.Age
	student.Email = updatedData.Email

	database.DB.Save(&student)
	c.JSON(200, student)
}

func GetStudentByName(c *gin.Context) {
	name := c.Param("name")
	var student models.Student
	result := database.DB.Where("name = ?", name).First(&student)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(200, student)
}
