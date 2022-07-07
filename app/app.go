package app

import (
	"active/app/dao"
	"fmt"
	"github.com/gin-gonic/gin"

	v1 "active/app/controller/api/v1"
	. "active/common"
	"active/config"
	"active/lib"
)

func Run(cfg *config.Config) {
	r := gin.Default()
	v1.NewRouter(r)

	// 初始化数据库
	InitMysql(cfg)

	DoSomeRoutine()
	UniqueIdWorker = lib.NewWorker(1) // 唯一id生成器

	_ = r.Run(cfg.Addr)
}

// 初始化数据库
func InitMysql(cfg *config.Config) {
	if cfg.Host == "" {
		panic(fmt.Errorf("get mysql conf error: %+v", cfg))
	}
	dbConf := DbConnConf{
		DriverName:      "mysql",
		ConnMaxLifetime: 86400,
		Prefix:          cfg.Prefix,
		Conn: MysqlConf{
			Host:     cfg.Host,
			Username: cfg.Username,
			Password: cfg.Password,
			Database: cfg.Database,
		},
	}
	dbConf.InitDbEngine()
	GetDbEngineIns()
	if cfg.ShowSql == "true" {
		GetDbEngineIns().ShowSQL(true)
	}
}

func DoSomeRoutine() {
	// 请求日志记录通道
	go func() {
		for {
			select {
			case reqLog := <-ReqLogChan:
				WG.Add(1)
				go func() {
					dao.UpdateReqLog(reqLog)
				}()
				WG.Wait()
			}
		}
	}()
}
