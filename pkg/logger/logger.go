package logger

import (
	"ZWebsite/pkg/constant"
	"ZWebsite/pkg/setting"
	"ZWebsite/pkg/utils"
	"context"
	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strconv"
	"time"
)

var (
	levelMap = map[string]zapcore.Level{
		"debug":  zapcore.DebugLevel,
		"info":   zapcore.InfoLevel,
		"warn":   zapcore.WarnLevel,
		"error":  zapcore.ErrorLevel,
		"dpanic": zapcore.DPanicLevel,
		"panic":  zapcore.PanicLevel,
		"fatal":  zapcore.FatalLevel,
	}
)

type Factory struct {
	logger *zap.Logger
	ctx    context.Context
}

var factory *Factory

func SetUp(){
	InitLogger(setting.Setting.LogLevel,setting.Setting.ServiceName)
}

func InitLogger(logLevel ,serviceName string)  {
	fileName := constant.LogDir + "/" + serviceName + ".%Y%m%d%H" + ".log"
	maxAgeHours := time.Hour * time.Duration(getLogRotateDays() * 24)
	// 日志切割 rotateLogs 写入流
	timeRotateWriter, _ := rotateLogs.New(
		fileName,
		rotateLogs.WithMaxAge(maxAgeHours),
		rotateLogs.WithRotationTime(24*time.Hour), )

	sync := zapcore.AddSync(timeRotateWriter)

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(getEncoderConfig()),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), sync),
		zap.NewAtomicLevelAt(getLoggerLevel(logLevel)))

	factory = &Factory{
		// 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数
		logger: zap.New(core,zap.AddCaller(),zap.AddCallerSkip(1)),
	}
}

func getLogRotateDays() int {
	envDays := os.Getenv(constant.LogRotateDaysEnvKey)
	if envDays == "" {
		return constant.DefaultLogRotateDays
	}

	days, err := strconv.Atoi(envDays)
	if err != nil {
		return constant.DefaultLogRotateDays
	}
	return utils.Max(days,constant.MinLogRotateDays)
}
func getEncoderConfig() zapcore.EncoderConfig {
	// 如果不指定对应key的name的话，对应key的信息不处理，即不会写入到文件中
	// 如MessageKey为空的话，内容主体不处理，即看不到log内容
	return zapcore.EncoderConfig{
		NameKey:      "name",
		MessageKey:   "msg",
		LevelKey:     "level",
		EncodeLevel:  func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {},
		TimeKey:      "s",
		EncodeTime:   func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeName: func(n string, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(n)
		},
	}
}
func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level

	}
	return zapcore.InfoLevel
}

func setLoggerName(l *Factory, level string) *zap.Logger {
	t := time.Now()
	suffixTimeStamp := t.Format("2006-01-02 15:04:05.000")
	// log 最前面设置log time
	name := suffixTimeStamp
	// check request_uuid
	traceId := l.ctx.Value(constant.RequestKey)
	if traceId != nil {
		name += "|" + traceId.(string) + "|" + levelMap[level].CapitalString()
	} else {
		name += "|" + levelMap[level].CapitalString()
	}
	return l.logger.Named(name)
}


func For(ctx context.Context) *Factory {
	return &Factory{
		logger: factory.logger,
		ctx:    ctx,
	}
}

func (l *Factory) Debug(msg string, fields ...zapcore.Field) {
	setLoggerName(l, "debug").Debug(msg, fields...)
}

func (l *Factory) Info(msg string, fields ...zapcore.Field) {
	setLoggerName(l, "info").Info(msg, fields...)
}

func (l *Factory) Infof(template string, args ...interface{}) {
	setLoggerName(l, "info").Sugar().Infof(template, args...)
}

func (l *Factory) Error(msg string, fields ...zapcore.Field) {
	setLoggerName(l, "error").Error(msg, fields...)
}

func (l *Factory) Errorf(template string, args ...interface{}) {
	setLoggerName(l, "error").Sugar().Errorf(template, args...)
}

func (l *Factory) Warn(msg string, fields ...zapcore.Field) {
	setLoggerName(l, "warn").Warn(msg, fields...)
}

func (l *Factory) Warnf(template string, args ...interface{}) {
	setLoggerName(l, "warn").Sugar().Warnf(template, args...)
}

func (l *Factory) Fatal(msg string, fields ...zapcore.Field) {
	setLoggerName(l, "fatal").Fatal(msg, fields...)
}

func (l *Factory) Fatalf(template string, args ...interface{}) {
	setLoggerName(l, "fatal").Sugar().Fatalf(template, args...)
}