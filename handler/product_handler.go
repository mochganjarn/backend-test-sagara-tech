package handler

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mochganjarn/go-template-project/external/db/model"
	"github.com/mochganjarn/go-template-project/service"
)

func CreateProduct(appDependencies *service.ClientConnection) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, "get form err: %s", err.Error())
			return
		}

		extention := filepath.Ext(file.Filename)
		acceptedExtention := []string{".jpg", ".jpeg", ".png"}

		if !service.In_array(extention, acceptedExtention) {
			c.PureJSON(http.StatusBadRequest, gin.H{
				"Status":  false,
				"Message": "File type is denied",
			})
			return
		}

		newFileName := uuid.New().String() + extention

		name := c.PostForm("name")
		price, err := strconv.Atoi(c.PostForm("price"))
		if err != nil {
			c.PureJSON(http.StatusInternalServerError, gin.H{
				"Status": false,
				"Error":  err.Error(),
			})
			return
		}
		stock, err := strconv.Atoi(c.PostForm("stock"))
		if err != nil {
			c.PureJSON(http.StatusInternalServerError, gin.H{
				"Status": false,
				"Error":  err.Error(),
			})
			return
		}

		product := model.Product{
			Name:     name,
			Price:    price,
			Stock:    stock,
			Filename: newFileName,
		}

		if err := c.SaveUploadedFile(file, "uploads/"+newFileName); err != nil {
			c.PureJSON(http.StatusBadRequest, gin.H{
				"Status": false,
				"Error":  err,
			})
			return
		}

		dbconn := appDependencies.DbClient.DbConnection
		dbconn.AutoMigrate(&product)
		result := dbconn.Create(&product)

		if result.RowsAffected > 0 {
			c.PureJSON(http.StatusCreated, gin.H{
				"Data":   product,
				"Status": true,
				"Result": "Successfully created data",
				"Size":   file.Size,
			})
		} else {
			c.PureJSON(http.StatusBadRequest, gin.H{
				"Status": false,
				"Error":  result.Error,
			})
		}
	}
}

func GetProduct(appDependencies *service.ClientConnection) gin.HandlerFunc {
	return func(c *gin.Context) {
		dbconn := appDependencies.DbClient.DbConnection
		dbconn.AutoMigrate(&model.Product{})
		products := []model.Product{}
		result := dbconn.Find(&products)
		if result.RowsAffected == 0 {
			c.PureJSON(http.StatusNotFound, gin.H{
				"Status":  false,
				"Message": "not found",
			})
			return
		}

		c.PureJSON(http.StatusOK, gin.H{
			"Status": true,
			"Result": &products,
		})
	}
}

func ShowProduct(appDependencies *service.ClientConnection) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appDependencies.DbClient.DbConnection
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.PureJSON(http.StatusBadRequest, gin.H{
				"Status": false,
				"Error":  err.Error(),
			})
		}
		product := model.Product{ID: uint(id)}

		db.AutoMigrate(&product)
		result := db.Where(&product).First(&product)

		if result.RowsAffected == 0 {
			c.PureJSON(http.StatusNotFound, gin.H{
				"Status":  false,
				"Message": "not found",
			})
			return
		}

		c.PureJSON(http.StatusOK, gin.H{
			"Status": true,
			"Result": product,
		})
	}
}

func UpdateProduct(appDependencies *service.ClientConnection) gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, "get form err: %s", err.Error())
			return
		}

		extention := filepath.Ext(file.Filename)
		acceptedExtention := []string{".jpg", ".jpeg", ".png"}

		if !service.In_array(extention, acceptedExtention) {
			c.PureJSON(http.StatusBadRequest, gin.H{
				"Status":  false,
				"Message": "File type is denied",
			})
			return
		}

		newFileName := uuid.New().String() + extention

		name := c.PostForm("name")
		price, err := strconv.Atoi(c.PostForm("price"))
		if err != nil {
			c.PureJSON(http.StatusInternalServerError, gin.H{
				"Status": false,
				"Error":  err.Error(),
			})
			return
		}
		stock, err := strconv.Atoi(c.PostForm("stock"))
		if err != nil {
			c.PureJSON(http.StatusInternalServerError, gin.H{
				"Status": false,
				"Error":  err.Error(),
			})
			return
		}
		db := appDependencies.DbClient.DbConnection
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.PureJSON(http.StatusBadRequest, gin.H{
				"Status": false,
				"Error":  err.Error(),
			})
		}

		product := model.Product{ID: uint(id)}
		db.AutoMigrate(&product)
		result := db.First(&product)
		if result.RowsAffected == 0 {
			c.PureJSON(http.StatusNotFound, gin.H{
				"Status":  false,
				"Message": "Product Not Found",
			})
			return
		}

		oldFile := product.Filename
		dir := "uploads/"
		err = os.Remove(dir + oldFile)

		if err != nil {
			c.PureJSON(http.StatusInternalServerError, gin.H{
				"Status": false,
				"Error":  err.Error(),
			})
			return
		}

		db.Model(&product).Updates(model.Product{
			Name:     name,
			Price:    price,
			Stock:    stock,
			Filename: newFileName,
		})

		if err := c.SaveUploadedFile(file, "uploads/"+newFileName); err != nil {
			c.PureJSON(http.StatusBadRequest, gin.H{
				"Status": false,
				"Error":  err,
			})
			return
		}

		c.PureJSON(http.StatusOK, gin.H{
			"Status":   true,
			"Result":   product,
			"Message":  "Successfuly updated",
			"Old file": oldFile,
		})

	}
}

func DeleteProduct(appDependencies *service.ClientConnection) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appDependencies.DbClient.DbConnection
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.PureJSON(http.StatusBadRequest, gin.H{
				"Status": false,
				"Error":  err.Error(),
			})
		}
		product := model.Product{ID: uint(id)}

		db.AutoMigrate(&product)
		result := db.First(&product)
		if result.RowsAffected == 0 {
			c.PureJSON(http.StatusNotFound, gin.H{
				"Status":  false,
				"Message": "Product Not Found",
			})
			return
		}

		File := product.Filename
		dir := "uploads/"
		err = os.Remove(dir + File)

		if err != nil {
			c.PureJSON(http.StatusInternalServerError, gin.H{
				"Status": false,
				"Error":  err.Error(),
			})
			return
		}

		db.Delete(&product)

		c.PureJSON(http.StatusOK, gin.H{
			"Status":  true,
			"Result":  product,
			"Message": "Product Succesfully Deleted",
		})
	}
}
