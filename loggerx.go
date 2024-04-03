package loggerx

// author：黄新云
// lastTime:2024年4月3日21:18:05
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

	// 验证文件夹权限
	if !checkDir(opt.dir) {
		panic("文件夹权限不足")
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

	// 日志删除
	go l.delete()

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

// 验证文件夹权限
// 根文件夹如果不存在则创建
func checkDir(dir string) bool {
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				log.Println("创建文件夹失败", err)
				return false
			}
		}
	}
	return true
}
