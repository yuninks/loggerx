package loggerx

// author：黄新云
// lastTime:2023年6月30日21:28:04
// desc: 日志封装类

// 优化方向
// 可以写盘的时候添加缓存，但是有个难点就是如果采用缓存的方式必须保证这个是最后关闭的，如果不是则退出的时候会丢失日志（写在缓存里还没刷盘）

// TODO:自动清除过期日志

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
	file     *os.File
	fileName string
	mu       sync.Mutex
	// channel  string
}

func NewLogger(opts ...Option) *Logger {
	opt := defaultOptions()
	for _, apply := range opts {
		apply(&opt)
	}

	l := &Logger{}

	// 打开文件
	err := l.createNewFile(true)
	if err != nil {
		panic(err)
	}
	log.SetOutput(l)
	log.SetFlags(log.LstdFlags | log.Llongfile | log.Lmicroseconds) // log.Lshortfile  | log.LUTC

	// 保存Gin日志写入到文件+控制台
	gin.DefaultWriter = io.MultiWriter(l, os.Stdout)
	gin.DefaultErrorWriter = io.MultiWriter(l, os.Stdout)

	// 赋值前缀
	if prefix != "" {
		log.SetPrefix(fmt.Sprintf("[%s]", prefix))
	}
	return l
}

// func (l *Logger) Channel(ch string) (r *Logger) {
// 	rr := l
// 	rr.channel = ch
// 	return rr
// }

// 超时删除

func (l *Logger) Write(b []byte) (n int, err error) {

	if l.file == nil {
		// 新建一个file连接
		l.createNewFile(false)
	}
	if l.fileName != nowFileName() {
		l.createNewFile(false)
	}

	n, err = l.file.Write(b)
	if err == nil && n < len(b) {
		err = io.ErrShortWrite
	}
	if err != nil {
		// 强制更新
		l.createNewFile(true)
	}

	return n, err
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
