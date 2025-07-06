package controller

import (
	"database/sql"
	"net/http"
	"weeklytickits/dto"
	"weeklytickits/models"
	"weeklytickits/utils"

	"github.com/gin-gonic/gin"
)

func MapProfileRequestToDB(req dto.ProfileRequest) dto.Profile {
	return dto.Profile{
		Fullname: toNullString(req.Fullname),
		Phone:    toNullString(req.Phone),
		Image:    toNullString(req.Image),
	}
}

func toNullString(s *string) sql.NullString {
	if s != nil {
		return sql.NullString{String: *s, Valid: true}
	}
	return sql.NullString{Valid: false}
}

func GetUserProfile(ctx *gin.Context) {
	userId, _ := ctx.Get("userID")
	data, err := models.GetProfileByUserId(userId.(int))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid User ID",
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "OK",
		Results: data,
	})
}
func UpdateProfileByUserId(ctx *gin.Context) {
	userId, _ := ctx.Get("userID")
	var req dto.ProfileRequest

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid Input",
			Error:   err.Error(),
		})
		return
	}
	profile := MapProfileRequestToDB(req)

	err = models.UpdateProfile(userId.(int), profile)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.Response{
			Success: false,
			Message: "ERROR",
			Error:   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, utils.Response{
		Success: true,
		Message: "Sucessfully Update Profile",
		Results: req,
	})
}
