package common

import (
	"sync"

	"active/config"
	"active/lib"
)

type MapItf map[string]interface{}

var Cfg config.Config

var UniqueIdWorker *lib.Worker

type ReqLogForChan struct {
	ReqNo      string
	UserId     string
	Router     string
	Method     string
	Agent      string
	Param      string
	HttpCode   int
	Code       int
	Msg        string
	Data       string
	Ip         string
	Server     string
	Cost       float64
	CreateTime string
	UpdateTime string
}

var ReqLogChan = make(chan *ReqLogForChan, 100)

var WG sync.WaitGroup
