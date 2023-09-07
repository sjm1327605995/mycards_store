package resp

import "github.com/gin-gonic/gin"

type SuccessResp struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, SuccessResp{
		Status: 200,
		Data:   data,
	})
}

type FailResp struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

func Fail(c *gin.Context, msg string) {
	c.JSON(200, FailResp{
		Status: 401,
		Error:  msg,
	})
}
