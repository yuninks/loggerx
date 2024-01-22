package loggerx

import (
	"context"
	"encoding/json"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func (l *Logger) logger(ctx context.Context, event string, v ...any) {
	pc, file, line, _ := runtime.Caller(2)
	// fmt.Println("runtime.Caller", pc, file, line, ok)

	basePath, _ := filepath.Abs("./")
	basePath = strings.ReplaceAll(basePath, "\\", "/")
	// fmt.Println("basePath", basePath)

	file = strings.TrimPrefix(file, basePath)

	funcName := runtime.FuncForPC(pc).Name()
	funcName = filepath.Ext(funcName)
	funcName = strings.TrimPrefix(funcName, ".")

	by, _ := json.Marshal(v)

	nowTime := time.Now().Local().Format("2006-01-02 15:04:05.000000")

	traceId, _ := ctx.Value(l.option.traceField).(string)

	writeStr := "[" + event + "]" + nowTime + " " + file + ":" + fmt.Sprintf("%d", line) + " " + funcName + " gid:" + getGID() + " " + traceId + " @data@: " + string(by) + "\n\n"

	l.write(event, []byte(writeStr))

	if l.option.errorToInfo && event == "error" {
		l.write("info", []byte(writeStr))
	}

	// log.Println("" + string(by))
}
