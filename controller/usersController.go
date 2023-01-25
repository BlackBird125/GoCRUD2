package controllers

import (
	"net/http"
	"strconv"

	"github.com/BlackBird125/GoCRUD2/models"
	"github.com/BlackBird125/GoCRUD2/models/repository"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

// GetAllUser ...                                                        
// @Summary User一覧を配列で返す                                          
// @Tags User                                                          
// @Produce  json                                                      
// @Success 200 {object} []models.User 
// @Router /users [get]                                                
func (pc UserController) Index(c *gin.Context) {
    var u repository.UserRepository
    users := u.GetAll()
    c.JSON(http.StatusOK, users)
}

// GetUser ...                                                        
// @Summary Userを返す     
// @Description get string by ID
// @ID get-string-by-int                                     
// @Tags User                                                          
// @Produce  json                    
// @Param  id path int true "User ID"                                  
// @Success 200 {object} models.User 
// @Router /users/{id} [get]        
func (pc UserController) Show(c *gin.Context) {
    var u repository.UserRepository
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	user := u.GetByID(idInt)
	c.JSON(http.StatusOK, user)
}

// PostTodo ...
// @Summary 新規Todoを作成
// @Tags Todo
// @Accept  json                         
// @Produce  json
// @Param title body string true "title" 
// @Param body body string true "body"
// @Success 200 {object} models.User
// @Router /users [post]
func (pc UserController) Create(c *gin.Context) {
    var u repository.UserRepository
	var user models.User
    if err := c.BindJSON(&user); err != nil {
        c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
    }
	createdUser, err :=  u.Create(user)
	if err != nil{
		c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
	}
	c.JSON(http.StatusOK, createdUser)
}

// DeleteUser ...                                                        
// @Summary Userを削除する                                          
// @Tags User                                                          
// @Produce  json                                                      
// @Success 200 {object} models.User
// @Router /users/{id} [delete]             
func (pc UserController) Delete(c *gin.Context) {
    var u repository.UserRepository
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	user, err := u.Delete(idInt)
	if err != nil{
		c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
	}
	c.JSON(http.StatusOK, user)
}

func (pc UserController) Update(c *gin.Context) {
    var u repository.UserRepository
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	var user models.User
	if err := c.BindJSON(&user); err != nil {
        c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
    }
	user, err := u.Update(idInt,user)
	if err != nil{
		c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
	}
	c.JSON(http.StatusOK, user)
}