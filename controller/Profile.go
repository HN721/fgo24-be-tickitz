package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"weeklytickits/dto"
	"weeklytickits/models"
	"weeklytickits/utils"

	"github.com/gin-gonic/gin"
)

func toNullString(s *string) sql.NullString {
	if s != nil {
		return sql.NullString{String: *s, Valid: true}
	}
	return sql.NullString{Valid: false}
}

// GetUserProfile godoc
// @Summary Get User Profile
// @Description Retrieve the profile data of the authenticated user
// @Tags Profile
// @Produce json
// @Success 200 {object} utils.Response{results=dto.Profile}
// @Failure 400 {object} utils.Response
// @Security Token
// @Router /profile [get]
func GetUserProfile(ctx *gin.Context) {
	userId, _ := ctx.Get("userID")
	fmt.Println(userId.(int))
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

// UpdateProfileByUserId godoc
// @Summary Update User Profile
// @Description Update the profile of the authenticated user
// @Tags Profile
// @Accept json
// @Produce json
// @Param profile body dto.ProfileRequest true "Profile Request Body"
// @Success 200 {object} utils.Response{results=dto.ProfileRequest}
// @Failure 400 {object} utils.Response
// @Failure 500 {object} utils.Response
// @Security Token
// @Router /profile [patch]
func UpdateProfileByUserId(ctx *gin.Context) {
	userId, _ := ctx.Get("userID")
	var req dto.Profile

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.Response{
			Success: false,
			Message: "Invalid Input",
			Error:   err.Error(),
		})
		return
	}

	err = models.UpdateProfile(userId.(int), req)
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
