package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"time"
)

var sugaredLogger *zap.SugaredLogger

func init() {
	encoderCfg := zapcore.EncoderConfig{
		MessageKey:   "msg",
		LevelKey:     "level",
		CallerKey:    "caller",
		TimeKey:      "time",
		EncodeLevel:  zapcore.CapitalColorLevelEncoder,
		EncodeTime:   timeEncoder,
		EncodeCaller: abusulteCallerEncoder,
	}
	encoderOptios := make([]zap.Option, 0)
	{
		encoderOptios = append(encoderOptios, zap.Development())
		encoderOptios = append(encoderOptios, zap.AddCaller())
		encoderOptios = append(encoderOptios, zap.AddCallerSkip(1))
	}
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), os.Stdout, zap.DebugLevel)
	sugaredLogger = zap.New(core, encoderOptios...).Sugar()
}
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

//caller位置为相对项目路径地址，打印出来能够跳转，方便调试
func abusulteCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	fullPath := caller.FullPath()
	index := strings.Index(fullPath, "/huobi_Golang")
	if index < 0 {
		enc.AppendString(caller.TrimmedPath())
	} else {
		enc.AppendString(fullPath[index+14:])
	}

}
func Fatal(template string, args ...interface{}) {
	sugaredLogger.Fatalf(template, args...)
}

func Error(template string, args ...interface{}) {
	sugaredLogger.Errorf(template, args...)
}

func Panic(template string, args ...interface{}) {
	sugaredLogger.Panicf(template, args...)
}

func Warn(template string, args ...interface{}) {
	sugaredLogger.Warnf(template, args...)
}

func Info(template string, args ...interface{}) {
	sugaredLogger.Infof(template, args...)
}

func Debug(template string, args ...interface{}) {
	sugaredLogger.Debugf(template, args...)
}
