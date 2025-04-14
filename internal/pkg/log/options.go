package log

import "go.uber.org/zap/zapcore"

type Options struct {
	DisableCaller     bool
	DisableStacktrace bool
	Level             string
	Format            string
	OutputPaths       []string
}

// NewOptions 创建一个带有默认参数的 Options 对象.
func NewOptions() *Options {
	return &Options{
		DisableCaller:     false,
		DisableStacktrace: false,
		Level:             zapcore.InfoLevel.String(),
		Format:            "console",
		OutputPaths:       []string{"stdout"},
	}
}
