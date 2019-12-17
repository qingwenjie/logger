package logger

import (
	"fmt"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

type Level int

const (
	PanicLevel   = logrus.PanicLevel
	FatalLevel   = logrus.FatalLevel
	ErrorLevel   = logrus.ErrorLevel
	WarnLevel    = logrus.WarnLevel
	InfoLevel    = logrus.InfoLevel
	DebugLevel   = logrus.DebugLevel
	TraceLevel   = logrus.TraceLevel
	DefaultLevel = TraceLevel
)

var (
	Log *logger
)

type logger struct {
	lock       sync.Mutex
	newLog     *logrus.Logger
	ErrorNew   error
	runtime    string
	stackDepth int
}

type Options struct {
	Level       logrus.Level
	LogFilePath string
	Depth       int
}

func (l *logger) Info(args ... interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.newLog.WithField("call", stackGet(l.stackDepth)).Info(args ...)
}

func (l *logger) Infof(format string, args ... interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.newLog.WithField("call", stackGet(l.stackDepth)).Infof(format, args...)
}

func (l *logger) Warn(args ... interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.newLog.WithField("call", stackGet(l.stackDepth)).Warn(args ...)
}

func (l *logger) Warnf(format string, args ... interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.newLog.WithField("call", stackGet(l.stackDepth)).Warnf(format, args...)
}

func (l *logger) Error(args ... interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.newLog.WithField("call", stackGet(l.stackDepth)).Error(args)
}

func (l *logger) Errorf(format string, args ... interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.newLog.WithField("call", stackGet(l.stackDepth)).Errorf(format, args...)
}

func (l *logger) Debug(args ... interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.newLog.WithField("call", stackGet(l.stackDepth)).Debug(args)
}

func (l *logger) Debugf(format string, args ... interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.newLog.WithField("call", stackGet(l.stackDepth)).Debugf(format, args...)
}

func (l *logger) Trace(args ... interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.newLog.WithField("call", stackGet(l.stackDepth)).Trace(args)
}

func (l *logger) Tracef(format string, args ... interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.newLog.WithField("call", stackGet(l.stackDepth)).Tracef(format, args...)
}

func (l *logger) Fatal(args ... interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.newLog.WithField("call", stackGet(l.stackDepth)).Fatal(args)
}

func (l *logger) Fatalf(format string, args ... interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.newLog.WithField("call", stackGet(l.stackDepth)).Fatalf(format, args...)
}

func (l *logger) Panic(args ... interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.newLog.WithField("call", stackGet(l.stackDepth)).Panic(args)
}

func (l *logger) Panicf(format string, args ... interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.newLog.WithField("call", stackGet(l.stackDepth)).Panicf(format, args...)
}

func (l *logger) newFile() *logrus.Logger {
	Log := logrus.New()
	if l.runtime == "" {
		if p, e := os.Getwd(); e != nil {
			l.ErrorNew = e
			return nil
		} else {
			//l.runtime = path.Join(p, "logs")
			l.runtime = p
		}
	} else {
		//l.runtime = path.Join(l.runtime, "logs")
	}
	if file, err := os.Stat(l.runtime); err != nil && os.IsNotExist(err) {
		e := os.MkdirAll(l.runtime, 0755)
		if e != nil {
			l.ErrorNew = e
			return nil
		}
	} else if file.Mode() != 0755 {
		e := os.Chmod(l.runtime, 0755)
		if e != nil {
			l.ErrorNew = e
			return nil
		}
	}
	//infoWriter, _ := rotatelogs.New(
	//	path.Join(l.runtime, "info-%Y%m%d.log"),
	//	//rotatelogs.WithLinkName(p+".info"), //把当前日志文件软链到 p+".info"
	//	//rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
	//	rotatelogs.WithRotationCount(10),                              //保留30天的数据
	//	rotatelogs.WithRotationTime(time.Duration(86400)*time.Second), //每天切割一次
	//)
	//errorWriter, _ := rotatelogs.New(
	//	path.Join(l.runtime, "error-%Y%m%d.log"),
	//	//rotatelogs.WithLinkName(p+".error"),
	//	//rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
	//	rotatelogs.WithRotationCount(10),
	//	rotatelogs.WithRotationTime(time.Duration(86400)*time.Second),
	//)
	//debugWriter, _ := rotatelogs.New(
	//	path.Join(l.runtime, "debug-%Y%m%d.log"),
	//	//rotatelogs.WithLinkName(p+".debug"),
	//	//rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
	//	rotatelogs.WithRotationCount(10),
	//	rotatelogs.WithRotationTime(time.Duration(86400)*time.Second),
	//)
	//traceWriter, _ := rotatelogs.New(
	//	path.Join(l.runtime, "trace-%Y%m%d.log"),
	//	//rotatelogs.WithLinkName(p+".trace"),
	//	//rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
	//	rotatelogs.WithRotationCount(10),
	//	rotatelogs.WithRotationTime(time.Duration(86400)*time.Second),
	//)
	//fatalWriter, _ := rotatelogs.New(
	//	path.Join(l.runtime, "fatal-%Y%m%d.log"),
	//	//rotatelogs.WithLinkName(p+".fatal"),
	//	//rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
	//	rotatelogs.WithRotationCount(10),
	//	rotatelogs.WithRotationTime(time.Duration(86400)*time.Second),
	//)
	//panicWriter, _ := rotatelogs.New(
	//	path.Join(l.runtime, "panic-%Y%m%d.log"),
	//	//rotatelogs.WithLinkName(p+".panic"),
	//	//rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
	//	rotatelogs.WithRotationCount(10),
	//	rotatelogs.WithRotationTime(time.Duration(86400)*time.Second),
	//)

	logsWriter, _ := rotatelogs.New(
		path.Join(l.runtime, "logs-%Y%m%d.log"),
		rotatelogs.WithRotationCount(10),                              //保留x天的数据
		rotatelogs.WithRotationTime(time.Duration(86400)*time.Second), //每天分隔一次
	)
	Log.Hooks.Add(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  logsWriter,
			logrus.ErrorLevel: logsWriter,
			logrus.DebugLevel: logsWriter,
			logrus.TraceLevel: logsWriter,
			logrus.FatalLevel: logsWriter,
			logrus.PanicLevel: logsWriter,
		}, &logrus.TextFormatter{},
	))
	return Log
}

func New(options *Options) *logger {
	if options == nil {
		options = DefaultOptions()
	}
	logger := logger{}
	logger.runtime = options.LogFilePath
	logger.newLog = logger.newFile()
	logger.newLog.SetLevel(options.Level)
	if logger.ErrorNew != nil {
		fmt.Println("err", logger.ErrorNew.Error())
		return nil
	}
	return &logger
}

func DefaultOptions() *Options {
	return &Options{
		Level:       DefaultLevel,
		LogFilePath: "logs",
		Depth:       1,
	}
}

func (l *logger) setStackDepth(depth int) {
	l.stackDepth = depth
}

func stackGet(depth int) string {
	pc, _, line, ok := runtime.Caller(depth + 1)
	if !ok {
		return ""
	}
	return fmt.Sprintf("%v:%v", filepath.Base(runtime.FuncForPC(pc).Name()), line)
}
