package loggerx

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// 文件操作

// 获取最新的文件名
func (l *Logger) nowFileName(event string) string {
	// ioc, _ := time.LoadLocation("Asia/Shanghai")
	// timeDir := fmt.Sprint(time.Now().In(ioc).Format("2006/01/02/15")) // 2006-01-02 15:04:05

	prefix := ""

	switch l.option.fileSplit {
	case FileSplitTimeA:
		// （年/月/日/时）
		prefix = fmt.Sprint(time.Now().In(l.option.timeZone).Format("2006/01/02/15")) // 2006-01-02 15:04:05
	case FileSplitTimeB:
		// （年/月/日）
		prefix = fmt.Sprint(time.Now().In(l.option.timeZone).Format("2006/01/02")) // 2006-01-02 15:04:05
	case FileSplitTimeC:
		// （年/月-日）
		prefix = fmt.Sprint(time.Now().In(l.option.timeZone).Format("2006/01-02")) // 2006-01-02 15:04:05
	case FileSplitTimeD:
		// （年-月-日-时）
		prefix = fmt.Sprint(time.Now().In(l.option.timeZone).Format("2006-01-02-15")) // 2006-01-02 15:04:05
	case FileSplitTimeE:
		// （年-月-日）
		prefix = fmt.Sprint(time.Now().In(l.option.timeZone).Format("2006-01-02")) // 2006-01-02 15:04:05
	}

	if prefix != "" {
		prefix = prefix + "_"
	}

	// timeDir := fmt.Sprint(time.Now().Local().Format("2006/01/02")) // 2006-01-02 15:04:05
	if l.channel != "" {
		prefix = l.channel + "/" + prefix
	}
	path := l.option.dir + "/" + prefix + event + ".log"
	// fmt.Println(filepath.Abs(path))
	return path
}

// 新建文件
func (l *Logger) getFile(event string, isRefresh bool) (*os.File, error) {
	f := l.loadFile(event)
	if f != nil && !isRefresh {
		return f, nil
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	fileName := l.nowFileName(event)

	dir, _ := filepath.Split(fileName) // 识别目录与文件
	os.MkdirAll(dir, os.ModePerm)      // 创建多层目录，如果存在不会报错

	// 打开该文件，如果不存在则创建
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		// 打开失败，尝试创建
		fmt.Println("打开日志文件失败")
		return nil, err
	}
	// 关闭原来的
	if f != nil {
		closeFile(f)
	}

	l.filePath.Store(l.channel, &filePath{
		file:     file,
		fileName: fileName,
	})

	return file, nil
}

// 加载文件
func (l *Logger) loadFile(event string) *os.File {
	val, ok := l.filePath.Load(l.channel)
	if !ok {
		return nil
	}
	f := val.(*filePath)
	if f == nil {
		return nil
	}
	if f.fileName != l.nowFileName(event) {
		// 原来的文件需关闭
		closeFile(f.file)
		return nil
	}
	return f.file
}

// 关闭文件
func closeFile(f *os.File) error {
	f.Sync()
	return f.Close()
}
