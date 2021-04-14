package err

import (
	"errors"
	"github.com/gofiber/fiber/v2"
)

var (
	ConfigFilePathIsInvalid      = errors.New("config file path is invalid")
	MainProcessContextTerminated = errors.New("main process context done received")
)

const ERROR = "error"
const SUCCESS = "success"

func Result(ctx *fiber.Ctx, status string, statusCode int, message string) error {
	//return fiber.Map{"status": status, "code": statusCode, "message": message, "data": nil}
	return ctx.Status(statusCode).JSON(fiber.Map{"status": status, "code": statusCode, "message": message, "data": nil})
}
