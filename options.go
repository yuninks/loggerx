package loggerx

type loggerOption struct {
	prefix      string // 日志前缀
	format      string // text json
	dir         string // 文件目录
	isGinLog    bool
	isGid       bool
	traceField  string // trace字段
	errorToInfo bool   //
}

func defaultOptions() loggerOption {
	return loggerOption{
		format:     "json",
		dir:        "./log",
		traceField: "trace_id",
	}
}

type Option func(*loggerOption)

// trace字段
func SetTraceField(traceField string) Option {
	return func(o *loggerOption) {
		o.traceField = traceField
	}
}

// 错误日志是否写入info日志
func SetErrorToInfo() Option {
	return func(o *loggerOption) {
		o.errorToInfo = true
	}
}

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

// 文件路径
func SetDir(dir string) Option {
	return func(o *loggerOption) {
		o.dir = dir
	}
}

// 保存goroutine的ID信息
func SetGID() Option {
	return func(o *loggerOption) {
		o.isGid = true
	}
}

// 文件切割规则
// 1.文件大小
// 2.时间A（年/月/日/时）
// 3.时间B（年/月-日）
// 4.时间C（年-月-日-时）
// 5.时间D（年-月-日）
// func SetFileSplit()

//
