package services

import (
	"errors"
	"fmt"

	g "github.com/bdn/jeker/db"
	"github.com/bdn/jeker/dto/requests"
	"github.com/bdn/jeker/models"
)

var (
	errNotFound = errors.New("User Profile Not Found")
)

func CreateUpdateProfileService(params models.Profile) (models.Profile, error) {
	var targetProfile models.Profile
	var err error
	var db = g.GetConn()
	result := db.Where("user_id = ?", params.UserId).Find(&targetProfile)
	err = result.Error
	if err != nil {
		fmt.Println("Error Here")
	}

	if targetProfile.ID == 0 {
		newProfile, err := CreateNewProfileService(params)
		if err != nil {
			return params, err
		}
		return newProfile, nil
	}
	params.ID = targetProfile.ID
	updatedProfile, err := UpdateProfileService(params)
	if err != nil {
		return params, err
	}
	return updatedProfile, nil
}

func CreateNewProfileService(params models.Profile) (models.Profile, error) {
	var err error
	var db = g.GetConn()
	var Profile models.Profile
	result := db.Model(&Profile).Create(&params)
	err = result.Error
	if err != nil {
		return Profile, err
	}
	return params, nil
}

func UpdateProfileService(params models.Profile) (models.Profile, error) {
	var err error
	var db = g.GetConn()
	Profile := models.Profile{
		ID: params.ID,
	}
	result := db.Model(&Profile).Updates(params)
	err = result.Error
	if err != nil {
		return Profile, err
	}
	return params, nil
}

func GetOneProfileService(userId int) (models.Profile, error) {
	var targetUser models.User
	var targetProfile models.Profile
	var err error
	var db = g.GetConn()
	result := db.Joins("User").Where("user_id = ?", userId).Find(&targetProfile)
	err = result.Error
	if err != nil {
		return targetProfile, errNotFound
	}
	if targetProfile.ID != 0 {
		return targetProfile, nil
	}
	result = db.Where("id = ?", userId).Find(&targetUser)
	err = result.Error
	if err != nil || targetUser.ID == 0 {
		return targetProfile, errNotFound
	}
	targetProfile.User = targetUser
	return targetProfile, nil
}

func GetListProfileService(pagination requests.GeneralPaginationQuery, filter requests.ProfileFilter) ([]models.Profile, error) {
	var db = g.GetConn()
	var listProfiles []models.Profile
	var err error
	offset := (pagination.Page - 1) * pagination.Limit
	result := db.Joins("User").
		Limit(pagination.Limit).
		Offset(offset).
		Order(fmt.Sprintf("%s %s", pagination.OrderWith, pagination.OrderBy))

	if filter.Username != "" {
		result.Where("username = ?", filter.Username)
	}
	result.Find(&listProfiles)
	fmt.Println(listProfiles)
	err = result.Error
	if err != nil {
		return listProfiles, err
	}
	return listProfiles, nil
}
