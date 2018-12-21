package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ikasamt/zapp/zapp"
)

func UserLogout(c *gin.Context) {
	zapp.SetSession(c, `user_id`, nil)
	c.Redirect(301, `/`)
}