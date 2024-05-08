package loggerx

import "io"

type loggerOption struct {
	prefix      string // 日志前缀
	format      string // text json
	dir         string // 文件目录
	isGinLog    bool
	isGid       bool
	traceField  string      // trace字段
	errorToInfo bool        // 错误日志是否写入info日志
	days        int         // 日志保存天数
	drivers     []io.Writer // 文件落盘驱动器
	fileSplit   FileSplit   // 文件切割规则
	sizeSplit   int         // 根据文件大小切割
}

func defaultOptions() loggerOption {
	return loggerOption{
		isGinLog:   true,
		isGid:      true,
		format:     "json",
		dir:        "./log",
		traceField: "trace_id",
		days:       7,
		fileSplit:  FileSplitTimeE,
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
func SetGinLog(open bool) Option {
	return func(o *loggerOption) {
		o.isGinLog = open
	}
}

// 文件路径
func SetDir(dir string) Option {
	return func(o *loggerOption) {
		if dir != "" {
			o.dir = dir
		}
	}
}

// 保存goroutine的ID信息
func SetGID(open bool) Option {
	return func(o *loggerOption) {
		o.isGid = open
	}
}

// 日志保存天数
func SetDays(days int) Option {
	return func(o *loggerOption) {
		o.days = days
	}
}

// 设置时区
func SetTimeZone() Option {
	return func(o *loggerOption) {
		// o.timeZone = timeZone
	}
}

// 文件额外的驱动
func SetExtraDriver(ds ...io.Writer) Option {
	return func(o *loggerOption) {
		for _, d := range ds {
			if d != nil {
				o.drivers = append(o.drivers, d)
			}
		}
	}
}

// 文件切割规则
// 1.文件大小
// 2.时间A（年/月/日/时）
// 3.时间B（年/月-日）
// 4.时间C（年-月-日-时）
// 5.时间D（年-月-日）
func SetFileSplit(split FileSplit) Option {
	return func(o *loggerOption) {
		o.fileSplit = split
	}
}

type FileSplit string

const (
	FileSplitNone  FileSplit = "none"  // 不切割
	FileSplitTimeA FileSplit = "timeA" // （年/月/日/时）
	FileSplitTimeB FileSplit = "timeB" // （年/月/日）
	FileSplitTimeC FileSplit = "timeC" // （年/月-日）
	FileSplitTimeD FileSplit = "timeD" // （年-月-日-时）
	FileSplitTimeE FileSplit = "timeE" // （年-月-日）

)

// 根据文件大小切割(暂时未生效)
func SetSizeSplit(m int) Option {
	return func(o *loggerOption) {
		o.sizeSplit = m
	}
}
