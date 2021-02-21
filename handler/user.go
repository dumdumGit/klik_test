package handler

import (
	"klik_test/auth"
	"klik_test/helper"
	"klik_test/user"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var inputUser user.RegisterUserInput

	err := c.ShouldBindJSON(&inputUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponse(
			"Oops, Something went wrong",
			http.StatusUnprocessableEntity,
			"failed",
			helper.FormatValidationError(err),
		))

		return
	}

	newUser, err := h.userService.RegisterUser(inputUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponse(
			"Oops, Something went wrong",
			http.StatusBadRequest,
			"failed",
			err.Error(),
		))

		return
	}

	token, err := h.authService.GenerateToken(newUser.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponse(
			"Oops, Invalid Token",
			http.StatusBadRequest,
			"failed",
			err.Error(),
		))

		return
	}

	formatter := user.FormatUser(newUser, token)

	c.JSON(http.StatusOK, helper.APIResponse(
		"Account has been Registered",
		http.StatusOK,
		"success",
		formatter,
	))
}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponse(
			"Oops, Something went wrong",
			http.StatusUnprocessableEntity,
			"failed",
			helper.FormatValidationError(err),
		))

		return
	}

	login, err := h.userService.Login(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponse(
			"Oops, Something went wrong",
			http.StatusBadRequest,
			"failed",
			err.Error(),
		))

		return
	}

	token, err := h.authService.GenerateToken(login.Id)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponse(
			"Oops, Invalid Token",
			http.StatusBadRequest,
			"failed",
			err.Error(),
		))

		return
	}

	session := sessions.Default(c)
	session.Set("id", login.Id)
	session.Set("email", login.Email)
	session.Save()

	formatter := user.FormatUser(login, token)

	c.JSON(http.StatusOK, helper.APIResponse(
		"Successfull Loggedin",
		http.StatusOK,
		"success",
		formatter,
	))
}

func (h *userHandler) AvailabilityEmail(c *gin.Context) {
	var input user.EmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponse(
			"Oops, Something went wrong",
			http.StatusUnprocessableEntity,
			"failed",
			helper.FormatValidationError(err),
		))

		return
	}

	IsEmailExist, err := h.userService.IsEmailExist(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.APIResponse(
			"Oops, Something went wrong",
			http.StatusUnprocessableEntity,
			"failed",
			helper.FormatValidationError(err),
		))

		return
	}

	data := gin.H{
		"is_available": IsEmailExist,
	}

	var meta string
	if IsEmailExist {
		meta = "Email is Available"
	} else {
		meta = "Email is already exists"
	}

	c.JSON(http.StatusOK, helper.APIResponse(
		meta,
		http.StatusOK,
		"success",
		data,
	))
}
