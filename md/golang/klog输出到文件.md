# klog输出到文件
可以通过自定义writer的方式，将klog的日志输出到文件。下面是一个示例：
将不同日志输出到不同的文件。
```go
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
	// Write all logs to respective files
	switch string(p[:1]) {
	case "I": // Info log starts with "I"
		return w.infoLog.Write(p)
	case "W": // Warning log starts with "W"
		return w.warnLog.Write(p)
	case "E": // Error log starts with "E"
		return w.errorLog.Write(p)
	case "F": // Fatal log starts with "F"
		return w.fatalLog.Write(p)
	default: // Default write logs to info
		return w.infoLog.Write(p)
	}
}

func Close() {
	_ = writer.infoLog.Close()
	_ = writer.warnLog.Close()
	_ = writer.errorLog.Close()
	_ = writer.fatalLog.Close()
}

func InitLogger() {
	err := os.Mkdir(constants.DefaultLogPath, 0777)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}
	// Open log files
	infoLog, err := os.OpenFile(constants.DefaultLogPath+"/info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	warnLog, err := os.OpenFile(constants.DefaultLogPath+"/warn.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	errorLog, err := os.OpenFile(constants.DefaultLogPath+"/error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	fatalLog, err := os.OpenFile(constants.DefaultLogPath+"/fatal.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
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
	err = flag.Set("one_output", "true")
	if err != nil {
		panic(err)
	}

	err = flag.Set("logtostderr", "false")
	if err != nil {
		panic(err)
	}

	if os.Getenv("ENV") != "live" {
		if err := flag.Set("alsologtostderr", "true"); err != nil {
			panic(err)
		}
	}
}
```