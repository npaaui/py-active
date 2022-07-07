package dao

import (
	"log"
	"runtime"

	"active/app/model"
	. "active/common"
)

func UpdateReqLog(args *ReqLogForChan) {
	defer func() {
		if r := recover(); r != nil {
			const size = 64 << 10
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			log.Printf("reqLogWorker: panic: %v\n%s", r, buf)
		}
		WG.Done()
	}()
	reqLog := model.NewReqLogModel().SetReqNo(args.ReqNo)
	if !reqLog.Info() {
		(&model.ReqLog{
			ReqNo:      args.ReqNo,
			UserId:     args.UserId,
			Router:     args.Router,
			Method:     args.Method,
			Agent:      args.Agent,
			Param:      args.Param,
			HttpCode:   args.HttpCode,
			Code:       args.Code,
			Msg:        args.Msg,
			Data:       args.Data,
			Ip:         args.Ip,
			ServerAddr: args.Server,
			Cost:       args.Cost,
		}).Insert()
	} else {
		reqLog.Update(&model.ReqLog{
			UserId:     args.UserId,
			Router:     args.Router,
			Method:     args.Method,
			Agent:      args.Agent,
			Param:      args.Param,
			HttpCode:   args.HttpCode,
			Code:       args.Code,
			Msg:        args.Msg,
			Data:       args.Data,
			Ip:         args.Ip,
			ServerAddr: args.Server,
			Cost:       args.Cost,
		})
	}
}
