package util

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func init() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
	}

	AppMode = cfg.Section("server").Key("AppMode").MustString("debug")
	HttpPort = cfg.Section("server").Key("HttpPort").MustString("debug")

	Db = cfg.Section("database").Key("Db").MustString("mysql")
	DbHost = cfg.Section("database").Key("DbHost").MustString("localhost")
	DbPort = cfg.Section("database").Key("DbPort").MustString("3306")
	DbUser = cfg.Section("database").Key("DbUser").MustString("root")
	DbPassWord = cfg.Section("database").Key("DbPassWord").MustString("test")
	DbName = cfg.Section("database").Key("DbName").MustString("test")
}
