package loggerx

type loggerOption struct {
	prefix   string
	format   string
	isGinLog bool
	isGid    bool
}

func defaultOptions() loggerOption {
	return loggerOption{
		format: "json",
	}
}

type Option func(*loggerOption)

// 日志的前缀
func SetPrefix(prefix string) Option {
	return func(o *loggerOption) {
		o.prefix = prefix
	}
}

// 日志格式(默认json)
func SetFormat(format string) Option {
	return func(o *loggerOption) {
		o.format = format
	}
}

// 是否保存gin的日志
func SetGinLog() Option {
	return func(o *loggerOption) {
		o.isGinLog = true
	}
}

// 文件规则
func SetFilePath(path string) Option {
	return func(o *loggerOption) {
		// o.isGinLog = true
	}
}

// 保存goroutine的ID信息
func SetGID() Option {
	return func(o *loggerOption) {
		o.isGid = true
	}
}

// 文件切割规则
// func SetFileSplit()
