package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// _, tokenStr, _ := c.Request.BasicAuth()
		// if _, err := util.ParseTokenStr(tokenStr); err != nil {
		// 	// c.AbortWithStatusJSON()
		// 	// fmt.Println("2222")
		// 	// c.AbortWithStatusJSON(500, gin.H{"status": false, "message": "错误"})
		// 	c.Abort()
		// 	// // c.Status(500)
		// 	// util.FailJsonResp(c, 401, "验证未通过")
		// } else {
		// 	fmt.Println("1111")
		// 	c.Next()
		// }
			c.Next()
	}
}