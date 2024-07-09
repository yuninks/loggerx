package loggerx

import (
	"io"
)

func (l *Logger) write(event string, b []byte) (n int, err error) {

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
