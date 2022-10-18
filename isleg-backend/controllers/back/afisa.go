package controllers

import (
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OneAfisa struct {
	ID           string                    `json:"id"`
	Image        string                    `json:"image"`
	Translations []models.TranslationAfisa `json:"translations"`
}

func CreateAfisa(c *gin.Context) {

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

	var fileName string

	// GET ALL LANGUAGE
	languages, err := GetAllLanguageWithIDAndNameShort()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// FILE UPLOAD
	file, errFile := c.FormFile("image")
	if errFile != nil {
		fileName = ""
	} else {
		extension := filepath.Ext(file.Filename)
		// VALIDATE IMAGE
		if extension != ".jpg" && extension != ".jpeg" && extension != ".png" && extension != ".gif" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "the file must be an image.",
			})
			return
		}

		newFileName := uuid.New().String() + extension
		fileName = "uploads/afisa/" + newFileName
	}

	dataNames := []string{"title", "description"}

	// VALIDATE DATA
	if err = models.ValidateAfisaData(languages, dataNames, c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// create afisa
	resultAFisa, err := db.Query("INSERT INTO afisa (image) VALUES ($1)", fileName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultAFisa.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	if fileName != "" {
		c.SaveUploadedFile(file, "./"+fileName)
	}

	// get id of added afisa
	lastAfisaID, err := db.Query("SELECT id FROM afisa WHERE deleted_at IS NULL ORDER BY created_at DESC LIMIT 1")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := lastAfisaID.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var afisaID string

	for lastAfisaID.Next() {
		if err := lastAfisaID.Scan(&afisaID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	// create translation afisa
	for _, v := range languages {
		resultTRAfisa, err := db.Query("INSERT INTO translation_afisa (afisa_id,lang_id,title,description) VALUES ($1,$2,$3,$4)", afisaID, v.ID, c.PostForm("title_"+v.NameShort), c.PostForm("description"+v.NameShort))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := resultTRAfisa.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "afisa successfully added",
	})

}

func UpdateAfisaByID(c *gin.Context) {

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

	ID := c.Param("id")
	var fileName string

	rowAfisa, err := db.Query("SELECT id,image FROM afisa WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowAfisa.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var afisaID, image string

	for rowAfisa.Next() {
		if err := rowAfisa.Scan(&afisaID, &image); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if afisaID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	languages, err := GetAllLanguageWithIDAndNameShort()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	dataNames := []string{"title", "description"}

	// VALIDATE DATA
	if err = models.ValidateAfisaData(languages, dataNames, c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}

	// FILE UPLOAD
	file, errFile := c.FormFile("image")
	if errFile != nil {
		fileName = image
	} else {
		extension := filepath.Ext(file.Filename)
		// VALIDATE IMAGE
		if extension != ".jpg" && extension != ".jpeg" && extension != ".png" && extension != ".gif" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": "the file must be an image.",
			})
			return
		}

		newFileName := uuid.New().String() + extension
		fileName = "uploads/afisa/" + newFileName

		if image != "" {
			if err := os.Remove("./" + image); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}
	}

	currentTime := time.Now()

	resultAfisa, err := db.Query("UPDATE afisa SET image = $1 , updated_at = $2 WHERE id = $3", fileName, currentTime, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultAfisa.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	if fileName != "" {
		c.SaveUploadedFile(file, "./"+fileName)
	}

	for _, v := range languages {
		resultTRAfisa, err := db.Query("UPDATE translation_afisa SET title = $1 , description = $2 , updated_at = $3 WHERE afisa_id = $4 AND lang_id = $5", c.PostForm("title_"+v.NameShort), c.PostForm("description_"+v.NameShort), currentTime, ID, v.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := resultTRAfisa.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "afisa successfully updated",
	})

}

func GetAfisaByID(c *gin.Context) {

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

	ID := c.Param("id")

	rowAfisa, err := db.Query("SELECT id,image FROM afisa WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowAfisa.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var afisa OneAfisa

	for rowAfisa.Next() {
		if err := rowAfisa.Scan(&afisa.ID, &afisa.Image); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if afisa.ID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	rowsTrAfisa, err := db.Query("SELECT lang_id,title,description FROM translation_afisa WHERE afisa_id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowsTrAfisa.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var translations []models.TranslationAfisa

	for rowsTrAfisa.Next() {
		var translation models.TranslationAfisa
		if err := rowsTrAfisa.Scan(&translation.LangID, &translation.Title, &translation.Description); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		translations = append(translations, translation)
	}

	afisa.Translations = translations

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"afisa":  afisa,
	})

}

func GetAfisas(c *gin.Context) {

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

	rowsAfisa, err := db.Query("SELECT id,image FROM afisa WHERE deleted_at IS NULL")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowsAfisa.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var afisas []OneAfisa

	for rowsAfisa.Next() {
		var afisa OneAfisa
		if err := rowsAfisa.Scan(&afisa.ID, &afisa.Image); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}

		rowsTrAfisa, err := db.Query("SELECT lang_id,title,description FROM translation_afisa WHERE deleted_at IS NULL AND afisa_id = $1", afisa.ID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
		defer func() {
			if err := rowsTrAfisa.Close(); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
		}()

		var translations []models.TranslationAfisa

		for rowsTrAfisa.Next() {
			var translation models.TranslationAfisa
			if err := rowsTrAfisa.Scan(&translation.LangID, &translation.Title, &translation.Description); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  false,
					"message": err.Error(),
				})
				return
			}
			translations = append(translations, translation)
		}

		afisa.Translations = translations

		afisas = append(afisas, afisa)

	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"afisas": afisas,
	})

}

func DeleteAfisaByID(c *gin.Context) {

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

	ID := c.Param("id")

	rowAfisa, err := db.Query("SELECT id FROM afisa WHERE id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowAfisa.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var afisaID string

	for rowAfisa.Next() {
		if err := rowAfisa.Scan(&afisaID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if afisaID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	currentTime := time.Now()

	resultAfisa, err := db.Query("UPDATE afisa SET deleted_at = $1 WHERE id = $2", currentTime, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultAfisa.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	resultTrAfisa, err := db.Query("UPDATE translation_afisa SET deleted_at = $1 WHERE afisa_id = $2", currentTime, ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultTrAfisa.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "afisa successfully deleted",
	})

}

func RestoreAfisaByID(c *gin.Context) {

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

	ID := c.Param("id")

	rowAfisa, err := db.Query("SELECT id FROM afisa WHERE id = $1 AND deleted_at IS NOT NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := rowAfisa.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	var afisaID string

	for rowAfisa.Next() {
		if err := rowAfisa.Scan(&afisaID); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if afisaID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	resultAfisa, err := db.Query("UPDATE afisa SET deleted_at = NULL WHERE id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultAfisa.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	resultTrAfisa, err := db.Query("UPDATE translation_afisa SET deleted_at = NULL WHERE afisa_id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer func() {
		if err := resultTrAfisa.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "afisa successfully restored",
	})

}

func DeletePermanentlyAfisaByID(c *gin.Context) {

	db, err := config.ConnDB()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer db.Close()

	ID := c.Param("id")

	rowAfisa, err := db.Query("SELECT id,image FROM afisa WHERE id = $1 AND deleted_at IS NOT NULL", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer rowAfisa.Close()

	var afisaID, image string

	for rowAfisa.Next() {
		if err := rowAfisa.Scan(&afisaID, &image); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	if afisaID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": "record not found",
		})
		return
	}

	if image != "" {
		if err := os.Remove("./" + image); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  false,
				"message": err.Error(),
			})
			return
		}
	}

	resultAfisa, err := db.Query("DELETE FROM afisa WHERE id = $1", ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err.Error(),
		})
		return
	}
	defer resultAfisa.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "afisa successfully deleted",
	})

}
