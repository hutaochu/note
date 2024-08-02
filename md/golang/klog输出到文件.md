# klog输出日志到文件
可以通过自定义writer的方式，使klog的日志输出到指定的文件中，方便采集和配置相关的告警策略。
## 自定义writer
```go
type LogWriter struct {
	mu       sync.Mutex
	infoLog  *os.File
	warnLog  *os.File
	errorLog *os.File
	fatalLog *os.File
}

var writer *LogWriter

func (w *LogWriter) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	// Simple example: Write all logs to respective files
	if string(p[:1]) == "I" { // Info log starts with "I"
		return w.infoLog.Write(p)
	} else if string(p[:1]) == "W" { // Warning log starts with "W"
		return w.warnLog.Write(p)
	} else if string(p[:1]) == "E" { // Error log starts with "E"
		return w.errorLog.Write(p)
	} else if string(p[:1]) == "F" { // Fatal log starts with "F"
		return w.fatalLog.Write(p)
	}
	return len(p), nil
}
```
通过klog日志的起始字符，判断日志的level，然后输出到指定的文件。

## 初始化文件
```go
	// Open log files
	infoLog, err := os.OpenFile("info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	warnLog, err := os.OpenFile("warn.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	errorLog, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	fatalLog, err := os.OpenFile("fatal.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
```

## 设置klog的writer
```go
    writer = &LogWriter{
		infoLog:  infoLog,
		warnLog:  warnLog,
		errorLog: errorLog,
		fatalLog: fatalLog,
	}

	// Set klog output to custom writer
	klog.SetOutput(writer)
```

## 设置klog配置
```go
    // 一定要先调用InitFlags方法，然后才使用flag.Set方法修改klog的配置
    klog.InitFlags(nil)
	klog.LogToStderr(false)
	flag.Set("one_output", "true")
    flag.Set("logtostderr", "false")
```

## 完整代码
```go
import (
	"flag"
	"k8s.io/klog/v2"
	"os"
	"sync"
)

// LogWriter wraps file descriptors for different log levels
type LogWriter struct {
	mu       sync.Mutex
	infoLog  *os.File
	warnLog  *os.File
	errorLog *os.File
	fatalLog *os.File
}

var writer *LogWriter

func (w *LogWriter) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	// Simple example: Write all logs to respective files
	if string(p[:1]) == "I" { // Info log starts with "I"
		return w.infoLog.Write(p)
	} else if string(p[:1]) == "W" { // Warning log starts with "W"
		return w.warnLog.Write(p)
	} else if string(p[:1]) == "E" { // Error log starts with "E"
		return w.errorLog.Write(p)
	} else if string(p[:1]) == "F" { // Fatal log starts with "F"
		return w.fatalLog.Write(p)
	}
	return len(p), nil
}

func Close() {
	_ = writer.infoLog.Close()
	_ = writer.warnLog.Close()
	_ = writer.errorLog.Close()
	_ = writer.fatalLog.Close()
}

func InitLogger() {
	// Open log files
	infoLog, err := os.OpenFile("info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	warnLog, err := os.OpenFile("warn.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	errorLog, err := os.OpenFile("error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	fatalLog, err := os.OpenFile("fatal.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}

	writer = &LogWriter{
		infoLog:  infoLog,
		warnLog:  warnLog,
		errorLog: errorLog,
		fatalLog: fatalLog,
	}

	// Set klog output to custom writer
	klog.SetOutput(writer)
	klog.InitFlags(nil)
	klog.LogToStderr(false)
	flag.Set("one_output", "true")
    flag.Set("logtostderr", "false")
}
```