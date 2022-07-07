package main

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"

	"github.com/npaaui/helper-go/db"
	"github.com/npaaui/helper-go/gen"

	. "active/common"
)

func main() {
	err := cleanenv.ReadConfig("./config/config.yml", &Cfg)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	(&db.Conf{
		DriverName:      "mysql",
		ConnMaxLifetime: 86400,
		Prefix:          Cfg.Prefix,
		Conn: db.MysqlConf{
			Host:     Cfg.Host,
			Username: Cfg.Username,
			Password: Cfg.Password,
			Database: Cfg.Database,
		},
	}).InitDbEngine()

	(&gen.Conf{
		ModelFolder: "app/model/",
		TplFile:     "tmp/gen/model.tpl",
		TableNames:  "",
		DbName:      "my_active",
	}).InitGenConf()
	gen.GenerateModelFile()
}
