package controllers

import (
	"context"
	"github/abbgo/isleg/isleg-backend/config"
	"github/abbgo/isleg/isleg-backend/helpers"
	"github/abbgo/isleg/isleg-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateNotification(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var notification models.Notification
	if err := c.BindJSON(&notification); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	for _, v := range notification.Translations {
		var lang_id string
		if err := db.QueryRow(context.Background(), "SELECT id FROM languages WHERE id = $1 AND deleted_at IS NULL", v.LangID).Scan(&lang_id); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		if lang_id == "" {
			helpers.HandleError(c, 404, "langauge not found")
			return
		}
	}

	var notification_id string
	if err := db.QueryRow(context.Background(), "INSERT INTO notifications (name) VALUES ($1) RETURNING id", notification.Name).Scan(&notification_id); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	for _, v := range notification.Translations {
		_, err := db.Exec(context.Background(), "INSERT INTO translation_notification (notification_id,lang_id,translation) VALUES ($1,$2,$3)", notification_id, v.LangID, v.Translation)
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

func UpdateNotificationByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var notification models.Notification
	if err := c.BindJSON(&notification); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var notification_id string
	if err := db.QueryRow(context.Background(), "SELECT id FROM notifications WHERE id = $1 AND deleted_at IS NULL", notification.ID).Scan(&notification_id); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if notification_id == "" {
		helpers.HandleError(c, 404, "notification not found")
		return
	}

	_, err = db.Exec(context.Background(), "UPDATE notifications SET name = $1 WHERE id = $2", notification.Name, notification.ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	for _, v := range notification.Translations {
		_, err := db.Exec(context.Background(), "UPDATE translation_notification SET translation = $1 WHERE lang_id = $2 AND notification_id = $3", v.Translation, v.LangID, notification_id)
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

func GetNotificationByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	ID := c.Param("id")

	var notification models.Notification
	rowNotification, err := db.Query(context.Background(), "SELECT id,name FROM notifications WHERE id = $1", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	for rowNotification.Next() {
		if err := rowNotification.Scan(&notification.ID, &notification.Name); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		if notification.ID == "" {
			helpers.HandleError(c, 404, "notification not found")
			return
		}

		rowsTrNotification, err := db.Query(context.Background(), "SELECT translation FROM translation_notification WHERE notification_id = $1 AND deleted_at IS NULL", notification.ID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		var trNotifications []models.TranslationNotification
		for rowsTrNotification.Next() {
			var trNotification models.TranslationNotification
			if err := rowsTrNotification.Scan(&trNotification.Translation); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
			trNotifications = append(trNotifications, trNotification)
		}
		notification.Translations = trNotifications
	}

	c.JSON(http.StatusOK, gin.H{
		"status":       true,
		"notification": notification,
	})
}

func GetNotifications(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	rowsNotification, err := db.Query(context.Background(), "SELECT id,name FROM notifications WHERE deleted_at IS NULL")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var notifications []models.Notification
	for rowsNotification.Next() {
		var notification models.Notification
		if err := rowsNotification.Scan(&notification.ID, &notification.Name); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		rowsTrNotification, err := db.Query(context.Background(), "SELECT translation FROM translation_notification WHERE notification_id = $1 AND deleted_at IS NULL", notification.ID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		var trNotifications []models.TranslationNotification
		for rowsTrNotification.Next() {
			var trNotification models.TranslationNotification
			if err := rowsTrNotification.Scan(&trNotification.Translation); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
			trNotifications = append(trNotifications, trNotification)
		}
		notification.Translations = trNotifications
		notifications = append(notifications, notification)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":        true,
		"notifications": notifications,
	})
}

func DeleteNotificationByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	ID := c.Param("id")

	var notification_id string
	if err := db.QueryRow(context.Background(), "SELECT id FROM notifications WHERE id = $1 AND deleted_at IS NULL", ID).Scan(&notification_id); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if notification_id == "" {
		helpers.HandleError(c, 404, "notification not found")
		return
	}

	_, err = db.Exec(context.Background(), "CALL delete_notification($1)", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully deleted",
	})
}

func RestoreNotificationByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	ID := c.Param("id")

	var notification_id string
	if err := db.QueryRow(context.Background(), "SELECT id FROM notifications WHERE id = $1 AND deleted_at IS NOT NULL", ID).Scan(&notification_id); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if notification_id == "" {
		helpers.HandleError(c, 404, "notification not found")
		return
	}

	_, err = db.Exec(context.Background(), "CALL restore_notification($1)", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully restored",
	})
}

func DeletePermanentlyNotificationByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	ID := c.Param("id")

	var notification_id string
	if err := db.QueryRow(context.Background(), "SELECT id FROM notifications WHERE id = $1 AND deleted_at IS NOT NULL", ID).Scan(&notification_id); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if notification_id == "" {
		helpers.HandleError(c, 404, "notification not found")
		return
	}

	_, err = db.Exec(context.Background(), "DELETE FROM notifications WHERE id = $1", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully deleted",
	})
}

func GetNotificationByLangID(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// GET DATA FROM ROUTE PARAMETER
	langShortName := c.Param("lang")

	// GET language id
	langID, err := GetLangID(langShortName)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	rowsNotification, err := db.Query(context.Background(), "SELECT id,name FROM notifications WHERE deleted_at IS NULL")
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var notifications []models.Notification
	for rowsNotification.Next() {
		var notification models.Notification
		if err := rowsNotification.Scan(&notification.ID, &notification.Name); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		rowTrNotification, err := db.Query(context.Background(), "SELECT translation FROM translation_notification WHERE notification_id = $1 AND lang_id = $2 AND deleted_at IS NULL", notification.ID, langID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		var trNotifications []models.TranslationNotification
		for rowTrNotification.Next() {
			var trNotification models.TranslationNotification
			if err := rowTrNotification.Scan(&trNotification.Translation); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
			trNotifications = append(trNotifications, trNotification)
		}
		notification.Translations = trNotifications
		notifications = append(notifications, notification)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":        true,
		"notifications": notifications,
	})
}
