package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context){
	msg := make(map[string]interface{})
	msg["status"] = "success"
	msg["username"] = "admin"
	msg["user_id"] = 1
	msg["user_power"] = 2
	c.JSON(http.StatusOK, gin.H{
		"errorcode":0,
		"msg": msg,
	})

}

func Logout(c *gin.Context) {
	msg := "退出成功"
	code := 0
	c.JSON(http.StatusOK, gin.H{
		"errorcode" : code,
		"msg" : msg,
	})
}
