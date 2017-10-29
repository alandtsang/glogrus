package logmgr

import (
	"errors"
	"path"
	"time"

	"../config"

	Log "../third_party/logrus"
)

var LOG *Log.Logger
var CACHE *Log.Logger

/*******************************************************************************
* Description: LOG初始化
*       Input: logPath  log文件路径
*      Output:
*      Return:
*      Others:
*******************************************************************************/
func InitLog() error {
	output := &Log.FileRotator{
		FileName:    path.Join(config.CONF.LogDir, "test.log"),
		MaxSize:     100 << 20,
		MaxDuration: 1 * time.Hour,
		TimeFormat:  "2006010215",
	}

	formatter := &Log.ClassicFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FieldsDelimiter: ", ",
	}

	LOG = &Log.Logger{
		Out:       output,
		Formatter: formatter,
		Hooks:     nil,
		Level:     Log.Level(config.CONF.LogConfig.LogLevel),
	}

	if LOG == nil {
		return errors.New("New logger failed")
	}

	cacheOut := &Log.FileRotator{
		FileName:    path.Join(config.CONF.LogDir, "vod_cache.log"),
		MaxSize:     100 << 20,
		MaxDuration: 1 * time.Hour,
		TimeFormat:  "2006010215",
	}

	CACHE = &Log.Logger{
		Out:       cacheOut,
		Formatter: formatter,
		Hooks:     nil,
		Level:     Log.Level(config.CONF.LogConfig.LogLevel),
	}

	if CACHE == nil {
		return errors.New("New cache logger failed")
	}

	return nil
}
