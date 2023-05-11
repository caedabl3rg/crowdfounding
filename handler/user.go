package handler

import (
	"crowfunding/helper"
	"crowfunding/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService: userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// tangkap input dari user
	// map input dari user ke struct RegisterInput
	// struct di atas kita passing sebagai paramete service

	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//token ,err := h.jwtService.GenerateTOken

	formatter := user.FormatUser(newUser, "tokentokentoken")

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context) {
	// USER MASUKKAN INPUT (EMAIL DAN PASSWORD)
	// INPUT DITANGKAP HANDLER
	// MAPPING DARI INPUT USER KE INPUT STRUCT
	// DISERVICE MENCARI DG BANTUAN REPOSITORY USER DENGA EMAIL X
	// MENCOCOKKAN PASSWORD

	var input user.LoginInput 
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors  }

		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity,"error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	loggedUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors" : err.Error()}
		response := helper.APIResponse("Login Failed",http.StatusUnprocessableEntity,"error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter := user.FormatUser(loggedUser, "tokentoken")
	response := helper.APIResponse("Success Login", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	return
}
