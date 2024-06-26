package loggerx

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// 监听dir文件夹的所有文件，循环查找是否有超过days天的文件，删除掉

func (l *Logger) delete() {

	tick := time.NewTicker(time.Hour)

	for {
		select {
		case <-tick.C:
			err := l.walkAndDel()
			if err != nil {
				fmt.Println(err)
			}
		case <-l.ctx.Done():
			return
		}
	}
}

func (l *Logger) walkAndDel() error {
	err := filepath.Walk(l.option.dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		// 没设置删除天数
		if l.option.days <= 0 {
			return nil
		}

		// 判断最后修改时间是否大于3天
		if info.ModTime().After(time.Now().AddDate(0, 0, -l.option.days)) {
			return nil
		}

		if info.IsDir() {
			if !isEmptyDir(path) {
				return nil
			}
		}
		ext := filepath.Ext(path)
		if ext != ".log" {
			return nil
		}
		// 删除文件
		fmt.Println("删除文件", path)
		return os.Remove(path)
	})

	return err
}

// 判断空文件夹
func isEmptyDir(path string) bool {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return len(files) == 0 // 文件夹为空
}
