# logger [![](https://github.com/fl1ckje/logger/workflows/build/badge.svg)](https://github.com/fl1ckje/logger/actions)

Pretty simple colorful logger based on slog package.

![Screenshot](https://github.com/user-attachments/assets/cf6a37f9-c9b8-4fca-b435-53765c70952d)

## Install

```
go get github.com/fl1ckje/logger
```

## Examples


```go
// Messages without JSON data
logger.Info("Starting server.")
logger.Debug("debug msg test")
logger.Info("info msg test")
logger.Warn("warn msg test")
logger.Error("error msg test")

// Messages with JSON data
logger.Debug("debug msg w json test", slog.String("string kv", "string value"))
logger.Info("info msg w json test", slog.Bool("bool kv", true))
logger.Warn("warn msg w json test", slog.Time("time kv", time.Now()))
logger.Error("error msg w json test", slog.Any("any kv", map[string]int{"user": 0, "id": 1, "ext_data": 2}))
```

## Credits

* Color support via @jwalton: [gchalk](https://github.com/jwalton/gchalk)
