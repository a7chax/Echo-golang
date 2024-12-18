package handler

import (
	"echo-golang/model"
	model_request "echo-golang/model/request"
	model_response "echo-golang/model/response"
	service "echo-golang/service/user"
	"echo-golang/utils"
	"echo-golang/validators"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type IUserHandler struct {
	service service.IUserService
}

func UserHandler(service service.IUserService) *IUserHandler {
	return &IUserHandler{service}
}

func (h *IUserHandler) GetAllUser(context echo.Context) error {
	user, metadata, err := h.service.GetAllUser()
	fmt.Println(err, "erorrnya")
	if err != nil {
		return context.JSON(http.StatusInternalServerError, model.BaseResponsePaginationNoData{
			IsSuccess: false,
			Message:   err.Error(),
			Metadata:  metadata,
		})
	}
	return context.JSON(http.StatusOK, model.BaseResponsePagination[model_response.User]{
		IsSuccess: true,
		Message:   "Get all user success",
		Data:      &user,
	})
}

func (h *IUserHandler) LoginUser(context echo.Context) error {
	var login model_request.Login

	err := context.Bind(&login)
	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	validator := validators.New()

	if err = validator.Validate(login); err != nil {
		return context.JSON(http.StatusBadRequest, model.BaseResponseNoData{
			IsSuccess: false,
			Message:   err.Error(),
		})
	}

	response, err := h.service.LoginUser(login)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}
	return context.JSON(http.StatusOK, response)

}

func (h *IUserHandler) RefreshToken(context echo.Context) error {
	token := context.Request().Header.Get("Authorization")
	fmt.Println(token)
	response, err := h.service.RefreshToken(token)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, model.BaseResponseNoData{
			IsSuccess: false,
			Message:   err.Error(),
		})
	}
	return context.JSON(http.StatusOK, response)
}

func (h *IUserHandler) RegisterUser(context echo.Context) error {
	var register model_request.Register

	err := context.Bind(&register)
	if err != nil {
		return context.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	validator := validators.New()

	if err = validator.Validate(register); err != nil {
		return context.JSON(http.StatusBadRequest, model.BaseResponseNoData{
			IsSuccess: false,
			Message:   err.Error(),
		})
	}

	response, err := h.service.RegisterUser(register)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, err)
	}
	return context.JSON(http.StatusOK, response)
}

func (h *IUserHandler) GetUser(context echo.Context) error {
	token := context.Request().Header.Get("Authorization")
	claims := &utils.JwtCustomClaims{}
	secret := os.Getenv("JWT_SECRET")

	jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	user, err := h.service.GetUser(claims.Id)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, model.BaseResponseNoData{
			IsSuccess: false,
			Message:   user.Message,
		})
	}
	return context.JSON(http.StatusOK, user)
}
