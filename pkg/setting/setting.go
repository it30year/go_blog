package setting

import (
	"github.com/go-ini/ini"
	"time"
	"log"
)

var (
	Cfg *ini.File
	RunMode string
	HttpPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	PageSize int
	JwtSectet string
)
func init(){
	var err error
	Cfg,err = ini.Load("conf/app.ini")
	if(err !=nil){
		log.Fatal("fail to parse 'conf.app.ini:%v'",err)
	}
	LoadBase()

}
func LoadBase(){
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}
