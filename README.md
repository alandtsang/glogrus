# glogrus

[![License](https://img.shields.io/badge/license-Apache%202-4EB1BA.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)

logrus common format

## Set config
Default log level is 4, meaning info, other level print detailed file location, line number, and the function name.

eg.
```
"log_dir":"/tmp",
"log_level":4
```

cat /tmp/test.log
```
2017-10-31 10:36:22 [info] Caught a signal: 12
```

## Log level

- Fatal: 1
- Error: 2
- Warn: 3
- Info: 4
- Debug: 5


```
"log_dir":"/tmp",
"log_level":5
```

```
2017-10-31 11:58:13 [error][main.go:33][main] Error Level
2017-10-31 11:58:13 [warning][main.go:34][main] Warn Level
2017-10-31 11:58:13 [info] Caught a signal: 12
2017-10-31 11:58:13 [debug][main.go:36][main] Debug value: 12
```

# Get Help

The fastest way to get response is to send email to my mail:
- <zengxianglong0@gmail.com>

# License

Please refer to [LICENSE](https://github.com/alandtsang/glogrus/blob/master/LICENSE) file.

