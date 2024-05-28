package loggerx

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// 监听dir文件夹的所有文件，循环查找是否有超过days天的文件，删除掉

func (l *Logger) delete() {
	err := filepath.Walk(l.option.dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		if l.option.days <= 0 {
			l.option.days = 1
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

	if err != nil {
		fmt.Println(err)
	}

	// 监听文件夹，获取新加入的文件
	

}

func isEmptyDir(path string) bool {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return len(files) == 0 // 文件夹为空
}
