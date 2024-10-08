package controllers

// func CreateDistrict(c *gin.Context) {

// 	// initialize database connection
// 	db, err := config.ConnDB()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	defer func() {
// 		if err := db.Close(); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}()

// 	// GET ALL LANGUAGE
// 	languages, err := GetAllLanguageWithIDAndNameShort()
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	// validate data from request
// 	priceStr := c.PostForm("price")
// 	price, err := strconv.ParseFloat(priceStr, 64)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	for _, v := range languages {
// 		if c.PostForm("name_"+v.NameShort) == "" {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": "name_" + v.NameShort + " is required",
// 			})
// 			return
// 		}
// 	}

// 	// create district
// 	resultDistrict, err := db.Query("INSERT INTO district (price) VALUES ($1) RETURNING id", price)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  false,
// 			"message": err.Error(),
// 		})
// 		return
// 	}
// 	defer func() {
// 		if err := resultDistrict.Close(); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}()

// 	var districtID string

// 	for resultDistrict.Next() {
// 		if err := resultDistrict.Scan(&districtID); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 	}

// 	// create translation afisa
// 	for _, v := range languages {
// 		resultTRDistrict, err := db.Query("INSERT INTO translation_district (lang_id,district_id,name) VALUES ($1,$2,$3)", v.ID, districtID, c.PostForm("name_"+v.NameShort))
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"status":  false,
// 				"message": err.Error(),
// 			})
// 			return
// 		}
// 		defer func() {
// 			if err := resultTRDistrict.Close(); err != nil {
// 				c.JSON(http.StatusBadRequest, gin.H{
// 					"status":  false,
// 					"message": err.Error(),
// 				})
// 				return
// 			}
// 		}()
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  true,
// 		"message": "data successfully added",
// 	})

// }
