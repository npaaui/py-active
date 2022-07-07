package middleware

import (
	"fmt"
	"runtime/debug"

	"github.com/gin-gonic/gin"

	. "active/common"
)

func RecoverDbError() gin.HandlerFunc {
	return func(g *gin.Context) {
		defer func() {
			p := recover()
			g.Abort()
			// 自定义异常
			if sysErr, ok := p.(SysErr); ok {
				ReturnErr(g, ErrSys, sysErr)
				return
			}
			if dbErr, ok := p.(DbErr); ok {
				fmt.Printf("panic recover! p: %v", p)
				debug.PrintStack()
				ReturnErr(g, ErrSysDbExec, dbErr)
				return
			}
			if validErr, ok := p.(ValidErr); ok {
				ReturnErr(g, ErrValidReq, validErr)
				return
			}
			if respErr, ok := p.(RespErr); ok {
				ReturnErrMsg(g, respErr.Code, respErr.Msg)
				return
			}

			// 其它异常
			if p != nil {
				fmt.Printf("panic recover! p: %v", p)
				debug.PrintStack()
			}
			if err, ok := p.(error); ok {
				ReturnErr(g, ErrSys, err)
				return
			}
		}()
		g.Next()
	}
}
