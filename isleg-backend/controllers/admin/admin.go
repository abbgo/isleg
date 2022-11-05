package controllers

import (
	"github/abbgo/isleg/isleg-backend/auth"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"

	frontController "github/abbgo/isleg/isleg-backend/controllers/front"
)

func RegisterAdmin(c *gin.Context) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var admin models.Admin

	if err := c.BindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := models.ValidateRegisterAdmin(admin.PhoneNumber, admin.Type); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	hashPassword, err := models.HashPassword(admin.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	resultAdmin, err := db.Query("INSERT INTO admins (full_name,phone_number,password,type) VALUES ($1,$2,$3,$4)", admin.FullName, admin.PhoneNumber, hashPassword, admin.Type)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultAdmin.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"status":       true,
		"phone_number": admin.PhoneNumber,
		"password":     admin.Password,
	})

}

func LoginAdmin(c *gin.Context) {

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var admin frontController.Login

	if err := c.BindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// check if email exists and password is correct
	row, err := db.Query("SELECT id,password,type FROM admins WHERE phone_number = $1 AND deleted_at IS NULL", admin.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := row.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var adminID, oldPassword, adminType string

	for row.Next() {
		if err := row.Scan(&adminID, &oldPassword, &adminType); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if adminID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "this admin does not exist",
		})
		return
	}

	credentialError := models.CheckPassword(admin.Password, oldPassword)
	if credentialError != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "invalid credentials",
		})
		return
	}

	accessTokenString, err := auth.GenerateAccessTokenForAdmin(admin.PhoneNumber, adminID, adminType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	refreshTokenString, err := auth.GenerateRefreshTokenForAdmin(admin.PhoneNumber, adminID, adminType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessTokenString,
		"refresh_token": refreshTokenString,
	})

}

func UpdateAdminInformation(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var admin models.Admin

	if err := c.BindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	rowAdmin, err := db.Query("SELECT id FROM admins WHERE id = $1 AND deleted_at IS NULL", admin.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowAdmin.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var admin_id string

	for rowAdmin.Next() {
		if err := rowAdmin.Scan(&admin_id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if admin_id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "admin not found",
		})
		return
	}

	resultCustomer, err := db.Query("UPDATE admins SET full_name = $1 , phone_number = $2, type = $3 WHERE id = $4", admin.FullName, admin.PhoneNumber, admin.Type, admin.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultCustomer.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully updated",
	})

}
