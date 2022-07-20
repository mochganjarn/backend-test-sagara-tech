package service

import (
	"github.com/mochganjarn/go-template-project/config"
	"github.com/mochganjarn/go-template-project/external/db"
	"github.com/mochganjarn/go-template-project/external/gcs"
	jwtclient "github.com/mochganjarn/go-template-project/external/jwt_client"
	redisclient "github.com/mochganjarn/go-template-project/external/redis_client"
)

//For Collecting Client Conection
type ClientConnection struct {
	DbClient    *db.Client
	GcsClient   *gcs.Client
	RedisClient *redisclient.Client
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
		// init connection to gcs
		GcsClient: gcs.InitGcsClient(),
		// init redis connection
		RedisClient: redisclient.InitRedisCacheClient(
			redisclient.RedisConfig{
				REDISHost: appConfig.REDISHost,
				REDISPort: appConfig.REDISPort,
			}),
		// save jwt secret to struct
		JwtSecret: jwtclient.JwtSecret{
			MySecret: appConfig.JWTSecret,
		},
	}
}
