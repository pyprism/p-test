package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/p-test/models"
	"log"
	"net/http"
	"strconv"
)

type UserController struct {}

type UserInput struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName string `json:"lastName" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (n *UserController) CreateUser(c *gin.Context) {
	var userInput UserInput

	err := c.BindJSON(&userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	pk := models.CreateUser(userInput.FirstName, userInput.LastName, userInput.Password)


	r := map[string]uint{"id": pk}
	c.JSON(http.StatusCreated, r)
}

func (n *UserController) GetUser(c *gin.Context) {
	uId := c.Param("userId")
	id, err := strconv.Atoi(uId)
	if err != nil {
		log.Fatalln(err)
	}
	user := models.GetUserById(id)
	name := user.FirstName + " " + user.LastName
	r := map[string]interface{}{"id": user.ID, "name": name}


	c.JSON(http.StatusFound, r)
}

type TagInput struct {
	UserId int `json:"user_id"`
	Name string `json:"name"`
}

func (n *UserController) SetTag(c *gin.Context) {

	var tagInput TagInput
	err := c.BindJSON(&tagInput)
	if err != nil {
		log.Fatalln(err)
	}
	models.UpdateUser(tagInput.UserId, tagInput.Name)

	r := map[string]interface{}{"id": "user.ID", "name": "name"}
	c.JSON(http.StatusOK, r)
}