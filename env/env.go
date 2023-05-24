package env

import (
	"log"

	"github.com/Unknwon/goconfig"
)

var Config *goconfig.ConfigFile
var err error

func init() {
	Config, err = goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		log.Fatalf("无法加载配置文件: %s", err)
	}
}
