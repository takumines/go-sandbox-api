package config

import (
	"github.com/rs/zerolog"
	"os"
)

type logLevel int

const (
	LevelError logLevel = iota
	LevelWarn
	LevelInfo
	LevelDebug
)

var logLevelMap = map[string]logLevel{
	"error": LevelError,
	"warn":  LevelWarn,
	"info":  LevelInfo,
	"debug": LevelDebug,
}

var zerologLevelMap = map[logLevel]zerolog.Level{
	LevelError: zerolog.ErrorLevel,
	LevelWarn:  zerolog.WarnLevel,
	LevelInfo:  zerolog.InfoLevel,
	LevelDebug: zerolog.DebugLevel,
}

func (level logLevel) ZerologLevel() zerolog.Level {
	if l, ok := zerologLevelMap[level]; ok {
		return l
	}
	return zerolog.InfoLevel
}

func NewLogger() zerolog.Logger {
	logger := zerolog.New(os.Stdout).Level(getZerologLevel()).With().Timestamp().Logger()

	return logger
}

func getLogLevel() logLevel {
	if level, ok := logLevelMap[os.Getenv("LOGLEVEL")]; ok {
		return level
	}
	return LevelInfo
}

func getZerologLevel() zerolog.Level {
	return getLogLevel().ZerologLevel()
}
