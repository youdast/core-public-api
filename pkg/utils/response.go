package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Metadata struct {
	Timestamp string `json:"timestamp"`
	Version   string `json:"version"`
}

type ApiResponse struct {
	Status   string      `json:"status"`
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
	Metadata Metadata    `json:"metadata"`
}

func SuccessResponse(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(ApiResponse{
		Status:  "success",
		Code:    fiber.StatusOK,
		Message: message,
		Data:    data,
		Metadata: Metadata{
			Timestamp: time.Now().Format(time.RFC3339),
			Version:   "v1",
		},
	})
}

func ErrorResponse(c *fiber.Ctx, code int, message string, data interface{}) error {
	return c.Status(code).JSON(ApiResponse{
		Status:  "error",
		Code:    code,
		Message: message,
		Data:    data,
		Metadata: Metadata{
			Timestamp: time.Now().Format(time.RFC3339),
			Version:   "v1",
		},
	})
}
