package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, code int, result interface{}) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    code,   // 自定义code
		"message": "OK",   // message
		"result":  result, // 数据
		"status":  "OK",
	})
}

func Err(c *gin.Context, httpCode int, code int, msg string) {
	c.JSON(httpCode, map[string]interface{}{
		"code":    code,
		"message": msg,
		"result":  gin.H{},
		"status":  "fail",
	})
}

func ErrWithResult(c *gin.Context, httpCode int, code int, msg string, result interface{}) {
	c.JSON(httpCode, map[string]interface{}{
		"code":    code,
		"message": msg,
		"result":  result,
		"status":  "fail",
	})
}
