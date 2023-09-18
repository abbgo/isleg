package controllers

import (
	"context"
	"errors"
	"github/abbgo/isleg/isleg-backend/auth"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/helpers"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	frontController "github/abbgo/isleg/isleg-backend/controllers/front"
)

func RegisterAdmin(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var admin models.Admin
	if err := c.BindJSON(&admin); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if err := models.ValidateRegisterAdmin(admin.PhoneNumber, admin.Type); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	hashPassword, err := models.HashPassword(admin.Password)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	_, err = db.Exec(context.Background(), "INSERT INTO admins (full_name,phone_number,password,type) VALUES ($1,$2,$3,$4)", admin.FullName, admin.PhoneNumber, hashPassword, admin.Type)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":       true,
		"phone_number": admin.PhoneNumber,
		"password":     admin.Password,
		"admin_type":   admin.Type,
	})
}

func LoginAdmin(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var admin frontController.Login

	if err := c.BindJSON(&admin); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// check if email exists and password is correct
	var adminID, oldPassword, adminType string
	db.QueryRow(context.Background(), "SELECT id,password,type FROM admins WHERE phone_number = $1 AND deleted_at IS NULL", admin.PhoneNumber).Scan(&adminID, &oldPassword, &adminType)
	if adminID == "" {
		helpers.HandleError(c, 400, "this admin does not exist")
		return
	}

	credentialError := models.CheckPassword(admin.Password, oldPassword)
	if credentialError != nil {
		helpers.HandleError(c, 400, "invalid credentials")
		return
	}

	accessTokenString, refreshTokenString, err := auth.GenerateAccessTokenForAdmin(admin.PhoneNumber, adminID, adminType)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	adm, err := GetAdminByID(adminID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessTokenString,
		"refresh_token": refreshTokenString,
		"admin_type":    adminType,
		"admin":         adm,
	})
}

func GetAdmin(c *gin.Context) {
	adminID, hasAdminID := c.Get("admin_id")
	if !hasAdminID {
		helpers.HandleError(c, 400, "adminID is required")
		return
	}

	var ok bool
	admin_id, ok := adminID.(string)
	if !ok {
		helpers.HandleError(c, 400, "admin_id must be uint")
		return
	}

	adm, err := GetAdminByID(admin_id)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"admin": adm,
	})
}

func GetAdminByID(id string) (models.Admin, error) {
	db, err := config.ConnDB()
	if err != nil {
		return models.Admin{}, err
	}
	defer db.Close()

	var admin models.Admin
	db.QueryRow(context.Background(), "SELECT full_name,phone_number FROM admins WHERE deleted_at IS NULL AND id = $1", id).Scan(&admin.FullName, &admin.PhoneNumber)
	if admin.PhoneNumber == "" {
		return models.Admin{}, errors.New("admin not found")
	}

	return admin, nil
}

func UpdateAdminInformation(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var admin models.Admin
	if err := c.BindJSON(&admin); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var admin_id string
	db.QueryRow(context.Background(), "SELECT id FROM admins WHERE id = $1 AND deleted_at IS NULL", admin.ID).Scan(&admin_id)
	if admin_id == "" {
		helpers.HandleError(c, 404, "admin not found")
		return
	}

	_, err = db.Exec(context.Background(), "UPDATE admins SET full_name = $1 , phone_number = $2, type = $3 WHERE id = $4", admin.FullName, admin.PhoneNumber, admin.Type, admin.ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully updated",
	})
}

func UpdateAdminPassword(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var admin models.Admin
	if err := c.BindJSON(&admin); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var admin_id string
	db.QueryRow(context.Background(), "SELECT id FROM admins WHERE id = $1 AND deleted_at IS NULL", admin.ID).Scan(&admin_id)
	if admin_id == "" {
		helpers.HandleError(c, 404, "admin not found")
		return
	}

	hashPassword, err := models.HashPassword(admin.Password)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	_, err = db.Exec(context.Background(), "UPDATE admins SET password = $1 WHERE id = $2", hashPassword, admin.ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "password of admin successfuly updated",
	})
}

func GetAdmins(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get limit from param
	limitStr := c.Param("limit")
	if limitStr == "" {
		helpers.HandleError(c, 400, "limit is required")
		return
	}
	limit, err := strconv.ParseUint(limitStr, 10, 64)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// get page from param
	pageStr := c.Param("page")
	if pageStr == "" {
		helpers.HandleError(c, 400, "page is required")
		return
	}
	page, err := strconv.ParseUint(pageStr, 10, 64)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	offset := limit * (page - 1)

	countOfAdmins := 0
	db.QueryRow(context.Background(), "SELECT COUNT(id) FROM admins WHERE deleted_at IS NULL").Scan(&countOfAdmins)
	var admins []models.Admin
	rowsAdmin, err := db.Query(context.Background(), "SELECT full_name,phone_number FROM admins WHERE deleted_at IS NULL LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowsAdmin.Close()

	var admin models.Admin
	for rowsAdmin.Next() {
		if err := rowsAdmin.Scan(&admin.FullName, &admin.PhoneNumber); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		admins = append(admins, admin)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":          true,
		"admins":          admins,
		"count_of_admins": countOfAdmins,
	})
}
