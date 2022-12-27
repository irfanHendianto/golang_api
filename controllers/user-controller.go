package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/golang_api/dto"
	"github.com/golang_api/helper"
	"github.com/golang_api/service"
)

// UserController contrains what user controller can do
type UserController interface {
	Update(context *gin.Context)
	Profile(context *gin.Context)
	Delete(context *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

// NewUserController is creating anew instance of UserControlller
func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (c *userController) Update(context *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO
	errDTO := context.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObject{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	if userUpdateDTO.Password != "" {
		if !helper.ComparedPassword(userUpdateDTO.Password, userUpdateDTO.ConfirmPassword) {
			response := helper.BuildErrorResponse("Password and Confirm Password must be match", "Password doesn't match", helper.EmptyObject{})
			context.JSON(http.StatusConflict, response)
			return
		}
	}
	claims := context.MustGet("Claims").(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	userUpdateDTO.ID = id
	u := c.userService.Update(userUpdateDTO)
	res := helper.BuildResponse(true, "OK!", u)
	context.JSON(http.StatusOK, res)
}

func (c *userController) Profile(context *gin.Context) {
	claims := context.MustGet("Claims").(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	user := c.userService.Profile(id)
	res := helper.BuildResponse(true, "OK", user)
	context.JSON(http.StatusOK, res)

}

func (c *userController) Delete(context *gin.Context) {
	var deleteUserDTO dto.DeleteUserDTO
	errDTO := context.ShouldBind(&deleteUserDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObject{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	data := c.userService.Delete(deleteUserDTO.ID)
	if data == 0 {
		res := helper.BuildErrorResponse("Cannot delete data", "Not Found", helper.EmptyObject{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	res := helper.BuildResponse(true, "OK", helper.EmptyObject{})
	context.JSON(http.StatusOK, res)
}
