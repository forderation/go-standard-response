package response

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"os"
)

type StandardResponse struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	HttpStatus int
	Data       interface{} `json:"data"`
}

func (sr *StandardResponse) StandardResponse(usingPrefixCode bool, c *fiber.Ctx) error {
	code := "00"
	if usingPrefixCode {
		code = fmt.Sprintf("%s_%s", os.Getenv("PREFIX_CODE"), sr.Code)
	}
	return c.Status(sr.HttpStatus).JSON(fiber.Map{
		"code":    code,
		"message": sr.Message,
		"data":    sr.Data,
	})
}
