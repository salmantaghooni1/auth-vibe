package logging

import (
	"fmt"

	"github.com/salmantaghooni/golang-car-web-api/config"
)

type Logger interface {
	Init()

	Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Infof(template string, args ...interface{})

	Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Debugf(template string, args ...interface{})

	Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Warnf(template string, args ...interface{})

	Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Errorf(template string, args ...interface{})

	Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{})
	Fatalf(template string, args ...interface{})
}

func NewLogger(cfg *config.Config) Logger {

	if cfg.Logger.Logger == "zap" {
		fmt.Println(cfg.Logger.Logger)
		return newZapLogger(cfg)
	} else if cfg.Logger.Logger == "zerolog" {
		return newZeroLogger(cfg)
	}
	panic("logger not supported")
}
