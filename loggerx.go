package loggerx

// author：黄新云
// lastTime:2023年6月30日21:28:04
// desc: 日志封装类

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"sync"

	// "sync_log/global"

	"github.com/gin-gonic/gin"
)

// 需要实现io.Writer接口
type Logger struct {
	filePath *sync.Map // filePath
	mu       *sync.Mutex
	option   loggerOption
	channel  string
}

type filePath struct {
	file     *os.File
	fileName string
}

func NewLogger(opts ...Option) *Logger {
	opt := defaultOptions()
	for _, apply := range opts {
		apply(&opt)
	}

	l := &Logger{
		filePath: &sync.Map{},
		mu:       &sync.Mutex{},
		option:   opt,
	}

	log.SetOutput(l)
	log.SetFlags(log.LstdFlags | log.Llongfile | log.Lmicroseconds) // log.Lshortfile  | log.LUTC

	// 保存Gin日志写入到文件+控制台
	if opt.isGinLog {
		gin.DefaultWriter = io.MultiWriter(l, os.Stdout)
		gin.DefaultErrorWriter = io.MultiWriter(l, os.Stdout)
	}
	return l
}

// 强制刷盘
func (l *Logger) MustSync() {
	l.filePath.Range(func(key, value any) bool {
		f := value.(*filePath)
		f.file.Sync()
		return true
	})
}

func (l *Logger) Channel(ch string) (r *Logger) {
	rr := *l
	rr.channel = ch
	return &rr
}

// 实现io.Writer接口
func (l *Logger) Write(b []byte) (n int, err error) {
	return l.write("info", b)
}

func (l *Logger) Info(ctx context.Context, v ...any) {
	l.logger(ctx, "info", v...)
}

func (l *Logger) Infof(ctx context.Context, format string, v ...any) {
	s := fmt.Sprintf(format, v...)
	l.logger(ctx, "info", s)
}

func (l *Logger) Error(ctx context.Context, v ...any) {
	l.logger(ctx, "error", v...)
}

func (l *Logger) Errorf(ctx context.Context, format string, v ...any) {
	s := fmt.Sprintf(format, v...)
	l.logger(ctx, "error", s)
}

// 添加固定的内容
// func (l *Logger) ContextWithFields(ctx context.Context, v ...any) {
// 	l.logger(ctx, "add", v...)
// }
// func (l *Logger) Field(key,val string) {
// 	l.logger(nil, "add", key,val)
// }

func getGID() string {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	return string(b)
}
