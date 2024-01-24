package services

import (
	"errors"

	g "github.com/bdn/jeker/db"
	"github.com/bdn/jeker/dto/responses"
	"github.com/bdn/jeker/models"
	"github.com/bdn/jeker/utils"
)

var (
	errPassword = errors.New("Invalid Password")
	errLogin    = errors.New("Invalid Password or Email/Username")
)

func UserRegisterService(params models.User) (models.User, error) {
	var User models.User
	var err error
	hashPassword, err := utils.HashPassword(params.Password)
	if err != nil {
		return User, errPassword
	}
	params.Password = hashPassword
	var db = g.GetConn()
	result := db.Model(&User).Create(&params)
	err = result.Error
	if err != nil {
		return User, err
	}
	return params, nil
}

func UserLoginService(params models.User) (string, error) {
	var targetUser models.User
	var err error
	var db = g.GetConn()
	result := db.Where("username = ?", params.Username).Or("email = ?", params.Email).Find(&targetUser)
	err = result.Error
	if err != nil || targetUser.ID == 0 {
		return "", errLogin
	}

	isValidPass := utils.VerifyPassword(params.Password, targetUser.Password)
	if !isValidPass {
		return "", errLogin
	}

	accessToken, err := utils.SignJWT(responses.JWTTokenFormat{
		Username: targetUser.Username,
		Email:    targetUser.Email,
	})
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
