package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg          *ini.File
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PageSize     int
	JwtSectet    string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatal("fail to parse 'conf.app.ini:%v'", err)
	}
	LoadBase()

}
func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}
func LoadService() {
	ser, err := Cfg.GetSection("Server")
	if err != nil {
		log.Fatalf("faile to get section 'server':%v", err)
	}
	Runmode = Cfg.Section("").Key("Run_Mode").MustString("debug")
	HttpPort = ser.Key("Http_Port").MustInt(8080)
	ReadTimeout = time.Duration(ser.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(ser.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}
func loadApp() {

}
