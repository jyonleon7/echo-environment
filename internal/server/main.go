package server

import (
	"fmt"
	"os"

	"github.com/jyonleon7/echo-environment/internal/api"
	"github.com/jyonleon7/echo-environment/internal/handler"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	middleware "github.com/oapi-codegen/echo-middleware"
)

func Run() error {
	addr := fmt.Sprintf(":%s", defaultEnv("PORT", "8080"))
	fmt.Println("listen", addr)
	fmt.Println("env", defaultEnv("APP_ENV", "dev"))

	swagger, err := api.GetSwagger()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	e := echo.New()
	h := handler.NewHandler()

	e.Use(echomiddleware.Logger())
	e.Use(middleware.OapiRequestValidator(swagger))

	api.RegisterHandlers(e, h)

	return e.Start(addr)
}

func defaultEnv(key, defaultValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return defaultValue
}
