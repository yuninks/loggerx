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

func (l *Logger) logger(ctx context.Context, action string, v ...any) {
	pc, file, line, _ := runtime.Caller(2)
	// fmt.Println("runtime.Caller", pc, file, line, ok)

	funcName := runtime.FuncForPC(pc).Name()
	funcName = filepath.Ext(funcName)
	funcName = strings.TrimPrefix(funcName, ".")

	by, _ := json.Marshal(v)

	nowTime := time.Now().Local().Format("20060102 15:04:05.000000")

	traceId, _ := ctx.Value("trace_id").(string)

	writeStr := "[" + action + "]" + nowTime + " " + file + ":" + fmt.Sprintf("%d", line) + " " + funcName + " gid:" + getGID() + " " + traceId + " @data@: " + string(by) + "\n\n"

	l.Write([]byte(writeStr))

	// log.Println("" + string(by))
}
