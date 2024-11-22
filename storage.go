package loggerx

import (
	"io"
	"sync"
)

// 写入，需要判断同步还是异步
func (l *Logger) write(event string, b []byte) (n int, err error) {
	if l.toAsync(event, b) {
		// fmt.Println("异步写入")
		return len(b), nil
	}
	return l.store(event, b)
}

// 实际的存储
func (l *Logger) store(event string, b []byte) (n int, err error) {

	if l.option.isPrintFile {
		f, err := l.getFile(event, false)
		if err != nil {
			return 0, err
		}

		n, err = f.Write(b)
		if err == nil && n < len(b) {
			err = io.ErrShortWrite
		}
		if err != nil {
			// 强制更新 & 再次写入
			f, err := l.getFile(event, true)
			if err == nil {
				f.Write(b)
			}
		}
	}

	if len(l.option.drivers) > 0 {
		io.MultiWriter(l.option.drivers...).Write(b)
	}
	return n, err
}

var chanStore = make(chan cacheData, 1000)
var chanOnce = sync.Once{}

type cacheData struct {
	logger *Logger
	Event  string
	Data   []byte
}

func (l *Logger) toAsync(event string, b []byte) bool {
	chanOnce.Do(func() {
		go func() {
			for val := range chanStore {
				val.logger.store(val.Event, val.Data)
			}
		}()
	})

	if l.writeType == writeTypeSync || // 指定同步模式
		(l.writeType == writeTypeDefault && l.option.writeType != writeTypeAsync) { // 默认同步模式
		return false
	}

	// 为了避免丢失，还是要阻塞等待
	chanStore <- cacheData{l, event, b}

	return true
}
