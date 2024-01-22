package loggerx

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// 日志文件的计算

// 获取最新的文件名
func nowFileName() string {
	// ioc, _ := time.LoadLocation("Asia/Shanghai")
	// timeDir := fmt.Sprint(time.Now().In(ioc).Format("2006/01/02/15")) // 2006-01-02 15:04:05
	timeDir := fmt.Sprint(time.Now().Local().Format("2006/01/02")) // 2006-01-02 15:04:05
	path := "./log/" + timeDir + ".log"
	// fmt.Println(filepath.Abs(path))
	return path
}

// 新建文件
func (l *Logger) createNewFile(isMust bool) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	fileName := nowFileName()

	if !isMust {
		if l.file != nil && l.fileName == fileName {
			return nil
		}
	}

	dir, _ := filepath.Split(fileName) // 识别目录与文件
	os.MkdirAll(dir, os.ModePerm)      // 创建多层目录，如果存在不会报错

	// 打开该文件，如果不存在则创建
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		// 打开失败，尝试创建
		fmt.Println("打开日志文件失败")
		return err
	}
	// 关闭原来的文件
	if l.file != nil {
		l.closeFile()
	}
	l.file = file
	l.fileName = fileName
	return nil
}

// 关闭文件
func (l *Logger) closeFile() error {
	l.file.Sync()
	return l.file.Close()
}
