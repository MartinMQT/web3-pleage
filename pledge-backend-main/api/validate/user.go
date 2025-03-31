package validate

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	"pledge-backend/api/common/statecode"
	"pledge-backend/api/models/request"
)

type User struct{}

func NewUser() *User {
	return &User{}
}

func (v *User) Login(c *gin.Context, req *request.Login) int {

	err := c.ShouldBind(req)
	if err == io.EOF {
		return statecode.ParameterEmptyErr
	} else if err != nil {
		var errs validator.ValidationErrors
		errors.As(err, &errs)
		for _, e := range errs {
			if e.Field() == "Name" && e.Tag() == "required" {
				return statecode.PNameEmpty
			}
			if e.Field() == "Password" && e.Tag() == "required" {
				return statecode.PNameEmpty
			}
		}
		return statecode.CommonErrServerErr
	}

	return statecode.CommonSuccess
}
