package logrus

import (
	"bytes"
	"fmt"
	"sort"
)

type ClassicFormatter struct {
	TimestampFormat string
	FieldsDelimiter string
}

func (f *ClassicFormatter) Format(entry *Entry) ([]byte, error) {
	b := &bytes.Buffer{}
	// write [%time] [%level] %message
	if f.TimestampFormat == "" {
		f.TimestampFormat = DefaultTimestampFormat
	}
	if f.FieldsDelimiter == "" {
		f.FieldsDelimiter = " "
	}

	if entry.Level.String() == "info" {
		fmt.Fprintf(b, "%s [%s] %s",
			entry.Time.Format(f.TimestampFormat), entry.Level.String(), entry.Message)
	} else {
		fmt.Fprintf(b, "%s [%s][%s:%d][%s] %s",
			entry.Time.Format(f.TimestampFormat), entry.Level.String(),
			entry.Data["file"], entry.Data["line"], entry.Data["func"],
			entry.Message)
	}

	// sort fields
	keys := make([]string, 0, len(entry.Data))
	for key := range entry.Data {
		if key != "func" && key != "file" && key != "line" {
			keys = append(keys, key)
		}
	}

	if len(keys) != 0 {
		b.WriteString(" .entry{")
		// append fields
		sort.Strings(keys)
		for idx := range keys {
			appendKeyValue(b, keys[idx], entry.Data[keys[idx]])
			if idx != (len(keys) - 1) {
				fmt.Fprint(b, f.FieldsDelimiter)
			}
		}
		b.WriteString("}")
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}

// 自定义格式的日志输出
func (f *ClassicFormatter) FormatEx(entry *Entry) ([]byte, error) {
	b := &bytes.Buffer{}
	// write [%time] [%level] %message
	if f.TimestampFormat == "" {
		f.TimestampFormat = DefaultTimestampFormat
	}
	if f.FieldsDelimiter == "" {
		f.FieldsDelimiter = " "
	}
	fmt.Fprintf(b, "[%s]%s", entry.Time.Format(f.TimestampFormat), entry.Message)

	// sort fields
	keys := make([]string, 0, len(entry.Data))
	for key := range entry.Data {
		if key != "func" && key != "file" && key != "line" {
			keys = append(keys, key)
		}
	}

	if len(keys) != 0 {
		b.WriteString(" .entry{")
		// append fields
		sort.Strings(keys)
		for idx := range keys {
			appendKeyValue(b, keys[idx], entry.Data[keys[idx]])
			if idx != (len(keys) - 1) {
				fmt.Fprint(b, f.FieldsDelimiter)
			}
		}
		b.WriteString("}")
	}

	b.WriteByte('\n')
	return b.Bytes(), nil
}

func appendKeyValue(b *bytes.Buffer, key, value interface{}) {
	switch value.(type) {
	case string:
		if needsQuoting(value.(string)) {
			fmt.Fprintf(b, "%v=%s", key, value)
		} else {
			fmt.Fprintf(b, "%v=%q", key, value)
		}
	case error:
		if needsQuoting(value.(error).Error()) {
			fmt.Fprintf(b, "%v=%s", key, value)
		} else {
			fmt.Fprintf(b, "%v=%q", key, value)
		}
	default:
		fmt.Fprintf(b, "%v=%v", key, value)
	}
}
