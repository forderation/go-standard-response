package response

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"runtime/debug"
)

func InvalidRequestResponse(prefixServiceName string, err error, c *fiber.Ctx) (errs error) {
	sr := StandardResponse{
		Code:       "INVALID-REQUEST-" + prefixServiceName,
		Message:    err.Error(),
		Data:       nil,
		HttpStatus: http.StatusBadRequest,
	}
	errs = sr.StandardResponse(true, c)
	return
}

func InternalServerErrorResponse(prefixServiceName string, err error, c *fiber.Ctx) (errs error) {
	fmt.Println("stacktrace from panic: \n" + string(debug.Stack()))
	sr := StandardResponse{
		Code:       "ERROR-" + prefixServiceName,
		Message:    err.Error(),
		Data:       nil,
		HttpStatus: http.StatusInternalServerError,
	}
	errs = sr.StandardResponse(true, c)
	return
}

const (
	DataObject = iota
	DataArray
	DataNil
)

func NotFoundResponse(message string, typeData int, c *fiber.Ctx) (errs error) {
	sr := StandardResponse{
		Code:       "ERR-NOT-FOUND",
		Message:    message,
		HttpStatus: http.StatusOK,
	}
	if typeData == DataObject {
		sr.Data = map[string]interface{}{}
	} else if typeData == DataArray {
		sr.Data = []map[string]interface{}{}
	} else {
		sr.Data = nil
	}
	return sr.StandardResponse(false, c)
}

func NotFoundResponseWithPagination(message string, c *fiber.Ctx) (errs error) {
	sr := StandardResponse{
		Code:    "ERR-NOT-FOUND",
		Message: message,
		Data: map[string]interface{}{
			"count": 0,
			"total": 0,
			"rows":  []interface{}{},
		},
		HttpStatus: http.StatusOK,
	}
	errs = sr.StandardResponse(false, c)
	return
}

func OkResponse(message string, data interface{}, c *fiber.Ctx) (errs error) {
	sr := StandardResponse{
		Code:       "00",
		Message:    message,
		Data:       data,
		HttpStatus: http.StatusOK,
	}
	errs = sr.StandardResponse(false, c)
	return
}

func OkResponseWithPagination(message string, count int64, total int, data interface{}, c *fiber.Ctx) (errs error) {
	sr := StandardResponse{
		Code:    "00",
		Message: message,
		Data: map[string]interface{}{
			"count": count,
			"total": total,
			"rows":  data,
		},
		HttpStatus: http.StatusOK,
	}
	errs = sr.StandardResponse(false, c)
	return
}
