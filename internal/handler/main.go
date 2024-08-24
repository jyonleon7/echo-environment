package handler

import (
	"fmt"
	"net/http"

	"github.com/jyonleon7/echo-environment/internal/api"
	"github.com/labstack/echo/v4"
)

type EchoEnvironmentHandler struct {
}

var (
	_ api.ServerInterface = (*EchoEnvironmentHandler)(nil)
)

func NewHandler() *EchoEnvironmentHandler {
	return &EchoEnvironmentHandler{}
}

func (c *EchoEnvironmentHandler) Ping(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "OK")
}

func (c *EchoEnvironmentHandler) Validate(ctx echo.Context) error {
	var v api.TestValid
	err := ctx.Bind(&v)
	if err != nil {
		return ctx.JSON(400, nil)
	}

	return ctx.String(http.StatusOK, fmt.Sprintf("w: %d, t: %s", v.Width, v.Title))
}
