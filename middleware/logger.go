package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"time"

	. "active/common"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (b bodyLogWriter) Write(data []byte) (int, error) {
	b.body.Write(data)
	return b.ResponseWriter.Write(data)
}

func (b bodyLogWriter) WriteString(s string) (int, error) {
	b.body.WriteString(s)
	return b.ResponseWriter.WriteString(s)
}

func ReqLog() gin.HandlerFunc {
	return func(g *gin.Context) {
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: g.Writer}
		g.Writer = bodyLogWriter

		start := time.Now().UnixNano() / 1e6
		reqNo := GetUniqueId()
		g.Set("req_no", reqNo)

		ReqLogChan <- &ReqLogForChan{
			ReqNo:  reqNo,
			Router: g.Request.URL.Path,
			Method: g.Request.Method,
			Agent:  g.Request.UserAgent(),
			Param:  g.Request.RequestURI,
			Ip:     g.ClientIP(),
			Server: Cfg.Addr,
		}

		g.Next()

		// 记录接口返回
		respBody := bodyLogWriter.body.String()
		resp := RespBody{}
		_ = json.Unmarshal([]byte(respBody), &resp)
		data := ""
		if g.Request.URL.Path != "/api/v1/lottery_user/list" { // 某些接口不记录返回数据
			dataByte, _ := json.Marshal(resp.Data)
			data = bytes.NewBuffer(dataByte).String()
		}
		cost := float64(time.Now().UnixNano()/1e6-start) / 1000
		ReqLogChan <- &ReqLogForChan{
			ReqNo:    reqNo,
			UserId:   g.GetString("user_id"),
			Cost:     cost,
			HttpCode: g.Writer.Status(),
			Code:     resp.Code,
			Msg:      resp.Msg,
			Data:     data,
		}
	}
}
