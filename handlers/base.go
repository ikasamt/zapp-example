package handlers

import (
	"fmt"
	"github.com/ikasamt/zapp-example/models"
	"github.com/ikasamt/zapp/zapp"
	"net/http"

	"github.com/gin-gonic/gin"
)

var ENV *zapp.Environment


func Pong(c *gin.Context) {
	c.String(http.StatusOK, `pong`)
}

func TopIndex(c *gin.Context) {
	zapp.Render(c, `app`, gin.H{})
}


func DummyCreateHandler(c *gin.Context) {
	instance := &models.User{}
	instance.Name = "aaa"
	instance.Value = "AAABBB"
	models.SaveUser(instance)
	message := fmt.Sprintf("%v 追加しました", instance)
	c.String(200, message)
}
