package logrus

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	defaultTimeFormat  = "20060102150405"
	defaultMaxSize     = uint64(100 << 20)
	defaultMaxDuration = time.Duration(1 * time.Hour)
)

type FileRotator struct {
	fd    *os.File
	id    uint64
	size  uint64
	ctime time.Time

	FileName    string
	MaxSize     uint64
	MaxDuration time.Duration
	TimeFormat  string
}

func (r *FileRotator) fileName() string {
	if r.FileName != "" {
		return r.FileName
	}
	// default: create logs in the working directory
	return filepath.Join(".", filepath.Base(os.Args[0])+".log")
}

func (r *FileRotator) maxSize() uint64 {
	if r.MaxSize > 0 {
		return r.MaxSize
	}
	return defaultMaxSize
}

func (r *FileRotator) maxDuration() time.Duration {
	if r.MaxDuration > 0 {
		return r.MaxDuration
	}
	return defaultMaxDuration
}

func (r *FileRotator) timeFormat() string {
	if r.TimeFormat != "" {
		return r.TimeFormat
	}
	return defaultTimeFormat
}

func (r *FileRotator) init() error {
	r.id = 0
	r.size = 0
	r.ctime = time.Now()
	return r.open()
}

func (r *FileRotator) rotate() error {
	r.id++
	r.size = 0
	r.ctime = time.Now()
	return r.open()
}

func (r *FileRotator) open() error {
	if err := r.close(); err != nil {
		return err
	}
	// open as a new file anyway
	name := fmt.Sprintf("%s_%s_%v", r.fileName(), r.ctime.Format(r.timeFormat()), r.id)
	flag := os.O_CREATE | os.O_WRONLY | os.O_APPEND
	mode := os.FileMode(0644)
	if fd, err := os.OpenFile(name, flag, mode); err != nil {
		return fmt.Errorf("can't open new logfile: %s", err)
	} else {
		r.fd = fd
		if _, err := os.Stat(r.fileName()); err == nil {
			os.Remove(r.fileName())
		}

		os.Symlink(name, r.fileName())
	}

	return nil
}

func (r *FileRotator) close() error {
	if r.fd == nil {
		return nil
	}
	err := r.fd.Close()
	r.fd = nil
	return err
}

func (r *FileRotator) Write(p []byte) (n int, err error) {
	// check write length
	writeLen := uint64(len(p))
	if writeLen > r.maxSize() {
		return 0, fmt.Errorf(
			"write length %d exceeds maximum file size %d", writeLen, r.maxSize(),
		)
	}
	// trigger the initial open
	if r.fd == nil {
		if err = r.init(); err != nil {
			return 0, err
		}
	}
	// trigger log rotate
	if r.size+writeLen > r.maxSize() || time.Now().Sub(r.ctime) > r.maxDuration() {
		if err := r.rotate(); err != nil {
			return 0, err
		}
	}
	// write
	n, err = r.fd.Write(p)
	r.size += uint64(n)

	return n, err
}

func (r *FileRotator) Close() error {
	return r.close()
}
