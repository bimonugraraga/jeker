package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/bdn/jeker/dto/requests"
	r "github.com/bdn/jeker/dto/responses"
	"github.com/bdn/jeker/middlewares"
	"github.com/bdn/jeker/models"
	services "github.com/bdn/jeker/services"
	"github.com/bdn/jeker/utils"
	"github.com/gofiber/fiber"
)

func CreateUpdateProfile(ctx *fiber.Ctx) {
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

	inputBody := new(requests.ProfileRequest)
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
	if inputBody.Bio == "" && inputBody.ProfilePictures == "" {
		response = r.ResponsesHTTP{
			StatusCode: 200,
			Message:    "Nothing To Update",
		}
		sendData, _ := json.Marshal(response)
		ctx.Status(200).Send(sendData)
		return
	}

	payload := models.Profile{
		UserId:          userData.ID,
		Bio:             &inputBody.Bio,
		ProfilePictures: &inputBody.ProfilePictures,
	}

	profile, err := services.CreateUpdateProfileService(payload)
	if err != nil {
		response = r.ResponsesHTTP{
			StatusCode: 400,
			Message:    "Failed To Update Or Create Profile",
			Data:       err.Error(),
		}
		sendData, _ := json.Marshal(response)
		ctx.Status(400).Send(sendData)
		return
	}
	response = r.ResponsesHTTP{
		StatusCode: 200,
		Message:    "Success To Update Or Create Profile",
		Data:       profile,
	}
	sendData, _ := json.Marshal(response)
	ctx.Status(200).Send(sendData)
	return
}

func GetOneProfile(ctx *fiber.Ctx) {
	var response r.ResponsesHTTP
	userId, err := strconv.Atoi(ctx.Params("user_id"))
	if err != nil {
		response = r.ResponsesHTTP{
			StatusCode: 400,
			Message:    "Failed To Process Request",
		}
		sendData, _ := json.Marshal(response)
		ctx.Status(400).Send(sendData)
		return
	}
	result, err := services.GetOneProfileService(userId)
	if err != nil || result.User.ID == 0 {
		response = r.ResponsesHTTP{
			StatusCode: 404,
			Message:    "User Profile Not Found",
		}
		sendData, _ := json.Marshal(response)
		ctx.Status(404).Send(sendData)
		return
	}
	readyResp := r.UserProfileResponses{
		UserId:          result.User.ID,
		Username:        result.User.Username,
		Bio:             result.Bio,
		ProfilePictures: result.ProfilePictures,
	}
	response = r.ResponsesHTTP{
		StatusCode: 200,
		Message:    "Success Get User Profile",
		Data:       readyResp,
	}
	sendData, _ := json.Marshal(response)
	ctx.Status(404).Send(sendData)
	return
}

func GetListOfProfile(ctx *fiber.Ctx) {
	var response r.ResponsesHTTP
	paginationQuery := new(requests.GeneralPaginationQuery)
	profileFilter := new(requests.ProfileFilter)
	err := ctx.QueryParser(paginationQuery)
	if err != nil {
		response = r.ResponsesHTTP{
			StatusCode: 400,
			Message:    "Failed To Process Request",
		}
		sendData, _ := json.Marshal(response)
		ctx.Status(400).Send(sendData)
		return
	}
	defaultPagination := utils.DefaultPagination(*paginationQuery)
	err = ctx.QueryParser(profileFilter)
	if err != nil {
		response = r.ResponsesHTTP{
			StatusCode: 400,
			Message:    "Failed To Process Request",
		}
		sendData, _ := json.Marshal(response)
		ctx.Status(400).Send(sendData)
		return
	}
	list, err := services.GetListProfileService(defaultPagination, *profileFilter)
	if err != nil {
		response = r.ResponsesHTTP{
			StatusCode: 500,
			Message:    "Internal Server Error",
		}
		sendData, _ := json.Marshal(response)
		ctx.Status(400).Send(sendData)
		return
	}

	var result []r.UserProfileResponses
	for _, val := range list {
		temp := r.UserProfileResponses{
			UserId:          val.UserId,
			Username:        val.User.Username,
			Bio:             val.Bio,
			ProfilePictures: val.ProfilePictures,
		}
		result = append(result, temp)
	}

	response = r.ResponsesHTTP{
		StatusCode: 200,
		Message:    "Success Fetch User Profile List",
		Data:       result,
	}
	sendData, _ := json.Marshal(response)
	ctx.Status(200).Send(sendData)
	return

}
