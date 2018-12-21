package models

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ikasamt/zapp/zapp"
)

func GetUser(userID int) (user User){
	// db.Debug().Where("id = ?", userID).First(&user) など
	return
}

// Auth Middleware
func UserAuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		// ログインしてなかったらログイン画面に飛ばす
		currentUserID := zapp.GetSession(c, "user_id", 0).(int)

		if currentUserID == 0 {
			url := fmt.Sprintf("/user/login?next_url=%s", c.Request.URL)
			http.Redirect(c.Writer, c.Request, url, http.StatusFound)
			c.Abort()
			return
		}

		// ユーザーを取得
		me := GetUser(currentUserID)

		// 設定
		c.Set("me", me)
		c.Next()
	}
}
