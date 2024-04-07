package logging


import (
	"fmt"
	"log"
	"os"
	"runtime"
)

func CreateLogger() {
	_, err := os.Stat("logger.txt")
	if os.IsNotExist(err) {
		file, err := os.Create("logger.txt")
		defer file.Close()
		if err != nil {
			log.Println("Failed to create a logger file")
			return
		}
	}
}

func Logger(message string) bool {
	successLog := false
	file, openErr := os.OpenFile("logger.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	defer file.Close()
	if openErr != nil {
		log.Println("Failed to open the logger")
		return successLog
	}

	pc, filename, line, ok := runtime.Caller(1)
	if ok {
		logMsg := fmt.Sprintf("Activity in %s[%s:%d]: %s", runtime.FuncForPC(pc).Name(), filename, line, message+"\n")
		_, err := fmt.Fprintln(file, logMsg)
		if err != nil {
			log.Println("Failed to write to the logger")
			return successLog
		}
		successLog = true
	}
	return successLog
}


// // LogLevel represents different levels of logging.
// type LogLevel int

// const (
// 	// LogLevelInfo represents informational messages.
// 	LogLevelInfo LogLevel = iota
// 	// LogLevelWarning represents warning messages.
// 	LogLevelWarning
// 	// LogLevelError represents error messages.
// 	LogLevelError
// )

// var (
// 	logFile     *os.File
// 	loggerMutex sync.Mutex
// )

// // ConfigureLogger configures the logger with the specified log file path.
// func ConfigureLogger(logFilePath string) error {
// 	loggerMutex.Lock()
// 	defer loggerMutex.Unlock()

// 	// Open the log file for writing
// 	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
// 	if err != nil {
// 		return fmt.Errorf("failed to open log file: %v", err)
// 	}
// 	logFile = file
// 	log.SetOutput(logFile)
// 	return nil
// }

// // CloseLogger closes the logger.
// func CloseLogger() {
// 	loggerMutex.Lock()
// 	defer loggerMutex.Unlock()

// 	if logFile != nil {
// 		logFile.Close()
// 	}
// }

// // Log logs the provided message with the given log level.
// // Log logs the provided message with the given log level to the log file.
// func Log(level LogLevel, message string) {
//     loggerMutex.Lock()
//     defer loggerMutex.Unlock()

//     // Get information about the caller
//     pc, filename, line, _ := runtime.Caller(1)
//     callerFunc := runtime.FuncForPC(pc).Name()

//     // Format the log message
//     logMsg := fmt.Sprintf("[%s] %s [%s:%d] - %s %s", levelToString(level), time.Now().Format(time.RFC3339), filename, line, message, callerFunc)

//     // Write the log message to the log file
//     if logFile != nil {
//         _, err := fmt.Fprintln(logFile, logMsg)
//         if err != nil {
//             log.Println("Failed to write to the logger file:", err)
//         }
//     } else {
//         log.Println("Log file is not configured. Cannot write log:", logMsg)
//     }

//     // Additionally, you can also log to stderr for terminal output
//     switch level {
//     case LogLevelInfo:
//         log.Println(logMsg)
//     case LogLevelWarning:
//         log.Println("WARNING:", logMsg)
//     case LogLevelError:
//         log.Println("ERROR:", logMsg)
//     }
// }


// // levelToString converts LogLevel to a string representation.
// func levelToString(level LogLevel) string {
// 	switch level {
// 	case LogLevelInfo:
// 		return "INFO"
// 	case LogLevelWarning:
// 		return "WARNING"
// 	case LogLevelError:
// 		return "ERROR"
// 	default:
// 		return "UNKNOWN"
// 	}
// }
