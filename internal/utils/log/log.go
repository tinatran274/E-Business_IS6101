package custom_log

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"strconv"
	"sync"
)

type Key string

const (
	timeFormat = "[2006-01-02 15:04:05.000]"
	reset      = "\033[0m"

	SlogFieldsKey Key = "slog_fields"

	black        = 30
	red          = 31
	green        = 32
	yellow       = 33
	blue         = 34
	magenta      = 35
	cyan         = 36
	lightGray    = 37
	darkGray     = 90
	lightRed     = 91
	lightGreen   = 92
	lightYellow  = 93
	lightBlue    = 94
	lightMagenta = 95
	lightCyan    = 96
	white        = 97
)

var (
	PostColor   string = fmt.Sprintf("\033[%sm%s%s", strconv.Itoa(lightGreen), "PST", "\033[0m")
	GetColor    string = fmt.Sprintf("\033[%sm%s%s", strconv.Itoa(lightBlue), "GET", "\033[0m")
	PutColor    string = fmt.Sprintf("\033[%sm%s%s", strconv.Itoa(lightYellow), "PUT", "\033[0m")
	DeleteColor string = fmt.Sprintf("\033[%sm%s%s", strconv.Itoa(lightRed), "DEL", "\033[0m")
	PatchColor  string = fmt.Sprintf("\033[%sm%s%s", strconv.Itoa(lightMagenta), "PTC", "\033[0m")

	SuccessColor string = fmt.Sprintf(
		"\033[%sm%%v%s",
		strconv.Itoa(green),
		"\033[0m",
	)

	FailColor string = fmt.Sprintf(
		"\033[%sm%%v%s",
		strconv.Itoa(yellow),
		"\033[0m",
	)

	ErrorColor string = fmt.Sprintf(
		"\033[%sm%%v%s",
		strconv.Itoa(red),
		"\033[0m",
	)

	validFields = map[string]bool{
		"request_id": true,
		"user_id":    true,
		"api_key":    true,
	}
)

type Handler struct {
	h slog.Handler
	b *bytes.Buffer
	m *sync.Mutex

	client LogClient
}

func NewHandler(opts *slog.HandlerOptions, client LogClient) *Handler {
	if opts == nil {
		opts = &slog.HandlerOptions{}
	}
	b := &bytes.Buffer{}
	return &Handler{
		b: b,
		h: slog.NewJSONHandler(b, &slog.HandlerOptions{
			Level:       opts.Level,
			AddSource:   opts.AddSource,
			ReplaceAttr: suppressDefaults(opts.ReplaceAttr),
		}),
		m: &sync.Mutex{},

		client: client,
	}
}

func NewLogHandler(h slog.Handler) *Handler {
	return &Handler{
		h: h,
		b: &bytes.Buffer{},
		m: &sync.Mutex{},
	}
}

func (h *Handler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.h.Enabled(ctx, level)
}

func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &Handler{h: h.h.WithAttrs(attrs), b: h.b, m: h.m}
}

func (h *Handler) WithGroup(name string) slog.Handler {
	return &Handler{h: h.h.WithGroup(name), b: h.b, m: h.m}
}

func (h *Handler) Handle(ctx context.Context, r slog.Record) error {
	var level string
	level = r.Level.String() + ":"
	switch r.Level {
	case slog.LevelDebug:
		level = "DEB"
		level = colorize(darkGray, level)
	case slog.LevelInfo:
		level = "INF"
		level = colorize(cyan, level)
	case slog.LevelWarn:
		level = "WRN"
		level = colorize(lightYellow, level)
	case slog.LevelError:
		level = "ERR"
		level = colorize(lightRed, level)
	}

	if attrs, ok := ctx.Value(SlogFieldsKey).([]slog.Attr); ok {
		r.AddAttrs(attrs...)
	}

	attrs, err := h.computeAttrs(ctx, r)
	if err != nil {
		return err
	}

	attrBytes, err := json.MarshalIndent(attrs, "", "  ")
	if err != nil {
		return fmt.Errorf("error when marshaling attrs: %w", err)
	}

	if _, err := fmt.Println(
		colorize(lightGray, r.Time.Format(timeFormat)),
		level,
		colorize(white, r.Message),
		colorize(darkGray, string(attrBytes)),
	); err != nil {
		return err
	}

	logMessage := fmt.Sprintf(
		"%s %s",
		colorize(lightGray, r.Time.Format(timeFormat)),
		colorize(white, r.Message),
	)

	type RequestInfo struct {
		RequestId string `json:"request_id"`
	}
	var requestInfo RequestInfo

	if err := json.Unmarshal(attrBytes, &requestInfo); err != nil {
		return fmt.Errorf("error when unmarshaling request info: %w", err)
	}

	if h.client != nil {
		go func() {
			h.client.Log(r.Level, logMessage, attrs)
		}()
	}

	return nil
}

func colorize(colorCode int, v string) string {
	return fmt.Sprintf("\033[%sm%s%s", strconv.Itoa(colorCode), v, reset)
}

func (h *Handler) computeAttrs(
	ctx context.Context,
	r slog.Record,
) (map[string]any, error) {
	h.m.Lock()
	defer func() {
		h.b.Reset()
		h.m.Unlock()
	}()
	if err := h.h.Handle(ctx, r); err != nil {
		return nil, fmt.Errorf("error when calling inner handler's Handle: %w", err)
	}

	var attrs map[string]any
	err := json.Unmarshal(h.b.Bytes(), &attrs)
	if err != nil {
		return nil, fmt.Errorf("error when unmarshaling inner handler's Handle result: %w", err)
	}

	return attrs, nil
}

func suppressDefaults(
	next func([]string, slog.Attr) slog.Attr,
) func([]string, slog.Attr) slog.Attr {
	return func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey ||
			a.Key == slog.LevelKey ||
			a.Key == slog.MessageKey {
			return slog.Attr{}
		}
		if next == nil {
			return a
		}
		return next(groups, a)
	}
}
