package main

import (
	"fmt"
	"log/slog"
	"os"

	"10.0.0.50/tuan.quang.tran/aioz-ads/config"
	_ "10.0.0.50/tuan.quang.tran/aioz-ads/docs"
	"10.0.0.50/tuan.quang.tran/aioz-ads/internal/server"
	custom_log "10.0.0.50/tuan.quang.tran/aioz-ads/internal/utils/log"
)

func main() {
	if os.Setenv("TZ", "UTC") != nil {
		panic("failed to set timezone")
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "app"
	}

	appConfig := config.MustNewAppConfig(fmt.Sprintf("./%s.env", env))

	defaultLogger := slog.New(
		custom_log.NewHandler(&slog.HandlerOptions{}, nil),
	)
	slog.SetDefault(defaultLogger)

	server := server.New(appConfig)

	server.Start()
}

//	@title			AIOZ ADS SERVICE
//	@version		1.0
//	@description	AIOZ ADS SERVICE API DOCUMENTATION
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@BasePath	/api/v1

//	@externalDocs.description	Find out more about Swagger
//	@externalDocs.url			https://swagger.io/resources/open-api/

//	@securityDefinitions.basic	BasicAuth

//	@securityDefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization
//	@description				Type "Bearer" followed by a space and JWT token.
