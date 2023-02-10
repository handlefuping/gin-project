package util


import "github.com/gin-gonic/gin"

//type Resp struct {
//	code int
//	data interface{}
//	count int
//	msg string
//
//}

func SuccessJsonResp(ctx *gin.Context, data interface{}, count int)  {
	ctx.JSON(200, gin.H{
		"code": 200,
		"data": data,
		"msg": "操作成功",
		"count": count,
	})

}

func FailJsonResp(ctx *gin.Context, code int, msg string)  {
	ctx.JSON(200, gin.H{
		"code": code,
		"data": nil,
		"msg": msg,
		"count": 0,
	})
}