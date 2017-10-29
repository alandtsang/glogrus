# glogrus
logrus

## Set config
eg.
```
"log_dir":"/tmp",
"log_level":4
```

cat /tmp/test.log
```
2017-10-29 23:17:18 [info][main.go:34][main] Caught a signal: 12
```

## Log level

debug: set `log_level` to 5

```
2017-10-29 23:24:36 [debug][main.go:33][main] Debug value: 12
2017-10-29 23:24:36 [info][main.go:34][main] Caught a signal: 12
```

