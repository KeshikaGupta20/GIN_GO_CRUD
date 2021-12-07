package controller

import (
	"fmt"
	"keshika/database"
	"keshika/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {

	db := database.DB

	emp := new(models.Employ)

	err := c.BindJSON(&emp)

	if err != nil {

		panic(err)
	}

	result := db.Create(&emp)

	if result != nil {
		fmt.Println("Post sucessfully uploaded")

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "emp": &emp})

	}else {
        c.JSON(http.StatusInternalServerError, 
            gin.H{"status": http.StatusInternalServerError, "error": "Failed to create the user"})
    } 

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})

}

func GetUser(c *gin.Context) {

	db := database.DB

	var emp []models.Employ

	result := db.Find(&emp)

	if result != nil {
		fmt.Println("Post sucessfully uploaded")

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "emp": &emp})

	}else {
        c.JSON(http.StatusInternalServerError, 
            gin.H{"status": http.StatusInternalServerError, "error": "Failed to create the user"})
    } 


	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})

}



func DeleteUser(c *gin.Context) {

	id := c.Param("empid")

	db := database.DB

	var emp models.Employ

	db.Find(&emp, id)

	result := db.Delete(&emp)
	if result != nil {
		fmt.Println("Post sucessfully uploaded")

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "emp": &emp})

	}else {
        c.JSON(http.StatusInternalServerError, 
            gin.H{"status": http.StatusInternalServerError, "error": "Failed to create the user"})
    } 


	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})


}

func GetUserbyid(c *gin.Context) {

	id := c.Param("empid")

	db := database.DB

	var emp models.Employ

	result := db.Find(&emp, id)

	if result != nil {
		fmt.Println("Post sucessfully uploaded")

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "emp": &emp})

	}else {
        c.JSON(http.StatusInternalServerError, 
            gin.H{"status": http.StatusInternalServerError, "error": "Failed to create the user"})
    } 


	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})


}

func UpdateUser(c *gin.Context){

	id := c.Param("empid")

	var emp = new(models.Employ)

	db := database.DB

	err := c.BindJSON(&emp)

	if err != nil {

		panic(err)
	}

	var e models.Employ

	db.Find(&emp, id)

	e.EmpName = emp.EmpName
	e.Address= emp.Address
	e.Email= emp.Email
	e.Phone= emp.Phone

	result := db.Save(&e)

	if result != nil {
		fmt.Println("Post sucessfully uploaded")
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "emp": &e})

	}else {
        c.JSON(http.StatusInternalServerError, 
            gin.H{"status": http.StatusInternalServerError, "error": "Failed to create the user"})
    } 


	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})



}

