package models

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/ikasamt/zapp/zapp"
)

// +jam ../clefs/datastore-handlers.go
type User struct {
	Name string
	Value string
	Errors []error
}


func (x *User) Kind() string { return `User`}
func (x User) NameKeyFormat() string {return fmt.Sprintf("%s", x.Name)}


func (x *User) Setter(c *gin.Context) {
	x.Name = zapp.GetParams(c, "name")
	x.Value = zapp.GetParams(c, "value")
}

func (x *User) Validations() error {
	return validation.ValidateStruct(&x,
		validation.Field(&x.Name, validation.Required),
		validation.Field(&x.Value, validation.Required),
	)
}