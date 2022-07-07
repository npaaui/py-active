package common

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type RespBody struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ReturnData(g *gin.Context, data interface{}) {
	g.JSON(200, RespBody{
		Code: SUCCESS,
		Msg:  GetMsg(SUCCESS),
		Data: data,
	})
	return
}

func ReturnErr(g *gin.Context, code int, err error) {
	msgF := fmt.Errorf(GetMsg(code) + ": " + err.Error()).Error()
	respErr := NewRespErr(code, msgF)
	g.JSON(200, RespBody{
		Code: respErr.Code,
		Msg:  respErr.Msg,
		Data: nil,
	})
	return
}

func ReturnErrMsg(g *gin.Context, code int, msg string) {
	respErr := NewRespErr(code, msg)
	g.JSON(200, RespBody{
		Code: respErr.Code,
		Msg:  respErr.Msg,
		Data: nil,
	})
	return
}

func ReturnByRespErr(g *gin.Context, respErr RespErr) {
	g.JSON(200, RespBody{
		Code: respErr.Code,
		Msg:  respErr.Msg,
		Data: nil,
	})
	return
}

func ReturnFile(c *gin.Context, fileName string, data []byte) {
	c.Writer.Header().Set("Content-Type", "application/octet-stream") //二进制流文件
	c.Writer.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s", fileName))
	c.Data(200, "application/octet-stream", data)
	c.Abort()
}
