package middlewares

import (
	"encoding/json"
	"errors"
	"strings"

	g "github.com/bdn/jeker/db"
	r "github.com/bdn/jeker/dto/responses"
	"github.com/bdn/jeker/models"
	"github.com/bdn/jeker/utils"
	"github.com/gofiber/fiber"
)

var (
	errAuthnUser = errors.New("Invalid Access Token")
)

type LocalUserData struct {
	ID       int
	Username string
	Email    string
}

func UserAuthentication(ctx *fiber.Ctx) {
	bearer := ctx.Get("Authorization")
	accessToken := strings.Split(bearer, " ")
	if len(accessToken) != 2 || accessToken[0] != "Bearer" {
		response := r.ResponsesHTTP{
			StatusCode: 401,
			Message:    "Unauthorized",
			Data:       errAuthnUser.Error(),
		}
		sendData, _ := json.Marshal(response)
		ctx.Status(401).Send(sendData)
		return
	}

	payload, err := utils.VerifyJWT(accessToken[1])
	if err != nil {
		response := r.ResponsesHTTP{
			StatusCode: 401,
			Message:    "Unauthorized",
			Data:       errAuthnUser.Error(),
		}
		sendData, _ := json.Marshal(response)
		ctx.Status(401).Send(sendData)
		return
	}

	var db = g.GetConn()
	var targetUser models.User
	result := db.Where("username = ?", payload.Username).Where("email = ?", payload.Email).Find(&targetUser)
	err = result.Error
	if err != nil || targetUser.ID == 0 {
		response := r.ResponsesHTTP{
			StatusCode: 401,
			Message:    "Unauthorized",
			Data:       errAuthnUser.Error(),
		}
		sendData, _ := json.Marshal(response)
		ctx.Status(401).Send(sendData)
		return
	}
	userData := LocalUserData{
		ID:       targetUser.ID,
		Username: targetUser.Username,
		Email:    targetUser.Email,
	}
	ctx.Locals("user", userData)
	ctx.Next()
}
