package controllers

import (
	"encoding/json"

	"github.com/bdn/jeker/dto/requests"
	r "github.com/bdn/jeker/dto/responses"
	"github.com/bdn/jeker/models"
	services "github.com/bdn/jeker/services"
	"github.com/gofiber/fiber"
)

func UserRegister(ctx *fiber.Ctx) {
	var response r.ResponsesHTTP
	inputBody := new(requests.UserRegLog)
	err := ctx.BodyParser(inputBody)
	if err != nil {
		response = r.ResponsesHTTP{
			StatusCode: 400,
			Message:    "Failed To Process Request",
		}
		sendData, _ := json.Marshal(response)
		ctx.Status(400).Send(sendData)
		return
	}
	if inputBody.Email == "" || inputBody.Username == "" || inputBody.Password == "" {
		response = r.ResponsesHTTP{
			StatusCode: 400,
			Message:    "Missing Input Field",
		}
		sendData, _ := json.Marshal(response)
		ctx.Status(400).Send(sendData)
		return
	}
	payload := models.User{
		Username: inputBody.Username,
		Email:    inputBody.Email,
		Password: inputBody.Password,
	}
	result, err := services.UserRegisterService(payload)
	if err != nil {
		response = r.ResponsesHTTP{
			StatusCode: 400,
			Message:    "Failed To Register User",
			Data:       err.Error(),
		}
		sendData, _ := json.Marshal(response)
		ctx.Status(400).Send(sendData)
		return
	}

	respUser := r.ResponseRegisterUser{
		ID:       result.ID,
		Username: result.Username,
		Email:    result.Email,
	}

	response = r.ResponsesHTTP{
		StatusCode: 201,
		Message:    "Success To Register User",
		Data:       respUser,
	}
	sendData, _ := json.Marshal(response)
	ctx.Status(201).Send(sendData)
	return
}

func UserLogin(ctx *fiber.Ctx) {
	var response r.ResponsesHTTP
	inputBody := new(requests.UserRegLog)
	err := ctx.BodyParser(inputBody)
	if err != nil {
		response = r.ResponsesHTTP{
			StatusCode: 400,
			Message:    "Failed To Process Request",
		}
		sendData, _ := json.Marshal(response)
		ctx.Status(400).Send(sendData)
		return
	}
	if inputBody.Password == "" || (inputBody.Username == "" && inputBody.Email == "") {
		response = r.ResponsesHTTP{
			StatusCode: 400,
			Message:    "Missing Input Field",
		}
		sendData, _ := json.Marshal(response)
		ctx.Status(400).Send(sendData)
		return
	}
	payload := models.User{
		Username: inputBody.Username,
		Email:    inputBody.Email,
		Password: inputBody.Password,
	}
	accessToken, err := services.UserLoginService(payload)
	if err != nil {
		response = r.ResponsesHTTP{
			StatusCode: 400,
			Message:    "Failed To Login",
			Data:       err.Error(),
		}
		sendData, _ := json.Marshal(response)
		ctx.Status(400).Send(sendData)
		return
	}
	response = r.ResponsesHTTP{
		StatusCode: 200,
		Message:    "Success To Login",
		Data: r.LoginUser{
			AccessToken: accessToken,
		},
	}
	sendData, _ := json.Marshal(response)
	ctx.Status(200).Send(sendData)
	return
}
