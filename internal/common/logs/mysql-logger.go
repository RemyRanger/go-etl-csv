package logs

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

// Colors
const (
	Reset       = "\033[0m"
	Red         = "\033[31m"
	Green       = "\033[32m"
	Yellow      = "\033[33m"
	Blue        = "\033[34m"
	Magenta     = "\033[35m"
	Cyan        = "\033[36m"
	White       = "\033[37m"
	BlueBold    = "\033[34;1m"
	MagentaBold = "\033[35;1m"
	RedBold     = "\033[31;1m"
	YellowBold  = "\033[33;1m"
)

// MysqlLogger : Mysql custom Logger
type MysqlLogger struct {
}

func NewMysqlLogger() *MysqlLogger {
	return &MysqlLogger{}
}

// LogMode log mode
func (l *MysqlLogger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

// Info print info
func (l MysqlLogger) Info(ctx context.Context, msg string, data ...any) {
	log.Info().Msgf(msg, data...)
}

// Warn print warn messages
func (l MysqlLogger) Warn(ctx context.Context, msg string, data ...any) {
	log.Warn().Msgf(msg, data...)
}

// Error print error messages
func (l MysqlLogger) Error(ctx context.Context, msg string, data ...any) {
	log.Error().Msgf(msg, data...)
}

// Trace print sql message
func (l MysqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)

	traceStr := "%s\n" + "[%.3fms] " + "[rows:%v]" + " %s"

	sql, rows := fc()
	if err != nil {
		if rows == -1 {
			log.Debug().Msgf(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			log.Debug().Msgf(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	} else {
		if rows == -1 {
			log.Trace().Msgf(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			log.Trace().Msgf(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
