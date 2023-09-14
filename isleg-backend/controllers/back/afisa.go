package controllers

import (
	"context"
	"fmt"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/helpers"
	"github/abbgo/isleg/isleg-backend/models"
	"github/abbgo/isleg/isleg-backend/pkg"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/jackc/pgx/v5"
)

type OneAfisa struct {
	ID           string                    `json:"id"`
	Image        string                    `json:"image"`
	Translations []models.TranslationAfisa `json:"translations"`
}

func CreateAfisa(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var afisa models.Afisa
	if err := c.BindJSON(&afisa); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// create afisa
	var afisaID string
	if err := db.QueryRow(context.Background(), "INSERT INTO afisa (image) VALUES ($1) RETURNING id", afisa.Image).Scan(&afisaID); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// create translation afisa
	for _, v := range afisa.TranslationAfisa {
		_, err := db.Exec(context.Background(), "INSERT INTO translation_afisa (afisa_id,lang_id,title,description,slug) VALUES ($1,$2,$3,$4,$5)", afisaID, v.LangID, v.Title, v.Description, slug.MakeLang(v.Title, "en"))
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully added",
	})
}

func UpdateAfisaByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from request parameter
	ID := c.Param("id")

	// check id and get image of afisa
	var afisaID, image string
	if err := db.QueryRow(context.Background(), "SELECT id,image FROM afisa WHERE id = $1 AND deleted_at IS NULL", ID).Scan(&afisaID, &image); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if afisaID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	var afisa models.Afisa
	if err := c.BindJSON(&afisa); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var fileName string
	if afisa.Image == "" {
		fileName = image
	} else {
		fileName = afisa.Image
	}

	_, err = db.Exec(context.Background(), "UPDATE afisa SET image = $1 WHERE id = $2", fileName, ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	for _, v := range afisa.TranslationAfisa {
		_, err := db.Exec(context.Background(), "UPDATE translation_afisa SET title = $1 , description = $2 , slug = $5 WHERE afisa_id = $3 AND lang_id = $4", v.Title, v.Description, ID, v.LangID, slug.MakeLang(v.Title, "en"))
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully updated",
	})
}

func GetAfisaByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from request parameter
	ID := c.Param("id")

	// check id and get data
	var afisa OneAfisa
	if err := db.QueryRow(context.Background(), "SELECT id,image FROM afisa WHERE id = $1 AND deleted_at IS NULL", ID).Scan(&afisa.ID, &afisa.Image); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if afisa.ID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	rowsTrAfisa, err := db.Query(context.Background(), "SELECT lang_id,title,description FROM translation_afisa WHERE afisa_id = $1 AND deleted_at IS NULL", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var translations []models.TranslationAfisa
	for rowsTrAfisa.Next() {
		var translation models.TranslationAfisa
		if err := rowsTrAfisa.Scan(&translation.LangID, &translation.Title, &translation.Description); err != nil {
			helpers.HandleError(c, 400, err.Error())
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
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	langID, err := GetLangID("tm")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

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
	var countOfAfisas uint

	searchQuery := c.Query("search")
	var searchStr, search string
	if searchQuery != "" {
		incomingsSarch := slug.MakeLang(searchQuery, "en")
		search = strings.ReplaceAll(incomingsSarch, "-", " | ")
		searchStr = fmt.Sprintf("%%%s%%", search)
	}

	statusQuery := c.DefaultQuery("status", "false")
	status, err := strconv.ParseBool(statusQuery)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var countAfisasQuery string
	if !status {
		if search == "" {
			countAfisasQuery = `SELECT COUNT(id) FROM afisa WHERE deleted_at IS NULL`
			if err = db.QueryRow(context.Background(), countAfisasQuery).Scan(&countOfAfisas); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		} else {
			countAfisasQuery = `SELECT COUNT(af.id) FROM afisa af INNER JOIN translation_afisa ta ON ta.afisa_id=af.id WHERE af.deleted_at IS NULL AND ta.deleted_at IS NULL AND ta.lang_id = $3 AND (to_tsvector(ta.slug) @@ to_tsquery($1) OR ta.slug LIKE $2)`
			if err = db.QueryRow(context.Background(), countAfisasQuery, search, searchStr, langID).Scan(&countOfAfisas); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		}
	} else {
		if search == "" {
			countAfisasQuery = `SELECT COUNT(id) FROM brends WHERE deleted_at IS NOT NULL`
			if err = db.QueryRow(context.Background(), countAfisasQuery).Scan(&countOfAfisas); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		} else {
			countAfisasQuery = `SELECT COUNT(af.id) FROM afisa af INNER JOIN translation_afisa ta ON ta.afisa_id=af.id WHERE af.deleted_at IS NOT NULL AND ta.deleted_at IS NOT NULL AND ta.lang_id = $3 AND (to_tsvector(ta.slug) @@ to_tsquery($1) OR ta.slug LIKE $2)`
			if err = db.QueryRow(context.Background(), countAfisasQuery, search, searchStr, langID).Scan(&countOfAfisas); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		}
	}

	var rowAfisasQuery string
	var rowAfisas pgx.Rows
	if !status {
		if search == "" {
			rowAfisasQuery = `SELECT id,image FROM afisa WHERE deleted_at IS NULL ORDER BY created_at DESC LIMIT $1 OFFSET $2`
			rowAfisas, err = db.Query(context.Background(), rowAfisasQuery, limit, offset)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		} else {
			rowAfisasQuery = `SELECT af.id,af.image FROM afisa af INNER JOIN translation_afisa ta ON ta.afisa_id=af.id WHERE af.deleted_at IS NULL AND ta.deleted_at IS NULL AND ta.lang_id = $5 AND (to_tsvector(ta.slug) @@ to_tsquery($3) OR ta.slug LIKE $4) ORDER BY af.created_at DESC LIMIT $1 OFFSET $2`
			rowAfisas, err = db.Query(context.Background(), rowAfisasQuery, limit, offset, search, searchStr, langID)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		}
	} else {
		if search == "" {
			rowAfisasQuery = `SELECT id,image FROM afisa WHERE deleted_at IS NOT NULL ORDER BY created_at DESC LIMIT $1 OFFSET $2`
			rowAfisas, err = db.Query(context.Background(), rowAfisasQuery, limit, offset)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		} else {
			rowAfisasQuery = `SELECT af.id,af.image FROM afisa af INNER JOIN translation_afisa ta ON ta.afisa_id=af.id WHERE af.deleted_at IS NOT NULL AND ta.deleted_at IS NOT NULL AND ta.lang_id = $5 AND (to_tsvector(ta.slug) @@ to_tsquery($3) OR ta.slug LIKE $4) ORDER BY af.created_at DESC LIMIT $1 OFFSET $2`
			rowAfisas, err = db.Query(context.Background(), rowAfisasQuery, limit, offset, search, searchStr, langID)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		}
	}

	var afisas []OneAfisa
	for rowAfisas.Next() {
		var afisa OneAfisa
		if err := rowAfisas.Scan(&afisa.ID, &afisa.Image); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		rowsTrAfisa, err := db.Query(context.Background(), "SELECT lang_id,title,description FROM translation_afisa WHERE lang_id = $2 AND afisa_id = $1", afisa.ID, langID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		var translations []models.TranslationAfisa

		for rowsTrAfisa.Next() {
			var translation models.TranslationAfisa
			if err := rowsTrAfisa.Scan(&translation.LangID, &translation.Title, &translation.Description); err != nil {
				helpers.HandleError(c, 400, err.Error())
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
		"total":  countOfAfisas,
	})
}

func DeleteAfisaByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from request parameter
	ID := c.Param("id")

	// check id
	var afisaID string
	if err := db.QueryRow(context.Background(), "SELECT id FROM afisa WHERE id = $1 AND deleted_at IS NULL", ID).Scan(&afisaID); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if afisaID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	_, err = db.Exec(context.Background(), "CALL delete_afisa($1)", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully deleted",
	})
}

func RestoreAfisaByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from request parameter
	ID := c.Param("id")

	// check id
	var afisaID string
	if err := db.QueryRow(context.Background(), "SELECT id FROM afisa WHERE id = $1 AND deleted_at IS NOT NULL", ID).Scan(&afisaID); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if afisaID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	_, err = db.Exec(context.Background(), "CALL restore_afisa($1)", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully restored",
	})
}

func DeletePermanentlyAfisaByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// get id from request parameter
	ID := c.Param("id")

	// check id and get image of afisa
	var afisaID, image string
	if err := db.QueryRow(context.Background(), "SELECT id,image FROM afisa WHERE id = $1 AND deleted_at IS NOT NULL", ID).Scan(&afisaID, &image); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if afisaID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	if image != "" {
		if err := os.Remove(pkg.ServerPath + image); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	_, err = db.Exec(context.Background(), "DELETE FROM afisa WHERE id = $1", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully deleted",
	})
}
