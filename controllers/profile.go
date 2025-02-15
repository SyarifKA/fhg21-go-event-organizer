package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/SyarifKA/fgh21-go-event-organizer/dtos"
	"github.com/SyarifKA/fgh21-go-event-organizer/lib"
	"github.com/SyarifKA/fgh21-go-event-organizer/models"
	"github.com/SyarifKA/fgh21-go-event-organizer/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func DataProfile(ctx *gin.Context) {
	userId := ctx.GetInt("userId")
	fmt.Println(userId)
	dataProfile := repository.FindProfileByUserId(userId)

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Detail Profile",
		Results: dataProfile,
	})
}

func UpdateProfile(ctx *gin.Context) {
	id := ctx.GetInt("userId")
	form := dtos.Profiles{}
	err := ctx.Bind(&form)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(profile)
	userUpdated, _ := repository.EditUser(models.UserUpdate{
		Email:    *form.Email,
		Username: form.Username,
	}, id)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	repository.EditProfile(models.UpdateProfile{
		FullName:      form.FullName,
		PhoneNumber:   form.PhoneNumber,
		Gender:        form.Gender,
		Profession:    *form.Profession,
		NationalityId: form.NationalityId,
		UserId:        id,
	})

	// if profileUpdated.Id == 0 {
	//     ctx.JSON(http.StatusNotFound, lib.Response{
	//         Success: false,
	//         Message: "Profile not found",
	//     })
	//     return
	// }
	// fmt.Println(profileUpdated)

	ctx.JSON(http.StatusOK, lib.Response{
		Success: true,
		Message: "Profile updated",
		Results: userUpdated,
	})
}

func UploadProfileImage(c *gin.Context) {
	id := c.GetInt("userId")

	maxFile := 500 * 1024
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, int64(maxFile))

	file, err := c.FormFile("profileImg")
	if err != nil {
		if err.Error() == "http: request body too large" {
			lib.HandlerMaxFile(c, "file size too large, max capacity 500 kb")
			return
		}
		lib.HandlerBadReq(c, "not file to upload")
		return
	}
	if id == 0 {
		lib.HandlerBadReq(c, "User not found")
		return
	}

	allowExt := map[string]bool{".jpg": true, ".jpeg": true, ".png": true}
	fileExt := strings.ToLower(filepath.Ext(file.Filename))
	if !allowExt[fileExt] {
		lib.HandlerBadReq(c, "extension file not validate")
		return
	}

	newFile := uuid.New().String() + fileExt

	uploadDir := "./img/profile/"
	if err := c.SaveUploadedFile(file, uploadDir+newFile); err != nil {
		lib.HandlerBadReq(c, "upload failed")
		return
	}

	tes := "/img/profile/" + newFile

	delImgBefore := repository.FindProfileByUserId(id)
	if delImgBefore.Picture != nil {
		// fileDel := strings.Split(*delImgBefore.Picture, "8888")[1]
		// os.Remove("." + fileDel)
		os.Remove(*delImgBefore.Picture)
	}

	profile, err := repository.UpdateProfileImage(models.Profile{Picture: &tes}, id)
	if err != nil {
		lib.HandlerBadReq(c, "upload failed")
		return
	}

	lib.HandlerOK(c, "Upload success", profile, nil)
}

func FindProfileByUserId(c *gin.Context) {
	id := c.GetInt("userId")
	if id == 0 {
		id, _ = strconv.Atoi(c.Param("id"))
	}

	profile := repository.FindProfileByUserId(id)
	fmt.Println(id)

	// if err != nil {
	// 	lib.HandlerBadReq(c, "Profile not found ")
	// 	return
	// }

	lib.HandlerOK(c, "Success Find Profile By UserId", profile, nil)
}
