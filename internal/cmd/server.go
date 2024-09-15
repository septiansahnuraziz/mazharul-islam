package cmd

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mazharul-islam/cacher"
	"github.com/mazharul-islam/config"
	"github.com/mazharul-islam/docs"
	"github.com/mazharul-islam/internal/controller/http"
	"github.com/mazharul-islam/internal/database"
	"github.com/mazharul-islam/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"strings"
)

var runServer = &cobra.Command{
	Use:   "server",
	Short: "run server",
	Long:  `This subcommand start the server`,
	Run:   server,
}

func init() {
	RootCmd.AddCommand(runServer)
}

func server(cmd *cobra.Command, args []string) {
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	db, err := database.InitializePostgresConnection()
	if err != nil {
		logrus.Fatal("err initialize db")
	}

	postgresDB, err := database.PostgreSQL.DB()
	defer utils.WrapCloser(postgresDB.Close)

	cacheManager := cacher.ConstructCacheManager()

	if config.EnableCaching() {
		redisDB, err := database.InitializeRedigoRedisConnectionPool(config.RedisCacheHost(), redisOptions)
		continueOrFatal(err)
		defer utils.WrapCloser(redisDB.Close)

		cacheManager.SetConnectionPool(redisDB)
	}

	cacheManager.SetDisableCaching(!config.EnableCaching())

	app := gin.Default()

	// get default url request
	app.UseRawPath = true
	app.UnescapePathValues = true
	app.RemoveExtraSlash = true

	// cors configuration
	corsConfig := cors.DefaultConfig()
	corsConfig.AddAllowHeaders("Authorization")
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowMethods("OPTIONS", "PUT", "POST", "GET", "DELETE")

	app.Use(cors.New(corsConfig))

	matchService := InitMatchService(db, cacheManager)

	http.RouteService(
		&app.RouterGroup,
		matchService,
	)

	initSwaggerDocs(&app.RouterGroup)

	if err := app.Run(); err != nil {
		logrus.Fatal(err)
	}
}

func initSwaggerDocs(app *gin.RouterGroup) {
	swaggerEndpoint := config.SwaggerEndpoint()
	swaggerSchemes := []string{"http"}

	// swagger configuration
	docs.SwaggerInfo.Title = config.AppName()
	docs.SwaggerInfo.Description = "POS-B2B API"
	docs.SwaggerInfo.Version = config.AppVersion()
	docs.SwaggerInfo.Host = swaggerEndpoint
	docs.SwaggerInfo.Schemes = swaggerSchemes

	swagConfig := &ginSwagger.Config{
		URL: swaggerEndpoint + "/docs/swagger/doc.json",
	}

	// swagger endpoint with authentication
	swaggerDocs := app.Group("/docs", gin.BasicAuth(gin.Accounts{config.SwaggerUsername(): config.SwaggerPassword()}))
	{
		swaggerDocs.GET("/swagger/*any", ginSwagger.CustomWrapHandler(swagConfig, swaggerFiles.Handler))
	}
}
