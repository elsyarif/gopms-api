package middleware

import (
	"errors"
	"fmt"
	"github.com/elsyarif/pms-api/pkg/common"
	"github.com/elsyarif/pms-api/pkg/helper"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

func ErrorHandler(c *gin.Context) {
	c.Next()
	errs := c.Errors

	if len(errs) > 0 {
		err, ok := errs[0].Err.(*common.AppError)

		if ok {
			response := helper.NewResponse()
			resError := response.Error("fail", err.Errors.Error(), nil)

			switch err.Type {
			case common.NotFoundError:
				c.JSON(http.StatusNotFound, resError)
				return
			case common.ValidationError:
				var errMessages []string
				var validation validator.ValidationErrors

				if errors.As(err.Errors, &validation) {
					for _, er := range validation {
						errMsg := fmt.Sprintf("an error in the %s field is %s", strings.ToLower(er.Field()), strings.ToLower(er.ActualTag()))
						errMessages = append(errMessages, errMsg)
					}
				}
				c.JSON(http.StatusBadRequest, response.Error("fail", err.Type, errMessages))
				return
			case common.NotAuthenticatedError:
				c.JSON(http.StatusUnauthorized, resError)
				return
			case common.NotAuthorizedError:
				c.JSON(http.StatusForbidden, resError)
				return
			case common.ResourceAlreadyExists:
				c.JSON(http.StatusBadRequest, response.Error("fail", err.Errors.Error(), nil))
				return
			case common.InvalidTokenError:
				c.JSON(http.StatusBadRequest, resError)
			default:
				c.JSON(http.StatusInternalServerError, resError)
				return
			}
		}
	}
}
