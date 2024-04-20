package logging

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

var stringToLogLevel = map[string]zerolog.Level{
	"debug": zerolog.DebugLevel,
	"info":  zerolog.InfoLevel,
	"warn":  zerolog.WarnLevel,
	"error": zerolog.ErrorLevel,
}

func LogLevelFromString(level string) zerolog.Level {
	return stringToLogLevel[level]
}

func GetLogger(w io.Writer, level string) zerolog.Logger {
	return zerolog.New(w).Level(LogLevelFromString(level)).With().Timestamp().Logger()
}

func CreateOrOpenLogFile() (*os.File, error) {
	currentDate := time.Now().Format("2006-01-02")
	filename := "logs/" + currentDate + ".log"
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func ChangeLogLevel(level string) {

}
