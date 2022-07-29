package service

import (
	"github.com/mochganjarn/go-template-project/config"
	"github.com/mochganjarn/go-template-project/external/db"
	jwtclient "github.com/mochganjarn/go-template-project/external/jwt_client"
)

//For Collecting Client Conection
type ClientConnection struct {
	DbClient *db.Client
	AppPort  string
	jwtclient.JwtSecret
}

//To instantiate application dependencies
func InstantiateDependencies(appConfig *config.Config) *ClientConnection {
	return &ClientConnection{
		// init database connection
		DbClient: db.InitDatabase(
			db.DBConfig{DBName: appConfig.DBName,
				DBUser:     appConfig.DBUser,
				DBHost:     appConfig.DBHost,
				DBPassword: appConfig.DBPassword,
				DBPort:     appConfig.DBPort,
			}),
		// init app port
		AppPort: appConfig.Port,
		// save jwt secret to struct
		JwtSecret: jwtclient.JwtSecret{
			MySecret: appConfig.JWTSecret,
		},
	}
}
