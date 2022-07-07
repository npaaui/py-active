package common

import (
	"errors"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

type DbConnConf struct {
	DriverName      string
	ConnMaxLifetime int64
	Prefix          string
	Conn
}
type Conn interface {
	GetDataSourceName() string
}

/**
 * mysql
 */
type MysqlConf struct {
	Host     string
	Username string
	Password string
	Database string
}

func (c MysqlConf) GetDataSourceName() (dataSourceName string) {
	dataSourceName = c.Username + ":" + c.Password + "@(" + c.Host + ")/" + c.Database + "?charset=utf8mb4&loc=Local"
	return
}

var ConfIns *DbConnConf

func (d *DbConnConf) InitDbEngine() {
	ConfIns = d
	SetDbEngine()
}

var EngineIns *xorm.Engine

var SetDbEngineOnce sync.Once

func GetDbEngineIns() *xorm.Engine {
	SetDbEngineOnce.Do(SetDbEngine)
	return EngineIns
}

func SetDbEngine() {
	if ConfIns == nil {
		panic(errors.New("[danger] DbConfIns is nil"))
	}

	DbEngine, err := xorm.NewEngine(ConfIns.DriverName, ConfIns.Conn.GetDataSourceName())
	if err != nil {
		panic(fmt.Errorf("[danger] NewEngine error: %w", err))
	}

	if ConfIns.ConnMaxLifetime > 0 {
		DbEngine.DB().SetConnMaxLifetime(time.Duration(ConfIns.ConnMaxLifetime) * time.Second)
	}

	if ConfIns.Prefix != "" {
		tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, ConfIns.Prefix)
		DbEngine.SetTableMapper(tbMapper)
	}

	EngineIns = DbEngine
}