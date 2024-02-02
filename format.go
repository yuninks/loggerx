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

	// writeStr := "[" + event + "]" + nowTime + " " + file + ":" + fmt.Sprintf("%d", line) + " " + funcName + " gid:" + getGID() + " " + traceId + " @data@: " + string(by) + "\n\n"

	fd := FormatData{
		Time:    nowTime,
		File:    file + ":" + fmt.Sprintf("%d", line),
		Func:    funcName,
		Gid:     getGID(),
		Content: string(by),
		TraceId: traceId,
	}
	fdb, _ := json.Marshal(fd)

	ff := []byte("\n[" + event + "]")
	fdb = append(ff, fdb...)
	

	l.write(event, fdb)

	if l.option.errorToInfo && event == "error" {
		l.write("info", fdb)
	}

	// log.Println("" + string(by))
}

type FormatData struct {
	Time    string `json:"time,omitempty"`
	File    string `json:"file,omitempty"`
	Func    string `json:"func,omitempty"`
	Gid     string `json:"gid,omitempty"`
	Content string `json:"content,omitempty"`
	TraceId string `json:"traceId,omitempty"`
}
