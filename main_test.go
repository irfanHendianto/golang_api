package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
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
		Username: "test@gmail.com",
		Password: "pass@word123456a",
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
		FullName:        "test",
		Username:        "test@gmail.com",
		Password:        "pass@word123456a",
		ConfirmPassword: "pass@word123456a",
	}
	jsonValue, _ := json.Marshal(&registerData)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/auth/register", bytes.NewReader(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetProfile(t *testing.T) {
	r := SetUpRouter()
	r.POST("/api/auth/login", authController.Login)
	r.Use(middleware.AuthorizeJWT(jwtService))
	r.GET("/api/user/profile", userController.Profile)

	getProfileData := dto.LoginDTO{
		Username: "test@gmail.com",
		Password: "pass@word123456a",
	}
	jsonValue, _ := json.Marshal(&getProfileData)
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
		Username: "test@gmail.com",
		Password: "pass@word123456a",
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
		FullName:        "test",
		Username:        "test@gmail.com",
		Password:        "pass@word123456a",
		ConfirmPassword: "pass@word123456a",
	}
	jsonValueUpdate, _ := json.Marshal(&UpdateData)
	req, _ := http.NewRequest("PUT", "/api/user/profile", bytes.NewBuffer(jsonValueUpdate))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
