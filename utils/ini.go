package utils

import (
	"gopkg.in/ini.v1"
)

var iniPath = "./gtp.ini"

func GetIni(section, key string) string {
	cfg, _ := ini.Load(iniPath)
	return cfg.Section(section).Key(key).String()
}

func SetIni(section, key, val string) error {
	cfg, err := ini.Load(iniPath)
	if err != nil {
		return err
	}
	cfg.Section(section).Key(key).SetValue(val)
	err = cfg.SaveTo(iniPath)
	if err != nil {
		return err
	}
	return nil
}
