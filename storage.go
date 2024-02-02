package loggerx

import (
	"io"
)

func (l *Logger) write(event string, b []byte) (n int, err error) {
	f, err := l.getFile(event, false)
	if err != nil {
		return 0, err
	}

	n, err = f.Write(b)
	if err == nil && n < len(b) {
		err = io.ErrShortWrite
	}
	if err != nil {
		// 强制更新
		l.getFile(event, true)
	}
	return f.Write(b)
}
