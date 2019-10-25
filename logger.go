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

const (
	PanicLevel    = logrus.PanicLevel
	FatalLevel    = logrus.FatalLevel
	ErrorLevel    = logrus.ErrorLevel
	InfoLevel     = logrus.InfoLevel
	DebugLevel    = logrus.DebugLevel
	TraceLevel    = logrus.TraceLevel
	DefaultLevel  = TraceLevel
	LOG_TYPE_FILE = "FILE"
	LOG_TYPE_TEXT = "TEXT"
)

var (
	TextLog *logger
	FileLog *logger
)

type logger struct {
	lock     sync.Mutex
	newLog   *logrus.Logger
	ErrorNew error
	runtime  string
}

type Options struct {
	Level       logrus.Level
	LogType     string
	LogFilePath string
}

func (l *logger) Info(args ... interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.newLog.Info(fileLine(), args)
}

func (l *logger) Infof(format string, args ... interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.newLog.Infof(format, args)
}

func (l *logger) Error(args ... interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.newLog.Error(fileLine(), args)
}

func (l *logger) Errorf(format string, args ... interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.newLog.Errorf(format, args)
}

func (l *logger) Debug(args ... interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.newLog.Debug(fileLine(), args)
}

func (l *logger) Debugf(format string, args ... interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.newLog.Debugf(format, args)
}

func (l *logger) Trace(args ... interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.newLog.Trace(fileLine(), args)
}

func (l *logger) Tracef(format string, args ... interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.newLog.Tracef(format, args)
}

func (l *logger) Fatal(args ... interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.newLog.Fatal(fileLine(), args)
}

func (l *logger) Fatalf(format string, args ... interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.newLog.Fatalf(format, args)
}

func (l *logger) Panic(args ... interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.newLog.Panic(fileLine(), args)
}

func (l *logger) Panicf(format string, args ... interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.newLog.Panicf(format, args)
}

func fileLine() string {
	_, file, line, _ := runtime.Caller(2)
	file = filepath.Base(file)
	return fmt.Sprintf("[%s:%d] ", file, line)
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
	infoWriter, _ := rotatelogs.New(
		path.Join(l.runtime, "info-%Y%m%d.log"),
		//rotatelogs.WithLinkName(p+".info"), //把当前日志文件软链到 p+".info"
		//rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
		rotatelogs.WithRotationCount(10),                              //保留30天的数据
		rotatelogs.WithRotationTime(time.Duration(86400)*time.Second), //每天切割一次
	)
	errorWriter, _ := rotatelogs.New(
		path.Join(l.runtime, "error-%Y%m%d.log"),
		//rotatelogs.WithLinkName(p+".error"),
		//rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
		rotatelogs.WithRotationCount(10),
		rotatelogs.WithRotationTime(time.Duration(86400)*time.Second),
	)
	debugWriter, _ := rotatelogs.New(
		path.Join(l.runtime, "debug-%Y%m%d.log"),
		//rotatelogs.WithLinkName(p+".debug"),
		//rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
		rotatelogs.WithRotationCount(10),
		rotatelogs.WithRotationTime(time.Duration(86400)*time.Second),
	)
	traceWriter, _ := rotatelogs.New(
		path.Join(l.runtime, "trace-%Y%m%d.log"),
		//rotatelogs.WithLinkName(p+".trace"),
		//rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
		rotatelogs.WithRotationCount(10),
		rotatelogs.WithRotationTime(time.Duration(86400)*time.Second),
	)
	fatalWriter, _ := rotatelogs.New(
		path.Join(l.runtime, "fatal-%Y%m%d.log"),
		//rotatelogs.WithLinkName(p+".fatal"),
		//rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
		rotatelogs.WithRotationCount(10),
		rotatelogs.WithRotationTime(time.Duration(86400)*time.Second),
	)
	panicWriter, _ := rotatelogs.New(
		path.Join(l.runtime, "panic-%Y%m%d.log"),
		//rotatelogs.WithLinkName(p+".panic"),
		//rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
		rotatelogs.WithRotationCount(10),
		rotatelogs.WithRotationTime(time.Duration(86400)*time.Second),
	)
	Log.Hooks.Add(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  infoWriter,
			logrus.ErrorLevel: errorWriter,
			logrus.DebugLevel: debugWriter,
			logrus.TraceLevel: traceWriter,
			logrus.FatalLevel: fatalWriter,
			logrus.PanicLevel: panicWriter,
		}, &logrus.TextFormatter{},
	))
	return Log
}

func New(options *Options) *logger {
	if options == nil {
		options = DefaultOptions()
	} else if options.LogType == "" {
		options.LogType = LOG_TYPE_FILE
	} else if options.LogFilePath == "" {
		//options.LogFilePath = "logs"
	}
	logger := logger{}
	if options.LogType == LOG_TYPE_FILE {
		logger.runtime = options.LogFilePath
		logger.newLog = logger.newFile()
	} else {
		logger.newLog = logrus.New()
	}
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
		LogType:     LOG_TYPE_FILE,
		LogFilePath: "logs",
	}
}
