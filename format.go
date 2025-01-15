package loggerx

import (
	"bytes"
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

	// by, _ := json.Marshal(v)

	nowTime := time.Now().In(l.option.timeZone).Format("2006-01-02 15:04:05.000000")

	traceId, _ := ctx.Value(l.option.traceField).(string)

	// writeStr := "[" + event + "]" + nowTime + " " + file + ":" + fmt.Sprintf("%d", line) + " " + funcName + " gid:" + getGID() + " " + traceId + " @data@: " + string(by) + "\n\n"

	for idx, val := range v {
		if _, ok := val.(error); ok {
			v[idx] = fmt.Sprintf("%+v", val)
		}
	}

	fd := FormatData{
		Time:    nowTime,
		File:    file + ":" + fmt.Sprintf("%d", line),
		Func:    funcName,
		Gid:     getGID(),
		Content: v,
		TraceId: traceId,
	}

	if event == "error" {
		// fd.Stack = string(debug.Stack())
	}

	var fdb []byte
	if l.option.escapeHTML {
		fdb, _ = json.Marshal(fd)
	} else {
		// 非转义
		var buf bytes.Buffer
		encoder := json.NewEncoder(&buf)
		encoder.SetEscapeHTML(false)
		encoder.Encode(fd)
		fdb = buf.Bytes()
	}
	fdb = bytes.TrimRight(fdb, "\n")

	ff := []byte("\n[" + event + "]")
	fdb = append(ff, fdb...)

	l.write(event, fdb)

	if l.option.errorToInfo && event == "error" {
		l.write("info", fdb)
	}

	// log.Println("" + string(by))
}

type FormatData struct {
	Time    string      `json:"time,omitempty"`
	File    string      `json:"file,omitempty"`
	Func    string      `json:"func,omitempty"`
	Gid     string      `json:"gid,omitempty"`
	Content interface{} `json:"content,omitempty"`
	TraceId string      `json:"traceId,omitempty"`
	Stack   string      `json:"stack,omitempty"`
}
