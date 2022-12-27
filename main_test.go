package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert"
	"github.com/golang_api/dto"
	"github.com/golang_api/helper"
	"github.com/golang_api/middleware"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestLogin(t *testing.T) {
	r := SetUpRouter()
	r.POST("/api/auth/login", authController.Login)
	loginData := dto.LoginDTO{
		Username: "irfan@gmail.com",
		Password: "pass@word1234568",
	}
	jsonValue, _ := json.Marshal(&loginData)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestRegister(t *testing.T) {
	r := SetUpRouter()
	r.POST("/api/auth/register", authController.Register)
	r.DELETE("/api/auth/delete", userController.Delete)
	registerData := dto.RegisterDTO{
		FullName:        "irfan a",
		Username:        "irfan022@gmail.com",
		Password:        "pass@word123456a",
		ConfirmPassword: "pass@word123456a",
	}
	jsonValue, _ := json.Marshal(&registerData)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewReader(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	var responseBody helper.Response
	json.NewDecoder(w.Body).Decode(&responseBody)
	temp := fmt.Sprintf("%.f", responseBody.Data.(map[string]interface{})["id"])
	ID, _ := strconv.ParseInt(temp, 10, 64)
	deleteUser := dto.DeleteUserDTO{
		ID: ID,
	}
	jsonValueUser, _ := json.Marshal(&deleteUser)
	wUser := httptest.NewRecorder()
	reqUser, _ := http.NewRequest("DELETE", "/api/auth/delete", bytes.NewReader(jsonValueUser))
	reqUser.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(wUser, reqUser)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetProfile(t *testing.T) {
	r := SetUpRouter()
	r.POST("/api/auth/login", authController.Login)
	r.Use(middleware.AuthorizeJWT(jwtService))
	r.GET("/api/user/profile", userController.Profile)

	loginData := dto.LoginDTO{
		Username: "irfan@gmail.com",
		Password: "pass@word1234568",
	}
	jsonValue, _ := json.Marshal(&loginData)
	wAuth := httptest.NewRecorder()
	reqAuth, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(jsonValue))
	reqAuth.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(wAuth, reqAuth)
	var responseBody helper.Response
	json.NewDecoder(wAuth.Body).Decode(&responseBody)
	token := fmt.Sprintf("%s", responseBody.Data.(map[string]interface{})["token"])

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/user/profile", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestUpdateProfile(t *testing.T) {
	r := SetUpRouter()
	r.POST("/api/auth/login", authController.Login)
	r.Use(middleware.AuthorizeJWT(jwtService))
	r.PUT("/api/user/profile", userController.Profile)

	loginData := dto.LoginDTO{
		Username: "irfan@gmail.com",
		Password: "pass@word1234568",
	}
	jsonValue, _ := json.Marshal(&loginData)
	wAuth := httptest.NewRecorder()
	reqAuth, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(jsonValue))
	reqAuth.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(wAuth, reqAuth)
	var responseBody helper.Response
	json.NewDecoder(wAuth.Body).Decode(&responseBody)
	token := fmt.Sprintf("%s", responseBody.Data.(map[string]interface{})["token"])

	w := httptest.NewRecorder()
	UpdateData := dto.UserUpdateDTO{
		FullName:        "irfan wijaya",
		Username:        "irfan@gmail.com",
		Password:        "pass@word1234568",
		ConfirmPassword: "pass@word1234568",
	}
	jsonValueUpdate, _ := json.Marshal(&UpdateData)
	req, _ := http.NewRequest("PUT", "/api/user/profile", bytes.NewBuffer(jsonValueUpdate))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
