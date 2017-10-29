package logrus

import (
	"errors"
	"runtime"
	"strings"
)

func Trace(depth int) (Fields, error) {
	depth = 2
	pc, file, line, ok := runtime.Caller(depth)
	if !ok {
		return nil, errors.New("runtime caller failed.")
	}

	path := strings.Split(file, "/")
	fname := path[len(path)-1]
	funcPath := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	funcName := funcPath[len(funcPath)-1]

	fields := make(Fields, 3)
	fields["func"] = funcName
	fields["file"] = fname
	fields["line"] = line

	return fields, nil
}
