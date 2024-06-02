package logger

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"runtime"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// See https://gorm.io/docs/logger.html
type gormLogger struct {
	slogger                   *slog.Logger
	IgnoreRecordNotFoundError bool
}

// logCommon This function creates slog messages with correct source code locations
func (l *gormLogger) logCommon(
	ctx context.Context,
	level slog.Level,
	msg string,
	args ...interface{},
) {
	if !l.slogger.Enabled(ctx, level) {
		return
	}

	var pcs [1]uintptr
	runtime.Callers(4, pcs[:]) // skip [Callers, Infof]

	r := slog.NewRecord(time.Now(), level, msg, pcs[0])
	r.Add(args...)
	_ = l.slogger.Handler().Handle(ctx, r)
}

// GORM uses these log messages in the form:
// Info(ctx, "wurst `%s` from %s\n", brot, utils.FileWithLineNum())
func (l *gormLogger) logMsg(
	ctx context.Context,
	level slog.Level,
	msg string,
	args ...interface{},
) {
	l.logCommon(ctx, level, fmt.Sprintf(msg, args...))
}

func (l *gormLogger) log(
	ctx context.Context,
	level slog.Level,
	msg string,
	args ...interface{},
) {
	l.logCommon(ctx, level, msg, args...)
}

func (l *gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	// The log level here is ignored because the one to use
	// is the level set on the slog instance used.
	return l
}

func (l *gormLogger) Info(ctx context.Context, msg string, args ...interface{}) {
	l.logMsg(ctx, slog.LevelInfo, msg, args...)
}

func (l *gormLogger) Warn(ctx context.Context, msg string, args ...interface{}) {
	l.logMsg(ctx, slog.LevelWarn, msg, args...)
}

func (l *gormLogger) Error(ctx context.Context, msg string, args ...interface{}) {
	l.logMsg(ctx, slog.LevelError, msg, args...)
}

func (l *gormLogger) Trace(
	ctx context.Context,
	begin time.Time,
	fc func() (sql string, rowsAffected int64),
	err error,
) {
	elapsedTime := time.Since(begin)

	if err != nil && (!errors.Is(err, gorm.ErrRecordNotFound) ||
		!l.IgnoreRecordNotFoundError) {
		sql, rows := fc()

		l.log(
			ctx,
			LevelTrace,
			err.Error(),
			slog.Any("error", err),
			slog.String("query", sql),
			slog.Duration("elapsed", elapsedTime),
			slog.Int64("rows", rows),
		)
		return
	}
	sql, rows := fc()

	l.log(
		ctx,
		LevelTrace,
		"SQL query executed",
		slog.String("query", sql),
		slog.Duration("elapsed", elapsedTime),
		slog.Int64("rows", rows),
	)
}

// NewGormLog use slog to inject the gorm logs on it
// sl is the slog to use to send the messages.
// ignoreRecordNotFound when it is true, no trace is logged
// when a record is not found.
// Return a logger.Interface initialized
func NewGormLog(sl *slog.Logger, ignoreRecordNotFound bool) logger.Interface {
	if sl == nil {
		sl = slog.Default()
	}
	return &gormLogger{
		slogger:                   sl,
		IgnoreRecordNotFoundError: ignoreRecordNotFound,
	}
}
