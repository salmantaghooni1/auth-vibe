package logging

import (
	"os"
	"sync"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/salmantaghooni/golang-car-web-api/config"
)

var once sync.Once
var zeroSinLogger *zerolog.Logger

type zeroLogger struct {
	cfg    *config.Config
	logger *zerolog.Logger
}

var zeroLogLevelMap = map[string]zerolog.Level{
	"debug": zerolog.DebugLevel,
	"info":  zerolog.InfoLevel,
	"warn":  zerolog.WarnLevel,
	"error": zerolog.ErrorLevel,
	"fatal": zerolog.FatalLevel,
}

func newZeroLogger(cfg *config.Config) *zeroLogger {
	logger := &zeroLogger{cfg: cfg}
	logger.Init()
	return logger
}

func (l *zeroLogger) getLogLevel() zerolog.Level {
	level, exist := zeroLogLevelMap[l.cfg.Logger.Level]
	if !exist {
		return zerolog.DebugLevel
	}
	return level
}

func (l *zeroLogger) Init() {
	once.Do(func() {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		file, err := os.OpenFile(l.cfg.Logger.FilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			panic("couldn't open logger file")
		}
		zerolog.SetGlobalLevel(l.getLogLevel())
		logger := zerolog.New(file).With().Timestamp().Str("App Name: ", "My App").Str("Logger: ", "Zerolog").Logger()
		zeroSinLogger = &logger
	})
	l.logger = zeroSinLogger

}

func (l *zeroLogger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	l.logger.Debug().Str("Category", string(cat)).Str("SubCategory", string(sub)).Fields(maptoZeroParams(extra)).Msg(msg)
}

func (l *zeroLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debug().Msgf(template, args...)
}

func (l *zeroLogger) Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	l.logger.Info().Str("Category", string(cat)).Str("SubCategory", string(sub)).Fields(maptoZeroParams(extra)).Msg(msg)
}

func (l *zeroLogger) Infof(template string, args ...interface{}) {
	l.logger.Info().Msgf(template, args...)
}

func (l *zeroLogger) Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	l.logger.Warn().Str("Category", string(cat)).Str("SubCategory", string(sub)).Fields(maptoZeroParams(extra)).Msg(msg)
}

func (l *zeroLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warn().Msgf(template, args...)
}

func (l *zeroLogger) Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	l.logger.Error().Str("Category", string(cat)).Str("SubCategory", string(sub)).Fields(maptoZeroParams(extra)).Msg(msg)
}

func (l *zeroLogger) Errorf(template string, args ...interface{}) {
	l.logger.Error().Msgf(template, args...)
}

func (l *zeroLogger) Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	l.logger.Fatal().Str("Category", string(cat)).Str("SubCategory", string(sub)).Fields(maptoZeroParams(extra)).Msg(msg)
}

func (l *zeroLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatal().Msgf(template, args...)
}
