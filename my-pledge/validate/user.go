package validate

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"my-pledge/common"
	"my-pledge/modules/request"
	"strings"
)

type User struct {
}

func NewUser() *User {
	return &User{}
}

func (u *User) LoginCheck(g *gin.Context, l *request.Login) int {
	err := g.ShouldBind(l)
	builder := strings.Builder{}
	if err != nil {
		var er validator.ValidationErrors
		errors.As(err, &er)
		for _, e := range er {
			builder.WriteString(e.Field())
			builder.WriteString(" ")
			builder.WriteString(e.Tag())
			builder.WriteString(" ")
		}
	}

	return common.Success
}
