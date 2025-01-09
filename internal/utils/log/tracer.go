package custom_log

import (
	"context"
	"fmt"
	"log/slog"
	"regexp"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/tracelog"
)

var (
	replaceTabs                      = regexp.MustCompile(`\t+`)
	replaceSpacesBeforeOpeningParens = regexp.MustCompile(`\s+\(`)
	replaceSpacesAfterOpeningParens  = regexp.MustCompile(`\(\s+`)
	replaceSpacesBeforeClosingParens = regexp.MustCompile(`\s+\)`)
	replaceSpacesAfterClosingParens  = regexp.MustCompile(`\)\s+`)
	replaceSpaces                    = regexp.MustCompile(`\s+`)
)

type LogTracer struct {
	logger            *Logger
	level             tracelog.LogLevel
	enableArgsLogging bool
	slowThreshold     time.Duration
}

func NewLogTracer(
	logger *Logger,
	level tracelog.LogLevel,
	enableArgsLogging bool,
	slowThreshold time.Duration,
) *LogTracer {
	return &LogTracer{
		logger:            logger,
		level:             level,
		enableArgsLogging: enableArgsLogging,
		slowThreshold:     slowThreshold,
	}
}

func (t *LogTracer) TraceQueryStart(
	c context.Context,
	_ *pgx.Conn,
	data pgx.TraceQueryStartData,
) context.Context {
	c = context.WithValue(c, "query", data.SQL)
	c = context.WithValue(c, "args", data.Args)
	c = context.WithValue(c, "start", time.Now())

	return c
}

// prettyPrintSQL removes empty lines and trims spaces.
func prettyPrintSQL(sql string) string {
	lines := strings.Split(sql, "\n")

	pretty := strings.Join(lines, " ")
	pretty = replaceTabs.ReplaceAllString(pretty, "")
	pretty = replaceSpacesBeforeOpeningParens.ReplaceAllString(pretty, "(")
	pretty = replaceSpacesAfterOpeningParens.ReplaceAllString(pretty, "(")
	pretty = replaceSpacesAfterClosingParens.ReplaceAllString(pretty, ")")
	pretty = replaceSpacesBeforeClosingParens.ReplaceAllString(pretty, ")")

	// Finally, replace multiple spaces with a single space
	pretty = replaceSpaces.ReplaceAllString(pretty, " ")

	return strings.TrimSpace(pretty)
}

func (t *LogTracer) TraceQueryEnd(
	c context.Context,
	_ *pgx.Conn,
	data pgx.TraceQueryEndData,
) {
	message := "query completed"
	sql, ok := c.Value("query").(string)
	if ok {
		sql = prettyPrintSQL(sql)
	}

	args, ok := c.Value("args").([]interface{})
	if !ok {
		args = nil
	}

	level := t.level
	var latencyString string
	var d time.Duration
	start, ok := c.Value("start").(time.Time)
	if !ok {
		d = 0
	} else {
		d = time.Since(start)
		if d >= t.slowThreshold {
			level = tracelog.LogLevelWarn
			message = "slow query completed"
		}

		switch {
		case d >= time.Second:
			latencyString = fmt.Sprintf(
				"%03ds",
				int64(d.Seconds()),
			)
		case d >= time.Millisecond:
			latencyString = fmt.Sprintf(
				"%03dms",
				int64(d.Milliseconds()),
			)
		case d >= time.Microsecond:
			latencyString = fmt.Sprintf(
				"%03dµs",
				int64(d.Microseconds()),
			)
		default:
			latencyString = fmt.Sprintf(
				"%03dns",
				d.Nanoseconds(),
			)
		}
	}

	dt := map[string]interface{}{
		"query":   sql,
		"runtime": latencyString,
	}

	if data.Err != nil {
		dt["error"] = data.Err
		dt["command_tag"] = data.CommandTag
		level = tracelog.LogLevelError
		message = data.Err.Error()
	}

	if t.enableArgsLogging {
		dt["args"] = args
	}

	t.logger.Log(c, level, message, dt)
}

type Logger struct {
	l               *slog.Logger
	invalidLevelKey string
}

type Option func(*Logger)

func WithInvalidLevelKey(key string) Option {
	return func(l *Logger) {
		l.invalidLevelKey = key
	}
}

func NewLogger(l *slog.Logger, options ...Option) *Logger {
	logger := &Logger{
		l:               l,
		invalidLevelKey: "INVALID_PGX_LOG_LEVEL",
	}

	for _, option := range options {
		option(logger)
	}

	return logger
}

func (l *Logger) Log(
	ctx context.Context,
	level tracelog.LogLevel,
	msg string,
	data map[string]interface{},
) {
	attrs := make([]slog.Attr, 0, len(data))
	for k, v := range data {
		attrs = append(attrs, slog.Any(k, v))
	}

	var lvl slog.Level
	switch level {
	case tracelog.LogLevelTrace:
		lvl = slog.LevelDebug - 1
		attrs = append(attrs, slog.Any("PGX_LOG_LEVEL", level))
	case tracelog.LogLevelDebug:
		lvl = slog.LevelDebug
	case tracelog.LogLevelInfo:
		lvl = slog.LevelInfo
	case tracelog.LogLevelWarn:
		lvl = slog.LevelWarn
	case tracelog.LogLevelError:
		lvl = slog.LevelError
	default:
		lvl = slog.LevelError
		attrs = append(
			attrs,
			slog.Any(l.invalidLevelKey, fmt.Errorf("invalid pgx log level: %v", level)),
		)
	}
	l.l.LogAttrs(ctx, lvl, msg, attrs...)
}
