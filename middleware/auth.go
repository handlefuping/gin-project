package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("AuthRequired middleware")
		//userId, tokenStr, _ := c.Request.BasicAuth()
		//fmt.Println(userId)
		//if _, err := util.ParseTokenStr(tokenStr); err != nil {
		//	util.FailJsonResp(c, 401, "验证未通过")
		//	c.Abort()
		//} else {
		//	c.Next()
		//}

		//c.Next()

	}
}