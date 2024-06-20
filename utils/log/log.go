package log

import (
	"fmt"
	"io"
	"log/syslog"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"gopkg.in/natefinch/lumberjack.v2"
)

var instance zerolog.Logger

type Syslog struct {
	Protocol string `env:"LOG_SYSLOG_PROTOCOL"`
	IP       string `env:"LOG_SYSLOG_IP"`
	Port     int    `env:"LOG_SYSLOG_PORT"`
	Tag      string `env:"LOG_SYSLOG_TAG"`
}

type Config struct {
	Level         string `env:"LOG_LEVEL" env-default:"debug"`
	Access        string `env:"LOG_ACCESS" env-default:"logs/access.log"`
	App           string `env:"LOG_APP" env-default:"logs/app.log"`
	Error         string `env:"LOG_ERROR" env-default:"logs/error.log"`
	Debug         string `env:"LOG_DEBUG" env-default:"logs/debug.log"`
	EnableConsole bool   `env:"LOG_ENABLE_CONSOLE" env-default:"true"`
	TimeFormat    string `env:"LOG_TIME_FORMAT" env-default:"2020-01-01 13:01:01"`
	Rotation      struct {
		MaxSize   int `env:"LOG_ROTATION_MAXSIZE" env-default:"20"`
		MaxBackup int `env:"LOG_ROTATION_MAX_BACKUP" env-default:"5"`
		MaxAge    int `env:"LOG_ROTATION_MAX_AGE" env-default:"0"`
	}
	Syslog Syslog
}

type CallerHook struct {
	Filter string
}

func (callerHook CallerHook) getFilePath(file string) string {
	pathList := strings.Split(file, "/")
	var i = 0
	for i = 0; i < len(pathList); i++ {
		if pathList[i] == callerHook.Filter {
			break
		}
	}
	return filepath.Join(pathList[i+1:]...)
}

func (callerHook CallerHook) Run(e *zerolog.Event, _ zerolog.Level, _ string) {
	for i := 0; i < 10; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok && strings.Contains(file, callerHook.Filter) && filepath.Base(file) != "log.go" && filepath.Base(file) != "gin_util.go" {
			e.Str("line", fmt.Sprintf("%s:%d", callerHook.getFilePath(file), line))
			break
		}
		if !ok {
			break
		}
	}
}

type FilteredWriter struct {
	w             zerolog.LevelWriter
	level         zerolog.Level
	isExactly     bool
	enableNoLevel bool
}

func (fw *FilteredWriter) Write(p []byte) (n int, err error) {
	return fw.w.Write(p)
}
func (fw *FilteredWriter) WriteLevel(level zerolog.Level, p []byte) (n int, err error) {
	if fw.isExactly {
		if level == fw.level {
			return fw.w.WriteLevel(level, p)
		}
	} else if level >= fw.level {
		if level != zerolog.NoLevel || (fw.level == zerolog.NoLevel && level == zerolog.NoLevel) {
			return fw.w.WriteLevel(level, p)
		} else if fw.enableNoLevel && level == zerolog.NoLevel {
			return fw.w.WriteLevel(level, p)
		}
	}
	return len(p), nil
}

func InitLog(config Config) error {
	// Log level
	levelStr := config.Level
	level, err := zerolog.ParseLevel(levelStr)
	if err != nil {
		return err
	}

	var writers []io.Writer

	// App log
	appLogPath := config.App
	appWriter, err := createRollingFileWriter(appLogPath, config)
	if err != nil {
		return err
	}
	writers = append(writers, &FilteredWriter{
		w:             zerolog.MultiLevelWriter(appWriter),
		level:         level,
		isExactly:     false,
		enableNoLevel: true,
	})

	// Error log
	if config.Error != "" {
		errorLogPath := config.Error
		errorWriter, err2 := createRollingFileWriter(errorLogPath, config)
		if err2 != nil {
			return err2
		}
		writers = append(writers, &FilteredWriter{
			w:         zerolog.MultiLevelWriter(errorWriter),
			level:     zerolog.WarnLevel,
			isExactly: false,
		})
	}

	// Console log
	if config.EnableConsole {
		consoleWriter := createConsoleWriterWithFormat(os.Stdout, true)
		writers = append(writers, &FilteredWriter{
			w:             zerolog.MultiLevelWriter(consoleWriter),
			level:         level,
			isExactly:     false,
			enableNoLevel: true,
		})
	}

	// Debug log
	if config.Debug != "" {
		debugLogPath := config.Debug
		debugWriter, err2 := createRollingFileWriter(debugLogPath, config)
		if err2 != nil {
			return err2
		}
		writers = append(writers, &FilteredWriter{
			w:         zerolog.MultiLevelWriter(createConsoleWriterWithFormat(debugWriter, true)),
			level:     zerolog.TraceLevel,
			isExactly: false,
		})
	}

	// Access log
	if config.Access != "" {
		accessLogPath := config.Access
		accessWriter, err2 := createRollingFileWriter(accessLogPath, config)
		if err2 != nil {
			return err2
		}
		writers = append(writers, &FilteredWriter{
			w:         zerolog.MultiLevelWriter(accessWriter),
			level:     zerolog.NoLevel,
			isExactly: true,
		})
	}

	if config.Syslog.IP != "" {
		syslogWriter, err2 := syslog.Dial(config.Syslog.Protocol, fmt.Sprintf("%s:%d", config.Syslog.IP, config.Syslog.Port),
			syslog.LOG_INFO, config.Syslog.Tag)
		if err2 != nil {
			return err2
		}
		writers = append(writers, &FilteredWriter{
			w:             zerolog.MultiLevelWriter(zerolog.SyslogLevelWriter(syslogWriter)),
			level:         level,
			isExactly:     false,
			enableNoLevel: true,
		})
	}
	w := zerolog.MultiLevelWriter(writers...)
	instance = zerolog.New(w).With().Timestamp().Logger()
	if level == zerolog.DebugLevel {
		instance = instance.Hook(CallerHook{Filter: "iot-platform-api"})
	}
	return nil
}

func createRollingFileWriter(filePath string, config Config) (io.Writer, error) {
	dirPath := filepath.Dir(filePath)
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err = os.MkdirAll(dirPath, 0744); err != nil {
			return nil, err
		}
	}
	return &lumberjack.Logger{
		Filename:   filePath,
		MaxBackups: config.Rotation.MaxBackup,
		MaxSize:    config.Rotation.MaxSize,
		MaxAge:     config.Rotation.MaxAge,
	}, nil
}

func createConsoleWriterWithFormat(out io.Writer, isNoColor bool) zerolog.ConsoleWriter {
	consoleWriter := zerolog.ConsoleWriter{Out: out, TimeFormat: time.RFC3339, NoColor: isNoColor}
	consoleWriter.FormatLevel = func(i interface{}) string {
		var l string
		if ll, ok := i.(string); ok {
			switch ll {
			case "trace":
				l = "TRC"
			case "debug":
				l = "DBG"
			case "info":
				l = "INF"
			case "warn":
				l = "WRN"
			case "error":
				l = "ERR"
			case "fatal":
				l = "FTL"
			case "panic":
				l = "PNC"
			default:
				l = "ACS"
			}
		} else {
			if i == nil {
				l = "ACS"
			} else {
				l = strings.ToUpper(fmt.Sprintf("%s", i))[0:3]
			}
		}
		return "| " + l
	}
	consoleWriter.FormatMessage = func(i interface{}) string {
		if i == nil {
			return "|" + " [] "
		}
		return "|" + fmt.Sprintf(" [%s] ", i)
	}
	consoleWriter.FormatFieldName = func(i interface{}) string {
		return "|" + fmt.Sprintf(" %s : ", i)
	}
	consoleWriter.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s ", i)
	}
	consoleWriter.FormatErrFieldName = func(i interface{}) string {
		return "|" + fmt.Sprintf(" %s : ", i)
	}
	consoleWriter.FormatErrFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s ", i)
	}
	return consoleWriter
}

func Info() *zerolog.Event {
	return instance.Info()
}

func Debug() *zerolog.Event {
	return instance.Debug()
}

func Error() *zerolog.Event {
	return instance.Error()
}

func Log() *zerolog.Event {
	return instance.Log()
}

func Trace() *zerolog.Event {
	return instance.Trace()
}

func Warn() *zerolog.Event {
	return instance.Warn()
}
