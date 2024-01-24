package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/bdn/jeker/dto/requests"
	r "github.com/bdn/jeker/dto/responses"
	"github.com/bdn/jeker/middlewares"
	"github.com/bdn/jeker/models"
	"github.com/bdn/jeker/services"
	"github.com/gofiber/fiber"
)

func CreateBook(ctx *fiber.Ctx) {
	var response r.ResponsesHTTP
	parsed := ctx.Locals("user")
	userData, ok := parsed.(middlewares.LocalUserData)
	if !ok {
		response = r.ResponsesHTTP{
			StatusCode: 400,
			Message:    "Failed To Process Request",
		}
		sendData, _ := json.Marshal(response)
		ctx.Status(400).Send(sendData)
		return
	}
	inputBody := new(requests.BookRequest)
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
	if inputBody.CategoryId == 0 ||
		inputBody.Title == "" ||
		inputBody.Cover == "" ||
		inputBody.Synopsis == "" {
		response = r.ResponsesHTTP{
			StatusCode: 400,
			Message:    "Missing Input Field",
		}
		sendData, _ := json.Marshal(response)
		ctx.Status(400).Send(sendData)
		return
	}
	bookPayload := models.Book{
		UserId:     userData.ID,
		CategoryId: inputBody.CategoryId,
		Title:      &inputBody.Title,
		Synopsis:   &inputBody.Synopsis,
		Cover:      &inputBody.Cover,
	}
	result, err := services.CreateBookService(bookPayload)
	if err != nil {
		response = r.ResponsesHTTP{
			StatusCode: 400,
			Message:    "Failed To Create Book",
			Data:       err.Error(),
		}
		sendData, _ := json.Marshal(response)
		ctx.Status(400).Send(sendData)
		return
	}
	readyData := r.BookResponse{
		UserId:     result.UserId,
		Author:     userData.Username,
		CategoryId: result.CategoryId,
		Category:   result.Category.Name,
		Title:      *result.Title,
		Cover:      *result.Cover,
		Synopsis:   *result.Synopsis,
		Status:     result.Status,
	}
	response = r.ResponsesHTTP{
		StatusCode: 201,
		Message:    "Success Create Book",
		Data:       readyData,
	}
	sendData, _ := json.Marshal(response)
	ctx.Status(201).Send(sendData)
	return
}

func UpdateBook(ctx *fiber.Ctx) {
	var response r.ResponsesHTTP
	parsed := ctx.Locals("user")
	userData, ok := parsed.(middlewares.LocalUserData)
	if !ok {
		response = r.ResponsesHTTP{
			StatusCode: 400,
			Message:    "Failed To Process Request",
		}
		sendData, _ := json.Marshal(response)
		ctx.Status(400).Send(sendData)
		return
	}
	bookId, err := strconv.Atoi(ctx.Params("book_id"))
	if err != nil {
		response = r.ResponsesHTTP{
			StatusCode: 400,
			Message:    "Failed To Process Request",
		}
		sendData, _ := json.Marshal(response)
		ctx.Status(400).Send(sendData)
		return
	}
	inputBody := new(requests.BookRequest)
	err = ctx.BodyParser(inputBody)
	if err != nil {
		response = r.ResponsesHTTP{
			StatusCode: 400,
			Message:    "Failed To Process Request",
		}
		sendData, _ := json.Marshal(response)
		ctx.Status(400).Send(sendData)
		return
	}

	if inputBody.Title == "" ||
		inputBody.Cover == "" ||
		inputBody.Synopsis == "" {
		response = r.ResponsesHTTP{
			StatusCode: 400,
			Message:    "Missing Input Field",
		}
		sendData, _ := json.Marshal(response)
		ctx.Status(400).Send(sendData)
		return
	}

	bookPayload := models.Book{
		ID:       bookId,
		UserId:   userData.ID,
		Title:    &inputBody.Title,
		Synopsis: &inputBody.Synopsis,
		Cover:    &inputBody.Cover,
	}

	result, err := services.UpdateBookService(bookPayload)
	if err != nil {
		response = r.ResponsesHTTP{
			StatusCode: 400,
			Message:    "Failed To Update Book",
			Data:       err.Error(),
		}
		sendData, _ := json.Marshal(response)
		ctx.Status(400).Send(sendData)
		return
	}
	readyData := r.BookResponse{
		UserId:     result.UserId,
		Author:     userData.Username,
		CategoryId: result.CategoryId,
		Category:   result.Category.Name,
		Title:      *result.Title,
		Cover:      *result.Cover,
		Synopsis:   *result.Synopsis,
	}
	response = r.ResponsesHTTP{
		StatusCode: 200,
		Message:    "Success Update Book",
		Data:       readyData,
	}
	sendData, _ := json.Marshal(response)
	ctx.Status(200).Send(sendData)
	return
}

func GetOneBook(ctx *fiber.Ctx) {

}
