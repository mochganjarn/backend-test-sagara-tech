package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mochganjarn/go-template-project/common/request"
	"github.com/mochganjarn/go-template-project/external/db/model"
	"github.com/mochganjarn/go-template-project/service"
)

func Login(appDependencies *service.ClientConnection) gin.HandlerFunc {
	return func(c *gin.Context) {

		db := appDependencies.DbClient.DbConnection
		var reqBody request.Login
		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//instantiate user model for query to database
		user := model.User{
			Username: reqBody.Username,
		}

		db.AutoMigrate(&user)
		result := db.Where(&user).First(&user)

		if result.RowsAffected == 0 {
			c.PureJSON(http.StatusNotFound, gin.H{
				"Status": false,
				"Error":  result.Error,
			})
		}

		//check if inputed password is match with password from database
		if !service.CheckPasswordHash(reqBody.Password, user.Password) {
			c.PureJSON(http.StatusUnauthorized, gin.H{
				"Status":  false,
				"Message": "Invalid Password",
			})
			return
		}

		token, err := service.JwtTokenGenerator(appDependencies, user.ID)
		if err != nil {
			c.PureJSON(http.StatusBadRequest, gin.H{
				"Result": "Failed Generate Token",
			})
			return
		}
		c.PureJSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}

func Register(appDependencies *service.ClientConnection) gin.HandlerFunc {
	return func(c *gin.Context) {

		var reqBody request.Login
		if err := c.ShouldBindJSON(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		hashedPassword, err := service.HashPassword(reqBody.Password)
		if err != nil {
			c.PureJSON(500, gin.H{
				"Status": false,
				"Error":  err,
			})
		}

		user := model.User{
			Username: reqBody.Username,
			Password: hashedPassword,
		}

		if err := model.CreateData(&user, appDependencies.DbClient); err != nil {
			c.PureJSON(400, gin.H{
				"Status": false,
				"Error":  err,
			})
		} else {
			c.PureJSON(201, gin.H{
				"Status": true,
				"Result": "Successfully created data",
			})
		}
	}
}
