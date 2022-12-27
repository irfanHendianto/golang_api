package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang_api/dto"
	"github.com/golang_api/entity"
	"github.com/golang_api/helper"
	"github.com/golang_api/service"
)

// AuthControllers contains what auth can do
type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Fail to process", errDTO.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Username, loginDTO.Password)
	if v, ok := authResult.(entity.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = generatedToken
		response := helper.BuildResponse(true, "Ok", v)
		ctx.JSON(http.StatusOK, response)
		return
	}
	response := helper.BuildErrorResponse("Please check again your password / username", "Invalid Credential", helper.EmptyObject{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	errDTO := ctx.ShouldBind(&registerDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Fail to process request", errDTO.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if !c.authService.IsDuplicateUsername(registerDTO.Username) {
		response := helper.BuildErrorResponse("Username Already Register !", "Duplicate username", helper.EmptyObject{})
		ctx.JSON(http.StatusConflict, response)
		return
	}
	if !helper.ComparedPassword(registerDTO.Password, registerDTO.ConfirmPassword) {
		response := helper.BuildErrorResponse("Password and Confirm Password must be match", "Password doesn't match", helper.EmptyObject{})
		ctx.JSON(http.StatusConflict, response)
		return
	}
	createdUser := c.authService.CreateUser(registerDTO)
	response := helper.BuildResponse(true, "Ok", createdUser)
	ctx.JSON(http.StatusCreated, response)
}
