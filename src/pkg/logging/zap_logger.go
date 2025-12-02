package logging

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/salmantaghooni/golang-car-web-api/config"
)

var zapSinLogger *zap.SugaredLogger

type zapLogger struct {
	cfg    *config.Config
	logger *zap.SugaredLogger
}

var zapLogLevelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
}

func newZapLogger(cfg *config.Config) *zapLogger {
	logger := &zapLogger{cfg: cfg}
	logger.Init()
	return logger
}

func (l *zapLogger) getLogLevel() zapcore.Level {
	level, exist := zapLogLevelMap[l.cfg.Logger.Level]
	if !exist {
		return zapcore.DebugLevel
	}
	return level
}

func (l *zapLogger) Init() {
	once.Do(func() {
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   l.cfg.Logger.FilePath,
			MaxSize:    1,
			MaxAge:     5,
			LocalTime:  true,
			MaxBackups: 10,
			Compress:   true,
		})

		config := zap.NewProductionEncoderConfig()
		config.EncodeTime = zapcore.ISO8601TimeEncoder

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(config),
			w,
			l.getLogLevel(),
		)

		logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()
		zapSinLogger = logger.With("App Name:", "My App", "Logger:", "zapcore")
	})
	l.logger = zapSinLogger
}

func (l *zapLogger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {

	params := prepareLogKeys(extra, cat, sub)
	l.logger.Debugw(msg, params...)
}

func (l *zapLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args)
}

func (l *zapLogger) Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {

	params := prepareLogKeys(extra, cat, sub)
	l.logger.Infow(msg, params...)
}

func (l *zapLogger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args)
}

func (l *zapLogger) Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {

	params := prepareLogKeys(extra, cat, sub)
	l.logger.Warnw(msg, params...)
}

func (l *zapLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args)
}

func (l *zapLogger) Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {

	params := prepareLogKeys(extra, cat, sub)
	l.logger.Errorw(msg, params...)
}

func (l *zapLogger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args)
}

func (l *zapLogger) Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {

	params := prepareLogKeys(extra, cat, sub)
	l.logger.Fatalw(msg, params...)
}

func (l *zapLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args)
}

func prepareLogKeys(extra map[ExtraKey]interface{}, cat Category, sub SubCategory) []interface{} {
	if extra == nil {
		extra = make(map[ExtraKey]interface{}, 0)
	}
	extra["Category"] = cat
	extra["SubCategory"] = sub
	params := maptoZapParams(extra)
	return params
}
